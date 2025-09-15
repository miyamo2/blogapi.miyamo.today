package usecase

import (
	"context"

	"blogapi.miyamo.today/article-service/internal/app/usecase/dto"
)

// ListAll provides the feature to list all articles.
type ListAll interface {
	// Execute lists all articles.
	Execute(ctx context.Context) (*dto.ListAllOutput, error)
}
