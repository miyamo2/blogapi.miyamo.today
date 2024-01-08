//go:generate mockgen -source=get_prev.go -destination=../../../../mock/if-adapter/controller/pb/usecase/mock_get_prev.go -package=mock_usecase
package usecase

import (
	"context"
)

// GetPrev is a use-case interface for getting articles.
type GetPrev[I GetPrevInDto, T Tag, A Article[T], O GetPrevOutDto[T, A]] interface {
	// Execute gets all articles.
	Execute(ctx context.Context, in I) (O, error)
}
