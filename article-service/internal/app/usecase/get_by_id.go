package usecase

import (
	"context"
	"database/sql"

	"github.com/newrelic/go-agent/v3/newrelic"

	"blogapi.miyamo.today/article-service/internal/app/usecase/dto"
	"blogapi.miyamo.today/article-service/internal/app/usecase/query"
	"github.com/cockroachdb/errors"
)

// GetByID implements usecase.GetByID
type GetByID struct {
	queries query.Queries
}

func (u *GetByID) Execute(ctx context.Context, in dto.GetByIDInput) (*dto.GetByIDOutput, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("Execute").End()

	row, err := u.queries.GetByID(ctx, in.ID())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.WithMessage(err, "article not found")
		}
		return nil, errors.WithStack(err)
	}

	result := dto.NewGetByIDOutput(
		row.ID,
		row.Title,
		row.Body,
		row.Thumbnail,
		row.CreatedAt,
		row.UpdatedAt,
		tagDtoFromQueryModel(row.Tags)...,
	)
	return &result, nil
}

// NewGetByID constructs GetByID
func NewGetByID(queries query.Queries) *GetByID {
	return &GetByID{queries: queries}
}
