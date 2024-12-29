package usecase

import (
	"context"
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	"reflect"
	"testing"

	"blogapi.miyamo.today/article-service/internal/app/usecase/dto"
	"blogapi.miyamo.today/article-service/internal/infra/rdb/query"
	mquery "blogapi.miyamo.today/article-service/internal/mock/app/usecase/query"
	mdb "blogapi.miyamo.today/article-service/internal/mock/core/db"
	"blogapi.miyamo.today/core/db"
	"github.com/cockroachdb/errors"
	"go.uber.org/mock/gomock"
)

func TestGetById_Execute(t *testing.T) {
	type args struct {
		ctx context.Context
		in  dto.GetByIdInDto
	}
	type want struct {
		out *dto.GetByIdOutDto
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
				in:  dto.NewGetByIdInDto("1"),
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
				queryService.EXPECT().GetById(gomock.Any(), "1", gomock.Any()).DoAndReturn(
					func(ctx context.Context, id string, out *db.SingleStatementResult[query.Article]) db.Statement {
						a := query.NewArticle(
							"1",
							"happy_path",
							"## happy_path",
							"thumbnail",
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							query.NewTag("1", "tag1"),
							query.NewTag("2", "tag2"))
						out.Set(a)
						return stmt
					}).Times(1)
			},
			want: func() want {
				o := dto.NewGetByIdOutDto(
					"1",
					"happy_path",
					"## happy_path",
					"thumbnail",
					synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					[]dto.Tag{
						dto.NewTag("1", "tag1"),
						dto.NewTag("2", "tag2"),
					})
				return want{out: &o, err: nil}
			}(),
		},
		"happy_path/article_has_no_tags": {
			args: args{
				ctx: context.Background(),
				in:  dto.NewGetByIdInDto("1"),
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
				queryService.EXPECT().GetById(gomock.Any(), "1", gomock.Any()).DoAndReturn(
					func(ctx context.Context, id string, out *db.SingleStatementResult[query.Article]) db.Statement {
						a := query.NewArticle(
							"1",
							"happy_path",
							"## happy_path",
							"thumbnail",
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0))
						out.Set(a)
						return stmt
					}).Times(1)
			},
			want: func() want {
				o := dto.NewGetByIdOutDto(
					"1",
					"happy_path",
					"## happy_path",
					"thumbnail",
					synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					make([]dto.Tag, 0))
				return want{out: &o, err: nil}
			}(),
		},
		"unhappy_path/transaction_managers_returns_error": {
			args: args{
				ctx: context.Background(),
				in:  dto.NewGetByIdInDto("1"),
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
				queryService.EXPECT().GetById(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
			},
			want: func() want {
				return want{out: nil, err: errTxmn}
			}(),
		},
		"unhappy_path/transaction_execute_statement_returns_error": {
			args: args{
				ctx: context.Background(),
				in:  dto.NewGetByIdInDto("1"),
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
				queryService.EXPECT().GetById(gomock.Any(), "1", gomock.Any()).DoAndReturn(
					func(ctx context.Context, id string, out *db.SingleStatementResult[query.Article]) db.Statement {
						a := query.NewArticle(
							"1",
							"unhappy_path/transaction_execute_statement_returns_error",
							"## unhappy_path/transaction_execute_statement_returns_error",
							"thumbnail",
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							query.NewTag("1", "tag1"),
							query.NewTag("2", "tag2"))
						out.Set(a)
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
				in:  dto.NewGetByIdInDto("1"),
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
				queryService.EXPECT().GetById(gomock.Any(), "1", gomock.Any()).DoAndReturn(
					func(ctx context.Context, id string, out *db.SingleStatementResult[query.Article]) db.Statement {
						a := query.NewArticle(
							"1",
							"happy_path/transaction_commit_returns_error",
							"## happy_path/transaction_commit_returns_error",
							"thumbnail",
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							query.NewTag("1", "tag1"),
							query.NewTag("2", "tag2"))
						out.Set(a)
						return stmt
					}).Times(1)
			},
			want: func() want {
				o := dto.NewGetByIdOutDto(
					"1",
					"happy_path/transaction_commit_returns_error",
					"## happy_path/transaction_commit_returns_error",
					"thumbnail",
					synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					[]dto.Tag{
						dto.NewTag("1", "tag1"),
						dto.NewTag("2", "tag2"),
					})
				return want{out: &o, err: nil}
			}(),
		},
		"happy_path/transaction_subscribe_error_receive_error": {
			args: args{
				ctx: context.Background(),
				in:  dto.NewGetByIdInDto("1"),
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
				queryService.EXPECT().GetById(gomock.Any(), "1", gomock.Any()).DoAndReturn(
					func(ctx context.Context, id string, out *db.SingleStatementResult[query.Article]) db.Statement {
						a := query.NewArticle(
							"1",
							"happy_path/transaction_subscribe_error_receive_error",
							"## happy_path/transaction_subscribe_error_receive_error",
							"thumbnail",
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							query.NewTag("1", "tag1"),
							query.NewTag("2", "tag2"))
						out.Set(a)
						return stmt
					}).Times(1)
			},
			want: func() want {
				o := dto.NewGetByIdOutDto(
					"1",
					"happy_path/transaction_subscribe_error_receive_error",
					"## happy_path/transaction_subscribe_error_receive_error",
					"thumbnail",
					synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					[]dto.Tag{
						dto.NewTag("1", "tag1"),
						dto.NewTag("2", "tag2"),
					})
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
			sut := NewGetById(transactionManager, queryService)
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
				t.Errorf("Execute() got = %v, want %v", got, tt.want)
			}
		})
	}
}
