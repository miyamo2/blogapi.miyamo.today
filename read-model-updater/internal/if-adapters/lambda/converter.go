package lambda

import (
	"context"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/app/usecase"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/if-adapters/model"
	"iter"
)

// ToSyncUsecaseInDtoConverter is an interface for the converter of the usecase.SyncUsecaseInDto
type ToSyncUsecaseInDtoConverter interface {
	ToSyncUsecaseInDtoSeq(ctx context.Context, records []model.Record) iter.Seq2[int, usecase.SyncUsecaseInDto]
}
