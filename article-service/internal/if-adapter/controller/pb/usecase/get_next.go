//go:generate mockgen -source=get_next.go -destination=../../../../mock/if-adapter/controller/pb/usecase/mock_get_next.go -package=mock_usecase
package usecase

import (
	"blogapi.miyamo.today/article-service/internal/app/usecase/dto"
	"context"
)

// GetNext is a use-case interface for getting articles.
type GetNext interface {
	// Execute gets all articles.
	Execute(ctx context.Context, in dto.GetNextInDto) (*dto.GetNextOutDto, error)
}
