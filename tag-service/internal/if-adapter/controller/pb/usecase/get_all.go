//go:generate mockgen -source=get_all.go -destination=../../../../mock/if-adapter/controller/pb/usecase/mock_get_all.go -package=mock_usecase
package usecase

import (
	"context"
)

// GetAll is a use-case interface for getting all tags.
type GetAll[A Article, T Tag[A], O GetAllOutDto[A, T]] interface {
	// Execute gets all tags.
	Execute(ctx context.Context) (O, error)
}
