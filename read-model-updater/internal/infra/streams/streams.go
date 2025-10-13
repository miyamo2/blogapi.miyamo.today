package streams

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodbstreams"
)

type Client interface {
	ListStreams(
		ctx context.Context, params *dynamodbstreams.ListStreamsInput, optFns ...func(*dynamodbstreams.Options),
	) (*dynamodbstreams.ListStreamsOutput, error)
	DescribeStream(
		ctx context.Context, params *dynamodbstreams.DescribeStreamInput, optFns ...func(*dynamodbstreams.Options),
	) (*dynamodbstreams.DescribeStreamOutput, error)
	GetRecords(
		ctx context.Context, params *dynamodbstreams.GetRecordsInput, optFns ...func(*dynamodbstreams.Options),
	) (*dynamodbstreams.GetRecordsOutput, error)
	GetShardIterator(
		ctx context.Context, params *dynamodbstreams.GetShardIteratorInput, optFns ...func(*dynamodbstreams.Options),
	) (*dynamodbstreams.GetShardIteratorOutput, error)
}
