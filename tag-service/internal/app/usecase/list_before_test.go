package usecase

import (
	"database/sql"
	"testing"

	"blogapi.miyamo.today/tag-service/internal/app/usecase/dto"
	"blogapi.miyamo.today/tag-service/internal/app/usecase/query"
	"blogapi.miyamo.today/tag-service/internal/infra/rdb/sqlc"
	"blogapi.miyamo.today/tag-service/internal/infra/rdb/types"
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	. "github.com/ovechkin-dm/mockio/v2/mock"
	"github.com/stretchr/testify/suite"
)

type ListBeforeTestSuite struct {
	suite.Suite
}

func TestListBeforeTestSuite(t *testing.T) {
	suite.Run(t, new(ListBeforeTestSuite))
}

func (s *ListBeforeTestSuite) TestListBefore_Execute() {
	s.Run(
		"happy_path/without-cursor/single/has-previous", func() {
			ctrl := NewMockController(s.T())
			queries := Mock[query.Queries](ctrl)
			WhenDouble(queries.ListBeforeWithLimit(AnyContext(), Exact(int32(2)))).
				ThenReturn(
					[]sqlc.ListBeforeWithLimitRow{
						{
							ID:        "2",
							Name:      "tag2",
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
						},
						{
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
						},
					}, nil,
				)

			u := NewListBefore(queries)

			out, err := u.Execute(s.T().Context(), dto.NewListBeforeInput(1))
			s.Require().NoError(err)
			s.Require().NotNil(out)
			s.Require().Equal(
				dto.NewListBeforeOutput(
					true,
					dto.NewTag(
						"2",
						"tag2",
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
					),
				), *out,
			)
		},
	)
	s.Run(
		"happy_path/without-cursor/single/end-of-page", func() {
			ctrl := NewMockController(s.T())
			queries := Mock[query.Queries](ctrl)
			WhenDouble(queries.ListBeforeWithLimit(AnyContext(), Exact(int32(2)))).
				ThenReturn(
					[]sqlc.ListBeforeWithLimitRow{
						{
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
						},
					}, nil,
				)

			u := NewListBefore(queries)

			out, err := u.Execute(s.T().Context(), dto.NewListBeforeInput(1))
			s.Require().NoError(err)
			s.Require().NotNil(out)
			s.Require().Equal(
				dto.NewListBeforeOutput(
					false,
					dto.NewTag(
						"1",
						"tag1",
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
					),
				), *out,
			)
		},
	)
	s.Run(
		"happy_path/without-cursor/multiple/has-previous", func() {
			ctrl := NewMockController(s.T())
			queries := Mock[query.Queries](ctrl)
			WhenDouble(queries.ListBeforeWithLimit(AnyContext(), Exact(int32(3)))).
				ThenReturn(
					[]sqlc.ListBeforeWithLimitRow{
						{
							ID:        "3",
							Name:      "tag3",
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
						},
						{
							ID:        "2",
							Name:      "tag2",
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
						},
						{
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
						},
					}, nil,
				)

			u := NewListBefore(queries)

			out, err := u.Execute(s.T().Context(), dto.NewListBeforeInput(2))
			s.Require().NoError(err)
			s.Require().NotNil(out)
			s.Require().Equal(
				dto.NewListBeforeOutput(
					true,
					dto.NewTag(
						"3",
						"tag3",
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
					),
					dto.NewTag(
						"2",
						"tag2",
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
					),
				), *out,
			)
		},
	)
	s.Run(
		"happy_path/without-cursor/multiple/end-of-page", func() {
			ctrl := NewMockController(s.T())
			queries := Mock[query.Queries](ctrl)
			WhenDouble(queries.ListBeforeWithLimit(AnyContext(), Exact(int32(3)))).
				ThenReturn(
					[]sqlc.ListBeforeWithLimitRow{
						{
							ID:        "2",
							Name:      "tag2",
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
						},
						{
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
						},
					}, nil,
				)

			u := NewListBefore(queries)

			out, err := u.Execute(s.T().Context(), dto.NewListBeforeInput(2))
			s.Require().NoError(err)
			s.Require().NotNil(out)
			s.Require().Equal(
				dto.NewListBeforeOutput(
					false,
					dto.NewTag(
						"2",
						"tag2",
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
					),
					dto.NewTag(
						"1",
						"tag1",
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
					),
				), *out,
			)
		},
	)
	s.Run(
		"unhappy_path/without-cursor/query-returns-error", func() {
			ctrl := NewMockController(s.T())
			queries := Mock[query.Queries](ctrl)
			WhenDouble(queries.ListBeforeWithLimit(AnyContext(), Exact(int32(2)))).
				ThenReturn(nil, sql.ErrConnDone)

			u := NewListBefore(queries)

			out, err := u.Execute(s.T().Context(), dto.NewListBeforeInput(1))
			s.Require().Error(err)
			s.Require().Nil(out)
			s.Require().ErrorIs(err, sql.ErrConnDone)
		},
	)

	s.Run(
		"happy_path/with-cursor/single/has-previous", func() {
			ctrl := NewMockController(s.T())
			queries := Mock[query.Queries](ctrl)
			WhenDouble(
				queries.ListBeforeWithLimitAndCursor(
					AnyContext(),
					Equal(sqlc.ListBeforeWithLimitAndCursorParams{ID: "0", Limit: 2}),
				),
			).
				ThenReturn(
					[]sqlc.ListBeforeWithLimitAndCursorRow{
						{
							ID:        "2",
							Name:      "tag2",
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
						},
						{
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
						},
					}, nil,
				)

			u := NewListBefore(queries)

			out, err := u.Execute(s.T().Context(), dto.NewListBeforeInput(1, dto.ListBeforeInputWithCursor("0")))
			s.Require().NoError(err)
			s.Require().NotNil(out)
			s.Require().Equal(
				dto.NewListBeforeOutput(
					true,
					dto.NewTag(
						"2",
						"tag2",
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
					),
				), *out,
			)
		},
	)
	s.Run(
		"happy_path/with-cursor/single/end-of-page", func() {
			ctrl := NewMockController(s.T())
			queries := Mock[query.Queries](ctrl)
			WhenDouble(
				queries.ListBeforeWithLimitAndCursor(
					AnyContext(),
					Equal(sqlc.ListBeforeWithLimitAndCursorParams{ID: "0", Limit: 2}),
				),
			).
				ThenReturn(
					[]sqlc.ListBeforeWithLimitAndCursorRow{
						{
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
						},
					}, nil,
				)

			u := NewListBefore(queries)

			out, err := u.Execute(s.T().Context(), dto.NewListBeforeInput(1, dto.ListBeforeInputWithCursor("0")))
			s.Require().NoError(err)
			s.Require().NotNil(out)
			s.Require().Equal(
				dto.NewListBeforeOutput(
					false,
					dto.NewTag(
						"1",
						"tag1",
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
					),
				), *out,
			)
		},
	)
	s.Run(
		"happy_path/with-cursor/multiple/has-previous", func() {
			ctrl := NewMockController(s.T())
			queries := Mock[query.Queries](ctrl)
			WhenDouble(
				queries.ListBeforeWithLimitAndCursor(
					AnyContext(),
					Equal(sqlc.ListBeforeWithLimitAndCursorParams{ID: "0", Limit: 3}),
				),
			).
				ThenReturn(
					[]sqlc.ListBeforeWithLimitAndCursorRow{
						{
							ID:        "3",
							Name:      "tag3",
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
						},
						{
							ID:        "2",
							Name:      "tag2",
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
						},
						{
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
						},
					}, nil,
				)

			u := NewListBefore(queries)

			out, err := u.Execute(s.T().Context(), dto.NewListBeforeInput(2, dto.ListBeforeInputWithCursor("0")))
			s.Require().NoError(err)
			s.Require().NotNil(out)
			s.Require().Equal(
				dto.NewListBeforeOutput(
					true,
					dto.NewTag(
						"3",
						"tag3",
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
					),
					dto.NewTag(
						"2",
						"tag2",
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
					),
				), *out,
			)
		},
	)
	s.Run(
		"happy_path/with-cursor/multiple/end-of-page", func() {
			ctrl := NewMockController(s.T())
			queries := Mock[query.Queries](ctrl)
			WhenDouble(
				queries.ListBeforeWithLimitAndCursor(
					AnyContext(),
					Equal(sqlc.ListBeforeWithLimitAndCursorParams{ID: "0", Limit: 3}),
				),
			).
				ThenReturn(
					[]sqlc.ListBeforeWithLimitAndCursorRow{
						{
							ID:        "2",
							Name:      "tag2",
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
						},
						{
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
						},
					}, nil,
				)

			u := NewListBefore(queries)

			out, err := u.Execute(s.T().Context(), dto.NewListBeforeInput(2, dto.ListBeforeInputWithCursor("0")))
			s.Require().NoError(err)
			s.Require().NotNil(out)
			s.Require().Equal(
				dto.NewListBeforeOutput(
					false,
					dto.NewTag(
						"2",
						"tag2",
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
					),
					dto.NewTag(
						"1",
						"tag1",
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
					),
				), *out,
			)
		},
	)
	s.Run(
		"unhappy_path/with-cursor/query-returns-error", func() {
			ctrl := NewMockController(s.T())
			queries := Mock[query.Queries](ctrl)
			WhenDouble(
				queries.ListBeforeWithLimitAndCursor(
					AnyContext(),
					Equal(sqlc.ListBeforeWithLimitAndCursorParams{ID: "0", Limit: 2}),
				),
			).
				ThenReturn(nil, sql.ErrConnDone)

			u := NewListBefore(queries)

			out, err := u.Execute(s.T().Context(), dto.NewListBeforeInput(1, dto.ListBeforeInputWithCursor("0")))
			s.Require().Error(err)
			s.Require().Nil(out)
			s.Require().ErrorIs(err, sql.ErrConnDone)
		},
	)
}
