package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	blogapictx "blogapi.miyamo.today/core/context"
	"blogapi.miyamo.today/read-model-updater/internal/configs/di"
	"github.com/avast/retry-go"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodbstreams"
	"github.com/aws/aws-sdk-go-v2/service/dynamodbstreams/types"
	"github.com/cockroachdb/errors"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"
	"golang.org/x/sync/errgroup"
)

var dependencies = di.GetDependecies()

type RecordsInfo struct {
	shardID       string
	shardIterator string
	records       []types.Record
}

func main() {
	recordsInfoCh := make(chan RecordsInfo, 1)
	errCh := make(chan error, 1)
	workerQueue := make(chan RecordsInfo, 1)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	go polling(ctx, recordsInfoCh, errCh)
	go work(ctx, workerQueue)

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
				slog.String("stream_arn", *dependencies.StreamARN),
				slog.String("error", err.Error()),
			)
		case recordsInfo := <-recordsInfoCh:
			workerQueue <- recordsInfo
		}
	}
}

func polling(ctx context.Context, recordsInfoCh chan<- RecordsInfo, errCh chan<- error) {
	for {
		select {
		case <-ctx.Done():
			err := context.Cause(ctx)
			if err != nil && !errors.Is(err, context.Canceled) {
				errCh <- errors.Wrap(err, "context done")
			}
			return
		default:
			describeStreamOutput, err := dependencies.StreamClient.DescribeStream(
				ctx, &dynamodbstreams.DescribeStreamInput{
					StreamArn: dependencies.StreamARN,
				},
			)
			if err != nil {
				errCh <- errors.Wrap(err, "failed to describe stream")
			}
			eg, egCtx := errgroup.WithContext(ctx)
			for _, shard := range describeStreamOutput.StreamDescription.Shards {
				select {
				case <-egCtx.Done():
					err := context.Cause(egCtx)
					errCh <- errors.Wrap(err, "error processing shards")
				default:
					eg.Go(
						func() error {
							shardID := *shard.ShardId
							slog.Default().InfoContext(
								egCtx,
								"processing shard",
								slog.String("shard_id", shardID),
							)
							defer slog.Default().InfoContext(
								egCtx,
								"processed shard",
								slog.String("shard_id", shardID),
							)

							getShardIteratorOutput, err := dependencies.StreamClient.GetShardIterator(
								egCtx, &dynamodbstreams.GetShardIteratorInput{
									StreamArn:         dependencies.StreamARN,
									ShardId:           &shardID,
									ShardIteratorType: types.ShardIteratorTypeTrimHorizon,
								},
							)
							if err != nil {
								return err
							}
							shardIterator := getShardIteratorOutput.ShardIterator

							for shardIterator != nil {
								slog.Default().InfoContext(
									egCtx,
									"shard iterator",
									slog.String("shard_id", shardID),
									slog.String("shard_iterator", *shardIterator),
								)
								getRecordsOutput, err := dependencies.StreamClient.GetRecords(
									egCtx, &dynamodbstreams.GetRecordsInput{
										ShardIterator: shardIterator,
										Limit:         aws.Int32(1000),
									},
								)
								if err != nil {
									return err
								}
								recordsInfoCh <- RecordsInfo{
									shardID:       shardID,
									shardIterator: *shardIterator,
									records:       getRecordsOutput.Records,
								}
								shardIterator = getRecordsOutput.NextShardIterator
							}
							return nil
						},
					)
				}
			}
			if err := eg.Wait(); err != nil {
				errCh <- errors.Wrap(err, "error processing shards")
			}
		}
	}
}

func work(ctx context.Context, queue <-chan RecordsInfo) {
	eg, egCtx := errgroup.WithContext(ctx)
	eg.SetLimit(5)
	for {
		select {
		case <-egCtx.Done():
			err := context.Cause(ctx)
			if err != nil && !errors.Is(err, context.Canceled) {
				slog.Default().ErrorContext(
					ctx,
					"worker context done",
					slog.String("error", err.Error()),
				)
			}
			return
		case msg := <-queue:
			eg.Go(
				func() error {
					return retry.Do(
						func() error {
							tx := dependencies.NewRelicApp.StartTransaction("stream-worker")
							defer tx.End()

							blogAPICtx := blogapictx.New(msg.shardID, msg.shardIterator, "queue", nil, nil)
							ctx := newrelic.NewContext(blogapictx.StoreToContext(egCtx, blogAPICtx), tx)

							err := dependencies.SyncHandler.Invoke(ctx, msg.records)
							if err != nil {
								slog.Default().ErrorContext(
									ctx,
									"failed to process stream in worker",
									slog.String("error", err.Error()),
								)
								tx.NoticeError(nrpkgerrors.Wrap(err))
							}
							return nil
						},
					)
				},
			)
		}
	}
}
