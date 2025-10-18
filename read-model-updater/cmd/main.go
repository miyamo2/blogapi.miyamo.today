package main

import (
	"context"
	"log/slog"
	"os"

	blogapictx "blogapi.miyamo.today/core/context"
	"blogapi.miyamo.today/read-model-updater/internal/configs/di"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodbstreams"
	"github.com/aws/aws-sdk-go-v2/service/dynamodbstreams/types"
	"github.com/cockroachdb/errors"
	"github.com/miyamo2/altnrslog"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"
	"golang.org/x/sync/errgroup"
)

func main() {
	var code int
	defer func() {
		os.Exit(code)
	}()

	dependencies := di.GetDependecies()

	tx := dependencies.NewRelicApp.StartTransaction(*dependencies.StreamARN)
	defer tx.End()
	ctx := newrelic.NewContext(context.Background(), tx)
	logger := slog.New(altnrslog.NewTransactionalHandler(dependencies.NewRelicApp, tx))
	slog.SetDefault(logger)

	if err := do(ctx, dependencies); err != nil {
		logger.ErrorContext(
			ctx,
			"failed to process stream",
			slog.String("stream_arn", *dependencies.StreamARN),
			slog.String("error", err.Error()),
		)
		tx.NoticeError(nrpkgerrors.Wrap(err))
		code = 1
	}
}

func do(ctx context.Context, dependencies *di.Dependencies) error {
	describeStreamOutput, err := dependencies.StreamClient.DescribeStream(
		ctx, &dynamodbstreams.DescribeStreamInput{
			StreamArn: dependencies.StreamARN,
		},
	)
	if err != nil {
		return errors.Wrap(err, "failed to describe stream")
	}
	eg, egCtx := errgroup.WithContext(ctx)
	for _, shard := range describeStreamOutput.StreamDescription.Shards {
		select {
		case <-egCtx.Done():
			err := context.Cause(egCtx)
			return errors.Wrap(err, "error processing shards")
		default:
			eg.Go(
				func() error {
					shardID := *shard.ShardId
					slog.Default().InfoContext(egCtx, "processing shard", slog.String("shard_id", shardID))
					defer slog.Default().InfoContext(egCtx, "processed shard", slog.String("shard_id", shardID))

					blogAPICtx := blogapictx.New(shardID, *dependencies.StreamARN, "queue", nil, nil)
					egCtx := blogapictx.StoreToContext(egCtx, blogAPICtx)

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

					for {
						getRecordsOutput, err := dependencies.StreamClient.GetRecords(
							egCtx, &dynamodbstreams.GetRecordsInput{
								ShardIterator: shardIterator,
								Limit:         aws.Int32(1000),
							},
						)
						if err != nil {
							return err
						}
						err = dependencies.SyncHandler.Invoke(egCtx, getRecordsOutput.Records)
						if err != nil {
							return err
						}
						shardIterator = getRecordsOutput.NextShardIterator
						if shardIterator == nil {
							break
						}
					}
					return nil
				},
			)
		}
	}
	if err := eg.Wait(); err != nil {
		return errors.Wrap(err, "error processing shards")
	}
	return nil
}
