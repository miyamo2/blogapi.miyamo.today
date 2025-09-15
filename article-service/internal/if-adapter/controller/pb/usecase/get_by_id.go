package usecase

import (
	"context"

	"blogapi.miyamo.today/article-service/internal/app/usecase/dto"
)

// GetByID provides the feature to get an article by id.
type GetByID interface {
	// Execute gets an article by id.
	Execute(ctx context.Context, in dto.GetByIDInput) (*dto.GetByIDOutput, error)
}
