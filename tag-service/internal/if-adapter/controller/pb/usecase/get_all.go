//go:generate mockgen -source=get_all.go -destination=../../../../mock/if-adapter/controller/pb/usecase/mock_get_all.go -package=mock_usecase
package usecase

import (
	"blogapi.miyamo.today/tag-service/internal/app/usecase/dto"
	"context"
)

// GetAll is a use-case interface for getting all tags.
type GetAll interface {
	// Execute gets all tags.
	Execute(ctx context.Context) (*dto.GetAllOutDto, error)
}
