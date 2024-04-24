package dynamodb

import (
	"context"
	"log/slog"
	"testing"

	dynamotypes "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/cockroachdb/errors"
	"github.com/miyamo2/blogapi.miyamo.today/core/db/dynamodb/client"
	"github.com/miyamo2/blogapi.miyamo.today/core/db/internal"
	mclient "github.com/miyamo2/blogapi.miyamo.today/core/internal/mock/infra/dynamodb/client"
	"go.uber.org/mock/gomock"
)

func TestTransaction_Commit(t *testing.T) {
	type expect struct {
		commitLength int
	}
	type testCase struct {
		commit  func() chan struct{}
		sut     func(commit chan struct{}) Transaction
		expect  expect
		wantErr bool
	}
	tests := map[string]testCase{
		"happy_path": {
			commit: func() chan struct{} {
				commit := make(chan struct{}, 1)
				return commit
			},
			sut: func(commit chan struct{}) Transaction {
				return Transaction{
					commit: commit,
				}
			},
			expect: expect{
				commitLength: 1,
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			commit := tt.commit()
			defer close(commit)
			tx := tt.sut(commit)

			if err := tx.Commit(context.Background()); (err != nil) != tt.wantErr {
				t.Errorf("Commit() error = %v, wantErr %v", err, tt.wantErr)
			}
			if len(commit) != tt.expect.commitLength {
				t.Errorf("Commit() rollback length = %v, want %v", len(commit), tt.expect.commitLength)
			}
		})
	}
}

func TestTransaction_ExecuteStatement(t *testing.T) {
	type testCase struct {
		stmtQueue      func() chan *internal.StatementRequest
		sut            func(commit chan *internal.StatementRequest) Transaction
		mockSubscriber func(chan *internal.StatementRequest)
	}

	tests := map[string]testCase{
		"happy_path": {
			stmtQueue: func() chan *internal.StatementRequest {
				stmtQueue := make(chan *internal.StatementRequest)
				return stmtQueue
			},
			sut: func(stmtQueue chan *internal.StatementRequest) Transaction {
				return Transaction{
					stmtQueue: stmtQueue,
				}
			},
			mockSubscriber: func(stmtQueue chan *internal.StatementRequest) {
				nx := <-stmtQueue
				slog.Info("receive", slog.Any("Statement", nx.Statement))
			},
		},
		"unhappy_path/returns_error": {
			stmtQueue: func() chan *internal.StatementRequest {
				stmtQueue := make(chan *internal.StatementRequest)
				return stmtQueue
			},
			sut: func(stmtQueue chan *internal.StatementRequest) Transaction {
				return Transaction{
					stmtQueue: stmtQueue,
				}
			},
			mockSubscriber: func(stmtQueue chan *internal.StatementRequest) {
				nx := <-stmtQueue
				slog.Info("receive", slog.Any("Statement", nx.Statement))
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			stmtQueue := tt.stmtQueue()
			defer close(stmtQueue)
			tx := tt.sut(stmtQueue)
			go tt.mockSubscriber(stmtQueue)

			if err := tx.ExecuteStatement(context.Background(), &Statement{}); err != nil {
				t.Errorf("ExecuteStatement() returns error = %v", err)
			}
		})
	}
}

func TestTransaction_Rollback(t *testing.T) {
	type expect struct {
		rollbackLength int
	}
	type testCase struct {
		rollback func() chan struct{}
		sut      func(commit chan struct{}) Transaction
		expect   expect
		wantErr  bool
	}
	tests := map[string]testCase{
		"happy_path": {
			rollback: func() chan struct{} {
				rollback := make(chan struct{}, 1)
				return rollback
			},
			sut: func(rollback chan struct{}) Transaction {
				return Transaction{
					rollback: rollback,
				}
			},
			expect: expect{
				rollbackLength: 1,
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			rollback := tt.rollback()
			defer close(rollback)
			tx := tt.sut(rollback)

			if err := tx.Rollback(context.Background()); (err != nil) != tt.wantErr {
				t.Errorf("Rollback() error = %v, wantErr %v", err, tt.wantErr)
			}
			if len(rollback) != tt.expect.rollbackLength {
				t.Errorf("Rollback() rollback length = %v, want %v", len(rollback), tt.expect.rollbackLength)
			}
		})
	}
}

