package query

import (
	"context"
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/miyamo2/blogapi.miyamo.today/tag-service/internal/infra/rdb/query/internal/entity"
	"regexp"
	"testing"

	"github.com/miyamo2/blogapi.miyamo.today/tag-service/internal/infra/rdb/query/model"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cockroachdb/errors"
	"github.com/miyamo2/blogapi.miyamo.today/core/db"
	gwrapper "github.com/miyamo2/blogapi.miyamo.today/core/db/gorm"
	"gorm.io/driver/postgres"
)

func TestArticleService_GetById(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
		out *db.SingleStatementResult[model.Tag]
	}
	type testCase struct {
		args        args
		execOpt     func() []db.ExecuteOption
		want        error
		wantErr     bool
		expectedOut *db.SingleStatementResult[model.Tag]
		exists      bool
	}
	tagTable := []string{
		"id",
		"name",
		"created_at",
		"updated_at",
		"articles",
	}
	tests := map[string]testCase{
		"happy_path": {
			args: args{
				ctx: context.Background(),
				id:  "tag1",
				out: &db.SingleStatementResult[model.Tag]{},
			},
			exists: true,
			execOpt: func() []db.ExecuteOption {
				sqlDB, mock, err := sqlmock.New()
				if err != nil {
					panic(err)
				}
				rows := sqlmock.NewRows(tagTable).
					AddRow(
						"tag1",
						"test",
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						entity.Articles([]entity.Article{
							{
								ID:        "1",
								Title:     "happy_path",
								Thumbnail: "01234567890",
								CreatedAt: synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
								UpdatedAt: synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
							},
						}))
				mq := mock.ExpectPrepare(regexp.QuoteMeta(
					`SELECT "t".*,  COALESCE(jsonb_agg(json_build_object('id', a.id, 'title', a.title, 'thumbnail', a.thumbnail, 'created_at', a.created_at, 'updated_at', a.updated_at)) FILTER (WHERE a.id IS NOT NULL), '[]'::json) AS "articles" FROM (SELECT * FROM "tags" WHERE "id" = $1) AS "t" LEFT OUTER JOIN "articles" AS "a" ON "t"."id" = "a"."tag_id" GROUP BY "t"."id"`))
				mq.ExpectQuery().
					WithArgs("tag1").WillReturnRows(rows)
				dialector := postgres.New(postgres.Config{
					Conn: sqlDB,
				})
				gwrapper.Invalidate()
				gwrapper.InvalidateDialector()
				gwrapper.InitializeDialector(&dialector)
				tx, err := gwrapper.Get(context.Background())
				if err != nil {
					panic(nil)
				}
				return []db.ExecuteOption{gwrapper.WithTransaction(tx)}
			},
			want: nil,
			expectedOut: func() *db.SingleStatementResult[model.Tag] {
				tag := model.NewTag(
					"tag1",
					"test",
					model.NewArticle(
						"1",
						"happy_path",
						"01234567890",
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
					))
				out := db.NewSingleStatementResult[model.Tag]()
				out.Set(tag)
				return out
			}(),
		},
		"happy_path/tag_has_no_article": {
			args: args{
				ctx: context.Background(),
				id:  "tag1",
				out: &db.SingleStatementResult[model.Tag]{},
			},
			exists: true,
			execOpt: func() []db.ExecuteOption {
				sqlDB, mock, err := sqlmock.New()
				if err != nil {
					panic(err)
				}
				rows := sqlmock.NewRows(tagTable).
					AddRow(
						"tag1",
						"test",
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						entity.Articles([]entity.Article{}))
				mq := mock.ExpectPrepare(regexp.QuoteMeta(
					`SELECT "t".*,  COALESCE(jsonb_agg(json_build_object('id', a.id, 'title', a.title, 'thumbnail', a.thumbnail, 'created_at', a.created_at, 'updated_at', a.updated_at)) FILTER (WHERE a.id IS NOT NULL), '[]'::json) AS "articles" FROM (SELECT * FROM "tags" WHERE "id" = $1) AS "t" LEFT OUTER JOIN "articles" AS "a" ON "t"."id" = "a"."tag_id" GROUP BY "t"."id"`))
				mq.ExpectQuery().
					WithArgs("tag1").WillReturnRows(rows)
				dialector := postgres.New(postgres.Config{
					Conn: sqlDB,
				})
				gwrapper.Invalidate()
				gwrapper.InvalidateDialector()
				gwrapper.InitializeDialector(&dialector)
				tx, err := gwrapper.Get(context.Background())
				if err != nil {
					panic(nil)
				}
				return []db.ExecuteOption{gwrapper.WithTransaction(tx)}
			},
			want: nil,
			expectedOut: func() *db.SingleStatementResult[model.Tag] {
				tag := model.NewTag(
					"tag1",
					"test",
				)
				out := db.NewSingleStatementResult[model.Tag]()
				out.Set(tag)
				return out
			}(),
		},
		"unhappy_path/not_found": {
			args: args{
				ctx: context.Background(),
				id:  "tag1",
				out: &db.SingleStatementResult[model.Tag]{},
			},
			execOpt: func() []db.ExecuteOption {
				sqlDB, mock, err := sqlmock.New()
				if err != nil {
					panic(err)
				}
				rows := sqlmock.NewRows(tagTable)
				mq := mock.ExpectPrepare(regexp.QuoteMeta(
					`SELECT "t".*,  COALESCE(jsonb_agg(json_build_object('id', a.id, 'title', a.title, 'thumbnail', a.thumbnail, 'created_at', a.created_at, 'updated_at', a.updated_at)) FILTER (WHERE a.id IS NOT NULL), '[]'::json) AS "articles" FROM (SELECT * FROM "tags" WHERE "id" = $1) AS "t" LEFT OUTER JOIN "articles" AS "a" ON "t"."id" = "a"."tag_id" GROUP BY "t"."id"`))
				mq.ExpectQuery().
					WithArgs("tag1").WillReturnRows(rows)
				dialector := postgres.New(postgres.Config{
					Conn: sqlDB,
				})
				gwrapper.Invalidate()
				gwrapper.InvalidateDialector()
				gwrapper.InitializeDialector(&dialector)
				tx, err := gwrapper.Get(context.Background())
				if err != nil {
					panic(nil)
				}
				return []db.ExecuteOption{gwrapper.WithTransaction(tx)}
			},
			want:    ErrNotFound,
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			s := NewTagService()
			err := s.GetById(tt.args.ctx, tt.args.id, tt.args.out).Execute(tt.args.ctx, tt.execOpt()...)
			if !errors.Is(err, tt.want) {
				t.Errorf("want error %+v but got %+v", tt.want, err)
				return
			}
			if tt.exists {
				if diff := cmp.Diff(tt.expectedOut.StrictGet(), tt.args.out.StrictGet(), cmp.AllowUnexported(model.Article{}), cmp.AllowUnexported(model.Tag{}), cmpopts.EquateEmpty()); diff != "" {
					t.Errorf(`unexpected output (-want +got): %s`, diff)
					return
				}
			}
		})
	}
}

