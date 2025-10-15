package handler

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodbstreams/types"
)

// SyncHandler is a handler for the Sync usecase.
type SyncHandler struct {
	syncUsecaseConverter ToSyncUsecaseInDtoConverter
	syncUsecase          SyncUsecase
}

// Invoke invokes the Sync usecase.
func (h *SyncHandler) Invoke(ctx context.Context, records []types.Record) error {
	dtoSeq := h.syncUsecaseConverter.ToSyncUsecaseInDtoSeq(ctx, records)
	return h.syncUsecase.SyncBlogSnapshotWithEvents(ctx, dtoSeq)
}

// NewSyncHandler creates a new SyncHandler.
func NewSyncHandler(
	syncUsecaseConverter ToSyncUsecaseInDtoConverter,
	syncUsecase SyncUsecase,
) *SyncHandler {
	return &SyncHandler{
		syncUsecaseConverter: syncUsecaseConverter,
		syncUsecase:          syncUsecase,
	}
}
