package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"

	blogapictx "blogapi.miyamo.today/core/context"
	"blogapi.miyamo.today/read-model-updater/internal/configs/di"
	"blogapi.miyamo.today/read-model-updater/internal/if-adapters/handler"
	"blogapi.miyamo.today/read-model-updater/internal/infra/queue"
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	"github.com/avast/retry-go"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/cockroachdb/errors"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"
)

var dependencies = di.GetDependecies()

func main() {
	messageCh := make(chan types.Message, 1)
	errCh := make(chan error, 1)
	workerQueue := make(chan types.Message, 1)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	go polling(ctx, messageCh, errCh, dependencies.QueueClient, dependencies.QueueURL)
	go work(ctx, workerQueue, dependencies.NewRelicApp, dependencies.QueueClient, dependencies.QueueURL, dependencies.SyncHandler)

	for {
		select {
		case <-ctx.Done():
			err := context.Cause(ctx)
			if err != nil && !errors.Is(err, context.Canceled) {
				os.Exit(1)
			}
		case err := <-errCh:
			slog.Default().ErrorContext(
				ctx,
				"error occurred during polling",
				slog.String("queue_url", *dependencies.QueueURL),
				slog.String("error", err.Error()),
			)
		case message := <-messageCh:
			workerQueue <- message
		}
	}
}

func polling(
	ctx context.Context,
	messageCh chan<- types.Message,
	errCh chan<- error,
	queueClient queue.Client,
	queueURL *string,
) {
	c := time.Tick(time.Second)
	for range c {
		select {
		case <-ctx.Done():
			err := context.Cause(ctx)
			if err != nil && !errors.Is(err, context.Canceled) {
				errCh <- errors.Wrap(err, "context done")
			}
			return
		default:
			message, err := queueClient.ReceiveMessage(
				ctx, &sqs.ReceiveMessageInput{
					MessageSystemAttributeNames: []types.MessageSystemAttributeName{
						types.MessageSystemAttributeNameSentTimestamp,
					},
					MessageAttributeNames: []string{"dynamodb.NewImage"},
					QueueUrl:              queueURL,
					MaxNumberOfMessages:   10,
				},
			)
			if err != nil {
				errCh <- errors.Wrap(err, "failed to receive message from queue")
				continue
			}
			for _, msg := range message.Messages {
				messageCh <- msg
			}
		}
	}

}

func work(
	ctx context.Context,
	queue <-chan types.Message,
	nr *newrelic.Application,
	queueClient queue.Client,
	queueURL *string,
	syncHandler *handler.SyncHandler,
) {
	var wg sync.WaitGroup
	for {
		select {
		case <-ctx.Done():
			err := context.Cause(ctx)
			if err != nil && !errors.Is(err, context.Canceled) {
				slog.Default().ErrorContext(
					ctx,
					"worker context canceled with error",
					slog.String("error", err.Error()),
				)
			}
			wg.Wait()
			return
		case msg := <-queue:
			wg.Go(
				func() {
					tx := nr.StartTransaction("stream-worker")
					defer tx.End()

					blogAPICtx := blogapictx.New(*msg.MessageId, "", "queue", nil, nil)
					ctx := newrelic.NewContext(blogapictx.StoreToContext(context.Background(), blogAPICtx), tx)

					if msg.Body == nil {
						tx.NoticeError(nrpkgerrors.Wrap(errors.New("message body is nil")))
						return
					}
					body := []byte(*msg.Body)

					sentTs, err := strconv.ParseInt(
						msg.Attributes[string(types.MessageSystemAttributeNameSentTimestamp)], 10, 64,
					)
					if err != nil {
						tx.NoticeError(nrpkgerrors.Wrap(errors.Wrap(err, "failed to parse sent timestamp")))
						return
					}
					eventAt := synchro.UnixMilli[tz.UTC](sentTs)

					err = retry.Do(
						func() error {
							err := syncHandler.Invoke(ctx, body, eventAt)
							if err != nil {
								return errors.WithStack(err)
							}
							return nil
						},
					)
					if err != nil {
						tx.NoticeError(nrpkgerrors.Wrap(errors.Wrap(err, "failed to process message in worker")))
						return
					}
					err = retry.Do(
						func() error {
							_, err = queueClient.DeleteMessage(
								ctx, &sqs.DeleteMessageInput{
									QueueUrl:      queueURL,
									ReceiptHandle: msg.ReceiptHandle,
								},
							)
							if err != nil {
								return errors.WithStack(err)
							}
							return nil
						},
					)
					if err != nil {
						tx.NoticeError(nrpkgerrors.Wrap(errors.Wrap(err, "failed to delete message from queue")))
					}
					return
				},
			)
		}
	}
}
