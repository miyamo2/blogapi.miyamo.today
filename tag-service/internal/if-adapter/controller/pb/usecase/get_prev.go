//go:generate mockgen -source=get_prev.go -destination=../../../../mock/if-adapter/controller/pb/usecase/mock_get_prev.go -package=mock_usecase
package usecase

import (
	"blogapi.miyamo.today/tag-service/internal/app/usecase/dto"
	"context"
)

// GetPrev is a use-case interface for getting previous tags.
type GetPrev interface {
	// Execute gets all tags.
	Execute(ctx context.Context, in dto.GetPrevInDto) (*dto.GetPrevOutDto, error)
}
