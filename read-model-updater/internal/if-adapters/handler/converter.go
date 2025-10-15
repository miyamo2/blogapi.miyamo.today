package handler

import (
	"context"
	"iter"

	"blogapi.miyamo.today/read-model-updater/internal/app/usecase"
	"github.com/aws/aws-sdk-go-v2/service/dynamodbstreams/types"
)

// ToSyncUsecaseInDtoConverter is an interface for the converter of the usecase.SyncUsecaseInDto
type ToSyncUsecaseInDtoConverter interface {
	ToSyncUsecaseInDtoSeq(
		ctx context.Context, records []types.Record,
	) iter.Seq2[int, usecase.SyncUsecaseInDto]
}
