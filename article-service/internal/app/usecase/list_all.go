package usecase

import (
	"context"

	"blogapi.miyamo.today/article-service/internal/app/usecase/query"
	"github.com/newrelic/go-agent/v3/newrelic"

	"blogapi.miyamo.today/article-service/internal/app/usecase/dto"
)

// ListAll implements usecase.ListAll
type ListAll struct {
	queries query.Queries
}

func (u *ListAll) Execute(ctx context.Context) (*dto.ListAllOutput, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("Execute").End()

	rows, err := u.queries.ListAfter(ctx)
	if err != nil {
		return nil, err
	}
	articles := make([]dto.Article, 0, len(rows))
	for _, row := range rows {
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
	result := dto.NewListAllOutput(articles...)
	return &result, nil
}

// NewListAll constructs ListAll.
func NewListAll(queries query.Queries) *ListAll {
	return &ListAll{queries: queries}
}
