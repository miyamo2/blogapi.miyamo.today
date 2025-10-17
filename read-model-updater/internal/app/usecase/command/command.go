package command

import (
	"context"

	"blogapi.miyamo.today/read-model-updater/internal/infra/rdb/sqlc/article"
	"blogapi.miyamo.today/read-model-updater/internal/infra/rdb/sqlc/tag"
	"github.com/jackc/pgx/v5"
)

// Article provides commands for Article.
type Article interface {
	PreAttachTags(ctx context.Context, arg []article.PreAttachTagsParams) (int64, error)
	WithTx(tx pgx.Tx) *article.Queries
	AttachTags(ctx context.Context, articleID string) error
	CreateTempTagsTable(ctx context.Context) error
	PutArticle(ctx context.Context, arg article.PutArticleParams) error
}

// Tag provides commands for Tag.
type Tag interface {
	PrePutArticle(ctx context.Context, arg []tag.PrePutArticleParams) (int64, error)
	PrePutTags(ctx context.Context, arg []tag.PrePutTagsParams) (int64, error)
	CreateTempArticlesTable(ctx context.Context) error
	CreateTempTagsTable(ctx context.Context) error
	PutArticle(ctx context.Context) error
	PutTags(ctx context.Context) error
}

// ArticleTx provides transaction for Article.
type ArticleTx interface {
	Begin(tx pgx.Tx) Article
}

type articleTx struct {
	queries *article.Queries
}

func (a *articleTx) Begin(tx pgx.Tx) Article {
	return a.queries.WithTx(tx)
}

func NewArticleTx(queries *article.Queries) ArticleTx {
	return &articleTx{queries: queries}
}

// TagTx provides transaction for Tag.
type TagTx interface {
	Begin(tx pgx.Tx) Tag
}

type tagTx struct {
	queries *tag.Queries
}

func (t *tagTx) Begin(tx pgx.Tx) Tag {
	return t.queries.WithTx(tx)
}

func NewTagTx(queries *tag.Queries) TagTx {
	return &tagTx{queries: queries}
}
