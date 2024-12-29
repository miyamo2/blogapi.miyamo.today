//go:generate mockgen -source=get_by_id.go -destination=../../../../mock/if-adapter/controller/pb/usecase/mock_get_by_id.go -package=mock_usecase
package usecase

import (
	"blogapi.miyamo.today/tag-service/internal/app/usecase/dto"
	"context"
)

// GetById is a use-case interface for getting a tag by id.
type GetById interface {
	// Execute gets a tag by id.
	Execute(ctx context.Context, in dto.GetByIdInDto) (*dto.GetByIdOutDto, error)
}
