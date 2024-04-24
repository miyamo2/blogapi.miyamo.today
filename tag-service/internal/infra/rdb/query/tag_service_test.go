package query

import (
	"context"
	"reflect"
	"regexp"
	"testing"

	"github.com/miyamo2/api.miyamo.today/tag-service/internal/infra/rdb/query/model"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cockroachdb/errors"
	"github.com/miyamo2/api.miyamo.today/core/db"
	gwrapper "github.com/miyamo2/api.miyamo.today/core/db/gorm"
	"gorm.io/driver/postgres"
)

func TestArticleService_GetById(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
		out *db.SingleStatementResult[*model.Tag]
	}
	type testCase struct {
		args        args
		execOpt     func() []db.ExecuteOption
		want        error
		wantErr     bool
		expectedOut *db.SingleStatementResult[*model.Tag]
	}
	tagTable := []string{
		"id",
		"name",
		"created_at",
		"updated_at",
		"article_id",
		"article_title",
		"article_thumbnail",
		"article_created_at",
		"article_updated_at",
	}
	tests := map[string]testCase{
		"happy_path": {
			args: args{
				ctx: context.Background(),
				id:  "tag1",
				out: &db.SingleStatementResult[*model.Tag]{},
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
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						nil,
						nil,
						nil,
						nil,
						nil).
					AddRow(
						"tag1",
						"test",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						"1",
						"happy_path",
						"01234567890",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00")
				mq := mock.ExpectPrepare(regexp.QuoteMeta(
					`SELECT "tags".*, "articles"."id" AS "article_id", "articles"."title" AS "article_title", "articles"."thumbnail" AS "article_thumbnail", "articles"."created_at" AS "article_created_at", "articles"."updated_at" AS "article_updated_at" FROM (SELECT "tags".* FROM "tags" WHERE "tags"."id" = $1) AS "tags" LEFT OUTER JOIN "articles" ON "tags"."id" = "articles"."tag_id"`))
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
			expectedOut: func() *db.SingleStatementResult[*model.Tag] {
				tag := model.NewTag(
					"tag1",
					"test",
					model.WithTagsSize(1))
				tag.AddArticle(model.NewArticle(
					"1",
					"happy_path",
					"01234567890",
					"2021-01-01 00:00:00",
					"2021-01-01 00:00:00",
				))
				out := db.NewSingleStatementResult[*model.Tag]()
				out.Set(&tag)
				return out
			}(),
		},
		"happy_path/tag_has_no_article": {
			args: args{
				ctx: context.Background(),
				id:  "tag1",
				out: &db.SingleStatementResult[*model.Tag]{},
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
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						nil,
						nil,
						nil,
						nil,
						nil)
				mq := mock.ExpectPrepare(regexp.QuoteMeta(
					`SELECT "tags".*, "articles"."id" AS "article_id", "articles"."title" AS "article_title", "articles"."thumbnail" AS "article_thumbnail", "articles"."created_at" AS "article_created_at", "articles"."updated_at" AS "article_updated_at" FROM (SELECT "tags".* FROM "tags" WHERE "tags"."id" = $1) AS "tags" LEFT OUTER JOIN "articles" ON "tags"."id" = "articles"."tag_id"`))
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
			expectedOut: func() *db.SingleStatementResult[*model.Tag] {
				tag := model.NewTag(
					"tag1",
					"test",
				)
				out := db.NewSingleStatementResult[*model.Tag]()
				out.Set(&tag)
				return out
			}(),
		},
		"unhappy_path/not_found": {
			args: args{
				ctx: context.Background(),
				id:  "tag1",
				out: &db.SingleStatementResult[*model.Tag]{},
			},
			execOpt: func() []db.ExecuteOption {
				sqlDB, mock, err := sqlmock.New()
				if err != nil {
					panic(err)
				}
				rows := sqlmock.NewRows(tagTable)
				mq := mock.ExpectPrepare(regexp.QuoteMeta(
					`SELECT "tags".*, "articles"."id" AS "article_id", "articles"."title" AS "article_title", "articles"."thumbnail" AS "article_thumbnail", "articles"."created_at" AS "article_created_at", "articles"."updated_at" AS "article_updated_at" FROM (SELECT "tags".* FROM "tags" WHERE "tags"."id" = $1) AS "tags" LEFT OUTER JOIN "articles" ON "tags"."id" = "articles"."tag_id"`))
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

func TestTagService_GetAll(t *testing.T) {
	type args struct {
		ctx              context.Context
		out              *db.MultipleStatementResult[*model.Tag]
		paginationOption []db.PaginationOption
	}
	type testCase struct {
		args        args
		execOpt     func() []db.ExecuteOption
		want        error
		wantErr     bool
		expectedOut *db.MultipleStatementResult[*model.Tag]
	}
	tagTable := []string{
		"id",
		"name",
		"created_at",
		"updated_at",
		"article_id",
		"article_title",
		"article_thumbnail",
		"article_created_at",
		"article_updated_at",
	}
	cursor := "1"
	zValCursor := ""
	tests := map[string]testCase{
		"happy_path/with_out_paging": {
			args: args{
				ctx:              context.Background(),
				out:              &db.MultipleStatementResult[*model.Tag]{},
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
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						nil,
						nil,
						nil,
						nil,
						nil).
					AddRow(
						"tag1",
						"test",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						"1",
						"with_out_paging",
						"01234567890",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00").
					AddRow(
						"tag2",
						"test",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						nil,
						nil,
						nil,
						nil,
						nil).
					AddRow(
						"tag2",
						"test",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						"1",
						"with_out_paging",
						"01234567890",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00")
				mq := mock.ExpectPrepare(regexp.QuoteMeta(
					`SELECT "tags".*, "articles"."id" AS "article_id", "articles"."title" AS "article_title", "articles"."thumbnail" AS "article_thumbnail", "articles"."created_at" AS "article_created_at", "articles"."updated_at" AS "article_updated_at" FROM (SELECT * FROM "tags") AS "tags" LEFT OUTER JOIN "articles" ON "tags"."id" = "articles"."tag_id" ORDER BY "tags"."id", "articles"."id" NULLS FIRST`))
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
			expectedOut: func() *db.MultipleStatementResult[*model.Tag] {
				tag1 := model.NewTag(
					"tag1",
					"test",
					model.WithTagsSize(1))
				tag1.AddArticle(model.NewArticle(
					"1",
					"with_out_paging",
					"01234567890",
					"2021-01-01 00:00:00",
					"2021-01-01 00:00:00",
				))
				tag2 := model.NewTag(
					"tag2",
					"test",
					model.WithTagsSize(1))
				tag2.AddArticle(model.NewArticle(
					"1",
					"with_out_paging",
					"01234567890",
					"2021-01-01 00:00:00",
					"2021-01-01 00:00:00",
				))
				tags := []*model.Tag{&tag1, &tag2}
				out := db.NewMultipleStatementResult[*model.Tag]()
				out.Set(tags)
				return out
			}(),
		},
		"happy_path/with_prev_paging_limit": {
			args: args{
				ctx: context.Background(),
				out: &db.MultipleStatementResult[*model.Tag]{},
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
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						nil,
						nil,
						nil,
						nil,
						nil).
					AddRow(
						"tag1",
						"test",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						"1",
						"with_prev_paging_limit",
						"01234567890",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00").
					AddRow(
						"tag2",
						"test",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						nil,
						nil,
						nil,
						nil,
						nil).
					AddRow(
						"tag2",
						"test",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						"1",
						"with_prev_paging_limit",
						"01234567890",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00")
				mq := mock.ExpectPrepare(regexp.QuoteMeta(
					`SELECT "tags".*, "articles"."id" AS "article_id", "articles"."title" AS "article_title", "articles"."thumbnail" AS "article_thumbnail", "articles"."created_at" AS "article_created_at", "articles"."updated_at" AS "article_updated_at" FROM (SELECT * FROM "tags" ORDER BY "id" DESC LIMIT 2) AS "tags" LEFT OUTER JOIN "articles" ON "tags"."id" = "articles"."tag_id" ORDER BY "tags"."id" DESC, "articles"."id" NULLS FIRST`))
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
			expectedOut: func() *db.MultipleStatementResult[*model.Tag] {
				tag1 := model.NewTag(
					"tag1",
					"test",
					model.WithTagsSize(1))
				tag1.AddArticle(model.NewArticle(
					"1",
					"with_prev_paging_limit",
					"01234567890",
					"2021-01-01 00:00:00",
					"2021-01-01 00:00:00",
				))
				tag2 := model.NewTag(
					"tag2",
					"test",
					model.WithTagsSize(1))
				tag2.AddArticle(model.NewArticle(
					"1",
					"with_prev_paging_limit",
					"01234567890",
					"2021-01-01 00:00:00",
					"2021-01-01 00:00:00",
				))
				out := db.NewMultipleStatementResult[*model.Tag]()
				out.Set([]*model.Tag{&tag1, &tag2})
				return out
			}(),
		},
		"happy_path/with_prev_paging_limit_and_cursor": {
			args: args{
				ctx: context.Background(),
				out: &db.MultipleStatementResult[*model.Tag]{},
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
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						nil,
						nil,
						nil,
						nil,
						nil).
					AddRow(
						"tag1",
						"test",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						"1",
						"with_prev_paging_limit_and_cursor",
						"01234567890",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00").
					AddRow(
						"tag2",
						"test",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						nil,
						nil,
						nil,
						nil,
						nil).
					AddRow(
						"tag2",
						"test",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						"1",
						"with_prev_paging_limit_and_cursor",
						"01234567890",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00")
				mq := mock.ExpectPrepare(regexp.QuoteMeta(
					`SELECT "tags".*, "articles"."id" AS "article_id", "articles"."title" AS "article_title", "articles"."thumbnail" AS "article_thumbnail", "articles"."created_at" AS "article_created_at", "articles"."updated_at" AS "article_updated_at" FROM (SELECT * FROM "tags" WHERE EXISTS(SELECT id FROM "tags" WHERE "id" = $1) AND "id" < $2 ORDER BY "id" DESC LIMIT 2) AS "tags" LEFT OUTER JOIN "articles" ON "tags"."id" = "articles"."tag_id" ORDER BY "tags"."id" DESC, "articles"."id" NULLS FIRST`))
				mq.ExpectQuery().
					WithArgs(cursor, cursor).
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
			expectedOut: func() *db.MultipleStatementResult[*model.Tag] {
				tag1 := model.NewTag(
					"tag1",
					"test",
					model.WithTagsSize(1))
				tag1.AddArticle(model.NewArticle(
					"1",
					"with_prev_paging_limit_and_cursor",
					"01234567890",
					"2021-01-01 00:00:00",
					"2021-01-01 00:00:00",
				))
				tag2 := model.NewTag(
					"tag2",
					"test",
					model.WithTagsSize(1))
				tag2.AddArticle(model.NewArticle(
					"1",
					"with_prev_paging_limit_and_cursor",
					"01234567890",
					"2021-01-01 00:00:00",
					"2021-01-01 00:00:00",
				))
				out := db.NewMultipleStatementResult[*model.Tag]()
				out.Set([]*model.Tag{&tag1, &tag2})
				return out
			}(),
		},
		"happy_path/with_prev_paging_invalid_limit": {
			args: args{
				ctx: context.Background(),
				out: &db.MultipleStatementResult[*model.Tag]{},
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
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						nil,
						nil,
						nil,
						nil,
						nil).
					AddRow(
						"tag1",
						"test",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						"1",
						"with_prev_paging_invalid_limit",
						"01234567890",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00").
					AddRow(
						"tag2",
						"test",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						nil,
						nil,
						nil,
						nil,
						nil).
					AddRow(
						"tag2",
						"test",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						"1",
						"with_prev_paging_invalid_limit",
						"01234567890",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00")
				mq := mock.ExpectPrepare(regexp.QuoteMeta(
					`SELECT "tags".*, "articles"."id" AS "article_id", "articles"."title" AS "article_title", "articles"."thumbnail" AS "article_thumbnail", "articles"."created_at" AS "article_created_at", "articles"."updated_at" AS "article_updated_at" FROM (SELECT * FROM "tags") AS "tags" LEFT OUTER JOIN "articles" ON "tags"."id" = "articles"."tag_id" ORDER BY "tags"."id", "articles"."id" NULLS FIRST`))
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
			expectedOut: func() *db.MultipleStatementResult[*model.Tag] {
				tag1 := model.NewTag(
					"tag1",
					"test",
					model.WithTagsSize(1))
				tag1.AddArticle(model.NewArticle(
					"1",
					"with_prev_paging_invalid_limit",
					"01234567890",
					"2021-01-01 00:00:00",
					"2021-01-01 00:00:00",
				))
				tag2 := model.NewTag(
					"tag2",
					"test",
					model.WithTagsSize(1))
				tag2.AddArticle(model.NewArticle(
					"1",
					"with_prev_paging_invalid_limit",
					"01234567890",
					"2021-01-01 00:00:00",
					"2021-01-01 00:00:00",
				))
				out := db.NewMultipleStatementResult[*model.Tag]()
				out.Set([]*model.Tag{&tag1, &tag2})
				return out
			}(),
		},
		"happy_path/with_prev_paging_zero_value_cursor": {
			args: args{
				ctx: context.Background(),
				out: &db.MultipleStatementResult[*model.Tag]{},
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
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						nil,
						nil,
						nil,
						nil,
						nil).
					AddRow(
						"tag1",
						"test",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						"1",
						"with_prev_paging_zero_value_cursor",
						"01234567890",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00").
					AddRow(
						"tag2",
						"test",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						nil,
						nil,
						nil,
						nil,
						nil).
					AddRow(
						"tag2",
						"test",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						"1",
						"with_prev_paging_zero_value_cursor",
						"01234567890",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00")
				mq := mock.ExpectPrepare(regexp.QuoteMeta(
					`SELECT "tags".*, "articles"."id" AS "article_id", "articles"."title" AS "article_title", "articles"."thumbnail" AS "article_thumbnail", "articles"."created_at" AS "article_created_at", "articles"."updated_at" AS "article_updated_at" FROM (SELECT * FROM "tags" ORDER BY "id" DESC LIMIT 2) AS "tags" LEFT OUTER JOIN "articles" ON "tags"."id" = "articles"."tag_id" ORDER BY "tags"."id" DESC, "articles"."id" NULLS FIRST`))
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
			expectedOut: func() *db.MultipleStatementResult[*model.Tag] {
				tag1 := model.NewTag(
					"tag1",
					"test",
					model.WithTagsSize(1))
				tag1.AddArticle(model.NewArticle(
					"1",
					"with_prev_paging_zero_value_cursor",
					"01234567890",
					"2021-01-01 00:00:00",
					"2021-01-01 00:00:00",
				))
				tag2 := model.NewTag(
					"tag2",
					"test",
					model.WithTagsSize(1))
				tag2.AddArticle(model.NewArticle(
					"1",
					"with_prev_paging_zero_value_cursor",
					"01234567890",
					"2021-01-01 00:00:00",
					"2021-01-01 00:00:00",
				))
				out := db.NewMultipleStatementResult[*model.Tag]()
				out.Set([]*model.Tag{&tag1, &tag2})
				return out
			}(),
		},
		"happy_path/with_next_paging_limit": {
			args: args{
				ctx: context.Background(),
				out: &db.MultipleStatementResult[*model.Tag]{},
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
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						nil,
						nil,
						nil,
						nil,
						nil).
					AddRow(
						"tag1",
						"test",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						"1",
						"with_next_paging_limit",
						"01234567890",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00").
					AddRow(
						"tag2",
						"test",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						nil,
						nil,
						nil,
						nil,
						nil).
					AddRow(
						"tag2",
						"test",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						"1",
						"with_next_paging_limit",
						"01234567890",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00")
				mq := mock.ExpectPrepare(regexp.QuoteMeta(
					`SELECT "tags".*, "articles"."id" AS "article_id", "articles"."title" AS "article_title", "articles"."thumbnail" AS "article_thumbnail", "articles"."created_at" AS "article_created_at", "articles"."updated_at" AS "article_updated_at" FROM (SELECT * FROM "tags" ORDER BY "id" LIMIT 2) AS "tags" LEFT OUTER JOIN "articles" ON "tags"."id" = "articles"."tag_id" ORDER BY "tags"."id", "articles"."id" NULLS FIRST`))
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
			expectedOut: func() *db.MultipleStatementResult[*model.Tag] {
				tag1 := model.NewTag(
					"tag1",
					"test",
					model.WithTagsSize(1))
				tag1.AddArticle(model.NewArticle(
					"1",
					"with_next_paging_limit",
					"01234567890",
					"2021-01-01 00:00:00",
					"2021-01-01 00:00:00",
				))
				tag2 := model.NewTag(
					"tag2",
					"test",
					model.WithTagsSize(1))
				tag2.AddArticle(model.NewArticle(
					"1",
					"with_next_paging_limit",
					"01234567890",
					"2021-01-01 00:00:00",
					"2021-01-01 00:00:00",
				))
				out := db.NewMultipleStatementResult[*model.Tag]()
				out.Set([]*model.Tag{&tag1, &tag2})
				return out
			}(),
		},
		"happy_path/with_next_paging_limit_and_cursor": {
			args: args{
				ctx: context.Background(),
				out: &db.MultipleStatementResult[*model.Tag]{},
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
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						nil,
						nil,
						nil,
						nil,
						nil).
					AddRow(
						"tag1",
						"test",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						"1",
						"with_next_paging_limit_and_cursor",
						"01234567890",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00").
					AddRow(
						"tag2",
						"test",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						nil,
						nil,
						nil,
						nil,
						nil).
					AddRow(
						"tag2",
						"test",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						"1",
						"with_next_paging_limit_and_cursor",
						"01234567890",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00")
				mq := mock.ExpectPrepare(regexp.QuoteMeta(
					`SELECT "tags".*, "articles"."id" AS "article_id", "articles"."title" AS "article_title", "articles"."thumbnail" AS "article_thumbnail", "articles"."created_at" AS "article_created_at", "articles"."updated_at" AS "article_updated_at" FROM (SELECT * FROM "tags" WHERE EXISTS(SELECT id FROM "tags" WHERE "id" = $1) AND "id" > $2 ORDER BY "id" LIMIT 2) AS "tags" LEFT OUTER JOIN "articles" ON "tags"."id" = "articles"."tag_id" ORDER BY "tags"."id", "articles"."id" NULLS FIRST`))
				mq.ExpectQuery().
					WithArgs(cursor, cursor).
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
			expectedOut: func() *db.MultipleStatementResult[*model.Tag] {
				tag1 := model.NewTag(
					"tag1",
					"test",
					model.WithTagsSize(1))
				tag1.AddArticle(model.NewArticle(
					"1",
					"with_next_paging_limit_and_cursor",
					"01234567890",
					"2021-01-01 00:00:00",
					"2021-01-01 00:00:00",
				))
				tag2 := model.NewTag(
					"tag2",
					"test",
					model.WithTagsSize(1))
				tag2.AddArticle(model.NewArticle(
					"1",
					"with_next_paging_limit_and_cursor",
					"01234567890",
					"2021-01-01 00:00:00",
					"2021-01-01 00:00:00",
				))
				out := db.NewMultipleStatementResult[*model.Tag]()
				out.Set([]*model.Tag{&tag1, &tag2})
				return out
			}(),
		},
		"happy_path/with_next_paging_invalid_limit": {
			args: args{
				ctx: context.Background(),
				out: &db.MultipleStatementResult[*model.Tag]{},
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
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						nil,
						nil,
						nil,
						nil,
						nil).
					AddRow(
						"tag1",
						"test",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						"1",
						"with_next_paging_invalid_limit",
						"01234567890",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00").
					AddRow(
						"tag2",
						"test",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						nil,
						nil,
						nil,
						nil,
						nil).
					AddRow(
						"tag2",
						"test",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						"1",
						"with_next_paging_invalid_limit",
						"01234567890",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00")
				mq := mock.ExpectPrepare(regexp.QuoteMeta(
					`SELECT "tags".*, "articles"."id" AS "article_id", "articles"."title" AS "article_title", "articles"."thumbnail" AS "article_thumbnail", "articles"."created_at" AS "article_created_at", "articles"."updated_at" AS "article_updated_at" FROM (SELECT * FROM "tags") AS "tags" LEFT OUTER JOIN "articles" ON "tags"."id" = "articles"."tag_id" ORDER BY "tags"."id", "articles"."id" NULLS FIRST`))
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
			expectedOut: func() *db.MultipleStatementResult[*model.Tag] {
				tag1 := model.NewTag(
					"tag1",
					"test",
					model.WithTagsSize(1))
				tag1.AddArticle(model.NewArticle(
					"1",
					"with_next_paging_invalid_limit",
					"01234567890",
					"2021-01-01 00:00:00",
					"2021-01-01 00:00:00",
				))
				tag2 := model.NewTag(
					"tag2",
					"test",
					model.WithTagsSize(1))
				tag2.AddArticle(model.NewArticle(
					"1",
					"with_next_paging_invalid_limit",
					"01234567890",
					"2021-01-01 00:00:00",
					"2021-01-01 00:00:00",
				))
				out := db.NewMultipleStatementResult[*model.Tag]()
				out.Set([]*model.Tag{&tag1, &tag2})
				return out
			}(),
		},
		"happy_path/with_next_paging_zero_value_cursor": {
			args: args{
				ctx: context.Background(),
				out: &db.MultipleStatementResult[*model.Tag]{},
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
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						nil,
						nil,
						nil,
						nil,
						nil).
					AddRow(
						"tag1",
						"test",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						"1",
						"with_next_paging_zero_value_cursor",
						"01234567890",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00").
					AddRow(
						"tag2",
						"test",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						nil,
						nil,
						nil,
						nil,
						nil).
					AddRow(
						"tag2",
						"test",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00",
						"1",
						"with_next_paging_zero_value_cursor",
						"01234567890",
						"2021-01-01 00:00:00",
						"2021-01-01 00:00:00")
				mq := mock.ExpectPrepare(regexp.QuoteMeta(
					`SELECT "tags".*, "articles"."id" AS "article_id", "articles"."title" AS "article_title", "articles"."thumbnail" AS "article_thumbnail", "articles"."created_at" AS "article_created_at", "articles"."updated_at" AS "article_updated_at" FROM (SELECT * FROM "tags" ORDER BY "id" LIMIT 2) AS "tags" LEFT OUTER JOIN "articles" ON "tags"."id" = "articles"."tag_id" ORDER BY "tags"."id", "articles"."id" NULLS FIRST`))
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
			expectedOut: func() *db.MultipleStatementResult[*model.Tag] {
				tag1 := model.NewTag(
					"tag1",
					"test",
					model.WithTagsSize(1))
				tag1.AddArticle(model.NewArticle(
					"1",
					"with_next_paging_zero_value_cursor",
					"01234567890",
					"2021-01-01 00:00:00",
					"2021-01-01 00:00:00",
				))
				tag2 := model.NewTag(
					"tag2",
					"test",
					model.WithTagsSize(1))
				tag2.AddArticle(model.NewArticle(
					"1",
					"with_next_paging_zero_value_cursor",
					"01234567890",
					"2021-01-01 00:00:00",
					"2021-01-01 00:00:00",
				))
				out := db.NewMultipleStatementResult[*model.Tag]()
				out.Set([]*model.Tag{&tag1, &tag2})
				return out
			}(),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			s := NewTagService()
			err := s.GetAll(tt.args.ctx, tt.args.out, tt.args.paginationOption...).Execute(tt.args.ctx, tt.execOpt()...)
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
