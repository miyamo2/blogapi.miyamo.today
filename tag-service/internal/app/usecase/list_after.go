package usecase

import (
	"context"

	"blogapi.miyamo.today/tag-service/internal/app/usecase/query"
	"github.com/newrelic/go-agent/v3/newrelic"

	"blogapi.miyamo.today/tag-service/internal/app/usecase/dto"
)

// ListAfter implements usecase.ListAfter
type ListAfter struct {
	queries query.Queries
}

func (u *ListAfter) Execute(ctx context.Context, in dto.ListAfterInput) (*dto.ListAfterOutput, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("Execute").End()

	first := min(max(in.First(), 1), 100) // TODO: config

	var (
		tags    = make([]dto.Tag, 0, first)
		hasNext bool
	)

	switch {
	case in.Cursor() != nil:
		rows, err := u.queries.ListAfterWithLimitAndCursor(
			ctx,
			query.NewListAfterWithLimitAndCursorParams(int32(first+1), *in.Cursor()),
		)
		if err != nil {
			return nil, err
		}
		hasNext = len(rows) > first
		for row := range getPage(rows, first, 1) {
			tags = append(
				tags, dto.NewTag(
					row.ID,
					row.Name,
					articleDtoFromQueryModel(row.Articles)...,
				),
			)
		}
	default:
		rows, err := u.queries.ListAfterWithLimit(ctx, int32(first+1))
		if err != nil {
			return nil, err
		}
		hasNext = len(rows) > first
		for row := range getPage(rows, first, 1) {
			tags = append(
				tags, dto.NewTag(
					row.ID,
					row.Name,
					articleDtoFromQueryModel(row.Articles)...,
				),
			)
		}
	}
	result := dto.NewListAfterOutput(hasNext, tags...)
	return &result, nil
}

// NewListAfter constructs ListAfter
func NewListAfter(queries query.Queries) *ListAfter {
	return &ListAfter{queries: queries}
}
