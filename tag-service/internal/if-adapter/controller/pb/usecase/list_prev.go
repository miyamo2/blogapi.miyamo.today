package usecase

import (
	"context"

	"blogapi.miyamo.today/tag-service/internal/app/usecase/dto"
)

// ListBefore provides a use-case interface for getting previous tags.
type ListBefore interface {
	// Execute gets all tags.
	Execute(ctx context.Context, in dto.ListBeforeInput) (*dto.ListBeforeOutput, error)
}
