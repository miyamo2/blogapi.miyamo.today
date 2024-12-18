//go:generate mockgen -source=$GOFILE -destination=../../../../../../mock/if-adapter/controller/graphql/resolver/presenter/converter/$GOFILE -package=$GOPACKAGE
package converters

import (
	"context"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/app/usecase/dto"

	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/if-adapter/presenters/graphql/model"
)

type ArticleConverter interface {
	ToArticle(ctx context.Context, from dto.ArticleOutDTO) (*model.ArticleNode, bool)
}

type ArticlesConverter interface {
	ToArticles(ctx context.Context, from dto.ArticlesOutDTO) (*model.ArticleConnection, bool)
}

type TagConverter interface {
	ToTag(ctx context.Context, from dto.TagOutDTO) (*model.TagNode, error)
}

type TagsConverter interface {
	ToTags(ctx context.Context, from dto.TagsOutDTO) (*model.TagConnection, error)
}

type CreateArticleConverter interface {
	ToCreateArticle(ctx context.Context, from dto.CreateArticleOutDTO) (*model.CreateArticlePayload, error)
}

type UpdateArticleTitleConverter interface {
	ToUpdateArticleTitle(ctx context.Context, from dto.UpdateArticleTitleOutDTO) (*model.UpdateArticleTitlePayload, error)
}

type UpdateArticleBodyConverter interface {
	ToUpdateArticleBody(ctx context.Context, from dto.UpdateArticleBodyOutDTO) (*model.UpdateArticleBodyPayload, error)
}
