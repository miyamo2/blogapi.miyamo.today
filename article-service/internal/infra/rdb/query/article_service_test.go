package query

import (
	"context"
	"reflect"
	"regexp"
	"testing"

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
		out *db.SingleStatementResult[*Article]
	}
	type testCase struct {
		args        args
		execOpt     func() []db.ExecuteOption
		want        error
		wantErr     bool
		expectedOut *db.SingleStatementResult[*Article]
	}
	articleTable := []string{
		"id",
		"title",
		"body",
		"thumbnail",
		"created_at",
		"updated_at",
		"tag_id",
		"tag_name",
	}
	tests := map[string]testCase{
		"happy_path": {
			args: args{
				ctx: context.Background(),
				id:  "1",
				out: &db.SingleStatementResult[*Article]{},
			},
			execOpt: func() []db.ExecuteOption {
				sqlDB, mock, err := sqlmock.New()
				if err != nil {
					panic(err)
				}
				rows := sqlmock.NewRows(articleTable).
					AddRow(
						"1",
						"happy_path",
						"## happy_path",
						"01234567890",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						nil,
						nil).
					AddRow(
						"1",
						"happy_path",
						"## happy_path",
						"01234567890",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						"tag1",
						"test")
				mq := mock.ExpectPrepare(regexp.QuoteMeta(
					`SELECT "articles".*, "tags"."id" AS "tag_id", "tags"."name" AS "tag_name" FROM (SELECT * FROM "articles" WHERE "id" = $1) AS "articles" LEFT OUTER JOIN "tags" ON "articles"."id" = "tags"."article_id"`))
				mq.ExpectQuery().
					WithArgs("1").WillReturnRows(rows)
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
			expectedOut: func() *db.SingleStatementResult[*Article] {
				article := NewArticle(
					"1",
					"happy_path",
					"## happy_path",
					"01234567890",
					"2021-01-01 00:00:00",
					"2021-01-01 00:00:00",
					WithTagsSize(1))
				article.AddTag(NewTag("tag1", "test"))
				out := db.NewSingleStatementResult[*Article]()
				out.Set(&article)
				return out
			}(),
		},
		"happy_path/article_has_no_tag": {
			args: args{
				ctx: context.Background(),
				id:  "1",
				out: &db.SingleStatementResult[*Article]{},
			},
			execOpt: func() []db.ExecuteOption {
				sqlDB, mock, err := sqlmock.New()
				if err != nil {
					panic(err)
				}
				rows := sqlmock.NewRows(articleTable).AddRow(
					"1",
					"happy_path",
					"## happy_path",
					"01234567890",
					"2021-01-01 00:00:00",
					"2021-01-01 00:00:00",
					nil,
					nil)
				mq := mock.ExpectPrepare(regexp.QuoteMeta(
					`SELECT "articles".*, "tags"."id" AS "tag_id", "tags"."name" AS "tag_name" FROM (SELECT * FROM "articles" WHERE "id" = $1) AS "articles" LEFT OUTER JOIN "tags" ON "articles"."id" = "tags"."article_id"`))
				mq.ExpectQuery().
					WithArgs("1").WillReturnRows(rows)
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
			expectedOut: func() *db.SingleStatementResult[*Article] {
				article := NewArticle(
					"1",
					"happy_path",
					"## happy_path",
					"01234567890",
					"2021-01-01 00:00:00",
					"2021-01-01 00:00:00")
				out := db.NewSingleStatementResult[*Article]()
				out.Set(&article)
				return out
			}(),
		},
		"unhappy_path/not_found": {
			args: args{
				ctx: context.Background(),
				id:  "1",
				out: &db.SingleStatementResult[*Article]{},
			},
			execOpt: func() []db.ExecuteOption {
				sqlDB, mock, err := sqlmock.New()
				if err != nil {
					panic(err)
				}
				rows := sqlmock.NewRows(articleTable)
				mq := mock.ExpectPrepare(regexp.QuoteMeta(
					`SELECT "articles".*, "tags"."id" AS "tag_id", "tags"."name" AS "tag_name" FROM (SELECT * FROM "articles" WHERE "id" = $1) AS "articles" LEFT OUTER JOIN "tags" ON "articles"."id" = "tags"."article_id"`))
				mq.ExpectQuery().
					WithArgs("1").WillReturnRows(rows)
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
			a := NewArticleService()
			err := a.GetById(tt.args.ctx, tt.args.id, tt.args.out).Execute(tt.args.ctx, tt.execOpt()...)
			if tt.wantErr {
				if err == nil {
					t.Errorf("want error but got nil")
					return
				}
				if !errors.Is(err, tt.want) {
					t.Errorf("want error %+v but got %+v", tt.want, err)
					return
				}
				return
			}
			if err != nil {
				t.Errorf("want nil but got error %+v", err)
				return
			}
			if !reflect.DeepEqual(tt.expectedOut.StrictGet(), tt.args.out.StrictGet()) {
				t.Errorf("want %+v but got %+v", tt.expectedOut.StrictGet(), tt.args.out.StrictGet())
				return
			}
		})
	}
}

func TestArticleService_GetAll(t *testing.T) {
	type args struct {
		ctx              context.Context
		paginationOption []db.PaginationOption
		out              *db.MultipleStatementResult[*Article]
	}
	type testCase struct {
		args        args
		execOpt     func() []db.ExecuteOption
		want        error
		wantErr     bool
		expectedOut *db.MultipleStatementResult[*Article]
	}
	articleTable := []string{
		"id",
		"title",
		"body",
		"thumbnail",
		"created_at",
		"updated_at",
		"tag_id",
		"tag_name",
	}
	cursor := "1"
	zValCursor := ""
	tests := map[string]testCase{
		"happy_path/with_out_paging": {
			args: args{
				ctx: context.Background(),
				out: &db.MultipleStatementResult[*Article]{},
			},
			execOpt: func() []db.ExecuteOption {
				sqlDB, mock, err := sqlmock.New()
				if err != nil {
					panic(err)
				}
				rows := sqlmock.NewRows(articleTable).
					AddRow(
						"1",
						"with_out_paging",
						"## with_out_paging",
						"01234567890",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						nil,
						nil).
					AddRow(
						"2",
						"with_out_paging_2",
						"## with_out_paging_2",
						"01234567890",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						nil,
						nil).
					AddRow(
						"2",
						"with_out_paging_2",
						"## with_out_paging_2",
						"01234567890",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						"tag1",
						"test")
				mq := mock.ExpectPrepare(regexp.QuoteMeta(
					`SELECT "articles".*, "tags"."id" AS "tag_id", "tags"."name" AS "tag_name" FROM (SELECT * FROM "articles") AS "articles" LEFT OUTER JOIN "tags" ON "articles"."id" = "tags"."article_id" ORDER BY "articles"."id", "tags"."id" NULLS FIRST`))
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
			expectedOut: func() *db.MultipleStatementResult[*Article] {
				articles := []*Article{
					{
						id:        "1",
						title:     "with_out_paging",
						body:      "## with_out_paging",
						thumbnail: "01234567890",
						createdAt: "2021-01-01 00:00:00",
						updatedAt: "2021-01-01 00:00:00",
						tags:      []Tag{},
					},
					{
						id:        "2",
						title:     "with_out_paging_2",
						body:      "## with_out_paging_2",
						thumbnail: "01234567890",
						createdAt: "2021-01-01 00:00:00",
						updatedAt: "2021-01-01 00:00:00",
						tags: []Tag{
							NewTag("tag1", "test"),
						},
					},
				}
				out := db.NewMultipleStatementResult[*Article]()
				out.Set(articles)
				return out
			}(),
		},
		"happy_path/with_prev_paging_limit": {
			args: args{
				ctx: context.Background(),
				paginationOption: []db.PaginationOption{
					db.WithPreviousPaging(1, nil),
				},
				out: &db.MultipleStatementResult[*Article]{},
			},
			execOpt: func() []db.ExecuteOption {
				sqlDB, mock, err := sqlmock.New()
				if err != nil {
					panic(err)
				}
				rows := sqlmock.NewRows(articleTable).
					AddRow(
						"1",
						"with_prev_paging_limit",
						"## with_prev_paging_limit",
						"01234567890",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						nil,
						nil).
					AddRow(
						"1",
						"with_prev_paging_limit",
						"## with_prev_paging_limit",
						"01234567890",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						"tag1",
						"tag")
				mq := mock.ExpectPrepare(
					regexp.QuoteMeta(
						`SELECT "articles".*, "tags"."id" AS "tag_id", "tags"."name" AS "tag_name" FROM (SELECT * FROM "articles" ORDER BY "id" DESC LIMIT $1) AS "articles" LEFT OUTER JOIN "tags" ON "articles"."id" = "tags"."article_id" ORDER BY "articles"."id" DESC, "tags"."id" NULLS FIRST`))
				mq.ExpectQuery().WithArgs(2).WillReturnRows(rows)
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
			expectedOut: func() *db.MultipleStatementResult[*Article] {
				articles := []*Article{
					{
						id:        "1",
						title:     "with_prev_paging_limit",
						body:      "## with_prev_paging_limit",
						thumbnail: "01234567890",
						createdAt: "2021-01-01 00:00:00",
						updatedAt: "2021-01-01 00:00:00",
						tags: []Tag{
							NewTag("tag1", "tag"),
						},
					},
				}
				out := db.NewMultipleStatementResult[*Article]()
				out.Set(articles)
				return out
			}(),
		},
		"happy_path/with_prev_paging_limit_and_cursor": {
			args: args{
				ctx: context.Background(),
				paginationOption: []db.PaginationOption{
					db.WithPreviousPaging(1, &cursor),
				},
				out: &db.MultipleStatementResult[*Article]{},
			},
			execOpt: func() []db.ExecuteOption {
				sqlDB, mock, err := sqlmock.New()
				if err != nil {
					panic(err)
				}
				rows := sqlmock.NewRows(articleTable).
					AddRow(
						"0",
						"with_prev_paging_limit_and_cursor",
						"## with_prev_paging_limit_and_cursor",
						"01234567890",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						nil,
						nil).
					AddRow(
						"0",
						"with_prev_paging_limit_and_cursor",
						"## with_prev_paging_limit_and_cursor",
						"01234567890",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						"tag1",
						"tag")
				mq := mock.ExpectPrepare(
					regexp.QuoteMeta(
						`SELECT "articles".*, "tags"."id" AS "tag_id", "tags"."name" AS "tag_name" FROM (SELECT * FROM "articles" WHERE EXISTS(SELECT id FROM "articles" WHERE "id" = $1) AND  "id" < $2 ORDER BY "id" DESC LIMIT $3) AS "articles" LEFT OUTER JOIN "tags" ON "articles"."id" = "tags"."article_id" ORDER BY "articles"."id" DESC, "tags"."id" NULLS FIRST`))
				mq.ExpectQuery().
					WithArgs("1", "1", 2).
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
			expectedOut: func() *db.MultipleStatementResult[*Article] {
				articles := []*Article{
					{
						id:        "0",
						title:     "with_prev_paging_limit_and_cursor",
						body:      "## with_prev_paging_limit_and_cursor",
						thumbnail: "01234567890",
						createdAt: "2021-01-01 00:00:00",
						updatedAt: "2021-01-01 00:00:00",
						tags: []Tag{
							NewTag("tag1", "tag"),
						},
					},
				}
				out := db.NewMultipleStatementResult[*Article]()
				out.Set(articles)
				return out
			}(),
		},
		"happy_path/with_prev_paging_invalid_limit": {
			args: args{
				ctx: context.Background(),
				paginationOption: []db.PaginationOption{
					db.WithPreviousPaging(0, nil),
				},
				out: &db.MultipleStatementResult[*Article]{},
			},
			execOpt: func() []db.ExecuteOption {
				sqlDB, mock, err := sqlmock.New()
				if err != nil {
					panic(err)
				}
				rows := sqlmock.NewRows(articleTable).
					AddRow(
						"1",
						"with_out_paging",
						"## with_out_paging",
						"01234567890",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						nil,
						nil).
					AddRow(
						"1",
						"with_out_paging",
						"## with_out_paging",
						"01234567890",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						"tag1",
						"tag")
				mq := mock.ExpectPrepare(
					regexp.QuoteMeta(
						`SELECT "articles".*, "tags"."id" AS "tag_id", "tags"."name" AS "tag_name" FROM (SELECT * FROM "articles") AS "articles" LEFT OUTER JOIN "tags" ON "articles"."id" = "tags"."article_id" ORDER BY "articles"."id", "tags"."id" NULLS FIRST`))
				mq.ExpectQuery().WillReturnRows(rows)
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
			expectedOut: func() *db.MultipleStatementResult[*Article] {
				articles := []*Article{
					{
						id:        "1",
						title:     "with_out_paging",
						body:      "## with_out_paging",
						thumbnail: "01234567890",
						createdAt: "2021-01-01 00:00:00",
						updatedAt: "2021-01-01 00:00:00",
						tags: []Tag{
							NewTag("tag1", "tag"),
						},
					},
				}
				out := db.NewMultipleStatementResult[*Article]()
				out.Set(articles)
				return out
			}(),
		},
		"happy_path/with_prev_paging_zero_value_cursor": {
			args: args{
				ctx: context.Background(),
				paginationOption: []db.PaginationOption{
					db.WithPreviousPaging(1, &zValCursor),
				},
				out: &db.MultipleStatementResult[*Article]{},
			},
			execOpt: func() []db.ExecuteOption {
				sqlDB, mock, err := sqlmock.New()
				if err != nil {
					panic(err)
				}
				rows := sqlmock.NewRows(articleTable).
					AddRow(
						"1",
						"with_prev_paging_zero_value_cursor",
						"## with_prev_paging_zero_value_cursor",
						"01234567890",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						nil,
						nil).
					AddRow(
						"1",
						"with_prev_paging_zero_value_cursor",
						"## with_prev_paging_zero_value_cursor",
						"01234567890",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						"tag1",
						"tag")
				mq := mock.ExpectPrepare(
					regexp.QuoteMeta(
						`SELECT "articles".*, "tags"."id" AS "tag_id", "tags"."name" AS "tag_name" FROM (SELECT * FROM "articles" ORDER BY "id" DESC LIMIT $1) AS "articles" LEFT OUTER JOIN "tags" ON "articles"."id" = "tags"."article_id" ORDER BY "articles"."id" DESC, "tags"."id" NULLS FIRST`))
				mq.ExpectQuery().WithArgs(2).WillReturnRows(rows)
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
			expectedOut: func() *db.MultipleStatementResult[*Article] {
				articles := []*Article{
					{
						id:        "1",
						title:     "with_prev_paging_zero_value_cursor",
						body:      "## with_prev_paging_zero_value_cursor",
						thumbnail: "01234567890",
						createdAt: "2021-01-01 00:00:00",
						updatedAt: "2021-01-01 00:00:00",
						tags: []Tag{
							NewTag("tag1", "tag"),
						},
					},
				}
				out := db.NewMultipleStatementResult[*Article]()
				out.Set(articles)
				return out
			}(),
		},
		"happy_path/with_next_paging_limit": {
			args: args{
				ctx: context.Background(),
				paginationOption: []db.PaginationOption{
					db.WithNextPaging(1, nil),
				},
				out: &db.MultipleStatementResult[*Article]{},
			},
			execOpt: func() []db.ExecuteOption {
				sqlDB, mock, err := sqlmock.New()
				if err != nil {
					panic(err)
				}
				rows := sqlmock.NewRows(articleTable).
					AddRow(
						"1",
						"with_next_paging_limit",
						"## with_next_paging_limit",
						"01234567890",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						nil,
						nil).
					AddRow(
						"1",
						"with_next_paging_limit",
						"## with_next_paging_limit",
						"01234567890",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						"tag1",
						"tag")
				mq := mock.ExpectPrepare(
					regexp.QuoteMeta(
						`SELECT "articles".*, "tags"."id" AS "tag_id", "tags"."name" AS "tag_name" FROM (SELECT * FROM "articles" ORDER BY "id" LIMIT $1) AS "articles" LEFT OUTER JOIN "tags" ON "articles"."id" = "tags"."article_id" ORDER BY "articles"."id", "tags"."id" NULLS FIRST`))
				mq.ExpectQuery().WithArgs(2).WillReturnRows(rows)
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
			expectedOut: func() *db.MultipleStatementResult[*Article] {
				articles := []*Article{
					{
						id:        "1",
						title:     "with_next_paging_limit",
						body:      "## with_next_paging_limit",
						thumbnail: "01234567890",
						createdAt: "2021-01-01 00:00:00",
						updatedAt: "2021-01-01 00:00:00",
						tags: []Tag{
							NewTag("tag1", "tag"),
						},
					},
				}
				out := db.NewMultipleStatementResult[*Article]()
				out.Set(articles)
				return out
			}(),
		},
		"happy_path/with_next_paging_limit_and_cursor": {
			args: args{
				ctx: context.Background(),
				paginationOption: []db.PaginationOption{
					db.WithNextPaging(1, &cursor),
				},
				out: &db.MultipleStatementResult[*Article]{},
			},
			execOpt: func() []db.ExecuteOption {
				sqlDB, mock, err := sqlmock.New()
				if err != nil {
					panic(err)
				}
				rows := sqlmock.NewRows(articleTable).
					AddRow(
						"0",
						"with_next_paging_limit_and_cursor",
						"## with_next_paging_limit_and_cursor",
						"01234567890",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						nil,
						nil).
					AddRow(
						"0",
						"with_next_paging_limit_and_cursor",
						"## with_next_paging_limit_and_cursor",
						"01234567890",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						"tag1",
						"tag")
				mq := mock.ExpectPrepare(
					regexp.QuoteMeta(
						`SELECT "articles".*, "tags"."id" AS "tag_id", "tags"."name" AS "tag_name" FROM (SELECT * FROM "articles" WHERE EXISTS(SELECT id FROM "articles" WHERE "id" = $1) AND  "id" > $2 ORDER BY "id" LIMIT $3) AS "articles" LEFT OUTER JOIN "tags" ON "articles"."id" = "tags"."article_id" ORDER BY "articles"."id", "tags"."id" NULLS FIRST`))
				mq.ExpectQuery().
					WithArgs("1", "1", 2).
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
			expectedOut: func() *db.MultipleStatementResult[*Article] {
				articles := []*Article{
					{
						id:        "0",
						title:     "with_next_paging_limit_and_cursor",
						body:      "## with_next_paging_limit_and_cursor",
						thumbnail: "01234567890",
						createdAt: "2021-01-01 00:00:00",
						updatedAt: "2021-01-01 00:00:00",
						tags: []Tag{
							NewTag("tag1", "tag"),
						},
					},
				}
				out := db.NewMultipleStatementResult[*Article]()
				out.Set(articles)
				return out
			}(),
		},
		"happy_path/with_next_paging_invalid_limit": {
			args: args{
				ctx: context.Background(),
				paginationOption: []db.PaginationOption{
					db.WithNextPaging(0, nil),
				},
				out: &db.MultipleStatementResult[*Article]{},
			},
			execOpt: func() []db.ExecuteOption {
				sqlDB, mock, err := sqlmock.New()
				if err != nil {
					panic(err)
				}
				rows := sqlmock.NewRows(articleTable).
					AddRow(
						"1",
						"with_out_paging",
						"## with_out_paging",
						"01234567890",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						nil,
						nil).
					AddRow(
						"1",
						"with_out_paging",
						"## with_out_paging",
						"01234567890",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						"tag1",
						"tag")
				mq := mock.ExpectPrepare(
					regexp.QuoteMeta(
						`SELECT "articles".*, "tags"."id" AS "tag_id", "tags"."name" AS "tag_name" FROM (SELECT * FROM "articles") AS "articles" LEFT OUTER JOIN "tags" ON "articles"."id" = "tags"."article_id" ORDER BY "articles"."id", "tags"."id" NULLS FIRST`))
				mq.ExpectQuery().WillReturnRows(rows)
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
			expectedOut: func() *db.MultipleStatementResult[*Article] {
				articles := []*Article{
					{
						id:        "1",
						title:     "with_out_paging",
						body:      "## with_out_paging",
						thumbnail: "01234567890",
						createdAt: "2021-01-01 00:00:00",
						updatedAt: "2021-01-01 00:00:00",
						tags: []Tag{
							NewTag("tag1", "tag"),
						},
					},
				}
				out := db.NewMultipleStatementResult[*Article]()
				out.Set(articles)
				return out
			}(),
		},
		"happy_path/with_next_paging_zero_value_cursor": {
			args: args{
				ctx: context.Background(),
				paginationOption: []db.PaginationOption{
					db.WithNextPaging(1, &zValCursor),
				},
				out: &db.MultipleStatementResult[*Article]{},
			},
			execOpt: func() []db.ExecuteOption {
				sqlDB, mock, err := sqlmock.New()
				if err != nil {
					panic(err)
				}
				rows := sqlmock.NewRows(articleTable).
					AddRow(
						"1",
						"with_next_paging_zero_value_cursor",
						"## with_next_paging_zero_value_cursor",
						"01234567890",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						nil,
						nil).
					AddRow(
						"1",
						"with_next_paging_zero_value_cursor",
						"## with_next_paging_zero_value_cursor",
						"01234567890",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						"tag1",
						"tag")
				mq := mock.ExpectPrepare(
					regexp.QuoteMeta(
						`SELECT "articles".*, "tags"."id" AS "tag_id", "tags"."name" AS "tag_name" FROM (SELECT * FROM "articles" ORDER BY "id" LIMIT $1) AS "articles" LEFT OUTER JOIN "tags" ON "articles"."id" = "tags"."article_id" ORDER BY "articles"."id", "tags"."id" NULLS FIRST`))
				mq.ExpectQuery().WithArgs(2).WillReturnRows(rows)
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
			expectedOut: func() *db.MultipleStatementResult[*Article] {
				articles := []*Article{
					{
						id:        "1",
						title:     "with_next_paging_zero_value_cursor",
						body:      "## with_next_paging_zero_value_cursor",
						thumbnail: "01234567890",
						createdAt: "2021-01-01 00:00:00",
						updatedAt: "2021-01-01 00:00:00",
						tags: []Tag{
							NewTag("tag1", "tag"),
						},
					},
				}
				out := db.NewMultipleStatementResult[*Article]()
				out.Set(articles)
				return out
			}(),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			a := NewArticleService()
			err := a.GetAll(tt.args.ctx, tt.args.out, tt.args.paginationOption...).Execute(tt.args.ctx, tt.execOpt()...)
			if tt.wantErr {
				if err == nil {
					t.Errorf("want error but got nil")
					return
				}
				if !errors.Is(err, tt.want) {
					t.Errorf("want error %+v but got %+v", tt.want, err)
					return
				}
				return
			}
			if err != nil {
				t.Errorf("want nil but got error %+v", err)
				return
			}
			if !reflect.DeepEqual(tt.expectedOut.StrictGet(), tt.args.out.StrictGet()) {
				t.Errorf("want %+v but got %+v", tt.expectedOut.StrictGet(), tt.args.out.StrictGet())
				return
			}
		})
	}
}
