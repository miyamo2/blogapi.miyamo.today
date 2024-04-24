//go:generate mockgen -source=get_next.go -destination=../../../../mock/if-adapter/controller/pb/usecase/mock_get_next.go -package=mock_usecase
package usecase

import (
	"context"
)

// GetNext is a use-case interface for getting next tags.
type GetNext[I GetNextInDto, A Article, T Tag[A], O GetNextOutDto[A, T]] interface {
	// Execute gets all tags.
	Execute(ctx context.Context, in I) (O, error)
}
