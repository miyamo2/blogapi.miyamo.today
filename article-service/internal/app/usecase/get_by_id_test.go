package usecase

import (
	"database/sql"
	"testing"

	"blogapi.miyamo.today/article-service/internal/app/usecase/query"
	"blogapi.miyamo.today/article-service/internal/infra/rdb/sqlc"
	"blogapi.miyamo.today/article-service/internal/infra/rdb/types"
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	"github.com/stretchr/testify/suite"

	"blogapi.miyamo.today/article-service/internal/app/usecase/dto"
	. "github.com/ovechkin-dm/mockio/v2/mock"
)

type GetByIdTestSuite struct {
	suite.Suite
}

func TestGetByIdTestSuite(t *testing.T) {
	suite.Run(t, new(GetByIdTestSuite))
}

func (s *GetByIdTestSuite) TestGetById_Execute() {
	s.Run(
		"happy_path", func() {
			ctrl := NewMockController(s.T())
			queries := Mock[query.Queries](ctrl)
			WhenDouble(queries.GetByID(AnyContext(), Exact("1"))).
				ThenReturn(
					sqlc.GetByIDRow{
						ID:        "1",
						Title:     "happy_path",
						Body:      "## happy_path",
						Thumbnail: "thumbnail",
						CreatedAt: synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						UpdatedAt: synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						Tags: []types.Tag{
							{
								ID:   "1",
								Name: "tag1",
							},
							{
								ID:   "2",
								Name: "tag2",
							},
						},
					}, nil,
				)

			u := NewGetByID(queries)
			out, err := u.Execute(s.T().Context(), dto.NewGetByIDInput("1"))
			s.Require().NoError(err)
			s.Require().NotNil(out)
			s.Require().Equal(
				dto.NewGetByIDOutput(
					"1",
					"happy_path",
					"## happy_path",
					"thumbnail",
					synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					dto.NewTag("1", "tag1"),
					dto.NewTag("2", "tag2"),
				), *out,
			)
		},
	)
	s.Run(
		"happy_path/article_has_no_tags", func() {
			ctrl := NewMockController(s.T())
			queries := Mock[query.Queries](ctrl)
			WhenDouble(queries.GetByID(AnyContext(), Exact("1"))).
				ThenReturn(
					sqlc.GetByIDRow{
						ID:        "1",
						Title:     "happy_path",
						Body:      "## happy_path",
						Thumbnail: "thumbnail",
						CreatedAt: synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						UpdatedAt: synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					}, nil,
				)

			u := NewGetByID(queries)
			out, err := u.Execute(s.T().Context(), dto.NewGetByIDInput("1"))
			s.Require().NoError(err)
			s.Require().NotNil(out)
			s.Require().Equal(
				dto.NewGetByIDOutput(
					"1",
					"happy_path",
					"## happy_path",
					"thumbnail",
					synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
				), *out,
			)
		},
	)
	s.Run(
		"unhappy_path/query_returns_error", func() {
			ctrl := NewMockController(s.T())
			queries := Mock[query.Queries](ctrl)
			WhenDouble(queries.GetByID(AnyContext(), Exact("1"))).
				ThenReturn(sqlc.GetByIDRow{}, sql.ErrNoRows)

			u := NewGetByID(queries)
			out, err := u.Execute(s.T().Context(), dto.NewGetByIDInput("1"))
			s.Require().Error(err)
			s.Require().Nil(out)
			s.Require().ErrorIs(err, sql.ErrNoRows)
		},
	)
}
