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

const TagDBName = "tag"

var (
	_ schema.Tabler = (*tag)(nil)
	_ schema.Tabler = (*tagArticle)(nil)
)

type tag struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	CreatedAt synchro.Time[tz.UTC]
	UpdatedAt synchro.Time[tz.UTC]
}

func (t *tag) TableName() string {
	return "tags"
}

type tagArticle struct {
	ID        string `gorm:"primaryKey"`
	TagID     string `gorm:"primaryKey"`
	Title     string
	Thumbnail string
	CreatedAt synchro.Time[tz.UTC]
	UpdatedAt synchro.Time[tz.UTC]
}

func (t *tagArticle) TableName() string {
	return "articles"
}

type TagCommandService struct{}

func (c *TagCommandService) ExecuteTagCommand(ctx context.Context, in model.ArticleCommand) db.Statement {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ArticleCommandService#ExecuteTagCommand").End()
	return gw.NewStatement(func(ctx context.Context, tx *gorm.DB, out db.StatementResult) error {
		nrtx := newrelic.FromContext(ctx)
		defer nrtx.StartSegment("ArticleCommandService#ExecuteTagCommand#Execute").End()
		tx = tx.WithContext(ctx)
		now := synchro.Now[tz.UTC]()

		for _, ti := range in.Tags() {
			t := tag{
				ID:        ti.ID(),
				Name:      ti.Name(),
				CreatedAt: now,
				UpdatedAt: now,
			}
			tx.Clauses(clause.OnConflict{DoNothing: true}).Create(t)

			a := &tagArticle{
				ID:        in.ID(),
				TagID:     ti.ID(),
				Title:     in.Title(),
				Thumbnail: in.Thumbnail(),
				CreatedAt: now,
				UpdatedAt: now,
			}

			tx.Clauses(clause.OnConflict{
				DoUpdates: clause.Assignments(map[string]interface{}{
					"title":      a.Title,
					"thumbnail":  a.Thumbnail,
					"updated_at": now,
				}),
			}).Create(a)
		}

		tx.Where("id = ?", in.ID()).
			Where("tag_id NOT IN (?)", func() []string {
				var ids []string
				for _, ti := range in.Tags() {
					ids = append(ids, ti.ID())
				}
				return ids
			}).
			Delete(&tagArticle{})
		return nil
	}, nil)
}

func NewTagCommandService() *TagCommandService {
	return &TagCommandService{}
}
