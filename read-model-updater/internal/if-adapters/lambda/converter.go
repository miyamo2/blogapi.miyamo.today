package lambda

import (
	"blogapi.miyamo.today/read-model-updater/internal/app/usecase"
	"context"
	"github.com/aws/aws-lambda-go/events"
	"iter"
)

// ToSyncUsecaseInDtoConverter is an interface for the converter of the usecase.SyncUsecaseInDto
type ToSyncUsecaseInDtoConverter interface {
	ToSyncUsecaseInDtoSeq(ctx context.Context, records []events.DynamoDBEventRecord) iter.Seq2[int, usecase.SyncUsecaseInDto]
}
