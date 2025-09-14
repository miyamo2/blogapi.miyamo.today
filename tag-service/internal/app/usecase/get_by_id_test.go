package usecase

import (
	"database/sql"
	"testing"

	"blogapi.miyamo.today/tag-service/internal/infra/rdb/sqlc"
	"blogapi.miyamo.today/tag-service/internal/infra/rdb/types"
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	"github.com/stretchr/testify/suite"

	"blogapi.miyamo.today/tag-service/internal/app/usecase/dto"
	"blogapi.miyamo.today/tag-service/internal/app/usecase/query"
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
						Name:      "tag1",
						CreatedAt: synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						UpdatedAt: synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						Articles: []types.Article{
							{
								ID:        "1",
								Title:     "happy_path",
								Thumbnail: "thumbnail",
								CreatedAt: synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								UpdatedAt: synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							},
							{
								ID:        "2",
								Title:     "happy_path2",
								Thumbnail: "thumbnail",
								CreatedAt: synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								UpdatedAt: synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							},
						},
					}, nil,
				)

			u := NewGetById(queries)
			out, err := u.Execute(s.T().Context(), dto.NewGetByIdInput("1"))
			s.Require().NoError(err)
			s.Require().NotNil(out)
			s.Require().Equal(
				dto.NewTag(
					"1",
					"tag1",
					[]dto.Article{
						dto.NewArticle(
							"1",
							"happy_path",
							"thumbnail",
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						),
						dto.NewArticle(
							"2",
							"happy_path2",
							"thumbnail",
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						),
					}...,
				), *out,
			)
		},
	)
	s.Run(
		"happy_path/tag_has_no_articles", func() {
			ctrl := NewMockController(s.T())
			queries := Mock[query.Queries](ctrl)
			WhenDouble(queries.GetByID(AnyContext(), Exact("1"))).
				ThenReturn(
					sqlc.GetByIDRow{
						ID:        "1",
						Name:      "tag1",
						CreatedAt: synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						UpdatedAt: synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					}, nil,
				)

			u := NewGetById(queries)
			out, err := u.Execute(s.T().Context(), dto.NewGetByIdInput("1"))
			s.Require().NoError(err)
			s.Require().NotNil(out)
			s.Require().Equal(
				dto.NewTag(
					"1",
					"tag1",
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

			u := NewGetById(queries)
			out, err := u.Execute(s.T().Context(), dto.NewGetByIdInput("1"))
			s.Require().Error(err)
			s.Require().Nil(out)
			s.Require().ErrorIs(err, sql.ErrNoRows)
		},
	)
}
