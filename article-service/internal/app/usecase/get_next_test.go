package usecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/cockroachdb/errors"
	"github.com/miyamo2/blogapi.miyamo.today/article-service/internal/app/usecase/dto"
	"github.com/miyamo2/blogapi.miyamo.today/article-service/internal/infra/rdb/query"
	mquery "github.com/miyamo2/blogapi.miyamo.today/article-service/internal/mock/app/usecase/query"
	mdb "github.com/miyamo2/blogapi.miyamo.today/article-service/internal/mock/core/db"
	"github.com/miyamo2/blogapi.miyamo.today/core/db"
	"go.uber.org/mock/gomock"
)

func TestGetNext_Execute(t *testing.T) {
	type args struct {
		ctx context.Context
		in  dto.GetNextInDto
	}
	type want struct {
		out *dto.GetNextOutDto
		err error
	}
	type testCase struct {
		args                    args
		setupTransaction        func(tx *mdb.MockTransaction, stmt *mdb.MockStatement)
		setupTransactionManager func(transactionManager *mdb.MockTransactionManager, tx *mdb.MockTransaction)
		setupArticleService     func(queryService *mquery.MockArticleService, stmt *mdb.MockStatement)
		want                    want
		wantErr                 bool
	}

	errTxmn := errors.New("transactionManager error")
	errStmt := errors.New("stmt error")
	errTxCommit := errors.New("tx commit error")
	errTxSubscribeError := errors.New("tx subscribe error")

	tests := map[string]testCase{
		"happy_path": {
			args: args{
				ctx: context.Background(),
				in:  dto.NewGetNextInDto(1, nil),
			},
			setupTransaction: func(tx *mdb.MockTransaction, stmt *mdb.MockStatement) {
				tx.EXPECT().SubscribeError().DoAndReturn(func() <-chan error {
					errCh := make(chan error, 1)
					errCh <- nil
					close(errCh)
					return errCh
				}).Times(1)
				tx.EXPECT().ExecuteStatement(gomock.Any(), stmt).Return(nil).Times(1)
				tx.EXPECT().Commit(gomock.Any()).Return(nil).Times(1)
			},
			setupTransactionManager: func(transactionManager *mdb.MockTransactionManager, tx *mdb.MockTransaction) {
				transactionManager.EXPECT().GetAndStart(gomock.Any()).Return(tx, nil).Times(1)
			},
			setupArticleService: func(queryService *mquery.MockArticleService, stmt *mdb.MockStatement) {
				queryService.EXPECT().GetAll(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
					func(ctx context.Context, out *db.MultipleStatementResult[*query.Article], paginationOption ...db.PaginationOption) db.Statement {
						a1 := query.NewArticle(
							"1",
							"happy_path",
							"## happy_path",
							"thumbnail",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z")
						a1.AddTag(query.NewTag("1", "tag1"))
						a1.AddTag(query.NewTag("2", "tag2"))
						a2 := query.NewArticle(
							"2",
							"happy_path2",
							"## happy_path2",
							"thumbnail",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z")
						a2.AddTag(query.NewTag("1", "tag1"))
						a2.AddTag(query.NewTag("2", "tag2"))
						result := []*query.Article{&a1, &a2}
						out.Set(result)
						return stmt
					}).Times(1)
			},
			want: func() want {
				o := dto.NewGetNextOutDto([]dto.Article{
					dto.NewArticle(
						"1",
						"happy_path",
						"## happy_path",
						"thumbnail",
						"2020-01-01T00:00:00Z",
						"2020-01-01T00:00:00Z",
						[]dto.Tag{
							dto.NewTag("1", "tag1"),
							dto.NewTag("2", "tag2"),
						}),
				}, true)
				return want{out: &o, err: nil}
			}(),
		},
		"happy_path/multiple": {
			args: args{
				ctx: context.Background(),
				in:  dto.NewGetNextInDto(2, nil),
			},
			setupTransaction: func(tx *mdb.MockTransaction, stmt *mdb.MockStatement) {
				tx.EXPECT().SubscribeError().DoAndReturn(func() <-chan error {
					errCh := make(chan error, 1)
					errCh <- nil
					close(errCh)
					return errCh
				}).Times(1)
				tx.EXPECT().ExecuteStatement(gomock.Any(), stmt).Return(nil).Times(1)
				tx.EXPECT().Commit(gomock.Any()).Return(nil).Times(1)
			},
			setupTransactionManager: func(transactionManager *mdb.MockTransactionManager, tx *mdb.MockTransaction) {
				transactionManager.EXPECT().GetAndStart(gomock.Any()).Return(tx, nil).Times(1)
			},
			setupArticleService: func(queryService *mquery.MockArticleService, stmt *mdb.MockStatement) {
				queryService.EXPECT().GetAll(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
					func(ctx context.Context, out *db.MultipleStatementResult[*query.Article], paginationOption ...db.PaginationOption) db.Statement {
						a1 := query.NewArticle(
							"1",
							"happy_path/multiple",
							"## happy_path/multiple",
							"thumbnail",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z")
						a1.AddTag(query.NewTag("1", "tag1"))
						a1.AddTag(query.NewTag("2", "tag2"))
						a2 := query.NewArticle(
							"2",
							"happy_path/multiple2",
							"## happy_path/multiple2",
							"thumbnail",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z")
						a2.AddTag(query.NewTag("1", "tag1"))
						a2.AddTag(query.NewTag("2", "tag2"))
						a3 := query.NewArticle(
							"3",
							"happy_path/multiple3",
							"## happy_path/multiple3",
							"thumbnail",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z")
						a3.AddTag(query.NewTag("1", "tag1"))
						a3.AddTag(query.NewTag("2", "tag2"))
						result := []*query.Article{&a1, &a2, &a3}
						out.Set(result)
						return stmt
					}).Times(1)
			},
			want: func() want {
				o := dto.NewGetNextOutDto([]dto.Article{
					dto.NewArticle(
						"1",
						"happy_path/multiple",
						"## happy_path/multiple",
						"thumbnail",
						"2020-01-01T00:00:00Z",
						"2020-01-01T00:00:00Z",
						[]dto.Tag{
							dto.NewTag("1", "tag1"),
							dto.NewTag("2", "tag2"),
						}),
					dto.NewArticle(
						"2",
						"happy_path/multiple2",
						"## happy_path/multiple2",
						"thumbnail",
						"2020-01-01T00:00:00Z",
						"2020-01-01T00:00:00Z",
						[]dto.Tag{
							dto.NewTag("1", "tag1"),
							dto.NewTag("2", "tag2"),
						}),
				}, true)
				return want{out: &o, err: nil}
			}(),
		},
		"happy_path/zero_article": {
			args: args{
				ctx: context.Background(),
				in:  dto.NewGetNextInDto(1, nil),
			},
			setupTransaction: func(tx *mdb.MockTransaction, stmt *mdb.MockStatement) {
				tx.EXPECT().SubscribeError().DoAndReturn(func() <-chan error {
					errCh := make(chan error, 1)
					errCh <- nil
					close(errCh)
					return errCh
				}).Times(1)
				tx.EXPECT().ExecuteStatement(gomock.Any(), stmt).Return(nil).Times(1)
				tx.EXPECT().Commit(gomock.Any()).Return(nil).Times(1)
			},
			setupTransactionManager: func(transactionManager *mdb.MockTransactionManager, tx *mdb.MockTransaction) {
				transactionManager.EXPECT().GetAndStart(gomock.Any()).Return(tx, nil).Times(1)
			},
			setupArticleService: func(queryService *mquery.MockArticleService, stmt *mdb.MockStatement) {
				queryService.EXPECT().GetAll(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
					func(ctx context.Context, out *db.MultipleStatementResult[*query.Article], paginationOption ...db.PaginationOption) db.Statement {
						result := make([]*query.Article, 0)
						out.Set(result)
						return stmt
					}).Times(1)
			},
			want: func() want {
				o := dto.NewGetNextOutDto(make([]dto.Article, 0), false)
				return want{out: &o, err: nil}
			}(),
		},
		"happy_path/article_has_no_tags": {
			args: args{
				ctx: context.Background(),
				in:  dto.NewGetNextInDto(1, nil),
			},
			setupTransaction: func(tx *mdb.MockTransaction, stmt *mdb.MockStatement) {
				tx.EXPECT().SubscribeError().DoAndReturn(func() <-chan error {
					errCh := make(chan error, 1)
					errCh <- nil
					close(errCh)
					return errCh
				}).Times(1)
				tx.EXPECT().ExecuteStatement(gomock.Any(), stmt).Return(nil).Times(1)
				tx.EXPECT().Commit(gomock.Any()).Return(nil).Times(1)
			},
			setupTransactionManager: func(transactionManager *mdb.MockTransactionManager, tx *mdb.MockTransaction) {
				transactionManager.EXPECT().GetAndStart(gomock.Any()).Return(tx, nil).Times(1)
			},
			setupArticleService: func(queryService *mquery.MockArticleService, stmt *mdb.MockStatement) {
				queryService.EXPECT().GetAll(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
					func(ctx context.Context, out *db.MultipleStatementResult[*query.Article], paginationOption ...db.PaginationOption) db.Statement {
						result := []*query.Article{
							func() *query.Article {
								a := query.NewArticle(
									"1",
									"happy_path/article_has_no_tags",
									"## happy_path/article_has_no_tags",
									"thumbnail",
									"2020-01-01T00:00:00Z",
									"2020-01-01T00:00:00Z")
								return &a
							}(),
						}
						out.Set(result)
						return stmt
					}).Times(1)
			},
			want: func() want {
				o := dto.NewGetNextOutDto(
					[]dto.Article{
						dto.NewArticle(
							"1",
							"happy_path/article_has_no_tags",
							"## happy_path/article_has_no_tags",
							"thumbnail",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z",
							[]dto.Tag{}),
					}, false)
				return want{out: &o, err: nil}
			}(),
		},
		"unhappy_path/transaction_managers_returns_error": {
			args: args{
				ctx: context.Background(),
				in:  dto.NewGetNextInDto(1, nil),
			},
			setupTransaction: func(tx *mdb.MockTransaction, stmt *mdb.MockStatement) {
				tx.EXPECT().SubscribeError().Times(0)
				tx.EXPECT().ExecuteStatement(gomock.Any(), gomock.Any()).Times(0)
				tx.EXPECT().Commit(gomock.Any()).Times(0)
			},
			setupTransactionManager: func(transactionManager *mdb.MockTransactionManager, tx *mdb.MockTransaction) {
				transactionManager.EXPECT().GetAndStart(gomock.Any()).Return(nil, errTxmn).Times(1)
			},
			setupArticleService: func(queryService *mquery.MockArticleService, stmt *mdb.MockStatement) {
				queryService.EXPECT().GetAll(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
			},
			want: func() want {
				return want{out: nil, err: errTxmn}
			}(),
		},
		"unhappy_path/transaction_execute_statement_returns_error": {
			args: args{
				ctx: context.Background(),
				in:  dto.NewGetNextInDto(1, nil),
			},
			setupTransaction: func(tx *mdb.MockTransaction, stmt *mdb.MockStatement) {
				tx.EXPECT().SubscribeError().DoAndReturn(func() <-chan error {
					errCh := make(chan error, 1)
					errCh <- nil
					close(errCh)
					return errCh
				}).Times(1)
				tx.EXPECT().ExecuteStatement(gomock.Any(), stmt).Return(errStmt).Times(1)
				tx.EXPECT().Commit(gomock.Any()).Times(0)
			},
			setupTransactionManager: func(transactionManager *mdb.MockTransactionManager, tx *mdb.MockTransaction) {
				transactionManager.EXPECT().GetAndStart(gomock.Any()).Return(tx, nil).Times(1)
			},
			setupArticleService: func(queryService *mquery.MockArticleService, stmt *mdb.MockStatement) {
				queryService.EXPECT().GetAll(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
					func(ctx context.Context, out *db.MultipleStatementResult[*query.Article], paginationOption ...db.PaginationOption) db.Statement {
						a := query.NewArticle(
							"1",
							"unhappy_path/transaction_execute_statement_returns_error",
							"## unhappy_path/transaction_execute_statement_returns_error",
							"thumbnail",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z")
						a.AddTag(query.NewTag("1", "tag1"))
						a.AddTag(query.NewTag("2", "tag2"))
						result := []*query.Article{&a}
						out.Set(result)
						return stmt
					}).Times(1)
			},
			want: func() want {
				return want{out: nil, err: errStmt}
			}(),
			wantErr: true,
		},
		"happy_path/transaction_commit_returns_error": {
			args: args{
				ctx: context.Background(),
				in:  dto.NewGetNextInDto(1, nil),
			},
			setupTransaction: func(tx *mdb.MockTransaction, stmt *mdb.MockStatement) {
				tx.EXPECT().SubscribeError().DoAndReturn(func() <-chan error {
					errCh := make(chan error, 1)
					errCh <- nil
					close(errCh)
					return errCh
				}).Times(1)
				tx.EXPECT().ExecuteStatement(gomock.Any(), stmt).Return(nil).Times(1)
				tx.EXPECT().Commit(gomock.Any()).Return(errTxCommit).Times(1)
			},
			setupTransactionManager: func(transactionManager *mdb.MockTransactionManager, tx *mdb.MockTransaction) {
				transactionManager.EXPECT().GetAndStart(gomock.Any()).Return(tx, nil).Times(1)
			},
			setupArticleService: func(queryService *mquery.MockArticleService, stmt *mdb.MockStatement) {
				queryService.EXPECT().GetAll(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
					func(ctx context.Context, out *db.MultipleStatementResult[*query.Article], paginationOption ...db.PaginationOption) db.Statement {
						a := query.NewArticle(
							"1",
							"happy_path/transaction_commit_returns_error",
							"## happy_path/transaction_commit_returns_error",
							"thumbnail",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z")
						a.AddTag(query.NewTag("1", "tag1"))
						a.AddTag(query.NewTag("2", "tag2"))
						result := []*query.Article{&a}
						out.Set(result)
						return stmt
					}).Times(1)
			},
			want: func() want {
				o := dto.NewGetNextOutDto(
					[]dto.Article{
						dto.NewArticle(
							"1",
							"happy_path/transaction_commit_returns_error",
							"## happy_path/transaction_commit_returns_error",
							"thumbnail",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z",
							[]dto.Tag{
								dto.NewTag("1", "tag1"),
								dto.NewTag("2", "tag2"),
							}),
					}, false)
				return want{out: &o, err: nil}
			}(),
		},
		"happy_path/transaction_subscribe_error_receive_error": {
			args: args{
				ctx: context.Background(),
				in:  dto.NewGetNextInDto(1, nil),
			},
			setupTransaction: func(tx *mdb.MockTransaction, stmt *mdb.MockStatement) {
				tx.EXPECT().SubscribeError().DoAndReturn(func() <-chan error {
					errCh := make(chan error, 1)
					errCh <- errTxSubscribeError
					close(errCh)
					return errCh
				}).Times(1)
				tx.EXPECT().ExecuteStatement(gomock.Any(), stmt).Return(nil).Times(1)
				tx.EXPECT().Commit(gomock.Any()).Return(errTxCommit).Times(1)
			},
			setupTransactionManager: func(transactionManager *mdb.MockTransactionManager, tx *mdb.MockTransaction) {
				transactionManager.EXPECT().GetAndStart(gomock.Any()).Return(tx, nil).Times(1)
			},
			setupArticleService: func(queryService *mquery.MockArticleService, stmt *mdb.MockStatement) {
				queryService.EXPECT().GetAll(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
					func(ctx context.Context, out *db.MultipleStatementResult[*query.Article], paginationOption ...db.PaginationOption) db.Statement {
						a := query.NewArticle(
							"1",
							"happy_path/transaction_subscribe_error_receive_error",
							"## happy_path/transaction_subscribe_error_receive_error",
							"thumbnail",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z",
						)
						a.AddTag(query.NewTag("1", "tag1"))
						a.AddTag(query.NewTag("2", "tag2"))
						result := []*query.Article{&a}
						out.Set(result)
						return stmt
					}).Times(1)
			},
			want: func() want {
				o := dto.NewGetNextOutDto(
					[]dto.Article{
						dto.NewArticle(
							"1",
							"happy_path/transaction_subscribe_error_receive_error",
							"## happy_path/transaction_subscribe_error_receive_error",
							"thumbnail",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z",
							[]dto.Tag{
								dto.NewTag("1", "tag1"),
								dto.NewTag("2", "tag2"),
							}),
					}, false)
				return want{out: &o, err: nil}
			}(),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			stmt := mdb.NewMockStatement(ctrl)
			tx := mdb.NewMockTransaction(ctrl)
			tt.setupTransaction(tx, stmt)
			transactionManager := mdb.NewMockTransactionManager(ctrl)
			tt.setupTransactionManager(transactionManager, tx)
			queryService := mquery.NewMockArticleService(ctrl)
			tt.setupArticleService(queryService, stmt)
			sut := NewGetNext(transactionManager, queryService)
			got, err := sut.Execute(tt.args.ctx, tt.args.in)
			if tt.wantErr {
				if err == nil {
					t.Errorf("Execute() expected to return an error, but it was nil. want: %+v", err)
					return
				}
				if !errors.Is(err, tt.want.err) {
					t.Errorf("Execute() error = %v, want %v", err, tt.want.err)
					return
				}
			}
			if !reflect.DeepEqual(got, tt.want.out) {
				t.Errorf("Execute() got = %+v, want %+v", got, tt.want.out)
			}
		})
	}
}
