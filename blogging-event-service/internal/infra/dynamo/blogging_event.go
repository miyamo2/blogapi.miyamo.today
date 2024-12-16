package dynamo

import (
	"context"
	"fmt"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/domain/model"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/pkg"
	"github.com/miyamo2/blogapi.miyamo.today/core/db"
	gw "github.com/miyamo2/blogapi.miyamo.today/core/db/gorm"
	"github.com/miyamo2/sqldav"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log/slog"
	"os"
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

type BloggingEventCommandService struct {
	ulidGen pkg.ULIDGenerator
}

func (s *BloggingEventCommandService) CreateArticle(ctx context.Context, in model.CreateArticleEvent, out *db.SingleStatementResult[*model.BloggingEventKey]) db.Statement {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("BloggingEventCommandService#AllEventsWithArticleID").End()
	return gw.NewStatement(func(ctx context.Context, tx *gorm.DB, out db.StatementResult) (err error) {
		nrtx := newrelic.FromContext(ctx)
		defer nrtx.StartSegment("BloggingEventCommandService#AllEventsWithArticleID").End()
		logger := slog.Default()
		logger.Info("START")

		tx = tx.WithContext(ctx)

		eventID := fmt.Sprintf("%s", s.ulidGen())
		articleID := fmt.Sprintf("%s", s.ulidGen())
		event := bloggingEvent{
			EventID:   eventID,
			ArticleID: articleID,
			Title:     toPtrString(in.Title()),
			Content:   toPtrString(in.Content()),
			Thumbnail: toPtrString(in.Thumbnail()),
			Tags:      in.Tags(),
		}
		result := tx.Create(&event)
		if err := result.Error; err != nil {
			return err
		}

		key := model.NewBloggingEventKey(eventID, articleID)
		out.Set(&key)
		logger.Info("END")
		return nil
	}, out)
}

func NewBloggingEventCommandService(ulidGen *pkg.ULIDGenerator) *BloggingEventCommandService {
	if ulidGen == nil {
		return &BloggingEventCommandService{
			ulidGen: ulid.Make,
		}
	}
	return &BloggingEventCommandService{
		ulidGen: *ulidGen,
	}
}

func toPtrString(s string) *string {
	return &s
}
