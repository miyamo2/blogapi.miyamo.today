package usecase

import (
	"context"

	"blogapi.miyamo.today/article-service/internal/app/usecase/dto"
)

// ListBefore provides the feature to list articles before a given cursor.
type ListBefore interface {
	// Execute lists articles before a given cursor.
	Execute(ctx context.Context, in dto.ListBeforeInput) (*dto.ListBeforeOutput, error)
}
