package usecase

import (
	"context"
	"fmt"

	"github.com/newrelic/go-agent/v3/newrelic"

	"blogapi.miyamo.today/article-service/internal/app/usecase/dto"
	"blogapi.miyamo.today/article-service/internal/app/usecase/query"
)

// ListBefore implements usecase.ListBefore
type ListBefore struct {
	queries query.Queries
}

func (u *ListBefore) Execute(ctx context.Context, in dto.ListBeforeInput) (*dto.ListBeforeOutput, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("Execute").End()

	var (
		articles    = make([]dto.Article, 0, in.Last())
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
	case in.Last() != 0:
		rows, err := u.queries.ListBeforeWithLimit(ctx, int32(in.Last()+1))
		if err != nil {
			return nil, err
		}
		hasPrevious = len(rows) > in.Last()
		for row := range getPage(rows, in.Last(), 1) {
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

	result := dto.NewListBeforeOutput(hasPrevious, articles...)
	return &result, nil
}

// NewListBefore constructs ListBeforePage.
func NewListBefore(queries query.Queries) *ListBefore {
	return &ListBefore{queries: queries}
}
