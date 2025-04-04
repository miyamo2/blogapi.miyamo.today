//go:generate mockgen -source=get_next.go -destination=../../../../mock/if-adapter/controller/pb/usecase/mock_get_next.go -package=mock_usecase
package usecase

import (
	"blogapi.miyamo.today/tag-service/internal/app/usecase/dto"
	"context"
)

// GetNext is a use-case interface for getting next tags.
type GetNext interface {
	// Execute gets all tags.
	Execute(ctx context.Context, in dto.GetNextInDto) (*dto.GetNextOutDto, error)
}
