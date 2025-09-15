package usecase

import (
	"context"

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

	last := min(max(in.Last(), 1), 100) // TODO: config

	var (
		articles    = make([]dto.Article, 0, last)
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
		rows, err := u.queries.ListBeforeWithLimit(ctx, int32(last+1))
		if err != nil {
			return nil, err
		}
		hasPrevious = len(rows) > last
		for row := range getPage(rows, last, 1) {
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
	}

	result := dto.NewListBeforeOutput(hasPrevious, articles...)
	return &result, nil
}

// NewListBefore constructs ListBeforePage.
func NewListBefore(queries query.Queries) *ListBefore {
	return &ListBefore{queries: queries}
}
