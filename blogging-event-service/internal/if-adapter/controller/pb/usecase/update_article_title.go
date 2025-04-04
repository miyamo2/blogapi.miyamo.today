//go:generate mockgen -source=$GOFILE -destination=../../../../mock/if-adapter/controller/pb/usecase/$GOFILE -package=$GOPACKAGE
package usecase

import (
	"blogapi.miyamo.today/blogging-event-service/internal/app/usecase/dto"
	"context"
)

// UpdateArticleTitle is a use-case interface for updating the title of an article.
type UpdateArticleTitle interface {
	// Execute updates the title of an article.
	Execute(ctx context.Context, in *dto.UpdateArticleTitleInDto) (*dto.UpdateArticleTitleOutDto, error)
}
