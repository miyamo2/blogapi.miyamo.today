package dynamo

import (
	"context"
	"fmt"
	"github.com/cockroachdb/errors"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/domain/model"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/pkg"
	"github.com/miyamo2/blogapi.miyamo.today/core/db"
	gw "github.com/miyamo2/blogapi.miyamo.today/core/db/gorm"
	"github.com/miyamo2/sqldav"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
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

var (
	_ schema.Tabler = (*bloggingEventCreateArticle)(nil)
	_ schema.Tabler = (*bloggingEventUpdateArticleTitle)(nil)
)

type bloggingEventCreateArticle struct {
	EventID   string `gorm:"primaryKey"`
	ArticleID string `gorm:"primaryKey"`
	Title     string
	Content   string
	Thumbnail string
	Tags      sqldav.Set[string]
}

func (b bloggingEventCreateArticle) TableName() string {
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

		event := bloggingEventCreateArticle{
			EventID:   eventID,
			ArticleID: articleID,
			Title:     in.Title(),
			Content:   in.Content(),
			Thumbnail: in.Thumbnail(),
			Tags:      sqldav.Set[string](in.Tags()),
		}

		if err := tx.Create(&event).Error; err != nil {
			err = errors.WithStack(err)
			nrtx.NoticeError(nrpkgerrors.Wrap(err))
		}

		key := model.NewBloggingEventKey(eventID, articleID)
		out.Set(&key)
		logger.Info("END")
		return nil
	}, out)
}

type bloggingEventUpdateArticleTitle struct {
	EventID   string `gorm:"primaryKey"`
	ArticleID string `gorm:"primaryKey"`
	Title     string
}

func (b bloggingEventUpdateArticleTitle) TableName() string {
	return os.Getenv("BLOGGING_EVENTS_TABLE_NAME")
}

func (s *BloggingEventCommandService) UpdateArticleTitle(ctx context.Context, in model.UpdateArticleTitleEvent, out *db.SingleStatementResult[*model.BloggingEventKey]) db.Statement {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("BloggingEventCommandService#UpdateArticleTitle").End()
	return gw.NewStatement(func(ctx context.Context, tx *gorm.DB, out db.StatementResult) (err error) {
		nrtx := newrelic.FromContext(ctx)
		defer nrtx.StartSegment("BloggingEventCommandService#UpdateArticleTitle").End()
		logger := slog.Default()
		logger.Info("START")

		tx = tx.WithContext(ctx)

		eventID := fmt.Sprintf("%s", s.ulidGen())
		articleID := in.ArticleID()

		event := bloggingEventUpdateArticleTitle{
			EventID:   eventID,
			ArticleID: articleID,
			Title:     in.Title(),
		}
		if err := tx.Create(&event).Error; err != nil {
			err = errors.WithStack(err)
			nrtx.NoticeError(nrpkgerrors.Wrap(err))
		}

		key := model.NewBloggingEventKey(eventID, articleID)
		out.Set(&key)
		logger.Info("END")
		return nil
	}, out)
}

type bloggingEventUpdateArticleBody struct {
	EventID   string `gorm:"primaryKey"`
	ArticleID string `gorm:"primaryKey"`
	Body      string
}

func (b bloggingEventUpdateArticleBody) TableName() string {
	return os.Getenv("BLOGGING_EVENTS_TABLE_NAME")
}

func (s *BloggingEventCommandService) UpdateArticleBody(ctx context.Context, in model.UpdateArticleBodyEvent, out *db.SingleStatementResult[*model.BloggingEventKey]) db.Statement {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("BloggingEventCommandService#UpdateArticleBody").End()
	return gw.NewStatement(func(ctx context.Context, tx *gorm.DB, out db.StatementResult) (err error) {
		nrtx := newrelic.FromContext(ctx)
		defer nrtx.StartSegment("BloggingEventCommandService#UpdateArticleBody").End()
		logger := slog.Default()
		logger.Info("START")

		tx = tx.WithContext(ctx)

		eventID := fmt.Sprintf("%s", s.ulidGen())
		articleID := in.ArticleID()

		event := bloggingEventUpdateArticleBody{
			EventID:   eventID,
			ArticleID: articleID,
			Body:      in.Body(),
		}
		if err := tx.Create(&event).Error; err != nil {
			err = errors.WithStack(err)
			nrtx.NoticeError(nrpkgerrors.Wrap(err))
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
