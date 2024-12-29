//go:generate mockgen -source=$GOFILE -destination=../../../../../mock/if-adapter/controller/graphql/resolver/usecase/$GOFILE -package=$GOPACKAGE
package usecase

import (
	"blogapi.miyamo.today/federator/internal/app/usecase/dto"
	"context"
)

// Article is a use-case of getting an article by id.
type Article interface {
	// Execute gets an article by id.
	Execute(ctx context.Context, in dto.ArticleInDTO) (dto.ArticleOutDTO, error)
}

// Articles is a use-case of getting an articles.
type Articles interface {
	// Execute gets articles.
	Execute(ctx context.Context, in dto.ArticlesInDTO) (dto.ArticlesOutDTO, error)
}