func TestManager_GetAndStart(t *testing.T) {
	type testCase struct {
		sut        func() *Transaction
		client     func(ctrl *gomock.Controller) client.Client
		action     func(sut *Transaction)
		beforeFunc func()
	}

	errTestManagerGetAndStart := errors.New("errTestManagerGetAndStart")

	sut := func() *Transaction {
		tx, _ := Manager().GetAndStart(context.Background())
		return tx.(*Transaction)
	}

	tests := map[string]testCase{
		"happy_path": {
			sut: sut,
			client: func(ctrl *gomock.Controller) client.Client {
				clt := mclient.NewMockClient(ctrl)
				clt.EXPECT().ExecuteTransaction(gomock.Any(), gomock.Any()).Return(nil, nil)
				return clt
			},
			action: func(sut *Transaction) {
				ctx := context.Background()
				errCh := sut.SubscribeError()
				err := sut.ExecuteStatement(ctx, NewStatement(make([]dynamotypes.ParameterizedStatement, 0)))
				if err != nil {
					t.Errorf("ExecuteStatement() returned error: %+v.", err)
					return
				}
				err = sut.Commit(ctx)
				if err != nil {
					t.Errorf("Commit() not returned error.")
				}
				err = <-errCh
				if err != nil {
					t.Errorf("Transaction has error: %+v.", err)
					return
				}
			},
		},
		"happy_path/rollback": {
			sut: sut,
			client: func(ctrl *gomock.Controller) client.Client {
				clt := mclient.NewMockClient(ctrl)
				clt.EXPECT().ExecuteTransaction(gomock.Any(), gomock.Any()).Return(nil, nil).Times(0)
				return clt
			},
			action: func(sut *Transaction) {
				ctx := context.Background()
				errCh := sut.SubscribeError()
				err := sut.ExecuteStatement(ctx, NewStatement(make([]dynamotypes.ParameterizedStatement, 0)))
				if err != nil {
					t.Errorf("ExecuteStatement() returned error: %+v", err)
				}
				err = sut.Rollback(ctx)
				if err != nil {
					t.Errorf("Rollback() returned error: %+v", err)
				}
				err = <-errCh
				if err != nil {
					t.Errorf("Transaction has error: %+v.", err)
					return
				}
			},
		},
		"unhappy_path/client_is_not_initialized": {
			sut: sut,
			client: func(ctrl *gomock.Controller) client.Client {
				clt := mclient.NewMockClient(ctrl)
				clt.EXPECT().ExecuteTransaction(gomock.Any(), gomock.Any()).Times(0)
				return clt
			},
			beforeFunc: func() {
				Invalidate()
			},
			action: func(sut *Transaction) {
				errCh := sut.SubscribeError()
				err := <-errCh
				if !errors.Is(err, ErrClientNotInitialized) {
					t.Errorf("Transaction has unexpected error: %+v.", err)
					return
				}
			},
		},
		"unhappy_path/dynamodb_transaction_returns_error": {
			sut: sut,
			client: func(ctrl *gomock.Controller) client.Client {
				clt := mclient.NewMockClient(ctrl)
				clt.EXPECT().ExecuteTransaction(gomock.Any(), gomock.Any()).Return(nil, errTestManagerGetAndStart).Times(1)
				return clt
			},
			action: func(sut *Transaction) {
				ctx := context.Background()
				errCh := sut.SubscribeError()
				err := sut.ExecuteStatement(ctx, NewStatement(make([]dynamotypes.ParameterizedStatement, 0)))
				if err != nil {
					t.Errorf("ExecuteStatement() returned error: %+v.", err)
					return
				}
				err = sut.Commit(ctx)
				if err != nil {
					t.Errorf("Commit() not returned error.")
				}
				err = <-errCh
				if !errors.Is(err, errTestManagerGetAndStart) {
					t.Errorf("Transaction has unexpected error: %+v.", err)
					return
				}
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			clt := tt.client(mockCtrl)
			Invalidate()
			Initialize(clt)
			if tt.beforeFunc != nil {
				tt.beforeFunc()
			}
			tx := tt.sut()
			tt.action(tx)
		})
	}
}
