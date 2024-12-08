package dynamo

import (
	"context"
	"fmt"
	"github.com/miyamo2/blogapi.miyamo.today/core/db"
	gw "github.com/miyamo2/blogapi.miyamo.today/core/db/gorm"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/domain/model"
	"github.com/miyamo2/sqldav"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"os"
	"slices"
)

const DBName = "dynamodb"

var _ schema.Tabler = (*bloggingEvent)(nil)

type bloggingEvent struct {
	EventID     string `gorm:"primaryKey"`
	ArticleID   string `gorm:"primaryKey"`
	Title       *string
	Content     *string
	Thumbnail   *string
	Tags        sqldav.TypedList[string]
	AttacheTags sqldav.TypedList[string]
	DetachTags  sqldav.TypedList[string]
	Invisible   *bool
}

func (b bloggingEvent) TableName() string {
	return fmt.Sprintf("blogging_events-_%s", os.Getenv("ENV"))
}

type BloggingEventQueryService struct{}

func (s *BloggingEventQueryService) AllEventsWithArticleID(ctx context.Context, articleId string, out *db.MultipleStatementResult[model.BloggingEvent]) db.Statement {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("BloggingEventQueryService#AllEventsWithArticleID").End()
	return gw.NewStatement(func(ctx context.Context, tx *gorm.DB, out db.StatementResult) (err error) {
		nrtx := newrelic.FromContext(ctx)
		defer nrtx.StartSegment("BloggingEventQueryService#AllEventsWithArticleID").End()
		tx = tx.Clauses(dbresolver.Use(DBName)).WithContext(ctx)

		rows := make([]bloggingEvent, 0)
		err = tx.Where("article_id = ?", articleId).Scan(&rows).Error
		if err != nil {
			return err
		}

		slices.SortFunc(rows, func(i, j bloggingEvent) int {
			iid := ulid.MustParseStrict(i.EventID)
			jid := ulid.MustParseStrict(j.EventID)
			return iid.Compare(jid)
		})
		// recover ulid.MustParseStrict
		defer func() {
			if rec := recover(); rec != nil {
				err = fmt.Errorf("recovered: %w", rec)
			}
		}()

		result := make([]model.BloggingEvent, 0)
		for _, r := range rows {
			result = append(result, model.NewBloggingEvent(r.EventID, r.ArticleID, r.Title, r.Content, r.Thumbnail, r.AttacheTags, r.DetachTags, r.Invisible))
		}
		out.Set(result)
		return nil
	}, out)
}

func NewBloggingEventQueryService() *BloggingEventQueryService {
	return &BloggingEventQueryService{}
}