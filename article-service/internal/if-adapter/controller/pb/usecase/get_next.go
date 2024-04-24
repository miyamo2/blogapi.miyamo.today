//go:generate mockgen -source=get_next.go -destination=../../../../mock/if-adapter/controller/pb/usecase/mock_get_next.go -package=mock_usecase
package usecase

import (
	"context"
)

// GetNext is a use-case interface for getting articles.
type GetNext[I GetNextInDto, T Tag, A Article[T], O GetNextOutDto[T, A]] interface {
	// Execute gets all articles.
	Execute(ctx context.Context, in I) (O, error)
}
