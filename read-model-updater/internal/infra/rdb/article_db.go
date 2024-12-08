package rdb

import (
	"context"
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	"github.com/miyamo2/blogapi.miyamo.today/core/db"
	gw "github.com/miyamo2/blogapi.miyamo.today/core/db/gorm"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/domain/model"
	"github.com/newrelic/go-agent/v3/newrelic"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

const ArticleDBName = "article"

var (
	_ schema.Tabler = (*article)(nil)
	_ schema.Tabler = (*articleTag)(nil)
)

type article struct {
	ID        string `gorm:"primaryKey"`
	Title     string
	Body      string
	Thumbnail string
	CreatedAt synchro.Time[tz.UTC]
	UpdatedAt synchro.Time[tz.UTC]
}

func (a *article) TableName() string {
	return "articles"
}

type articleTag struct {
	ID        string `gorm:"primaryKey"`
	ArticleID string `gorm:"primaryKey"`
	Name      string
	CreatedAt synchro.Time[tz.UTC]
	UpdatedAt synchro.Time[tz.UTC]
}

func (a *articleTag) TableName() string {
	return "tags"
}

type ArticleCommandService struct{}

func (s *ArticleCommandService) ExecuteArticleCommand(ctx context.Context, in model.ArticleCommand) db.Statement {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ArticleCommandService#ExecuteArticleCommand").End()
	return gw.NewStatement(func(ctx context.Context, tx *gorm.DB, out db.StatementResult) error {
		nrtx := newrelic.FromContext(ctx)
		defer nrtx.StartSegment("ArticleCommandService#ExecuteArticleCommand#Execute").End()
		tx = tx.WithContext(ctx)
		now := synchro.Now[tz.UTC]()

		a := &article{
			ID:        in.ID(),
			Title:     in.Title(),
			Body:      in.Body(),
			Thumbnail: in.Thumbnail(),
			CreatedAt: now,
			UpdatedAt: now,
		}

		tx.Clauses(clause.OnConflict{
			DoUpdates: clause.Assignments(map[string]interface{}{
				"title":      a.Title,
				"body":       a.Body,
				"thumbnail":  a.Thumbnail,
				"updated_at": now,
			}),
		}).Create(a)

		tx.Where("article_id = ?", in.ID()).
			Where("id NOT IN (?)", func() []string {
				var ids []string
				for _, ti := range in.Tags() {
					ids = append(ids, ti.ID())
				}
				return ids
			}).
			Delete(&articleTag{})

		for _, ti := range in.Tags() {
			t := &articleTag{
				ID:        ti.ID(),
				ArticleID: in.ID(),
				Name:      ti.Name(),
				CreatedAt: now,
				UpdatedAt: now,
			}
			tx.Clauses(clause.OnConflict{DoNothing: true}).Create(t)
		}
		return nil
	}, nil)
}

func NewArticleCommandService() *ArticleCommandService {
	return &ArticleCommandService{}
}