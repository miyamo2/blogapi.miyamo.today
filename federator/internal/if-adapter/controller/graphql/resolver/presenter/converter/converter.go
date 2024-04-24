//go:generate mockgen -source=converter.go -destination=../../../../../../mock/if-adapter/controller/graphql/resolver/presenter/converter/mock_converter.go -package=converter
package converter

import (
	"context"

	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/if-adapter/controller/graphql/resolver/usecase/dto"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/if-adapter/presenters/graphql/model"
)

type ArticleConverter[T dto.Tag, AT dto.ArticleTag[T], O dto.ArticleOutDto[T, AT]] interface {
	ToArticle(ctx context.Context, from O) (*model.ArticleNode, bool)
}

type ArticlesConverter[T dto.Tag, AT dto.ArticleTag[T], O dto.ArticlesOutDto[T, AT]] interface {
	ToArticles(ctx context.Context, from O) (*model.ArticleConnection, bool)
}

type TagConverter[A dto.Article, TA dto.TagArticle[A], O dto.TagOutDto[A, TA]] interface {
	ToTag(ctx context.Context, from O) (*model.TagNode, error)
}

type TagsConverter[A dto.Article, TA dto.TagArticle[A], O dto.TagsOutDto[A, TA]] interface {
	ToTags(ctx context.Context, from O) (*model.TagConnection, error)
}
