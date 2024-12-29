//go:generate mockgen -source=$GOFILE -destination=../../../../mock/if-adapter/controller/pb/usecase/$GOFILE -package=$GOPACKAGE
package usecase

import (
	"blogapi.miyamo.today/blogging-event-service/internal/app/usecase/dto"
	"context"
)

// UpdateArticleBody is a use-case interface for updating the body of an article.
type UpdateArticleBody interface {
	// Execute updates the body of an article.
	Execute(ctx context.Context, in *dto.UpdateArticleBodyInDto) (*dto.UpdateArticleBodyOutDto, error)
}
