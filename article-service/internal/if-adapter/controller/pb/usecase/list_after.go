package usecase

import (
	"context"

	"blogapi.miyamo.today/article-service/internal/app/usecase/dto"
)

// ListAfter provides the feature to list articles after a given cursor.
type ListAfter interface {
	// Execute list articles after a given cursor.
	Execute(ctx context.Context, in dto.ListAfterInput) (*dto.ListAfterOutput, error)
}