func TestTagService_GetAll(t *testing.T) {
	type args struct {
		ctx              context.Context
		out              *db.MultipleStatementResult[model.Tag]
		paginationOption []db.PaginationOption
	}
	type testCase struct {
		args        args
		execOpt     func() []db.ExecuteOption
		want        error
		wantErr     bool
		expectedOut *db.MultipleStatementResult[model.Tag]
	}
	tagTable := []string{
		"id",
		"name",
		"created_at",
		"updated_at",
		"articles",
	}
	cursor := "1"
	zValCursor := ""
	tests := map[string]testCase{
		"happy_path/with_out_paging": {
			args: args{
				ctx:              context.Background(),
				out:              &db.MultipleStatementResult[model.Tag]{},
				paginationOption: []db.PaginationOption{},
			},
			execOpt: func() []db.ExecuteOption {
				sqlDB, mock, err := sqlmock.New()
				if err != nil {
					panic(err)
				}
				rows := sqlmock.NewRows(tagTable).
					AddRow(
						"tag1",
						"test",
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						entity.Articles([]entity.Article{
							{
								ID:        "1",
								Title:     "with_out_paging",
								Thumbnail: "01234567890",
								CreatedAt: synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
								UpdatedAt: synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
							},
						})).
					AddRow(
						"tag2",
						"test",
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						entity.Articles([]entity.Article{
							{
								ID:        "1",
								Title:     "with_out_paging",
								Thumbnail: "01234567890",
								CreatedAt: synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
								UpdatedAt: synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
							},
						}))
				mq := mock.ExpectPrepare(regexp.QuoteMeta(
					`SELECT "t".*,  COALESCE(jsonb_agg(json_build_object('id', a.id, 'title', a.title, 'thumbnail', a.thumbnail, 'created_at', a.created_at, 'updated_at', a.updated_at)) FILTER (WHERE a.id IS NOT NULL), '[]'::json) AS "articles" FROM (SELECT * FROM "tags" ORDER BY "id") AS "t" LEFT OUTER JOIN "articles" AS "a" ON "t"."id" = "a"."tag_id" GROUP BY "t"."id"`))
				mq.ExpectQuery().
					WillReturnRows(rows)
				dialector := postgres.New(postgres.Config{
					Conn: sqlDB,
				})
				gwrapper.Invalidate()
				gwrapper.InvalidateDialector()
				gwrapper.InitializeDialector(&dialector)
				tx, err := gwrapper.Get(context.Background())
				if err != nil {
					panic(nil)
				}
				return []db.ExecuteOption{gwrapper.WithTransaction(tx)}
			},
			want: nil,
			expectedOut: func() *db.MultipleStatementResult[model.Tag] {
				tag1 := model.NewTag(
					"tag1",
					"test",
					model.NewArticle(
						"1",
						"with_out_paging",
						"01234567890",
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
					))
				tag2 := model.NewTag(
					"tag2",
					"test",
					model.NewArticle(
						"1",
						"with_out_paging",
						"01234567890",
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
					))
				tags := []model.Tag{tag1, tag2}
				out := db.NewMultipleStatementResult[model.Tag]()
				out.Set(tags)
				return out
			}(),
		},
		"happy_path/with_prev_paging_limit": {
			args: args{
				ctx: context.Background(),
				out: &db.MultipleStatementResult[model.Tag]{},
				paginationOption: []db.PaginationOption{
					db.WithPreviousPaging(1, nil),
				},
			},
			execOpt: func() []db.ExecuteOption {
				sqlDB, mock, err := sqlmock.New()
				if err != nil {
					panic(err)
				}
				rows := sqlmock.NewRows(tagTable).
					AddRow(
						"tag1",
						"test",
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						entity.Articles([]entity.Article{
							{
								ID:        "1",
								Title:     "with_prev_paging_limit",
								Thumbnail: "01234567890",
								CreatedAt: synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
								UpdatedAt: synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
							},
						})).
					AddRow(
						"tag2",
						"test",
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						entity.Articles([]entity.Article{
							{
								ID:        "1",
								Title:     "with_prev_paging_limit",
								Thumbnail: "01234567890",
								CreatedAt: synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
								UpdatedAt: synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
							},
						}))
				mq := mock.ExpectPrepare(regexp.QuoteMeta(
					`SELECT "t".*,  COALESCE(jsonb_agg(json_build_object('id', a.id, 'title', a.title, 'thumbnail', a.thumbnail, 'created_at', a.created_at, 'updated_at', a.updated_at)) FILTER (WHERE a.id IS NOT NULL), '[]'::json) AS "articles" FROM (SELECT * FROM "tags" ORDER BY "id" DESC LIMIT $1) AS "t" LEFT OUTER JOIN "articles" AS "a" ON "t"."id" = "a"."tag_id" GROUP BY "t"."id"`))
				mq.ExpectQuery().
					WithArgs(2).
					WillReturnRows(rows)
				dialector := postgres.New(postgres.Config{
					Conn: sqlDB,
				})
				gwrapper.Invalidate()
				gwrapper.InvalidateDialector()
				gwrapper.InitializeDialector(&dialector)
				tx, err := gwrapper.Get(context.Background())
				if err != nil {
					panic(nil)
				}
				return []db.ExecuteOption{gwrapper.WithTransaction(tx)}
			},
			want: nil,
			expectedOut: func() *db.MultipleStatementResult[model.Tag] {
				tag1 := model.NewTag(
					"tag1",
					"test",
					model.NewArticle(
						"1",
						"with_prev_paging_limit",
						"01234567890",
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0)))
				tag2 := model.NewTag(
					"tag2",
					"test",
					model.NewArticle("1",
						"with_prev_paging_limit",
						"01234567890",
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0)))
				out := db.NewMultipleStatementResult[model.Tag]()
				out.Set([]model.Tag{tag1, tag2})
				return out
			}(),
		},
		"happy_path/with_prev_paging_limit_and_cursor": {
			args: args{
				ctx: context.Background(),
				out: &db.MultipleStatementResult[model.Tag]{},
				paginationOption: []db.PaginationOption{
					db.WithPreviousPaging(1, &cursor),
				},
			},
			execOpt: func() []db.ExecuteOption {
				sqlDB, mock, err := sqlmock.New()
				if err != nil {
					panic(err)
				}
				rows := sqlmock.NewRows(tagTable).
					AddRow(
						"tag1",
						"test",
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						entity.Articles([]entity.Article{{
							ID:        "1",
							Title:     "with_prev_paging_limit_and_cursor",
							Thumbnail: "01234567890",
							CreatedAt: synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
							UpdatedAt: synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						}})).
					AddRow(
						"tag2",
						"test",
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						entity.Articles([]entity.Article{{
							ID:        "1",
							Title:     "with_prev_paging_limit_and_cursor",
							Thumbnail: "01234567890",
							CreatedAt: synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
							UpdatedAt: synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						}}))
				mq := mock.ExpectPrepare(regexp.QuoteMeta(
					`SELECT "t".*,  COALESCE(jsonb_agg(json_build_object('id', a.id, 'title', a.title, 'thumbnail', a.thumbnail, 'created_at', a.created_at, 'updated_at', a.updated_at)) FILTER (WHERE a.id IS NOT NULL), '[]'::json) AS "articles" FROM (SELECT * FROM "tags" WHERE EXISTS(SELECT id FROM "tags" WHERE "id" = $1) AND "id" < $2 ORDER BY "id" DESC LIMIT $3) AS "t" LEFT OUTER JOIN "articles" AS "a" ON "t"."id" = "a"."tag_id" GROUP BY "t"."id"`))
				mq.ExpectQuery().
					WithArgs(cursor, cursor, 2).
					WillReturnRows(rows)
				dialector := postgres.New(postgres.Config{
					Conn: sqlDB,
				})
				gwrapper.Invalidate()
				gwrapper.InvalidateDialector()
				gwrapper.InitializeDialector(&dialector)
				tx, err := gwrapper.Get(context.Background())
				if err != nil {
					panic(nil)
				}
				return []db.ExecuteOption{gwrapper.WithTransaction(tx)}
			},
			want: nil,
			expectedOut: func() *db.MultipleStatementResult[model.Tag] {
				tag1 := model.NewTag(
					"tag1",
					"test",
					model.NewArticle(
						"1",
						"with_prev_paging_limit_and_cursor",
						"01234567890",
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
					))
				tag2 := model.NewTag(
					"tag2",
					"test",
					model.NewArticle(
						"1",
						"with_prev_paging_limit_and_cursor",
						"01234567890",
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
					))
				out := db.NewMultipleStatementResult[model.Tag]()
				out.Set([]model.Tag{tag1, tag2})
				return out
			}(),
		},
		"happy_path/with_prev_paging_invalid_limit": {
			args: args{
				ctx: context.Background(),
				out: &db.MultipleStatementResult[model.Tag]{},
				paginationOption: []db.PaginationOption{
					db.WithPreviousPaging(0, nil),
				},
			},
			execOpt: func() []db.ExecuteOption {
				sqlDB, mock, err := sqlmock.New()
				if err != nil {
					panic(err)
				}
				rows := sqlmock.NewRows(tagTable).
					AddRow(
						"tag1",
						"test",
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						entity.Articles([]entity.Article{
							{
								ID:        "1",
								Title:     "with_prev_paging_invalid_limit",
								Thumbnail: "01234567890",
								CreatedAt: synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
								UpdatedAt: synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
							},
						})).
					AddRow(
						"tag2",
						"test",
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						entity.Articles([]entity.Article{
							{
								ID:        "1",
								Title:     "with_prev_paging_invalid_limit",
								Thumbnail: "01234567890",
								CreatedAt: synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
								UpdatedAt: synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
							},
						}))
				mq := mock.ExpectPrepare(regexp.QuoteMeta(
					`SELECT "t".*,  COALESCE(jsonb_agg(json_build_object('id', a.id, 'title', a.title, 'thumbnail', a.thumbnail, 'created_at', a.created_at, 'updated_at', a.updated_at)) FILTER (WHERE a.id IS NOT NULL), '[]'::json) AS "articles" FROM (SELECT * FROM "tags" ORDER BY "id" DESC) AS "t" LEFT OUTER JOIN "articles" AS "a" ON "t"."id" = "a"."tag_id" GROUP BY "t"."id"`))
				mq.ExpectQuery().
					WillReturnRows(rows)
				dialector := postgres.New(postgres.Config{
					Conn: sqlDB,
				})
				gwrapper.Invalidate()
				gwrapper.InvalidateDialector()
				gwrapper.InitializeDialector(&dialector)
				tx, err := gwrapper.Get(context.Background())
				if err != nil {
					panic(nil)
				}
				return []db.ExecuteOption{gwrapper.WithTransaction(tx)}
			},
			want: nil,
			expectedOut: func() *db.MultipleStatementResult[model.Tag] {
				tag1 := model.NewTag(
					"tag1",
					"test",
					model.NewArticle(
						"1",
						"with_prev_paging_invalid_limit",
						"01234567890",
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
					))
				tag2 := model.NewTag(
					"tag2",
					"test",
					model.NewArticle(
						"1",
						"with_prev_paging_invalid_limit",
						"01234567890",
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
					))
				out := db.NewMultipleStatementResult[model.Tag]()
				out.Set([]model.Tag{tag1, tag2})
				return out
			}(),
		},
		"happy_path/with_prev_paging_zero_value_cursor": {
			args: args{
				ctx: context.Background(),
				out: &db.MultipleStatementResult[model.Tag]{},
				paginationOption: []db.PaginationOption{
					db.WithPreviousPaging(1, &zValCursor),
				},
			},
			execOpt: func() []db.ExecuteOption {
				sqlDB, mock, err := sqlmock.New()
				if err != nil {
					panic(err)
				}
				rows := sqlmock.NewRows(tagTable).
					AddRow(
						"tag1",
						"test",
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						entity.Articles([]entity.Article{
							{
								ID:        "1",
								Title:     "with_prev_paging_zero_value_cursor",
								Thumbnail: "01234567890",
								CreatedAt: synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
								UpdatedAt: synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
							},
						})).
					AddRow(
						"tag2",
						"test",
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						entity.Articles([]entity.Article{
							{
								ID:        "1",
								Title:     "with_prev_paging_zero_value_cursor",
								Thumbnail: "01234567890",
								CreatedAt: synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
								UpdatedAt: synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
							},
						}))
				mq := mock.ExpectPrepare(regexp.QuoteMeta(
					`SELECT "t".*,  COALESCE(jsonb_agg(json_build_object('id', a.id, 'title', a.title, 'thumbnail', a.thumbnail, 'created_at', a.created_at, 'updated_at', a.updated_at)) FILTER (WHERE a.id IS NOT NULL), '[]'::json) AS "articles" FROM (SELECT * FROM "tags" ORDER BY "id" DESC LIMIT $1) AS "t" LEFT OUTER JOIN "articles" AS "a" ON "t"."id" = "a"."tag_id" GROUP BY "t"."id"`))
				mq.ExpectQuery().
					WithArgs(2).
					WillReturnRows(rows)
				dialector := postgres.New(postgres.Config{
					Conn: sqlDB,
				})
				gwrapper.Invalidate()
				gwrapper.InvalidateDialector()
				gwrapper.InitializeDialector(&dialector)
				tx, err := gwrapper.Get(context.Background())
				if err != nil {
					panic(nil)
				}
				return []db.ExecuteOption{gwrapper.WithTransaction(tx)}
			},
			want: nil,
			expectedOut: func() *db.MultipleStatementResult[model.Tag] {
				tag1 := model.NewTag(
					"tag1",
					"test",
					model.NewArticle(
						"1",
						"with_prev_paging_zero_value_cursor",
						"01234567890",
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
					))
				tag2 := model.NewTag(
					"tag2",
					"test",
					model.NewArticle(
						"1",
						"with_prev_paging_zero_value_cursor",
						"01234567890",
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
					))
				out := db.NewMultipleStatementResult[model.Tag]()
				out.Set([]model.Tag{tag1, tag2})
				return out
			}(),
		},
		"happy_path/with_next_paging_limit": {
			args: args{
				ctx: context.Background(),
				out: &db.MultipleStatementResult[model.Tag]{},
				paginationOption: []db.PaginationOption{
					db.WithNextPaging(1, nil),
				},
			},
			execOpt: func() []db.ExecuteOption {
				sqlDB, mock, err := sqlmock.New()
				if err != nil {
					panic(err)
				}
				rows := sqlmock.NewRows(tagTable).
					AddRow(
						"tag1",
						"test",
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						entity.Articles([]entity.Article{
							{
								ID:        "1",
								Title:     "with_next_paging_limit",
								Thumbnail: "01234567890",
								CreatedAt: synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
								UpdatedAt: synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
							},
						}),
					).
					AddRow(
						"tag2",
						"test",
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						entity.Articles([]entity.Article{
							{
								ID:        "1",
								Title:     "with_next_paging_limit",
								Thumbnail: "01234567890",
								CreatedAt: synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
								UpdatedAt: synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
							},
						}))
				mq := mock.ExpectPrepare(regexp.QuoteMeta(
					`SELECT "t".*,  COALESCE(jsonb_agg(json_build_object('id', a.id, 'title', a.title, 'thumbnail', a.thumbnail, 'created_at', a.created_at, 'updated_at', a.updated_at)) FILTER (WHERE a.id IS NOT NULL), '[]'::json) AS "articles" FROM (SELECT * FROM "tags" ORDER BY "id" LIMIT $1) AS "t" LEFT OUTER JOIN "articles" AS "a" ON "t"."id" = "a"."tag_id" GROUP BY "t"."id"`))
				mq.ExpectQuery().
					WithArgs(2).
					WillReturnRows(rows)
				dialector := postgres.New(postgres.Config{
					Conn: sqlDB,
				})
				gwrapper.Invalidate()
				gwrapper.InvalidateDialector()
				gwrapper.InitializeDialector(&dialector)
				tx, err := gwrapper.Get(context.Background())
				if err != nil {
					panic(nil)
				}
				return []db.ExecuteOption{gwrapper.WithTransaction(tx)}
			},
			want: nil,
			expectedOut: func() *db.MultipleStatementResult[model.Tag] {
				tag1 := model.NewTag(
					"tag1",
					"test",
					model.NewArticle(
						"1",
						"with_next_paging_limit",
						"01234567890",
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
					))
				tag2 := model.NewTag(
					"tag2",
					"test",
					model.NewArticle(
						"1",
						"with_next_paging_limit",
						"01234567890",
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
					))
				out := db.NewMultipleStatementResult[model.Tag]()
				out.Set([]model.Tag{tag1, tag2})
				return out
			}(),
		},
		"happy_path/with_next_paging_limit_and_cursor": {
			args: args{
				ctx: context.Background(),
				out: &db.MultipleStatementResult[model.Tag]{},
				paginationOption: []db.PaginationOption{
					db.WithNextPaging(1, &cursor),
				},
			},
			execOpt: func() []db.ExecuteOption {
				sqlDB, mock, err := sqlmock.New()
				if err != nil {
					panic(err)
				}
				rows := sqlmock.NewRows(tagTable).
					AddRow(
						"tag1",
						"test",
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						entity.Articles([]entity.Article{
							{
								ID:        "1",
								Title:     "with_next_paging_limit_and_cursor",
								Thumbnail: "01234567890",
								CreatedAt: synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
								UpdatedAt: synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
							},
						})).
					AddRow(
						"tag2",
						"test",
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						entity.Articles([]entity.Article{
							{
								ID:        "1",
								Title:     "with_next_paging_limit_and_cursor",
								Thumbnail: "01234567890",
								CreatedAt: synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
								UpdatedAt: synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
							},
						}))
				mq := mock.ExpectPrepare(regexp.QuoteMeta(
					`SELECT "t".*,  COALESCE(jsonb_agg(json_build_object('id', a.id, 'title', a.title, 'thumbnail', a.thumbnail, 'created_at', a.created_at, 'updated_at', a.updated_at)) FILTER (WHERE a.id IS NOT NULL), '[]'::json) AS "articles" FROM (SELECT * FROM "tags" WHERE EXISTS(SELECT id FROM "tags" WHERE "id" = $1) AND "id" > $2 ORDER BY "id" LIMIT $3) AS "t" LEFT OUTER JOIN "articles" AS "a" ON "t"."id" = "a"."tag_id" GROUP BY "t"."id"`))
				mq.ExpectQuery().
					WithArgs(cursor, cursor, 2).
					WillReturnRows(rows)
				dialector := postgres.New(postgres.Config{
					Conn: sqlDB,
				})
				gwrapper.Invalidate()
				gwrapper.InvalidateDialector()
				gwrapper.InitializeDialector(&dialector)
				tx, err := gwrapper.Get(context.Background())
				if err != nil {
					panic(nil)
				}
				return []db.ExecuteOption{gwrapper.WithTransaction(tx)}
			},
			want: nil,
			expectedOut: func() *db.MultipleStatementResult[model.Tag] {
				tag1 := model.NewTag(
					"tag1",
					"test",
					model.NewArticle(
						"1",
						"with_next_paging_limit_and_cursor",
						"01234567890",
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
					))
				tag2 := model.NewTag(
					"tag2",
					"test",
					model.NewArticle(
						"1",
						"with_next_paging_limit_and_cursor",
						"01234567890",
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
					))
				out := db.NewMultipleStatementResult[model.Tag]()
				out.Set([]model.Tag{tag1, tag2})
				return out
			}(),
		},
		"happy_path/with_next_paging_invalid_limit": {
			args: args{
				ctx: context.Background(),
				out: &db.MultipleStatementResult[model.Tag]{},
				paginationOption: []db.PaginationOption{
					db.WithNextPaging(0, nil),
				},
			},
			execOpt: func() []db.ExecuteOption {
				sqlDB, mock, err := sqlmock.New()
				if err != nil {
					panic(err)
				}
				rows := sqlmock.NewRows(tagTable).
					AddRow(
						"tag1",
						"test",
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						entity.Articles([]entity.Article{
							{
								ID:        "1",
								Title:     "with_next_paging_invalid_limit",
								Thumbnail: "01234567890",
								CreatedAt: synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
								UpdatedAt: synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
							},
						})).
					AddRow(
						"tag2",
						"test",
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						entity.Articles([]entity.Article{
							{
								ID:        "1",
								Title:     "with_next_paging_invalid_limit",
								Thumbnail: "01234567890",
								CreatedAt: synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
								UpdatedAt: synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
							},
						}))
				mq := mock.ExpectPrepare(regexp.QuoteMeta(
					`SELECT "t".*,  COALESCE(jsonb_agg(json_build_object('id', a.id, 'title', a.title, 'thumbnail', a.thumbnail, 'created_at', a.created_at, 'updated_at', a.updated_at)) FILTER (WHERE a.id IS NOT NULL), '[]'::json) AS "articles" FROM (SELECT * FROM "tags" ORDER BY "id") AS "t" LEFT OUTER JOIN "articles" AS "a" ON "t"."id" = "a"."tag_id" GROUP BY "t"."id"`))
				mq.ExpectQuery().
					WillReturnRows(rows)
				dialector := postgres.New(postgres.Config{
					Conn: sqlDB,
				})
				gwrapper.Invalidate()
				gwrapper.InvalidateDialector()
				gwrapper.InitializeDialector(&dialector)
				tx, err := gwrapper.Get(context.Background())
				if err != nil {
					panic(nil)
				}
				return []db.ExecuteOption{gwrapper.WithTransaction(tx)}
			},
			want: nil,
			expectedOut: func() *db.MultipleStatementResult[model.Tag] {
				tag1 := model.NewTag(
					"tag1",
					"test",
					model.NewArticle(
						"1",
						"with_next_paging_invalid_limit",
						"01234567890",
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
					))
				tag2 := model.NewTag(
					"tag2",
					"test",
					model.NewArticle(
						"1",
						"with_next_paging_invalid_limit",
						"01234567890",
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
					))
				out := db.NewMultipleStatementResult[model.Tag]()
				out.Set([]model.Tag{tag1, tag2})
				return out
			}(),
		},
		"happy_path/with_next_paging_zero_value_cursor": {
			args: args{
				ctx: context.Background(),
				out: &db.MultipleStatementResult[model.Tag]{},
				paginationOption: []db.PaginationOption{
					db.WithNextPaging(1, &zValCursor),
				},
			},
			execOpt: func() []db.ExecuteOption {
				sqlDB, mock, err := sqlmock.New()
				if err != nil {
					panic(err)
				}
				rows := sqlmock.NewRows(tagTable).
					AddRow(
						"tag1",
						"test",
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						entity.Articles([]entity.Article{
							{
								ID:        "1",
								Title:     "with_next_paging_zero_value_cursor",
								Thumbnail: "01234567890",
								CreatedAt: synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
								UpdatedAt: synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
							},
						})).
					AddRow(
						"tag2",
						"test",
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						entity.Articles([]entity.Article{
							{
								ID:        "1",
								Title:     "with_next_paging_zero_value_cursor",
								Thumbnail: "01234567890",
								CreatedAt: synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
								UpdatedAt: synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
							},
						}))
				mq := mock.ExpectPrepare(regexp.QuoteMeta(
					`SELECT "t".*,  COALESCE(jsonb_agg(json_build_object('id', a.id, 'title', a.title, 'thumbnail', a.thumbnail, 'created_at', a.created_at, 'updated_at', a.updated_at)) FILTER (WHERE a.id IS NOT NULL), '[]'::json) AS "articles" FROM (SELECT * FROM "tags" ORDER BY "id" LIMIT $1) AS "t" LEFT OUTER JOIN "articles" AS "a" ON "t"."id" = "a"."tag_id" GROUP BY "t"."id"`))
				mq.ExpectQuery().
					WithArgs(2).
					WillReturnRows(rows)
				dialector := postgres.New(postgres.Config{
					Conn: sqlDB,
				})
				gwrapper.Invalidate()
				gwrapper.InvalidateDialector()
				gwrapper.InitializeDialector(&dialector)
				tx, err := gwrapper.Get(context.Background())
				if err != nil {
					panic(nil)
				}
				return []db.ExecuteOption{gwrapper.WithTransaction(tx)}
			},
			want: nil,
			expectedOut: func() *db.MultipleStatementResult[model.Tag] {
				tag1 := model.NewTag(
					"tag1",
					"test",
					model.NewArticle(
						"1",
						"with_next_paging_zero_value_cursor",
						"01234567890",
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
					))
				tag2 := model.NewTag(
					"tag2",
					"test",
					model.NewArticle(
						"1",
						"with_next_paging_zero_value_cursor",
						"01234567890",
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2021, 1, 1, 0, 0, 0, 0),
					))
				out := db.NewMultipleStatementResult[model.Tag]()
				out.Set([]model.Tag{tag1, tag2})
				return out
			}(),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			s := NewTagService()
			err := s.GetAll(tt.args.ctx, tt.args.out, tt.args.paginationOption...).Execute(tt.args.ctx, tt.execOpt()...)
			if !errors.Is(err, tt.want) {
				t.Errorf("want error %+v but got %+v", tt.want, err)
				return
			}
			if diff := cmp.Diff(tt.expectedOut.StrictGet(), tt.args.out.StrictGet(), cmp.AllowUnexported(model.Article{}), cmp.AllowUnexported(model.Tag{}), cmpopts.EquateEmpty()); diff != "" {
				t.Errorf(`unexpected output (-want +got): %s`, diff)
				return
			}
		})
	}
}
