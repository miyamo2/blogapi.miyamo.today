//go:generate mockgen -source=$GOFILE -destination=../../../../mock/if-adapter/controller/pb/usecase/$GOFILE -package=$GOPACKAGE
package usecase

import (
	"context"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/app/usecase/dto"
)

// CreateArticle is a use-case interface for creating an article.
type CreateArticle interface {
	// Execute creates an article.
	Execute(ctx context.Context, in *dto.CreateArticleInDto) (*dto.CreateArticleOutDto, error)
}
