package usecase

import (
	"context"

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

	last := min(max(in.Last(), 1), 100) // TODO: config

	var (
		tags        = make([]dto.Tag, 0, last)
		hasPrevious bool
	)

	switch {
	case in.Cursor() != nil:
		rows, err := u.queries.ListBeforeWithLimitAndCursor(
			ctx,
			query.NewListBeforeWithLimitAndCursorParams(int32(last+1), *in.Cursor()),
		)
		if err != nil {
			return nil, err
		}
		hasPrevious = len(rows) > last
		for row := range getPage(rows, last, 1) {
			tags = append(
				tags, dto.NewTag(
					row.ID,
					row.Name,
					articleDtoFromQueryModel(row.Articles)...,
				),
			)
		}
	default:
		rows, err := u.queries.ListBeforeWithLimit(ctx, int32(last+1))
		if err != nil {
			return nil, err
		}
		hasPrevious = len(rows) > last
		for row := range getPage(rows, last, 1) {
			tags = append(
				tags, dto.NewTag(
					row.ID,
					row.Name,
					articleDtoFromQueryModel(row.Articles)...,
				),
			)
		}
	}

	result := dto.NewListBeforeOutput(hasPrevious, tags...)
	return &result, nil
}

// NewListBefore constructs ListBefore
func NewListBefore(queries query.Queries) *ListBefore {
	return &ListBefore{queries: queries}
}
