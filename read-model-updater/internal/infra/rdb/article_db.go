package rdb

import (
	"context"

	"blogapi.miyamo.today/core/db"
	gw "blogapi.miyamo.today/core/db/gorm"
	"blogapi.miyamo.today/read-model-updater/internal/domain/model"
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/oklog/ulid/v2"
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

func (s *ArticleCommandService) ExecuteArticleCommand(
	ctx context.Context, in model.ArticleCommand, eventAt synchro.Time[tz.UTC],
) db.Statement {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ArticleCommandService#ExecuteArticleCommand").End()
	return gw.NewStatement(
		func(ctx context.Context, tx *gorm.DB, out db.StatementResult) error {
			nrtx := newrelic.FromContext(ctx)
			defer nrtx.StartSegment("ArticleCommandService#ExecuteArticleCommand#Execute").End()

			tx = tx.WithContext(ctx)

			articleID, err := ulid.Parse(in.ID())
			if err != nil {
				return err
			}

			v := &article{
				ID:        in.ID(),
				Title:     in.Title(),
				Body:      in.Body(),
				Thumbnail: in.Thumbnail(),
				CreatedAt: synchro.UnixMilli[tz.UTC](int64(articleID.Time())),
				UpdatedAt: eventAt,
			}

			tx.Clauses(
				clause.OnConflict{
					Columns: []clause.Column{{Name: "id"}},
					DoUpdates: clause.Assignments(
						map[string]interface{}{
							"title":      v.Title,
							"body":       v.Body,
							"thumbnail":  v.Thumbnail,
							"updated_at": eventAt,
						},
					),
				},
			).Create(v)

			existsTagIDs := func() []string {
				var ids []string
				for _, ti := range in.Tags() {
					ids = append(ids, ti.ID())
				}
				return ids
			}()
			tx.Where("article_id = ?", in.ID()).
				Where("articleID NOT IN (?)", existsTagIDs).
				Delete(&articleTag{})

			var articleTags []*articleTag
			for _, ti := range in.Tags() {
				tagID, err := ulid.Parse(ti.ID())
				if err != nil {
					return err
				}
				articleTags = append(
					articleTags, &articleTag{
						ID:        ti.ID(),
						ArticleID: in.ID(),
						Name:      ti.Name(),
						CreatedAt: synchro.UnixMilli[tz.UTC](int64(tagID.Time())),
						UpdatedAt: eventAt,
					},
				)
			}
			tx.Clauses(clause.OnConflict{DoNothing: true}).Create(articleTags)
			return nil
		}, nil,
	)
}

func NewArticleCommandService() *ArticleCommandService {
	return &ArticleCommandService{}
}
