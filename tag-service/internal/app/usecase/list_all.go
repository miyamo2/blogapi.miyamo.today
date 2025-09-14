package usecase

import (
	"context"

	"github.com/newrelic/go-agent/v3/newrelic"

	"blogapi.miyamo.today/tag-service/internal/app/usecase/dto"
	"blogapi.miyamo.today/tag-service/internal/app/usecase/query"
)

// ListAll implements usecase.ListAll
type ListAll struct {
	queries query.Queries
}

func (u *ListAll) Execute(ctx context.Context) (*dto.ListAllOutput, error) {
	nrtx := newrelic.FromContext(ctx)
	seg := nrtx.StartSegment("Execute")
	defer seg.End()

	rows, err := u.queries.ListAfter(ctx)
	if err != nil {
		return nil, err
	}

	tags := make([]dto.Tag, 0, len(rows))
	for _, row := range rows {
		tags = append(
			tags, dto.NewTag(
				row.ID,
				row.Name,
				articleDtoFromQueryModel(row.Articles)...,
			),
		)
	}
	result := dto.NewListAllOutput(tags...)
	return &result, nil
}

// NewListAll constructs usecase.ListAll
func NewListAll(queries query.Queries) *ListAll {
	return &ListAll{queries: queries}
}
