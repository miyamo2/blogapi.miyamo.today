package dynamo

import (
	"blogapi.miyamo.today/core/db"
	gw "blogapi.miyamo.today/core/db/gorm"
	"blogapi.miyamo.today/read-model-updater/internal/domain/model"
	"context"
	"fmt"
	"github.com/cockroachdb/errors"
	"github.com/miyamo2/dynmgrm"
	"github.com/miyamo2/sqldav"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log/slog"
	"os"
	"slices"
)

type DB struct {
	*gorm.DB
}

var _ schema.Tabler = (*bloggingEvent)(nil)

type bloggingEvent struct {
	EventID    string `gorm:"primaryKey"`
	ArticleID  string `gorm:"primaryKey"`
	Title      *string
	Content    *string
	Thumbnail  *string
	Tags       sqldav.Set[string]
	AttachTags sqldav.Set[string]
	DetachTags sqldav.Set[string]
	Invisible  *bool
}

func (b bloggingEvent) TableName() string {
	return os.Getenv("BLOGGING_EVENTS_TABLE_NAME")
}

type BloggingEventQueryService struct{}

func (s *BloggingEventQueryService) AllEventsWithArticleID(ctx context.Context, articleId string, out *db.MultipleStatementResult[model.BloggingEvent]) db.Statement {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("BloggingEventQueryService#AllEventsWithArticleID").End()
	return gw.NewStatement(func(ctx context.Context, tx *gorm.DB, out db.StatementResult) (err error) {
		nrtx := newrelic.FromContext(ctx)
		defer nrtx.StartSegment("BloggingEventQueryService#AllEventsWithArticleID").End()
		logger := slog.Default()
		logger.Info("[RMU] START")
		tx = tx.WithContext(ctx)

		rows := make([]bloggingEvent, 0)
		err = tx.Select("event_id", "article_id", "title", "content", "thumbnail", "tags", "attach_tags", "detach_tags", "invisible").
			Table(bloggingEvent{}.TableName()).Clauses(
			dynmgrm.SecondaryIndex("article_id_event_id-Index")).
			Where("article_id = ?", articleId).Scan(&rows).Error
		if err != nil {
			err = errors.WithStack(err)
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
				err = errors.WithStack(err)
			}
		}()

		result := make([]model.BloggingEvent, 0)
		for _, r := range rows {
			result = append(result, model.NewBloggingEvent(r.EventID, r.ArticleID, r.Title, r.Content, r.Thumbnail, r.Tags, r.AttachTags, r.DetachTags, r.Invisible))
		}
		out.Set(result)
		logger.Info("[RMU] END", slog.Int("result count", len(result)))
		return nil
	}, out)
}

func NewBloggingEventQueryService() *BloggingEventQueryService {
	return &BloggingEventQueryService{}
}
