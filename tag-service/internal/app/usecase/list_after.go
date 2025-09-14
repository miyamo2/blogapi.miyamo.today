package usecase

import (
	"context"
	"fmt"

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

	var (
		tags    = make([]dto.Tag, 0, in.First())
		hasNext bool
	)

	switch {
	case in.Cursor() != nil && in.First() != 0:
		rows, err := u.queries.ListAfterWithLimitAndCursor(
			ctx,
			query.NewListAfterWithLimitAndCursorParams(int32(in.First()+1), *in.Cursor()),
		)
		if err != nil {
			return nil, err
		}
		hasNext = len(rows) > in.First()
		for row := range getPage(rows, in.First(), 1) {
			tags = append(
				tags, dto.NewTag(
					row.ID,
					row.Name,
					articleDtoFromQueryModel(row.Articles)...,
				),
			)
		}
	case in.First() != 0:
		rows, err := u.queries.ListAfterWithLimit(ctx, int32(in.First()+1))
		if err != nil {
			return nil, err
		}
		hasNext = len(rows) > in.First()
		for row := range getPage(rows, in.First(), 1) {
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
	result := dto.NewListAfterOutput(hasNext, tags...)
	return &result, nil
}

// NewListAfter constructs ListAfter
func NewListAfter(queries query.Queries) *ListAfter {
	return &ListAfter{queries: queries}
}
