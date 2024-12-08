package lambda

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/app/usecase"
	"iter"
)

// ToSyncUsecaseInDtoConverter is an interface for the converter of the usecase.SyncUsecaseInDto
type ToSyncUsecaseInDtoConverter interface {
	ToSyncUsecaseInDtoSeq(ctx context.Context, records []events.DynamoDBEventRecord) iter.Seq2[int, usecase.SyncUsecaseInDto]
}
