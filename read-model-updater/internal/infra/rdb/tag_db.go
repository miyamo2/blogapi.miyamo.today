package rdb

import (
	"context"
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	"github.com/miyamo2/altnrslog"
	"github.com/miyamo2/blogapi.miyamo.today/core/db"
	gw "github.com/miyamo2/blogapi.miyamo.today/core/db/gorm"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/domain/model"
	"github.com/newrelic/go-agent/v3/newrelic"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
	"log/slog"
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

		logger, err := altnrslog.FromContext(ctx)
		if err != nil {
			logger = slog.Default()
		}
		logger.Info("[RMU] START")

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
				Columns: []clause.Column{{Name: "id"}, {Name: "tag_id"}},
				DoUpdates: clause.Assignments(map[string]interface{}{
					"title":      a.Title,
					"thumbnail":  a.Thumbnail,
					"updated_at": now,
				}),
			}).Create(a)
		}

		tagIDToBeDeleted := func() []string {
			var ids []string
			for _, ti := range in.Tags() {
				ids = append(ids, ti.ID())
			}
			return ids
		}()
		tx.Where("id = ?", in.ID()).
			Where("tag_id NOT IN (?)", tagIDToBeDeleted).Delete(&tagArticle{})
		logger.Info("[RMU] END")

		tx.Where("NOT EXISTS (SELECT 1 FROM articles WHERE articles.tag_id = tags.id)").Delete(&tag{})
		return nil
	}, nil)
}

func NewTagCommandService() *TagCommandService {
	return &TagCommandService{}
}
