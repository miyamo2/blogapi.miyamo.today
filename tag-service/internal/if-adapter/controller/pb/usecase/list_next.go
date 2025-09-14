package usecase

import (
	"context"

	"blogapi.miyamo.today/tag-service/internal/app/usecase/dto"
)

// ListAfter provides a use-case interface for getting next tags.
type ListAfter interface {
	// Execute gets all tags.
	Execute(ctx context.Context, in dto.ListAfterInput) (*dto.ListAfterOutput, error)
}
