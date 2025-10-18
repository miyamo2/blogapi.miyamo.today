package usecase

import (
	"context"
	"iter"
	"log/slog"
	"slices"
	"time"

	"blogapi.miyamo.today/read-model-updater/internal/app/usecase/command"
	"blogapi.miyamo.today/read-model-updater/internal/app/usecase/externalapi"
	"blogapi.miyamo.today/read-model-updater/internal/app/usecase/query"
	"blogapi.miyamo.today/read-model-updater/internal/domain/model"
	"blogapi.miyamo.today/read-model-updater/internal/infra/rdb/sqlc/article"
	"blogapi.miyamo.today/read-model-updater/internal/infra/rdb/sqlc/tag"
	"github.com/cockroachdb/errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/newrelic/go-agent/v3/newrelic"
	"golang.org/x/sync/errgroup"
)

// Sync is an usecese of sync
type Sync struct {
	bloggingEventQueryService query.BloggingEventService
	articleTx                 command.ArticleTx
	tagTx                     command.TagTx
	blogAPIPublisher          externalapi.BlogPublisher
	articleDBPool             *pgxpool.Pool
	tagDBPool                 *pgxpool.Pool
}

type ArticleDBPool *pgxpool.Pool

type TagDBPool *pgxpool.Pool

// SyncBlogSnapshotWithEvents synchronized blog snapshot with event
func (u *Sync) SyncBlogSnapshotWithEvents(ctx context.Context, in iter.Seq2[int, SyncUsecaseInDto]) error {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("Sync#SyncBlogSnapshotWithEvents").End()
	for _, dto := range in {
		if err := u.executePerEvent(ctx, dto); err != nil {
			return errors.WithStack(err)
		}
	}
	if err := u.blogAPIPublisher.Publish(ctx); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (u *Sync) executePerEvent(ctx context.Context, dto SyncUsecaseInDto) error {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("Sync#executePerEvent").End()

	logger := slog.Default()
	logger.Info("[RMU] START", slog.Any("dto", dto))
	defer logger.Info("[RMU] END")

	bloggingEvents, err := u.bloggingEventQueryService.ListEventsByArticleID(ctx, dto.ArticleID)
	if err != nil {
		return errors.WithStack(err)
	}

	articleCommand := model.ArticleCommandFromBloggingEvents(bloggingEvents)
	if articleCommand == nil {
		logger.Warn("nil article command")
		return nil
	}
	articleTx, err := u.articleDBPool.BeginTx(
		ctx, pgx.TxOptions{
			IsoLevel:   pgx.ReadCommitted,
			AccessMode: pgx.ReadWrite,
		},
	)
	if err != nil {
		return errors.WithStack(err)
	}
	defer func() {
		_ = articleTx.Rollback(ctx)
	}()

	tagTx, err := u.tagDBPool.BeginTx(
		ctx, pgx.TxOptions{
			IsoLevel:   pgx.ReadCommitted,
			AccessMode: pgx.ReadWrite,
		},
	)
	if err != nil {
		return errors.WithStack(err)
	}
	defer func() {
		_ = tagTx.Rollback(ctx)
	}()

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	errGroup, egCtx := errgroup.WithContext(ctx)
	errGroup.Go(
		func() error {
			q := u.articleTx.Begin(articleTx)
			err := q.CreateTempTagsTable(egCtx)
			if err != nil {
				return errors.WithStack(err)
			}
			err = q.PutArticle(
				egCtx, article.PutArticleParams{
					ID:        articleCommand.ID(),
					Title:     articleCommand.Title(),
					Body:      articleCommand.Body(),
					Thumbnail: articleCommand.Thumbnail(),
					CreatedAt: dto.EventAt,
					UpdatedAt: dto.EventAt,
				},
			)
			if err != nil {
				return errors.WithStack(err)
			}
			_, err = q.PreAttachTags(
				egCtx,
				slices.Collect(
					func(yield func(article.PreAttachTagsParams) bool) {
						for _, v := range articleCommand.Tags() {
							if yield(
								article.PreAttachTagsParams{
									ID:        v.ID(),
									ArticleID: articleCommand.ID(),
									Name:      v.Name(),
									CreatedAt: dto.EventAt,
									UpdatedAt: dto.EventAt,
								},
							) {
								return
							}
						}
					},
				),
			)
			if err != nil {
				return errors.WithStack(err)
			}
			err = q.AttachTags(egCtx, articleCommand.ID())
			if err != nil {
				return errors.WithStack(err)
			}
			return articleTx.Commit(egCtx)
		},
	)
	errGroup.Go(
		func() error {
			q := u.tagTx.Begin(tagTx)
			err := q.CreateTempArticlesTable(egCtx)
			if err != nil {
				return errors.WithStack(err)
			}
			err = q.CreateTempTagsTable(egCtx)
			if err != nil {
				return errors.WithStack(err)
			}
			_, err = q.PrePutTags(
				egCtx, slices.Collect(
					func(yield func(tag.PrePutTagsParams) bool) {
						for _, v := range articleCommand.Tags() {
							if yield(
								tag.PrePutTagsParams{
									ID:        v.ID(),
									Name:      v.Name(),
									CreatedAt: dto.EventAt,
									UpdatedAt: dto.EventAt,
								},
							) {
								return
							}
						}
					},
				),
			)
			if err != nil {
				return errors.WithStack(err)
			}
			err = q.PutTags(
				egCtx, slices.Collect(
					func(yield func(string) bool) {
						for _, v := range articleCommand.Tags() {
							if yield(v.ID()) {
								return
							}
						}
					},
				),
			)
			if err != nil {
				return errors.WithStack(err)
			}
			_, err = q.PrePutArticle(
				egCtx, slices.Collect(
					func(yield func(tag.PrePutArticleParams) bool) {
						for _, v := range articleCommand.Tags() {
							if yield(
								tag.PrePutArticleParams{
									ID:        articleCommand.ID(),
									TagID:     v.ID(),
									Title:     articleCommand.Title(),
									Thumbnail: articleCommand.Thumbnail(),
									CreatedAt: dto.EventAt,
									UpdatedAt: dto.EventAt,
								},
							) {
								return
							}
						}
					},
				),
			)
			if err != nil {
				return errors.WithStack(err)
			}
			err = q.PutArticle(egCtx, articleCommand.ID())
			if err != nil {
				return errors.WithStack(err)
			}
			return tagTx.Commit(egCtx)
		},
	)
	if err := errGroup.Wait(); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// NewSync returns new Sync
func NewSync(
	bloggingEventQueryService query.BloggingEventService,
	articleTx command.ArticleTx,
	tagTx command.TagTx,
	articleDBPool ArticleDBPool,
	tagDBPool TagDBPool,
	blogAPIPublisher externalapi.BlogPublisher,
) *Sync {
	return &Sync{
		articleTx:                 articleTx,
		bloggingEventQueryService: bloggingEventQueryService,
		tagTx:                     tagTx,
		blogAPIPublisher:          blogAPIPublisher,
		articleDBPool:             articleDBPool,
		tagDBPool:                 tagDBPool,
	}
}
