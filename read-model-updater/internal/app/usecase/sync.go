package usecase

import (
	"context"
	"github.com/cockroachdb/errors"
	"github.com/miyamo2/altnrslog"
	"github.com/miyamo2/blogapi.miyamo.today/core/db"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/app/usecase/command"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/app/usecase/externalapi"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/app/usecase/query"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/domain/model"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/infra/rdb"
	"github.com/newrelic/go-agent/v3/newrelic"
	"iter"
	"log/slog"
)

// Sync is an usecese of sync
type Sync struct {
	transactionManager        db.TransactionManager
	bloggingEventQueryService query.BloggingEventService
	articleCommandService     command.ArticleService
	tagCommandService         command.TagService
	blogAPIPublisher          externalapi.BlogPublisher
}

// SyncBlogSnapshotWithEvents synchronized blog snapshot with event
func (u *Sync) SyncBlogSnapshotWithEvents(ctx context.Context, in iter.Seq2[int, SyncUsecaseInDto]) error {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("Sync#SyncBlogSnapshotWithEvents").End()
	for _, dto := range in {
		if err := u.executePerEvent(ctx, dto); err != nil {
			return err
		}
	}
	if err := u.blogAPIPublisher.Publish(ctx); err != nil {
		return err
	}
	return nil
}

func (u *Sync) executePerEvent(ctx context.Context, dto SyncUsecaseInDto) error {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("Sync#executePerEvent").End()

	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		logger = slog.Default()
	}
	logger.InfoContext(ctx, "START", slog.Any("dto", dto))
	defer logger.InfoContext(ctx, "END")

	bloggingEvents := db.NewMultipleStatementResult[model.BloggingEvent]()
	if err := u.bloggingEventQueryService.AllEventsWithArticleID(ctx, dto.ArticleId(), bloggingEvents).Execute(ctx); err != nil {
		return err
	}

	articleCommand := model.ArticleCommandFromBloggingEvents(bloggingEvents.StrictGet())
	if articleCommand == nil {
		logger.WarnContext(ctx, "nil article command")
		return nil
	}

	articleTx, err := u.transactionManager.GetAndStart(ctx, db.GetAndStartWithDBSource(rdb.ArticleDBName))
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	articleErrSub := articleTx.SubscribeError()

	tagTx, err := u.transactionManager.GetAndStart(ctx, db.GetAndStartWithDBSource(rdb.TagDBName))
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	tagErrSub := tagTx.SubscribeError()

	done := make(chan struct{})
	errCh := make(chan error)
	go func() {
		err = articleTx.ExecuteStatement(ctx, u.articleCommandService.ExecuteArticleCommand(ctx, *articleCommand))
		if err != nil {
			errCh <- err
			return
		}

		err = tagTx.ExecuteStatement(ctx, u.tagCommandService.ExecuteTagCommand(ctx, *articleCommand))
		if err != nil {
			errCh <- err
			return
		}

		if err := articleTx.Commit(ctx); err != nil {
			errCh <- err
			return
		}
		if err := tagTx.Commit(ctx); err != nil {
			errCh <- err
			return
		}
		done <- struct{}{}
		return
	}()
	select {
	case <-ctx.Done():
		articleTx.Rollback(ctx)
		tagTx.Rollback(ctx)
		return ctx.Err()
	case err := <-articleErrSub:
		articleTx.Rollback(ctx)
		tagTx.Rollback(ctx)
		return err
	case err := <-tagErrSub:
		articleTx.Rollback(ctx)
		tagTx.Rollback(ctx)
		return err
	case err := <-errCh:
		articleTx.Rollback(ctx)
		tagTx.Rollback(ctx)
		return err
	case <-done:
	}
	return nil
}

// NewSync returns new Sync
func NewSync(
	transactionManager db.TransactionManager,
	bloggingEventQueryService query.BloggingEventService,
	articleCommandService command.ArticleService,
	tagCommandService command.TagService,
	blogAPIPublisher externalapi.BlogPublisher,
) *Sync {
	return &Sync{
		transactionManager:        transactionManager,
		articleCommandService:     articleCommandService,
		bloggingEventQueryService: bloggingEventQueryService,
		tagCommandService:         tagCommandService,
		blogAPIPublisher:          blogAPIPublisher,
	}
}
