package main

import (
	"context"
	"log/slog"
	"os"

	"blogapi.miyamo.today/core/log"
	"blogapi.miyamo.today/read-model-updater/internal/configs/di"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodbstreams"
	"github.com/aws/aws-sdk-go-v2/service/dynamodbstreams/types"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"
	"golang.org/x/sync/errgroup"
)

func main() {
	dependencies := di.GetDependecies()
	slog.SetDefault(log.New())

	tx := dependencies.NewRelicApp.StartTransaction(*dependencies.StreamARN)
	ctx := newrelic.NewContext(context.Background(), tx)

	describeStreamOutput, err := dependencies.StreamClient.DescribeStream(
		ctx, &dynamodbstreams.DescribeStreamInput{
			StreamArn: dependencies.StreamARN,
		},
	)
	if err != nil {
		slog.Default().Error("failed to describe stream", slog.Any("error", err))
	}
	eg, egCtx := errgroup.WithContext(ctx)
	for _, shard := range describeStreamOutput.StreamDescription.Shards {
		select {
		case <-egCtx.Done():
			err := context.Cause(egCtx)
			slog.Default().Error("error processing shards", slog.Any("error", err))
			tx.NoticeError(nrpkgerrors.Wrap(err))
			os.Exit(1)
		default:
			eg.Go(
				func() error {
					shardID := *shard.ShardId
					slog.Default().Info("processing shard", slog.String("shard_id", shardID))

					getShardIteratorOutput, err := dependencies.StreamClient.GetShardIterator(
						ctx, &dynamodbstreams.GetShardIteratorInput{
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
						getRecordsOutput, err := dependencies.StreamClient.GetRecords(
							ctx, &dynamodbstreams.GetRecordsInput{
								ShardIterator: shardIterator,
								Limit:         aws.Int32(1000),
							},
						)
						if err != nil {
							return err
						}
						err = dependencies.SyncHandler.Invoke(ctx, getRecordsOutput.Records)
						if err != nil {
							return err
						}
						shardIterator = getRecordsOutput.NextShardIterator
					}
					return nil
				},
			)
		}
	}
	if err := eg.Wait(); err != nil {
		slog.Default().Error("error processing shards", slog.Any("error", err))
		tx.NoticeError(nrpkgerrors.Wrap(err))
		os.Exit(1)
	}
}
