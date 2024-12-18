//go:generate mockgen -source=$GOFILE -destination=../../../../../mock/if-adapter/controller/graphql/resolver/usecase/$GOFILE -package=$GOPACKAGE
package usecase

import (
	"context"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/app/usecase/dto"
)

type CreateArticle interface {
	Execute(ctx context.Context, in dto.CreateArticleInDTO) (dto.CreateArticleOutDTO, error)
}

type UpdateArticleTitle interface {
	Execute(ctx context.Context, in dto.UpdateArticleTitleInDTO) (dto.UpdateArticleTitleOutDTO, error)
}

type UpdateArticleBody interface {
	Execute(ctx context.Context, in dto.UpdateArticleBodyInDTO) (dto.UpdateArticleBodyOutDTO, error)
}
