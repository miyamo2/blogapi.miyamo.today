package query

import (
	"context"
	"fmt"
	"log/slog"

	"blogapi.miyamo.today/core/log"
	"github.com/miyamo2/altnrslog"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"

	"blogapi.miyamo.today/article-service/internal/infra/rdb/query/internal/entity"
	"blogapi.miyamo.today/core/db"
	gwrapper "blogapi.miyamo.today/core/db/gorm"
	"github.com/cockroachdb/errors"
	"gorm.io/gorm"
)

var ErrNotFound = errors.New("not found")

type ArticleService struct{}

func (a *ArticleService) GetById(ctx context.Context, id string, out *db.SingleStatementResult[Article]) db.Statement {
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
			slog.Any("out", out),
		))
	stmt := gwrapper.NewStatement(
		func(ctx context.Context, tx *gorm.DB, out db.StatementResult) error {
			nrtx := newrelic.FromContext(ctx)
			defer nrtx.StartSegment("GetById Statement").End()
			logger.InfoContext(ctx, "BEGIN",
				slog.Group("parameters",
					slog.Any("out", out),
				))
			tx = tx.WithContext(ctx)
			var rows []entity.Article

			q := tx.Select(`"a".*, COALESCE(jsonb_agg(json_build_object('id', t.id, 'name', t.name)) FILTER (WHERE t.id IS NOT NULL), '[]'::json) AS "tags"`).
				Table(`(?) AS "a"`, tx.Select(`*`).Table("articles").Where(`"id" = ?`, id)).
				Joins(`LEFT OUTER JOIN "tags" AS "t" ON "a"."id" = "t"."article_id"`, id).
				Group(`"a"."id"`)

			gwrapper.TraceableScan(nrtx, q, &rows)
			if len(rows) == 0 {
				err := errors.WithDetail(ErrNotFound, fmt.Sprintf("id: %v", id))
				nrtx.NoticeError(nrpkgerrors.Wrap(err))
				logger.WarnContext(ctx, "END",
					slog.Group("return",
						slog.Any("error", err)))
				return err
			}
			row := rows[0]
			article := NewArticle(
				row.ID,
				row.Title,
				row.Body,
				row.Thumbnail,
				row.CreatedAt,
				row.UpdatedAt,
				func() []Tag {
					v := make([]Tag, 0, len(row.Tags))
					for _, t := range row.Tags {
						v = append(v, NewTag(t.ID, t.Name))
					}
					return v
				}()...,
			)
			out.Set(article)
			logger.InfoContext(ctx, "END",

				slog.Group("return",
					slog.Any("error", nil)))
			return nil
		}, out)
	defer logger.InfoContext(ctx, "END",

		slog.Group("return",
			slog.Any("stmt", stmt)))
	return stmt
}

func (a *ArticleService) GetAll(ctx context.Context, out *db.MultipleStatementResult[Article], paginationOption ...db.PaginationOption) db.Statement {
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
			slog.Any("out", out),
		))
	stmt := gwrapper.NewStatement(
		func(ctx context.Context, tx *gorm.DB, out db.StatementResult) error {
			nrtx := newrelic.FromContext(ctx)
			defer nrtx.StartSegment("GetAll Statement").End()
			logger.InfoContext(ctx, "BEGIN",
				slog.Group("parameters",
					slog.Any("out", out),
				))
			var rows []entity.Article
			tx = tx.WithContext(ctx)

			articleQuery := tx.Select(`*`).Table("articles")
			q := tx.Select(`"a".*, COALESCE(jsonb_agg(json_build_object('id', t.id, 'name', t.name)) FILTER (WHERE t.id IS NOT NULL), '[]'::json) AS "tags"`).
				Joins(`LEFT OUTER JOIN "tags" AS "t" ON "a"."id" = "t"."article_id"`).
				Group(`"a"."id"`)

			func() {
				nextPaging := pagination.IsNextPaging()
				prevPaging := pagination.IsPreviousPaging()

				limit := pagination.Limit()
				// must fetch one more row to check if there is more page.
				if limit != 0 {
					articleQuery.Limit(limit + 1)
				}

				cursor := pagination.Cursor()
				if nextPaging {
					if cursor != "" {
						articleQuery.
							Where(
								`EXISTS(?)`,
								tx.Select(`id`).Table("articles").Where(`"id" = ?`, cursor),
							).
							Where(`"id" > ?`, cursor)
					}
					articleQuery.Order(`"id"`)
					return
				}
				if prevPaging {
					if cursor != "" {
						articleQuery.
							Where(
								`EXISTS(?)`,
								tx.Select(`id`).Table("articles").Where(`"id" = ?`, cursor),
							).
							Where(`"id" < ?`, cursor)
					}
					articleQuery.Order(`"id" DESC`)
					return
				}
				articleQuery.Order(`"id"`)
			}()
			q.Table(`(?) AS "a"`, articleQuery)
			gwrapper.TraceableScan(nrtx, q, &rows)
			result := make([]Article, 0, len(rows))
			for _, r := range rows {
				article := NewArticle(
					r.ID,
					r.Title,
					r.Body,
					r.Thumbnail,
					r.CreatedAt,
					r.UpdatedAt,
					func() []Tag {
						v := make([]Tag, 0, len(r.Tags))
						for _, t := range r.Tags {
							v = append(v, NewTag(t.ID, t.Name))
						}
						return v
					}()...,
				)
				result = append(result, article)
			}
			out.Set(result)
			logger.InfoContext(ctx, "END",
				slog.Group("return",
					slog.Any("error", nil)))
			return nil
		}, out)
	defer logger.InfoContext(ctx, "END",

		slog.Group("return",
			slog.Any("stmt", stmt)))
	return stmt
}

func NewArticleService() *ArticleService {
	return &ArticleService{}
}
