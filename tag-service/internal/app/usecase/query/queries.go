package query

import (
	"context"

	"blogapi.miyamo.today/tag-service/internal/infra/rdb"
)

// Queries provides all queries methods
type Queries interface {
	GetByID(ctx context.Context, id string) (rdb.GetByIDRow, error)
	ListAfter(ctx context.Context) ([]rdb.ListAfterRow, error)
	ListAfterWithLimit(ctx context.Context, limit int32) ([]rdb.ListAfterWithLimitRow, error)
	ListAfterWithLimitAndCursor(
		ctx context.Context, arg rdb.ListAfterWithLimitAndCursorParams,
	) ([]rdb.ListAfterWithLimitAndCursorRow, error)
	ListBefore(ctx context.Context) ([]rdb.ListBeforeRow, error)
	ListBeforeWithLimit(ctx context.Context, limit int32) ([]rdb.ListBeforeWithLimitRow, error)
	ListBeforeWithLimitAndCursor(
		ctx context.Context, arg rdb.ListBeforeWithLimitAndCursorParams,
	) ([]rdb.ListBeforeWithLimitAndCursorRow, error)
}

// NewListAfterWithLimitAndCursorParams constructs ListAfterWithLimitAndCursorParams
func NewListAfterWithLimitAndCursorParams(limit int32, cursor string) rdb.ListAfterWithLimitAndCursorParams {
	return rdb.ListAfterWithLimitAndCursorParams{
		ID:    cursor,
		Limit: limit,
	}
}

// NewListBeforeWithLimitAndCursorParams constructs ListBeforeWithLimitAndCursorParams
func NewListBeforeWithLimitAndCursorParams(limit int32, cursor string) rdb.ListBeforeWithLimitAndCursorParams {
	return rdb.ListBeforeWithLimitAndCursorParams{
		ID:    cursor,
		Limit: limit,
	}
}
