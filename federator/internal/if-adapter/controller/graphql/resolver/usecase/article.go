//go:generate mockgen -source=article.go -destination=../../../../../mock/if-adapter/controller/graphql/resolver/usecase/mock_article.go -package=usecase
package usecase

import (
	"context"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/app/usecase/dto"
)

// Article is a use-case of getting an article by id.
type Article interface {
	// Execute gets an article by id.
	Execute(ctx context.Context, in dto.ArticleInDto) (dto.ArticleOutDto, error)
}

// Articles is a use-case of getting an articles.
type Articles interface {
	// Execute gets articles.
	Execute(ctx context.Context, in dto.ArticlesInDto) (dto.ArticlesOutDto, error)
}
