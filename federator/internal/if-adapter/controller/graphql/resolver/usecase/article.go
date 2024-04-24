//go:generate mockgen -source=article.go -destination=../../../../../mock/if-adapter/controller/graphql/resolver/usecase/mock_article.go -package=usecase
package usecase

import (
	"context"

	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/if-adapter/controller/graphql/resolver/usecase/dto"
)

// Article is a use-case of getting an article by id.
type Article[I dto.ArticleInDto, T dto.Tag, AT dto.ArticleTag[T], O dto.ArticleOutDto[T, AT]] interface {
	// Execute gets an article by id.
	Execute(ctx context.Context, in I) (O, error)
}

// Articles is a use-case of getting an articles.
type Articles[I dto.ArticlesInDto, T dto.Tag, AT dto.ArticleTag[T], O dto.ArticlesOutDto[T, AT]] interface {
	// Execute gets articles.
	Execute(ctx context.Context, in I) (O, error)
}
