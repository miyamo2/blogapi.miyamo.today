package handler

import (
	"context"
	"iter"

	"blogapi.miyamo.today/read-model-updater/internal/app/usecase"
)

// SyncUsecase is an interface for the usecase of synchronize blog snapshot with events.
type SyncUsecase interface {
	// SyncBlogSnapshotWithEvents synchronizes blog snapshot with events.
	SyncBlogSnapshotWithEvents(ctx context.Context, in iter.Seq2[int, usecase.SyncUsecaseInDto]) error
}
