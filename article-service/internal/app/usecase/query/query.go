package query

import (
	"context"

	"blogapi.miyamo.today/article-service/internal/infra/rdb/sqlc"
)

type Queries interface {
	GetByID(ctx context.Context, id string) (sqlc.GetByIDRow, error)
	ListAfter(ctx context.Context) ([]sqlc.ListAfterRow, error)
	ListAfterWithLimit(ctx context.Context, limit int32) ([]sqlc.ListAfterWithLimitRow, error)
	ListAfterWithLimitAndCursor(
		ctx context.Context, arg sqlc.ListAfterWithLimitAndCursorParams,
	) ([]sqlc.ListAfterWithLimitAndCursorRow, error)
	ListBefore(ctx context.Context) ([]sqlc.ListBeforeRow, error)
	ListBeforeWithLimit(ctx context.Context, limit int32) ([]sqlc.ListBeforeWithLimitRow, error)
	ListBeforeWithLimitAndCursor(
		ctx context.Context, arg sqlc.ListBeforeWithLimitAndCursorParams,
	) ([]sqlc.ListBeforeWithLimitAndCursorRow, error)
}

// NewListAfterWithLimitAndCursorParams constructs ListAfterWithLimitAndCursorParams
func NewListAfterWithLimitAndCursorParams(limit int32, cursor string) sqlc.ListAfterWithLimitAndCursorParams {
	return sqlc.ListAfterWithLimitAndCursorParams{
		ID:    cursor,
		Limit: limit,
	}
}

// NewListBeforeWithLimitAndCursorParams constructs ListBeforeWithLimitAndCursorParams
func NewListBeforeWithLimitAndCursorParams(limit int32, cursor string) sqlc.ListBeforeWithLimitAndCursorParams {
	return sqlc.ListBeforeWithLimitAndCursorParams{
		ID:    cursor,
		Limit: limit,
	}
}
