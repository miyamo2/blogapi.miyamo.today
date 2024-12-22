//go:generate mockgen -source=$GOFILE -destination=../../../../../mock/if-adapter/controller/graphql/resolver/usecase/$GOFILE -package=$GOPACKAGE
package usecase

import (
	"context"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/app/usecase/dto"
)

// CreateArticle is the usecase for creating an article.
type CreateArticle interface {
	// Execute creates an article.
	Execute(ctx context.Context, in dto.CreateArticleInDTO) (dto.CreateArticleOutDTO, error)
}

// UpdateArticleTitle is the usecase for updating an article title.
type UpdateArticleTitle interface {
	// Execute updates an article title.
	Execute(ctx context.Context, in dto.UpdateArticleTitleInDTO) (dto.UpdateArticleTitleOutDTO, error)
}

// UpdateArticleBody is the usecase for updating an article body.
type UpdateArticleBody interface {
	// Execute updates an article body.
	Execute(ctx context.Context, in dto.UpdateArticleBodyInDTO) (dto.UpdateArticleBodyOutDTO, error)
}

// UpdateArticleThumbnail is the usecase for updating an article thumbnail.
type UpdateArticleThumbnail interface {
	// Execute updates an article thumbnail.
	Execute(ctx context.Context, in dto.UpdateArticleThumbnailInDTO) (dto.UpdateArticleThumbnailOutDTO, error)
}
