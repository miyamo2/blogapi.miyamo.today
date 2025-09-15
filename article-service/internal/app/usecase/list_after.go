package usecase

import (
	"context"
	"fmt"

	"github.com/newrelic/go-agent/v3/newrelic"

	"blogapi.miyamo.today/article-service/internal/app/usecase/dto"
	"blogapi.miyamo.today/article-service/internal/app/usecase/query"
)

// ListAfter implements usecase.ListAfter
type ListAfter struct {
	queries query.Queries
}

func (u *ListAfter) Execute(ctx context.Context, in dto.ListAfterInput) (*dto.ListAfterOutput, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("Execute").End()

	var (
		articles = make([]dto.Article, 0, in.First())
		hasNext  bool
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
			articles = append(
				articles, dto.NewArticle(
					row.ID,
					row.Title,
					row.Body,
					row.Thumbnail,
					row.CreatedAt,
					row.UpdatedAt,
					tagDtoFromQueryModel(row.Tags)...,
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
			articles = append(
				articles, dto.NewArticle(
					row.ID,
					row.Title,
					row.Body,
					row.Thumbnail,
					row.CreatedAt,
					row.UpdatedAt,
					tagDtoFromQueryModel(row.Tags)...,
				),
			)
		}
	default:
		return nil, fmt.Errorf("invalid arguments")
	}
	result := dto.NewListAfterOutput(hasNext, articles...)
	return &result, nil
}

// NewListAfter constructs ListAfterPage.
func NewListAfter(queries query.Queries) *ListAfter {
	return &ListAfter{queries: queries}
}
