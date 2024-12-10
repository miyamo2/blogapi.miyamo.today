package lambda

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"
	"log/slog"
)

// SyncHandler is a handler for the Sync usecase.
type SyncHandler struct {
	syncUsecaseConverter ToSyncUsecaseInDtoConverter
	syncUsecase          SyncUsecase
}

// Invoke invokes the Sync usecase.
func (h *SyncHandler) Invoke(ctx context.Context, stream events.DynamoDBEvent) error {
	nrtx := newrelic.FromContext(ctx)
	logger := slog.Default()
	logger.Info("[RMU] START")

	dtoSeq := h.syncUsecaseConverter.ToSyncUsecaseInDtoSeq(ctx, stream.Records)
	err := h.syncUsecase.SyncBlogSnapshotWithEvents(ctx, dtoSeq)
	if err != nil {
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger.Error("failed to sync", slog.Any("error", err))
		return err
	}
	logger.Info("[RMU] END")
	return nil
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
