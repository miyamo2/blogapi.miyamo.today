package lambda

import (
	"context"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/app/usecase"
	"iter"
)

// SyncUsecase is an interface for the usecase of synchronize blog snapshot with events.
type SyncUsecase interface {
	// SyncBlogSnapshotWithEvents synchronizes blog snapshot with events.
	SyncBlogSnapshotWithEvents(ctx context.Context, in iter.Seq2[int, usecase.SyncUsecaseInDto]) error
}
