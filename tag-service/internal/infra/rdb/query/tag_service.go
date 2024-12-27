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

func (t *TagService) GetById(ctx context.Context, id string, out *db.SingleStatementResult[model.Tag]) db.Statement {
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
			q := tx.
				Select(`"t".*,  COALESCE(jsonb_agg(json_build_object('id', a.id, 'title', a.title, 'thumbnail', a.thumbnail, 'created_at', a.created_at, 'updated_at', a.updated_at)) FILTER (WHERE a.id IS NOT NULL), '[]'::json) AS "articles"`).
				Table(`(?) AS "t"`, tx.Select(`*`).Table(`"tags"`).Where(`"id" = ?`, id)).
				Joins(`LEFT OUTER JOIN "articles" AS "a" ON "t"."id" = "a"."tag_id"`).
				Group(`"t"."id"`)

			var rows []entity.Tag
			gwrapper.TraceableScan(nrtx, q, &rows)
			if len(rows) == 0 {
				err := errors.WithDetail(ErrNotFound, fmt.Sprintf("id: %v", id))
				nrtx.NoticeError(nrpkgerrors.Wrap(err))
				return err
			}
			var tag model.Tag

			row := rows[0]
			tag = model.NewTag(row.ID, row.Name, func() []model.Article {
				articles := make([]model.Article, 0)
				for _, a := range row.Articles {
					articles = append(articles, model.NewArticle(
						a.ID,
						a.Title,
						a.Thumbnail,
						a.CreatedAt,
						a.UpdatedAt))
				}
				return articles
			}()...)
			out.Set(tag)
			return nil
		}, out)
	defer logger.InfoContext(ctx, "END",

		slog.Group("return",
			slog.String("stmt", fmt.Sprintf("%v", stmt))))
	return stmt
}

func (t *TagService) GetAll(ctx context.Context, out *db.MultipleStatementResult[model.Tag], paginationOption ...db.PaginationOption) db.Statement {
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

			var rows []entity.Tag
			tx = tx.WithContext(ctx)
			q := tx.Select(`"t".*,  COALESCE(jsonb_agg(json_build_object('id', a.id, 'title', a.title, 'thumbnail', a.thumbnail, 'created_at', a.created_at, 'updated_at', a.updated_at)) FILTER (WHERE a.id IS NOT NULL), '[]'::json) AS "articles"`).
				Joins(`LEFT OUTER JOIN "articles" AS "a" ON "t"."id" = "a"."tag_id"`).
				Group(`"t"."id"`)
			buildQuery(pagination, tx, q)

			gwrapper.TraceableScan(nrtx, q, &rows)

			result := make([]model.Tag, 0, len(rows))
			for _, r := range rows {
				tag := model.NewTag(r.ID, r.Name, func() []model.Article {
					articles := make([]model.Article, 0)
					for _, a := range r.Articles {
						articles = append(articles, model.NewArticle(
							a.ID,
							a.Title,
							a.Thumbnail,
							a.CreatedAt,
							a.UpdatedAt))
					}
					return articles
				}()...)
				result = append(result, tag)
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
	tagQuery := tx.Select(`*`).Table("tags")
	q.Table(`(?) AS "t"`, tagQuery)

	limit := pagination.Limit()
	// must fetch one more row to check if there is more page.
	if limit != 0 {
		tagQuery.Limit(limit + 1)
	}

	cursor := pagination.Cursor()
	if pagination.IsNextPaging() {
		tagQuery.Order(`"id"`)
		if cursor == "" {
			return
		}
		tagQuery.
			Where(
				`EXISTS(?)`,
				tx.Select(`id`).Table("tags").Where(`"id" = ?`, cursor),
			).
			Where(`"id" > ?`, cursor)
		return
	}
	if pagination.IsPreviousPaging() {
		tagQuery.Order(`"id" DESC`)
		if cursor == "" {
			return
		}
		tagQuery.
			Where(
				`EXISTS(?)`,
				tx.Select(`id`).Table("tags").Where(`"id" = ?`, cursor),
			).
			Where(`"id" < ?`, cursor)
		return
	}
	tagQuery.Order(`"id"`)
}

func NewTagService() *TagService {
	return &TagService{}
}
