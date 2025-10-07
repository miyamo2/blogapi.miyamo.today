package command

import (
	"context"

	"blogapi.miyamo.today/read-model-updater/internal/infra/rdb/sqlc/article"
	"blogapi.miyamo.today/read-model-updater/internal/infra/rdb/sqlc/tag"
	"github.com/jackc/pgx/v5"
)

// Article is a command service interface for the Article.
type Article interface {
	PreAttachTags(ctx context.Context, arg []article.PreAttachTagsParams) (int64, error)
	WithTx(tx pgx.Tx) *article.Queries
	AttachTags(ctx context.Context, articleID string) error
	CreateTempTagsTable(ctx context.Context) error
	PutArticle(ctx context.Context, arg article.PutArticleParams) error
}

type Tag interface {
	PreTagArticles(ctx context.Context, arg []tag.PreTagArticlesParams) (int64, error)
	CreateTempArticlesTable(ctx context.Context) error
	PutTag(ctx context.Context, arg tag.PutTagParams) error
	TagArticles(ctx context.Context, tagID string) error
	WithTx(tx pgx.Tx) *tag.Queries
}
