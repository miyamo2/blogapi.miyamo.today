package usecase

import (
	"context"

	"blogapi.miyamo.today/tag-service/internal/app/usecase/dto"
)

// ListAll provides a use-case interface for getting all tags.
type ListAll interface {
	// Execute gets all tags.
	Execute(ctx context.Context) (*dto.ListAllOutput, error)
}
