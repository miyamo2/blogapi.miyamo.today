//go:generate mockgen -source=get_all.go -destination=../../../../mock/if-adapter/controller/pb/usecase/mock_get_all.go -package=mock_usecase
package usecase

import (
	"context"
)

// GetAll is a use-case interface for getting all articles.
type GetAll[T Tag, A Article[T], O GetAllOutDto[T, A]] interface {
	// Execute gets all articles.
	Execute(ctx context.Context) (O, error)
}
