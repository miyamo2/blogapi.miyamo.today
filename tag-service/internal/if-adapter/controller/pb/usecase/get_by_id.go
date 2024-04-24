//go:generate mockgen -source=get_by_id.go -destination=../../../../mock/if-adapter/controller/pb/usecase/mock_get_by_id.go -package=mock_usecase
package usecase

import (
	"context"
)

// GetById is a use-case interface for getting a tag by id.
type GetById[I GetByIdInDto, A Article, T Tag[A]] interface {
	// Execute gets a tag by id.
	Execute(ctx context.Context, in I) (T, error)
}
