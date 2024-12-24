package query

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/miyamo2/altnrslog"
	"github.com/miyamo2/blogapi.miyamo.today/core/log"

	"github.com/miyamo2/blogapi.miyamo.today/tag-service/internal/infra/rdb/query/model"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"

	"github.com/cockroachdb/errors"
	"github.com/miyamo2/blogapi.miyamo.today/core/db"
	gwrapper "github.com/miyamo2/blogapi.miyamo.today/core/db/gorm"
	"github.com/miyamo2/blogapi.miyamo.today/tag-service/internal/infra/rdb/query/internal/entity"
	"gorm.io/gorm"
)

var ErrNotFound = errors.New("not found")

// TagService is an implementation of query.TagService
type TagService struct{}

func (t *TagService) GetById(ctx context.Context, id string, out *db.SingleStatementResult[*model.Tag]) db.Statement {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("GetById").End()
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters",
			slog.String("id", id),
			slog.String("out", fmt.Sprintf("%v", out)),
		))
	stmt := gwrapper.NewStatement(
		func(ctx context.Context, tx *gorm.DB, out db.StatementResult) error {
			nrtx := newrelic.FromContext(ctx)
			defer nrtx.StartSegment("GetById Execute").End()
			logger, err := altnrslog.FromContext(ctx)
			if err != nil {
				err = errors.WithStack(err)
				nrtx.NoticeError(nrpkgerrors.Wrap(err))
				logger = log.DefaultLogger()
			}
			logger.InfoContext(ctx, "BEGIN",
				slog.Group("bind",
					slog.String("id", id)))
			defer func() { logger.InfoContext(ctx, "END") }()
			tx = tx.WithContext(ctx)
			var rows []entity.TagArticle
			tagQuery := tx.
				Select(`"tags".*`).
				Table(`"tags"`).
				Where(`"tags"."id" = ?`, id)
			q := tx.
				Select(`"tags".*, "articles"."id" AS "article_id", "articles"."title" AS "article_title", "articles"."thumbnail" AS "article_thumbnail", "articles"."created_at" AS "article_created_at", "articles"."updated_at" AS "article_updated_at"`).
				Table(`(?) AS "tags"`, tagQuery).
				Joins(`LEFT OUTER JOIN "articles" ON "tags"."id" = "articles"."tag_id"`)
			gwrapper.TraceableScan(nrtx, q, &rows)
			if len(rows) == 0 {
				err := errors.WithDetail(ErrNotFound, fmt.Sprintf("id: %v", id))
				nrtx.NoticeError(nrpkgerrors.Wrap(err))
				return err
			}
			var tag model.Tag
			for i, r := range rows {
				if i == 0 {
					tag = model.NewTag(
						r.ID,
						r.Name,
						model.WithTagsSize(len(rows)),
					)
				}
				if r.ArticleID == nil || r.ArticleTitle == nil {
					continue
				}
				tag.AddArticle(model.NewArticle(
					*r.ArticleID,
					*r.ArticleTitle,
					*r.ArticleThumbnail,
					*r.ArticleCreatedAt,
					*r.ArticleUpdatedAt))
			}
			out.Set(&tag)
			return nil
		}, out)
	defer logger.InfoContext(ctx, "END",

		slog.Group("return",
			slog.String("stmt", fmt.Sprintf("%v", stmt))))
	return stmt
}

func (t *TagService) GetAll(ctx context.Context, out *db.MultipleStatementResult[*model.Tag], paginationOption ...db.PaginationOption) db.Statement {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("GetAll").End()
	pagination := db.Pagination{}
	for _, opt := range paginationOption {
		opt(&pagination)
	}
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters",
			slog.String("out", fmt.Sprintf("%v", out)),
		))
	stmt := gwrapper.NewStatement(
		func(ctx context.Context, tx *gorm.DB, out db.StatementResult) error {
			nrtx := newrelic.FromContext(ctx)
			defer nrtx.StartSegment("GetAll Execute").End()
			logger, err := altnrslog.FromContext(ctx)
			if err != nil {
				err = errors.WithStack(err)
				nrtx.NoticeError(nrpkgerrors.Wrap(err))
				logger = log.DefaultLogger()
			}
			logger.InfoContext(ctx, "BEGIN",
				slog.Group("bind",
					slog.String("pagination", fmt.Sprintf("%v", pagination)),
				))
			defer func() { logger.InfoContext(ctx, "END") }()
			var rows []entity.TagArticle
			tx = tx.WithContext(ctx)
			q := tx.Select(`"tags".*, "articles"."id" AS "article_id", "articles"."title" AS "article_title", "articles"."thumbnail" AS "article_thumbnail", "articles"."created_at" AS "article_created_at", "articles"."updated_at" AS "article_updated_at"`).
				Joins(`LEFT OUTER JOIN "articles" ON "tags"."id" = "articles"."tag_id"`)
			buildQuery(pagination, tx, q)
			gwrapper.TraceableScan(nrtx, q, &rows)
			tagMap := make(map[string]*model.Tag)
			result := make([]*model.Tag, 0)
			for _, r := range rows {
				v, ok := tagMap[r.ID]
				if !ok {
					m := model.NewTag(r.ID, r.Name)
					v = &m
					tagMap[r.ID] = v
					result = append(result, v)
				}
				if r.ArticleID == nil || r.ArticleTitle == nil {
					continue
				}
				v.AddArticle(model.NewArticle(
					*r.ArticleID,
					*r.ArticleTitle,
					*r.ArticleThumbnail,
					*r.ArticleCreatedAt,
					*r.ArticleUpdatedAt))
			}
			out.Set(result)
			return nil
		}, out)
	defer logger.InfoContext(ctx, "END",

		slog.Group("return",
			slog.String("stmt", fmt.Sprintf("%v", stmt))))
	return stmt
}

// buildQuery builds a query for pagination.
func buildQuery(pagination db.Pagination, tx *gorm.DB, q *gorm.DB) {
	nextPaging := pagination.IsNextPaging()
	prevPaging := pagination.IsPreviousPaging()
	tagQuery := tx.Select(`*`).Table("tags")
	q.Table(`(?) AS "tags"`, tagQuery)
	if !nextPaging && !prevPaging {
		// default
		q.Order(`"tags"."id", "articles"."id" NULLS FIRST`)
		return
	}
	l := pagination.Limit()
	if l <= 0 {
		// default
		q.Order(`"tags"."id", "articles"."id" NULLS FIRST`)
		return
	}
	// must fetch one more row to check if there is more page.
	tagQuery.Limit(l + 1)
	c := pagination.Cursor()
	if nextPaging {
		if c != "" {
			tagQuery.
				Where(
					`EXISTS(?)`,
					tx.Select(`id`).Table("tags").Where(`"id" = ?`, c),
				).
				Where(`"id" > ?`, c)
		}
		tagQuery.Order(`"id"`)
		q.Order(`"tags"."id", "articles"."id" NULLS FIRST`)
		return
	}
	if prevPaging {
		if c != "" {
			tagQuery.
				Where(
					`EXISTS(?)`,
					tx.Select(`id`).Table("tags").Where(`"id" = ?`, c),
				).
				Where(`"id" < ?`, c)
		}
		tagQuery.Order(`"id" DESC`)
		q.Order(`"tags"."id" DESC, "articles"."id" NULLS FIRST`)
		return
	}
}

func NewTagService() *TagService {
	return &TagService{}
}
