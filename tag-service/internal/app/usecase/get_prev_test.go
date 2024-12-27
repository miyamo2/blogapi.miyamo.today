package usecase

import (
	"context"
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	"reflect"
	"testing"

	"github.com/miyamo2/blogapi.miyamo.today/tag-service/internal/infra/rdb/query/model"
	mquery "github.com/miyamo2/blogapi.miyamo.today/tag-service/internal/mock/app/usecase/query"
	mdb "github.com/miyamo2/blogapi.miyamo.today/tag-service/internal/mock/core/db"
	"go.uber.org/mock/gomock"

	"github.com/cockroachdb/errors"
	"github.com/miyamo2/blogapi.miyamo.today/core/db"
	"github.com/miyamo2/blogapi.miyamo.today/tag-service/internal/app/usecase/dto"
)

func TestGetPrev_Execute(t *testing.T) {
	type args struct {
		ctx context.Context
		in  dto.GetPrevInDto
	}
	type want struct {
		out *dto.GetPrevOutDto
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
		"happy_path/multiple/has_prev": {
			args: args{
				ctx: context.Background(),
				in:  dto.NewGetPrevInDto(2, nil),
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
					Return(nil).Times(1)
			},
			setupTransactionManager: func(transactionManager *mdb.MockTransactionManager, tx *mdb.MockTransaction) {
				transactionManager.EXPECT().
					GetAndStart(gomock.Any()).
					Return(tx, nil).
					Times(1)
			},
			setupTagService: func(queryService *mquery.MockTagService, stmt *mdb.MockStatement) {
				queryService.EXPECT().
					GetAll(gomock.Any(), gomock.Any(), gomock.Any()).
					DoAndReturn(
						func(ctx context.Context, out *db.MultipleStatementResult[model.Tag], paginationOption ...db.PaginationOption) db.Statement {
							tg1 := model.NewTag("1", "tag1", model.NewArticle(
								"1",
								"happy_path/multiple/has_prev1",
								"thumbnail",
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								model.NewArticle(
									"2",
									"happy_path/multiple/has_prev2",
									"thumbnail",
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)))
							tg2 := model.NewTag("2", "tag2", model.NewArticle(
								"1",
								"happy_path/multiple/has_prev1",
								"thumbnail",
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								model.NewArticle(
									"2",
									"happy_path/multiple/has_prev2",
									"thumbnail",
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)))
							tg3 := model.NewTag("3", "tag3", model.NewArticle(
								"1",
								"happy_path/multiple/has_prev1",
								"thumbnail",
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								model.NewArticle(
									"2",
									"happy_path/multiple/has_prev2",
									"thumbnail",
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)))

							result := []model.Tag{tg1, tg2, tg3}
							out.Set(result)
							return stmt
						}).
					Times(1)
			},
			want: func() want {
				o := dto.NewGetPrevOutDto(true)
				o = o.WithTagDto(
					dto.NewTag(
						"1",
						"tag1",
						[]dto.Article{
							dto.NewArticle(
								"1",
								"happy_path/multiple/has_prev1",
								"thumbnail",
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
							dto.NewArticle(
								"2",
								"happy_path/multiple/has_prev2",
								"thumbnail",
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
						},
					))
				o = o.WithTagDto(
					dto.NewTag(
						"2",
						"tag2",
						[]dto.Article{
							dto.NewArticle(
								"1",
								"happy_path/multiple/has_prev1",
								"thumbnail",
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
							dto.NewArticle(
								"2",
								"happy_path/multiple/has_prev2",
								"thumbnail",
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
						},
					))
				return want{out: &o, err: nil}
			}(),
		},
		"happy_path/multiple/has_not_prev": {
			args: args{
				ctx: context.Background(),
				in:  dto.NewGetPrevInDto(2, nil),
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
					Return(nil).Times(1)
			},
			setupTransactionManager: func(transactionManager *mdb.MockTransactionManager, tx *mdb.MockTransaction) {
				transactionManager.EXPECT().
					GetAndStart(gomock.Any()).
					Return(tx, nil).
					Times(1)
			},
			setupTagService: func(queryService *mquery.MockTagService, stmt *mdb.MockStatement) {
				queryService.EXPECT().
					GetAll(gomock.Any(), gomock.Any(), gomock.Any()).
					DoAndReturn(
						func(ctx context.Context, out *db.MultipleStatementResult[model.Tag], paginationOption ...db.PaginationOption) db.Statement {
							tg1 := model.NewTag("1", "tag1", model.NewArticle(
								"1",
								"happy_path/multiple/has_not_prev1",
								"thumbnail",
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								model.NewArticle(
									"2",
									"happy_path/multiple/has_not_prev2",
									"thumbnail",
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)))
							tg2 := model.NewTag("2", "tag2", model.NewArticle(
								"1",
								"happy_path/multiple/has_not_prev1",
								"thumbnail",
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								model.NewArticle(
									"2",
									"happy_path/multiple/has_not_prev2",
									"thumbnail",
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)))
							result := []model.Tag{tg1, tg2}
							out.Set(result)
							return stmt
						}).
					Times(1)
			},
			want: func() want {
				o := dto.NewGetPrevOutDto(false)
				o = o.WithTagDto(
					dto.NewTag(
						"1",
						"tag1",
						[]dto.Article{
							dto.NewArticle(
								"1",
								"happy_path/multiple/has_not_prev1",
								"thumbnail",
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
							dto.NewArticle(
								"2",
								"happy_path/multiple/has_not_prev2",
								"thumbnail",
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
						},
					))
				o = o.WithTagDto(
					dto.NewTag(
						"2",
						"tag2",
						[]dto.Article{
							dto.NewArticle(
								"1",
								"happy_path/multiple/has_not_prev1",
								"thumbnail",
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
							dto.NewArticle(
								"2",
								"happy_path/multiple/has_not_prev2",
								"thumbnail",
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
						},
					))
				return want{out: &o, err: nil}
			}(),
		},
		"happy_path/tag_has_no_articles": {
			args: args{
				ctx: context.Background(),
				in:  dto.NewGetPrevInDto(1, nil),
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
					GetAll(gomock.Any(), gomock.Any(), gomock.Any()).
					DoAndReturn(
						func(ctx context.Context, out *db.MultipleStatementResult[model.Tag], paginationOption ...db.PaginationOption) db.Statement {
							tg := model.NewTag("1", "tag1")
							result := []model.Tag{tg}
							out.Set(result)
							return stmt
						}).
					Times(1)
			},
			want: func() want {
				o := dto.NewGetPrevOutDto(false)
				o = o.WithTagDto(dto.NewTag("1", "tag1", []dto.Article{}))
				return want{out: &o, err: nil}
			}(),
		},
		"unhappy_path/transaction_managers_returns_error": {
			args: args{
				ctx: context.Background(),
				in:  dto.NewGetPrevInDto(1, nil),
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
					GetAll(gomock.Any(), gomock.Any()).
					Times(0)
			},
			want: func() want {
				return want{out: nil, err: errTxmn}
			}(),
		},
		"unhappy_path/transaction_execute_statement_returns_error": {
			args: args{
				ctx: context.Background(),
				in:  dto.NewGetPrevInDto(1, nil),
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
					Return(errStmt).
					Times(1)
				tx.EXPECT().
					Commit(gomock.Any()).
					Times(0)
			},
			setupTransactionManager: func(transactionManager *mdb.MockTransactionManager, tx *mdb.MockTransaction) {
				transactionManager.EXPECT().
					GetAndStart(gomock.Any()).
					Return(tx, nil).
					Times(1)
			},
			setupTagService: func(queryService *mquery.MockTagService, stmt *mdb.MockStatement) {
				queryService.EXPECT().
					GetAll(gomock.Any(), gomock.Any(), gomock.Any()).
					DoAndReturn(
						func(ctx context.Context, out *db.MultipleStatementResult[model.Tag], paginationOption ...db.PaginationOption) db.Statement {
							tg := model.NewTag("1", "tag1", model.NewArticle(
								"1",
								"unhappy_path/transaction_execute_statement_returns_error",
								"thumbnail",
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)))
							result := []model.Tag{tg}
							out.Set(result)
							return stmt
						}).
					Times(1)
			},
			want: func() want {
				return want{out: nil, err: errStmt}
			}(),
			wantErr: true,
		},
		"happy_path/transaction_commit_returns_error": {
			args: args{
				ctx: context.Background(),
				in:  dto.NewGetPrevInDto(1, nil),
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
					Return(errTxCommit).
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
					GetAll(gomock.Any(), gomock.Any(), gomock.Any()).
					DoAndReturn(
						func(ctx context.Context, out *db.MultipleStatementResult[model.Tag], paginationOption ...db.PaginationOption) db.Statement {
							tg := model.NewTag("1", "tag1", model.NewArticle(
								"1",
								"happy_path/transaction_commit_returns_error",
								"thumbnail",
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)))
							result := []model.Tag{tg}
							out.Set(result)
							return stmt
						}).Times(1)
			},
			want: func() want {
				o := dto.NewGetPrevOutDto(false)
				o = o.WithTagDto(
					dto.NewTag("1", "tag1", []dto.Article{
						dto.NewArticle(
							"1",
							"happy_path/transaction_commit_returns_error",
							"thumbnail",
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
					}),
				)
				return want{out: &o, err: nil}
			}(),
		},
		"happy_path/transaction_subscribe_error_receive_error": {
			args: args{
				ctx: context.Background(),
				in:  dto.NewGetPrevInDto(1, nil),
			},
			setupTransaction: func(tx *mdb.MockTransaction, stmt *mdb.MockStatement) {
				tx.EXPECT().
					SubscribeError().
					DoAndReturn(func() <-chan error {
						errCh := make(chan error, 1)
						errCh <- errTxSubscribeError
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
					Return(errTxCommit).
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
					GetAll(gomock.Any(), gomock.Any(), gomock.Any()).
					DoAndReturn(
						func(ctx context.Context, out *db.MultipleStatementResult[model.Tag], paginationOption ...db.PaginationOption) db.Statement {
							tg1 := model.NewTag("1", "tag1", model.NewArticle(
								"1",
								"happy_path/transaction_subscribe_error_receive_error",
								"thumbnail",
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							))
							result := []model.Tag{tg1}
							out.Set(result)
							return stmt
						}).
					Times(1)
			},
			want: func() want {
				o := dto.NewGetPrevOutDto(false)
				o = o.WithTagDto(
					dto.NewTag(
						"1",
						"tag1",
						[]dto.Article{
							dto.NewArticle(
								"1",
								"happy_path/transaction_subscribe_error_receive_error",
								"thumbnail",
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
						}))
				return want{out: &o, err: nil}
			}(),
		},
		"happy_path/single/has_prev": {
			args: args{
				ctx: context.Background(),
				in:  dto.NewGetPrevInDto(1, nil),
			},
			setupTransaction: func(tx *mdb.MockTransaction, stmt *mdb.MockStatement) {
				tx.EXPECT().
					SubscribeError().
					DoAndReturn(func() <-chan error {
						errCh := make(chan error, 1)
						errCh <- errTxSubscribeError
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
					Return(errTxCommit).
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
					GetAll(gomock.Any(), gomock.Any(), gomock.Any()).
					DoAndReturn(
						func(ctx context.Context, out *db.MultipleStatementResult[model.Tag], paginationOption ...db.PaginationOption) db.Statement {
							tg1 := model.NewTag("1", "tag1", model.NewArticle(
								"1",
								"happy_path/single/has_prev",
								"thumbnail",
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							))
							tg2 := model.NewTag("2", "tag2", model.NewArticle(
								"1",
								"happy_path/single/has_prev",
								"thumbnail",
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							))
							result := []model.Tag{tg1, tg2}
							out.Set(result)
							return stmt
						}).
					Times(1)
			},
			want: func() want {
				o := dto.NewGetPrevOutDto(true)
				o = o.WithTagDto(
					dto.NewTag(
						"1",
						"tag1",
						[]dto.Article{
							dto.NewArticle(
								"1",
								"happy_path/single/has_prev",
								"thumbnail",
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
						}))
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
			sut := NewGetPrev(transactionManager, queryService)
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
				t.Errorf("Execute() got = %+v, want %+v", *got, *tt.want.out)
			}
		})
	}
}
