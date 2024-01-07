//go:generate mockgen -source=get_by_id.go -destination=../../../../mock/if-adapter/controller/pb/usecase/mock_get_by_id.go -package=mock_usecase
package usecase

import (
	"context"
)

// GetById is a use-case interface for getting an article by id.
type GetById[I GetByIdInDto, T Tag, A Article[T]] interface {
	// Execute gets an article by id.
	Execute(ctx context.Context, in I) (A, error)
}
