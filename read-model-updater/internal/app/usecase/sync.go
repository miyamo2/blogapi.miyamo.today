package usecase

import (
	"blogapi.miyamo.today/core/db"
	gw "blogapi.miyamo.today/core/db/gorm"
	"blogapi.miyamo.today/read-model-updater/internal/app/usecase/command"
	"blogapi.miyamo.today/read-model-updater/internal/app/usecase/externalapi"
	"blogapi.miyamo.today/read-model-updater/internal/app/usecase/query"
	"blogapi.miyamo.today/read-model-updater/internal/domain/model"
	"blogapi.miyamo.today/read-model-updater/internal/infra/dynamo"
	"blogapi.miyamo.today/read-model-updater/internal/infra/rdb"
	"context"
	"github.com/newrelic/go-agent/v3/newrelic"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"iter"
	"log/slog"
)

// Sync is an usecese of sync
type Sync struct {
	rdbGorm                   *rdb.DB
	dynamodbGorm              *dynamo.DB
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

	logger := slog.Default()
	logger.Info("[RMU] START", slog.Any("dto", dto))
	defer logger.Info("[RMU] END")

	bloggingEvents := db.NewMultipleStatementResult[model.BloggingEvent]()
	if err := u.bloggingEventQueryService.AllEventsWithArticleID(ctx, dto.ArticleId(), bloggingEvents).
		Execute(ctx, gw.WithTransaction(u.dynamodbGorm.Session(&gorm.Session{
			PrepareStmt:            false,
			SkipDefaultTransaction: true,
		}))); err != nil {
		return err
	}

	articleCommand := model.ArticleCommandFromBloggingEvents(bloggingEvents.StrictGet())
	if articleCommand == nil {
		logger.Warn("nil article command")
		return nil
	}

	articleTx := u.rdbGorm.Session(&gorm.Session{
		PrepareStmt:            false,
		SkipDefaultTransaction: true,
	}).Clauses(dbresolver.Use(rdb.ArticleDBName)).Begin()

	tagTx := u.rdbGorm.Session(&gorm.Session{
		PrepareStmt:            false,
		SkipDefaultTransaction: true,
	}).Clauses(dbresolver.Use(rdb.TagDBName)).Begin()

	done := make(chan struct{})
	errCh := make(chan error)
	go func() {
		err := u.articleCommandService.ExecuteArticleCommand(ctx, *articleCommand).Execute(ctx, gw.WithTransaction(articleTx))
		if err != nil {
			errCh <- err
			return
		}

		err = u.tagCommandService.ExecuteTagCommand(ctx, *articleCommand).Execute(ctx, gw.WithTransaction(tagTx))
		if err != nil {
			errCh <- err
			return
		}

		if err := articleTx.Commit().Error; err != nil {
			errCh <- err
			return
		}
		if err := tagTx.Commit().Error; err != nil {
			errCh <- err
			return
		}
		done <- struct{}{}
	}()
	select {
	case <-ctx.Done():
		articleTx.Rollback()
		tagTx.Rollback()
		return ctx.Err()
	case err := <-errCh:
		articleTx.Rollback()
		tagTx.Rollback()
		return err
	case <-done:
	}
	return nil
}

// NewSync returns new Sync
func NewSync(
	rdbGorm *rdb.DB,
	dynamodbGorm *dynamo.DB,
	bloggingEventQueryService query.BloggingEventService,
	articleCommandService command.ArticleService,
	tagCommandService command.TagService,
	blogAPIPublisher externalapi.BlogPublisher,
) *Sync {
	return &Sync{
		rdbGorm:                   rdbGorm,
		dynamodbGorm:              dynamodbGorm,
		articleCommandService:     articleCommandService,
		bloggingEventQueryService: bloggingEventQueryService,
		tagCommandService:         tagCommandService,
		blogAPIPublisher:          blogAPIPublisher,
	}
}
