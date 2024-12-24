package usecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/cockroachdb/errors"
	"github.com/miyamo2/blogapi.miyamo.today/core/db"
	"github.com/miyamo2/blogapi.miyamo.today/tag-service/internal/app/usecase/dto"
	"github.com/miyamo2/blogapi.miyamo.today/tag-service/internal/infra/rdb/query/model"
	mquery "github.com/miyamo2/blogapi.miyamo.today/tag-service/internal/mock/app/usecase/query"
	mdb "github.com/miyamo2/blogapi.miyamo.today/tag-service/internal/mock/core/db"
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
		setupTagService         func(queryService *mquery.MockTagService, stmt *mdb.MockStatement)
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
			setupTagService: func(queryService *mquery.MockTagService, stmt *mdb.MockStatement) {
				queryService.EXPECT().
					GetById(gomock.Any(), "1", gomock.Any()).
					DoAndReturn(
						func(ctx context.Context, id string, out *db.SingleStatementResult[*model.Tag]) db.Statement {
							tg := model.NewTag("1", "tag1", model.WithTagsSize(2))
							tg.AddArticle(model.NewArticle(
								"1",
								"happy_path",
								"thumbnail",
								"2020-01-01T00:00:00Z",
								"2020-01-01T00:00:00Z"))
							tg.AddArticle(model.NewArticle(
								"2",
								"happy_path2",
								"thumbnail",
								"2020-01-01T00:00:00Z",
								"2020-01-01T00:00:00Z"))

							out.Set(&tg)
							return stmt
						}).
					Times(1)
			},
			want: func() want {
				o := dto.NewGetByIdOutDto(
					"1",
					"tag1",
					[]dto.Article{
						dto.NewArticle(
							"1",
							"happy_path",
							"thumbnail",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z"),
						dto.NewArticle(
							"2",
							"happy_path2",
							"thumbnail",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z"),
					})
				return want{out: &o, err: nil}
			}(),
		},
		"happy_path/tag_has_no_articles": {
			args: args{
				ctx: context.Background(),
				in:  dto.NewGetByIdInDto("1"),
			},
			setupTransaction: func(tx *mdb.MockTransaction, stmt *mdb.MockStatement) {
				tx.EXPECT().
					SubscribeError().
					DoAndReturn(func() <-chan error {
						errCh := make(chan error, 1)
						errCh <- nil
						close(errCh)
						return errCh
					}).
					Times(1)
				tx.EXPECT().
					ExecuteStatement(gomock.Any(), stmt).
					Return(nil).
					Times(1)
				tx.EXPECT().
					Commit(gomock.Any()).
					Return(nil).
					Times(1)
			},
			setupTransactionManager: func(transactionManager *mdb.MockTransactionManager, tx *mdb.MockTransaction) {
				transactionManager.EXPECT().
					GetAndStart(gomock.Any()).
					Return(tx, nil).
					Times(1)
			},
			setupTagService: func(queryService *mquery.MockTagService, stmt *mdb.MockStatement) {
				queryService.EXPECT().
					GetById(gomock.Any(), "1", gomock.Any()).
					DoAndReturn(
						func(ctx context.Context, id string, out *db.SingleStatementResult[*model.Tag]) db.Statement {
							t := model.NewTag("1", "tag1", model.WithTagsSize(0))
							out.Set(&t)
							return stmt
						}).
					Times(1)
			},
			want: func() want {
				o := dto.NewGetByIdOutDto(
					"1",
					"tag1",
					make([]dto.Article, 0))
				return want{out: &o, err: nil}
			}(),
		},
		"unhappy_path/transaction_managers_returns_error": {
			args: args{
				ctx: context.Background(),
				in:  dto.NewGetByIdInDto("1"),
			},
			setupTransaction: func(tx *mdb.MockTransaction, stmt *mdb.MockStatement) {
				tx.EXPECT().
					SubscribeError().
					Times(0)
				tx.EXPECT().
					ExecuteStatement(gomock.Any(), gomock.Any()).
					Times(0)
				tx.EXPECT().
					Commit(gomock.Any()).
					Times(0)
			},
			setupTransactionManager: func(transactionManager *mdb.MockTransactionManager, tx *mdb.MockTransaction) {
				transactionManager.EXPECT().
					GetAndStart(gomock.Any()).
					Return(nil, errTxmn).
					Times(1)
			},
			setupTagService: func(queryService *mquery.MockTagService, stmt *mdb.MockStatement) {
				queryService.EXPECT().
					GetById(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(0)
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
				tx.EXPECT().
					ExecuteStatement(gomock.Any(), stmt).
					Return(errStmt).
					Times(1)
				tx.EXPECT().Commit(gomock.Any()).Times(0)
			},
			setupTransactionManager: func(transactionManager *mdb.MockTransactionManager, tx *mdb.MockTransaction) {
				transactionManager.EXPECT().GetAndStart(gomock.Any()).Return(tx, nil).Times(1)
			},
			setupTagService: func(queryService *mquery.MockTagService, stmt *mdb.MockStatement) {
				queryService.EXPECT().GetById(gomock.Any(), "1", gomock.Any()).DoAndReturn(
					func(ctx context.Context, id string, out *db.SingleStatementResult[*model.Tag]) db.Statement {
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
			setupTagService: func(queryService *mquery.MockTagService, stmt *mdb.MockStatement) {
				queryService.EXPECT().GetById(gomock.Any(), "1", gomock.Any()).DoAndReturn(
					func(ctx context.Context, id string, out *db.SingleStatementResult[*model.Tag]) db.Statement {
						tg := model.NewTag("1", "tag1", model.WithTagsSize(1))
						tg.AddArticle(model.NewArticle(
							"1",
							"happy_path/transaction_commit_returns_error",
							"thumbnail",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z"))
						out.Set(&tg)
						return stmt
					}).Times(1)
			},
			want: func() want {
				o := dto.NewGetByIdOutDto(
					"1",
					"tag1",
					[]dto.Article{
						dto.NewArticle(
							"1",
							"happy_path/transaction_commit_returns_error",
							"thumbnail",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z"),
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
			setupTagService: func(queryService *mquery.MockTagService, stmt *mdb.MockStatement) {
				queryService.EXPECT().GetById(gomock.Any(), "1", gomock.Any()).DoAndReturn(
					func(ctx context.Context, id string, out *db.SingleStatementResult[*model.Tag]) db.Statement {
						tg := model.NewTag("1", "tag1", model.WithTagsSize(1))
						tg.AddArticle(model.NewArticle(
							"1",
							"happy_path/transaction_subscribe_error_receive_error",
							"thumbnail",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z"))
						out.Set(&tg)
						out.Set(&tg)
						return stmt
					}).Times(1)
			},
			want: func() want {
				o := dto.NewGetByIdOutDto(
					"1",
					"tag1",
					[]dto.Article{
						dto.NewArticle(
							"1",
							"happy_path/transaction_subscribe_error_receive_error",
							"thumbnail",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z"),
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
			queryService := mquery.NewMockTagService(ctrl)
			tt.setupTagService(queryService, stmt)
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
