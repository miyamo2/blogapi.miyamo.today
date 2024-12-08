package lambda

import (
	"context"
	"github.com/cockroachdb/errors"
	"github.com/miyamo2/altnrslog"
	"github.com/miyamo2/blogapi.miyamo.today/core/log"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/if-adapters/model"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// SyncHandler is a handler for the Sync usecase.
type SyncHandler struct {
	syncUsecaseConverter ToSyncUsecaseInDtoConverter
	syncUsecase          SyncUsecase
}

// Invoke invokes the Sync usecase.
func (h *SyncHandler) Invoke(ctx context.Context, stream model.EventStream) error {
	nrtx := newrelic.FromContext(ctx)
	logger := log.New(log.WithAltNRSlogTransactionalHandler(nrtx.Application(), nrtx))
	ctx, err := altnrslog.StoreToContext(ctx, logger)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
	}

	dtoSeq := h.syncUsecaseConverter.ToSyncUsecaseInDtoSeq(ctx, stream.Records)
	err = h.syncUsecase.SyncBlogSnapshotWithEvents(ctx, dtoSeq)
	if err != nil {
		return err
	}
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