package handler

import (
	"context"

	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// SyncHandler is a handler for the Sync usecase.
type SyncHandler struct {
	syncUsecaseConverter ToSyncUsecaseInDtoConverter
	syncUsecase          SyncUsecase
}

// Invoke invokes the Sync usecase.
func (h *SyncHandler) Invoke(ctx context.Context, body []byte, eventAt synchro.Time[tz.UTC]) error {
	tx := newrelic.FromContext(ctx)
	defer tx.StartSegment("SyncHandler#Invoke").End()

	dto, err := h.syncUsecaseConverter.ToSyncUsecaseInDto(ctx, body, eventAt)
	if err != nil {
		return err
	}
	return h.syncUsecase.SyncBlogSnapshotWithEvents(ctx, dto)
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
