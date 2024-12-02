//go:generate mockgen -source=converter.go -destination=../../../../../../mock/if-adapter/controller/graphql/resolver/presenter/converter/mock_converter.go -package=converter
package converter

import (
	"context"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/app/usecase/dto"

	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/if-adapter/presenters/graphql/model"
)

type ArticleConverter interface {
	ToArticle(ctx context.Context, from dto.ArticleOutDto) (*model.ArticleNode, bool)
}

type ArticlesConverter interface {
	ToArticles(ctx context.Context, from dto.ArticlesOutDto) (*model.ArticleConnection, bool)
}

type TagConverter interface {
	ToTag(ctx context.Context, from dto.TagOutDto) (*model.TagNode, error)
}

type TagsConverter interface {
	ToTags(ctx context.Context, from dto.TagsOutDto) (*model.TagConnection, error)
}
