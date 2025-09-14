package usecase

import (
	"context"

	"blogapi.miyamo.today/tag-service/internal/app/usecase/dto"
)

// GetById provides a use-case interface for getting a tag by id.
type GetById interface {
	// Execute gets a tag by id.
	Execute(ctx context.Context, in dto.GetByIdInput) (*dto.GetByIdOutput, error)
}
