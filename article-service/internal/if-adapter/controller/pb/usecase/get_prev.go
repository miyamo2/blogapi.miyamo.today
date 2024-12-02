//go:generate mockgen -source=get_prev.go -destination=../../../../mock/if-adapter/controller/pb/usecase/mock_get_prev.go -package=mock_usecase
package usecase

import (
	"context"
	"github.com/miyamo2/blogapi.miyamo.today/article-service/internal/app/usecase/dto"
)

// GetPrev is a use-case interface for getting articles.
type GetPrev interface {
	// Execute gets all articles.
	Execute(ctx context.Context, in dto.GetPrevInDto) (*dto.GetPrevOutDto, error)
}
