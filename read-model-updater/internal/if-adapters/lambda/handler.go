package lambda

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/cockroachdb/errors"
	"github.com/miyamo2/altnrslog"
	"github.com/miyamo2/blogapi.miyamo.today/core/log"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// SyncHandler is a handler for the Sync usecase.
type SyncHandler struct {
	syncUsecaseConverter ToSyncUsecaseInDtoConverter
	syncUsecase          SyncUsecase
}

// Invoke invokes the Sync usecase.
func (h *SyncHandler) Invoke(ctx context.Context, stream events.DynamoDBEvent) error {
	nrtx := newrelic.FromContext(ctx)
	logger := log.New(log.WithAltNRSlogTransactionalHandler(nrtx.Application(), nrtx))
	ctx, err := altnrslog.StoreToContext(ctx, logger)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
	}
	logger.InfoContext(ctx, "START")

	dtoSeq := h.syncUsecaseConverter.ToSyncUsecaseInDtoSeq(ctx, stream.Records)
	err = h.syncUsecase.SyncBlogSnapshotWithEvents(ctx, dtoSeq)
	if err != nil {
		return err
	}
	logger.InfoContext(ctx, "END")
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
