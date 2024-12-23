package query

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/miyamo2/altnrslog"
	"github.com/miyamo2/blogapi.miyamo.today/core/log"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"

	"github.com/cockroachdb/errors"
	"github.com/miyamo2/blogapi.miyamo.today/article-service/internal/infra/rdb/query/internal/entity"
	"github.com/miyamo2/blogapi.miyamo.today/core/db"
	gwrapper "github.com/miyamo2/blogapi.miyamo.today/core/db/gorm"
	"gorm.io/gorm"
)

var ErrNotFound = errors.New("not found")

type ArticleService struct{}

func (a *ArticleService) GetById(ctx context.Context, id string, out *db.SingleStatementResult[*Article]) db.Statement {
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
			var rows []entity.ArticleTag
			articleQuery := tx.Select(`*`).Table("articles").Where(`"id" = ?`, id)
			q := tx.Select(`"articles".*, "tags"."id" AS "tag_id", "tags"."name" AS "tag_name"`).
				Table(`(?) AS "articles"`, articleQuery).
				Joins(`LEFT OUTER JOIN "tags" ON "articles"."id" = "tags"."article_id"`, id)
			gwrapper.TraceableScan(nrtx, q, &rows)
			if len(rows) == 0 {
				err := errors.WithDetail(ErrNotFound, fmt.Sprintf("id: %v", id))
				nrtx.NoticeError(nrpkgerrors.Wrap(err))
				logger.WarnContext(ctx, "END",

					slog.Group("return",
						slog.Any("error", err)))
				return err
			}
			var article Article
			for i, r := range rows {
				if i == 0 {
					article = NewArticle(
						r.ID,
						r.Title,
						r.Body,
						r.Thumbnail,
						r.CreatedAt,
						r.UpdatedAt,
						WithTagsSize(len(rows)),
					)
				}
				if r.TagID == nil || r.TagName == nil {
					continue
				}
				article.AddTag(NewTag(*r.TagID, *r.TagName))
			}
			out.Set(&article)
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

func (a *ArticleService) GetAll(ctx context.Context, out *db.MultipleStatementResult[*Article], paginationOption ...db.PaginationOption) db.Statement {
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
			var rows []entity.ArticleTag
			tx = tx.WithContext(ctx)
			q := tx.Select(`"articles".*, "tags"."id" AS "tag_id", "tags"."name" AS "tag_name"`).
				Joins(`LEFT OUTER JOIN "tags" ON "articles"."id" = "tags"."article_id"`)

			func() {
				nextPaging := pagination.IsNextPaging()
				prevPaging := pagination.IsPreviousPaging()
				articleQuery := tx.Select(`*`).Table("articles")
				q.Table(`(?) AS "articles"`, articleQuery)
				if !nextPaging && !prevPaging {
					// default
					q.Order(`"articles"."id", "tags"."id" NULLS FIRST`)
					return
				}
				l := pagination.Limit()
				if l <= 0 {
					// default
					q.Order(`"articles"."id", "tags"."id" NULLS FIRST`)
					return
				}
				// must fetch one more row to check if there is more page.
				articleQuery.Limit(l + 1)
				c := pagination.Cursor()
				if nextPaging {
					if c != "" {
						articleQuery.
							Where(
								`EXISTS(?)`,
								tx.Select(`id`).Table("articles").Where(`"id" = ?`, c),
							).
							Where(`"id" > ?`, c)
					}
					articleQuery.Order(`"id"`)
					q.Order(`"articles"."id", "tags"."id" NULLS FIRST`)
					return
				}
				if prevPaging {
					if c != "" {
						articleQuery.
							Where(
								`EXISTS(?)`,
								tx.Select(`id`).Table("articles").Where(`"id" = ?`, c),
							).
							Where(`"id" < ?`, c)
					}
					articleQuery.Order(`"id" DESC`)
					q.Order(`"articles"."id" DESC, "tags"."id" NULLS FIRST`)
					return
				}
			}()
			gwrapper.TraceableScan(nrtx, q, &rows)
			articleMap := make(map[string]*Article)
			result := make([]*Article, 0)
			for _, r := range rows {
				v, ok := articleMap[r.ID]
				if !ok {
					m := NewArticle(
						r.ID,
						r.Title,
						r.Body,
						r.Thumbnail,
						r.CreatedAt,
						r.UpdatedAt,
					)
					v = &m
					articleMap[r.ID] = v
					result = append(result, v)
				}
				if r.TagID == nil || r.TagName == nil {
					continue
				}
				v.AddTag(NewTag(*r.TagID, *r.TagName))
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
