package usecase

import (
	"database/sql"
	"testing"

	"blogapi.miyamo.today/article-service/internal/app/usecase/dto"
	"blogapi.miyamo.today/article-service/internal/app/usecase/query"
	"blogapi.miyamo.today/article-service/internal/infra/rdb/sqlc"
	"blogapi.miyamo.today/article-service/internal/infra/rdb/types"
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	. "github.com/ovechkin-dm/mockio/v2/mock"
	"github.com/stretchr/testify/suite"
)

type ListAfterTestSuite struct {
	suite.Suite
}

func TestListAfterTestSuite(t *testing.T) {
	suite.Run(t, new(ListAfterTestSuite))
}

func (s *ListAfterTestSuite) TestListAfter_Execute() {
	s.Run(
		"happy_path/without-cursor/single/has-next", func() {
			ctrl := NewMockController(s.T())
			queries := Mock[query.Queries](ctrl)
			WhenDouble(queries.ListAfterWithLimit(AnyContext(), Exact(int32(2)))).
				ThenReturn(
					[]sqlc.ListAfterWithLimitRow{
						{
							ID:        "1",
							Title:     "happy_path1",
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
						},
						{
							ID:        "2",
							Title:     "happy_path2",
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
						},
					}, nil,
				)

			u := NewListAfter(queries)

			out, err := u.Execute(s.T().Context(), dto.NewListAfterInput(1))
			s.Require().NoError(err)
			s.Require().NotNil(out)
			s.Require().Equal(
				dto.NewListAfterOutput(
					true,
					dto.NewArticle(
						"1",
						"happy_path1",
						"## happy_path",
						"thumbnail",
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						dto.NewTag("1", "tag1"),
						dto.NewTag("2", "tag2"),
					),
				), *out,
			)
		},
	)
	s.Run(
		"happy_path/without-cursor/single/end-of-page", func() {
			ctrl := NewMockController(s.T())
			queries := Mock[query.Queries](ctrl)
			WhenDouble(queries.ListAfterWithLimit(AnyContext(), Exact(int32(2)))).
				ThenReturn(
					[]sqlc.ListAfterWithLimitRow{
						{
							ID:        "1",
							Title:     "happy_path1",
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
						},
					}, nil,
				)

			u := NewListAfter(queries)

			out, err := u.Execute(s.T().Context(), dto.NewListAfterInput(1))
			s.Require().NoError(err)
			s.Require().NotNil(out)
			s.Require().Equal(
				dto.NewListAfterOutput(
					false,
					dto.NewArticle(
						"1",
						"happy_path1",
						"## happy_path",
						"thumbnail",
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						dto.NewTag("1", "tag1"),
						dto.NewTag("2", "tag2"),
					),
				), *out,
			)
		},
	)
	s.Run(
		"happy_path/without-cursor/multiple/has-next", func() {
			ctrl := NewMockController(s.T())
			queries := Mock[query.Queries](ctrl)
			WhenDouble(queries.ListAfterWithLimit(AnyContext(), Exact(int32(3)))).
				ThenReturn(
					[]sqlc.ListAfterWithLimitRow{
						{
							ID:        "1",
							Title:     "happy_path1",
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
						},
						{
							ID:        "2",
							Title:     "happy_path2",
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
						},
						{
							ID:        "3",
							Title:     "happy_path3",
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
						},
					}, nil,
				)

			u := NewListAfter(queries)

			out, err := u.Execute(s.T().Context(), dto.NewListAfterInput(2))
			s.Require().NoError(err)
			s.Require().NotNil(out)
			s.Require().Equal(
				dto.NewListAfterOutput(
					true,
					dto.NewArticle(
						"1",
						"happy_path1",
						"## happy_path",
						"thumbnail",
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						dto.NewTag("1", "tag1"),
						dto.NewTag("2", "tag2"),
					),
					dto.NewArticle(
						"2",
						"happy_path2",
						"## happy_path",
						"thumbnail",
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						dto.NewTag("1", "tag1"),
						dto.NewTag("2", "tag2"),
					),
				), *out,
			)
		},
	)
	s.Run(
		"happy_path/without-cursor/multiple/end-of-page", func() {
			ctrl := NewMockController(s.T())
			queries := Mock[query.Queries](ctrl)
			WhenDouble(queries.ListAfterWithLimit(AnyContext(), Exact(int32(3)))).
				ThenReturn(
					[]sqlc.ListAfterWithLimitRow{
						{
							ID:        "1",
							Title:     "happy_path1",
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
						},
						{
							ID:        "2",
							Title:     "happy_path2",
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
						},
					}, nil,
				)

			u := NewListAfter(queries)

			out, err := u.Execute(s.T().Context(), dto.NewListAfterInput(2))
			s.Require().NoError(err)
			s.Require().NotNil(out)
			s.Require().Equal(
				dto.NewListAfterOutput(
					false,
					dto.NewArticle(
						"1",
						"happy_path1",
						"## happy_path",
						"thumbnail",
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						dto.NewTag("1", "tag1"),
						dto.NewTag("2", "tag2"),
					),
					dto.NewArticle(
						"2",
						"happy_path2",
						"## happy_path",
						"thumbnail",
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						dto.NewTag("1", "tag1"),
						dto.NewTag("2", "tag2"),
					),
				), *out,
			)
		},
	)
	s.Run(
		"unhappy_path/without-cursor/query-returns-error", func() {
			ctrl := NewMockController(s.T())
			queries := Mock[query.Queries](ctrl)
			WhenDouble(queries.ListAfterWithLimit(AnyContext(), Exact(int32(2)))).
				ThenReturn(nil, sql.ErrConnDone)

			u := NewListAfter(queries)

			out, err := u.Execute(s.T().Context(), dto.NewListAfterInput(1))
			s.Require().Error(err)
			s.Require().Nil(out)
			s.Require().ErrorIs(err, sql.ErrConnDone)
		},
	)

	s.Run(
		"happy_path/with-cursor/single/has-next", func() {
			ctrl := NewMockController(s.T())
			queries := Mock[query.Queries](ctrl)
			WhenDouble(
				queries.ListAfterWithLimitAndCursor(
					AnyContext(),
					Equal(sqlc.ListAfterWithLimitAndCursorParams{ID: "0", Limit: 2}),
				),
			).
				ThenReturn(
					[]sqlc.ListAfterWithLimitAndCursorRow{
						{
							ID:        "1",
							Title:     "happy_path1",
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
						},
						{
							ID:        "2",
							Title:     "happy_path2",
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
						},
					}, nil,
				)

			u := NewListAfter(queries)

			out, err := u.Execute(s.T().Context(), dto.NewListAfterInput(1, dto.ListAfterInputWithCursor("0")))
			s.Require().NoError(err)
			s.Require().NotNil(out)
			s.Require().Equal(
				dto.NewListAfterOutput(
					true,
					dto.NewArticle(
						"1",
						"happy_path1",
						"## happy_path",
						"thumbnail",
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						dto.NewTag("1", "tag1"),
						dto.NewTag("2", "tag2"),
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
				queries.ListAfterWithLimitAndCursor(
					AnyContext(),
					Equal(sqlc.ListAfterWithLimitAndCursorParams{ID: "0", Limit: 2}),
				),
			).
				ThenReturn(
					[]sqlc.ListAfterWithLimitAndCursorRow{
						{
							ID:        "1",
							Title:     "happy_path1",
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
						},
					}, nil,
				)

			u := NewListAfter(queries)

			out, err := u.Execute(s.T().Context(), dto.NewListAfterInput(1, dto.ListAfterInputWithCursor("0")))
			s.Require().NoError(err)
			s.Require().NotNil(out)
			s.Require().Equal(
				dto.NewListAfterOutput(
					false,
					dto.NewArticle(
						"1",
						"happy_path1",
						"## happy_path",
						"thumbnail",
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						dto.NewTag("1", "tag1"),
						dto.NewTag("2", "tag2"),
					),
				), *out,
			)
		},
	)
	s.Run(
		"happy_path/with-cursor/multiple/has-next", func() {
			ctrl := NewMockController(s.T())
			queries := Mock[query.Queries](ctrl)
			WhenDouble(
				queries.ListAfterWithLimitAndCursor(
					AnyContext(),
					Equal(sqlc.ListAfterWithLimitAndCursorParams{ID: "0", Limit: 3}),
				),
			).
				ThenReturn(
					[]sqlc.ListAfterWithLimitAndCursorRow{
						{
							ID:        "1",
							Title:     "happy_path1",
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
						},
						{
							ID:        "2",
							Title:     "happy_path2",
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
						},
						{
							ID:        "3",
							Title:     "happy_path3",
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
						},
					}, nil,
				)

			u := NewListAfter(queries)

			out, err := u.Execute(s.T().Context(), dto.NewListAfterInput(2, dto.ListAfterInputWithCursor("0")))
			s.Require().NoError(err)
			s.Require().NotNil(out)
			s.Require().Equal(
				dto.NewListAfterOutput(
					true,
					dto.NewArticle(
						"1",
						"happy_path1",
						"## happy_path",
						"thumbnail",
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						dto.NewTag("1", "tag1"),
						dto.NewTag("2", "tag2"),
					),
					dto.NewArticle(
						"2",
						"happy_path2",
						"## happy_path",
						"thumbnail",
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						dto.NewTag("1", "tag1"),
						dto.NewTag("2", "tag2"),
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
				queries.ListAfterWithLimitAndCursor(
					AnyContext(),
					Equal(sqlc.ListAfterWithLimitAndCursorParams{ID: "0", Limit: 3}),
				),
			).
				ThenReturn(
					[]sqlc.ListAfterWithLimitAndCursorRow{
						{
							ID:        "1",
							Title:     "happy_path1",
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
						},
						{
							ID:        "2",
							Title:     "happy_path2",
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
						},
					}, nil,
				)

			u := NewListAfter(queries)

			out, err := u.Execute(s.T().Context(), dto.NewListAfterInput(2, dto.ListAfterInputWithCursor("0")))
			s.Require().NoError(err)
			s.Require().NotNil(out)
			s.Require().Equal(
				dto.NewListAfterOutput(
					false,
					dto.NewArticle(
						"1",
						"happy_path1",
						"## happy_path",
						"thumbnail",
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						dto.NewTag("1", "tag1"),
						dto.NewTag("2", "tag2"),
					),
					dto.NewArticle(
						"2",
						"happy_path2",
						"## happy_path",
						"thumbnail",
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						dto.NewTag("1", "tag1"),
						dto.NewTag("2", "tag2"),
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
				queries.ListAfterWithLimitAndCursor(
					AnyContext(),
					Equal(sqlc.ListAfterWithLimitAndCursorParams{ID: "0", Limit: 2}),
				),
			).
				ThenReturn(nil, sql.ErrConnDone)

			u := NewListAfter(queries)

			out, err := u.Execute(s.T().Context(), dto.NewListAfterInput(1, dto.ListAfterInputWithCursor("0")))
			s.Require().Error(err)
			s.Require().Nil(out)
			s.Require().ErrorIs(err, sql.ErrConnDone)
		},
	)
}
