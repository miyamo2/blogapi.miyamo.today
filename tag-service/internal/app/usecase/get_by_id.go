package usecase

import (
	"context"

	"github.com/newrelic/go-agent/v3/newrelic"

	"blogapi.miyamo.today/tag-service/internal/app/usecase/dto"
	"blogapi.miyamo.today/tag-service/internal/app/usecase/query"
	"github.com/cockroachdb/errors"
)

// GetById is an implementation of usecase.GetById
type GetById struct {
	queries query.Queries
}

func (u *GetById) Execute(ctx context.Context, in dto.GetByIdInput) (*dto.GetByIdOutput, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("Execute").End()

	tag, err := u.queries.GetByID(ctx, in.Id())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	result := dto.NewTag(
		tag.ID,
		tag.Name,
		articleDtoFromQueryModel(tag.Articles)...,
	)
	return &result, nil
}

// NewGetById is constructor of GetById
func NewGetById(queries query.Queries) *GetById {
	return &GetById{queries: queries}
}
