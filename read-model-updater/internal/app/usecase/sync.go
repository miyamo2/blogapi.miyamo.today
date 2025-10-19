package usecase

import (
	"context"
	"log/slog"
	"slices"
	"time"

	"blogapi.miyamo.today/read-model-updater/internal/app/usecase/command"
	"blogapi.miyamo.today/read-model-updater/internal/app/usecase/externalapi"
	"blogapi.miyamo.today/read-model-updater/internal/app/usecase/query"
	"blogapi.miyamo.today/read-model-updater/internal/domain/model"
	"blogapi.miyamo.today/read-model-updater/internal/infra/rdb/sqlc/article"
	"blogapi.miyamo.today/read-model-updater/internal/infra/rdb/sqlc/tag"
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	"github.com/cockroachdb/errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/oklog/ulid/v2"
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
func (u *Sync) SyncBlogSnapshotWithEvents(ctx context.Context, dto *SyncUsecaseInDto) error {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("Sync#SyncBlogSnapshotWithEvents").End()
	if err := u.execute(ctx, dto); err != nil {
		return errors.WithStack(err)
	}
	if err := u.blogAPIPublisher.Publish(ctx); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (u *Sync) execute(ctx context.Context, dto *SyncUsecaseInDto) error {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("Sync#execute").End()

	slog.Default().InfoContext(
		ctx,
		"Syncing blog snapshot",
		slog.String("article_id", dto.ArticleID),
		slog.Time("event_at", dto.EventAt.StdTime()),
	)
	defer slog.Default().InfoContext(
		ctx,
		"Synced blog snapshot",
		slog.String("article_id", dto.ArticleID),
		slog.Time("event_at", dto.EventAt.StdTime()),
	)

	bloggingEvents, err := u.bloggingEventQueryService.ListEventsByArticleID(ctx, dto.ArticleID)
	if err != nil {
		return errors.WithStack(err)
	}

	articleCommand := model.ArticleCommandFromBloggingEvents(bloggingEvents)

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

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	articleULID, err := ulid.Parse(articleCommand.ID())
	if err != nil {
		return errors.WithStack(err)
	}

	errGroup, egCtx := errgroup.WithContext(ctx)
	errGroup.Go(
		func() error {
			nrtx := nrtx.NewGoroutine()
			ctx := newrelic.NewContext(egCtx, nrtx)
			defer nrtx.StartSegment("ArticleTx").End()

			q := u.articleTx.Begin(articleTx)
			err := q.CreateTempTagsTable(ctx)
			if err != nil {
				return errors.WithStack(err)
			}
			err = q.PutArticle(
				ctx, article.PutArticleParams{
					ID:        articleCommand.ID(),
					Title:     articleCommand.Title(),
					Body:      articleCommand.Body(),
					Thumbnail: articleCommand.Thumbnail(),
					CreatedAt: synchro.In[tz.UTC](articleULID.Timestamp()),
					UpdatedAt: dto.EventAt,
				},
			)
			if err != nil {
				return errors.WithStack(err)
			}
			_, err = q.PreAttachTags(
				ctx,
				slices.Collect(
					func(yield func(article.PreAttachTagsParams) bool) {
						for _, v := range articleCommand.Tags() {
							if !yield(
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
			err = q.AttachTags(ctx, articleCommand.ID())
			if err != nil {
				return errors.WithStack(err)
			}
			return articleTx.Commit(ctx)
		},
	)
	errGroup.Go(
		func() error {
			nrtx := nrtx.NewGoroutine()
			ctx := newrelic.NewContext(egCtx, nrtx)
			defer nrtx.StartSegment("TagTx").End()

			q := u.tagTx.Begin(tagTx)
			err := q.CreateTempArticlesTable(ctx)
			if err != nil {
				return errors.WithStack(err)
			}
			err = q.CreateTempTagsTable(ctx)
			if err != nil {
				return errors.WithStack(err)
			}
			_, err = q.PrePutTags(
				ctx, slices.Collect(
					func(yield func(tag.PrePutTagsParams) bool) {
						for _, v := range articleCommand.Tags() {
							if !yield(
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
				ctx, slices.Collect(
					func(yield func(string) bool) {
						for _, v := range articleCommand.Tags() {
							if !yield(v.ID()) {
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
				ctx, slices.Collect(
					func(yield func(tag.PrePutArticleParams) bool) {
						for _, v := range articleCommand.Tags() {
							if !yield(
								tag.PrePutArticleParams{
									ID:        articleCommand.ID(),
									TagID:     v.ID(),
									Title:     articleCommand.Title(),
									Thumbnail: articleCommand.Thumbnail(),
									CreatedAt: synchro.In[tz.UTC](articleULID.Timestamp()),
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
			err = q.PutArticle(ctx, articleCommand.ID())
			if err != nil {
				return errors.WithStack(err)
			}
			return tagTx.Commit(ctx)
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
