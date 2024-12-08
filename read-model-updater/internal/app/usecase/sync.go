package usecase

import (
	"context"
	"github.com/miyamo2/blogapi.miyamo.today/core/db"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/app/usecase/command"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/app/usecase/externalapi"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/app/usecase/query"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/domain/model"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/infra/rdb"
	"github.com/newrelic/go-agent/v3/newrelic"
	"iter"
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
func (t *Sync) SyncBlogSnapshotWithEvents(ctx context.Context, in iter.Seq2[int, SyncUsecaseInDto]) error {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("Sync#SyncBlogSnapshotWithEvents").End()
	for _, dto := range in {
		if err := t.executePerEvent(ctx, dto); err != nil {
			return err
		}
	}
	return nil
}

func (t *Sync) executePerEvent(ctx context.Context, dto SyncUsecaseInDto) error {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("Sync#executePerEvent").End()

	bloggingEvents := db.NewMultipleStatementResult[model.BloggingEvent]()
	if err := t.bloggingEventQueryService.AllEventsWithArticleID(ctx, dto.ArticleId(), bloggingEvents).Execute(ctx); err != nil {
		return err
	}

	articleCommand := model.ArticleCommandFromBloggingEvents(bloggingEvents.StrictGet())
	if articleCommand == nil {
		return nil
	}

	articleTx, err := t.transactionManager.GetAndStart(ctx, db.GetAndStartWithDBSource(rdb.ArticleDBName))
	if err != nil {
		return err
	}
	articleErrSub := articleTx.SubscribeError()

	tagTx, err := t.transactionManager.GetAndStart(ctx, db.GetAndStartWithDBSource(rdb.TagDBName))
	if err != nil {
		return err
	}
	tagErrSub := tagTx.SubscribeError()

	errCh := make(chan error)
	go func() {
		err = articleTx.ExecuteStatement(ctx, t.articleCommandService.ExecuteArticleCommand(ctx, *articleCommand))
		if err != nil {
			errCh <- err
			return
		}

		err = tagTx.ExecuteStatement(ctx, t.tagCommandService.ExecuteTagCommand(ctx, *articleCommand))
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
	default:
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
