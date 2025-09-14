package usecase

import (
	"context"
	"fmt"

	"blogapi.miyamo.today/tag-service/internal/app/usecase/query"
	"github.com/newrelic/go-agent/v3/newrelic"

	"blogapi.miyamo.today/tag-service/internal/app/usecase/dto"
)

// ListBefore implements usecase.ListBefore
type ListBefore struct {
	queries query.Queries
}

func (u *ListBefore) Execute(ctx context.Context, in dto.ListBeforeInput) (*dto.ListBeforeOutput, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("Execute").End()

	var (
		tags        = make([]dto.Tag, 0, in.Last())
		hasPrevious bool
	)

	switch {
	case in.Cursor() != nil && in.Last() != 0:
		rows, err := u.queries.ListBeforeWithLimitAndCursor(
			ctx,
			query.NewListBeforeWithLimitAndCursorParams(int32(in.Last()+1), *in.Cursor()),
		)
		if err != nil {
			return nil, err
		}
		hasPrevious = len(rows) > in.Last()
		for row := range getPage(rows, in.Last(), 1) {
			tags = append(
				tags, dto.NewTag(
					row.ID,
					row.Name,
					articleDtoFromQueryModel(row.Articles)...,
				),
			)
		}
	case in.Last() != 0:
		rows, err := u.queries.ListBeforeWithLimit(ctx, int32(in.Last()+1))
		if err != nil {
			return nil, err
		}
		hasPrevious = len(rows) > in.Last()
		for row := range getPage(rows, in.Last(), 1) {
			tags = append(
				tags, dto.NewTag(
					row.ID,
					row.Name,
					articleDtoFromQueryModel(row.Articles)...,
				),
			)
		}
	default:
		return nil, fmt.Errorf("invalid arguments")
	}

	result := dto.NewListBeforeOutput(hasPrevious, tags...)
	return &result, nil
}

// NewListBefore constructs ListBefore
func NewListBefore(queries query.Queries) *ListBefore {
	return &ListBefore{queries: queries}
}
