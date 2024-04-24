//go:generate mockgen -source=get_prev.go -destination=../../../../mock/if-adapter/controller/pb/usecase/mock_get_prev.go -package=mock_usecase
package usecase

import (
	"context"
)

// GetPrev is a use-case interface for getting previous tags.
type GetPrev[I GetPrevInDto, A Article, T Tag[A], O GetPrevOutDto[A, T]] interface {
	// Execute gets all tags.
	Execute(ctx context.Context, in I) (O, error)
}
