package gorm

import (
	"context"
	"testing"

	"blogapi.miyamo.today/core/db"
	"blogapi.miyamo.today/core/db/internal"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cockroachdb/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	type expect struct {
		err error
	}
	type testCase struct {
		stmtQueue      func() chan *internal.StatementRequest
		sut            func(commit chan *internal.StatementRequest) Transaction
		mockSubscriber func(chan *internal.StatementRequest)
		expect         expect
		wantErr        bool
	}
	testErr := errors.New("test error")

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
				nx.ErrCh <- nil
			},
			expect: expect{
				err: nil,
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
				nx.ErrCh <- testErr
			},
			expect: expect{
				err: testErr,
			},
			wantErr: true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			stmtQueue := tt.stmtQueue()
			defer close(stmtQueue)
			tx := tt.sut(stmtQueue)
			go tt.mockSubscriber(stmtQueue)

			err := tx.ExecuteStatement(context.Background(), &Statement{})
			if (err != nil) == tt.wantErr {
				if !errors.Is(err, tt.expect.err) {
					t.Errorf("ExecuteStatement() error = %v, want %v", err, tt.expect.err)
				}
			} else {
				t.Errorf("ExecuteStatement() error = %v, wantErr %v", err, tt.wantErr)
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
		action     func(t *testing.T, sut *Transaction)
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
			action: func(t *testing.T, sut *Transaction) {
				ctx := context.Background()
				stmt := NewStatement(
					func(ctx context.Context, tx *gorm.DB, out db.StatementResult) error {
						out.Set("some string")
						return nil
					}, db.NewSingleStatementResult[string]())
				err := sut.ExecuteStatement(ctx, stmt)
				if err != nil {
					t.Errorf("ExecuteStatement() error = %v", err)
					return
				}
				if stmt.Result().Get() != "some string" {
					t.Errorf("StatementResult.Get() returned = %v, expected = %v", stmt.Result().Get(), "some string")
					return
				}
				err = sut.Commit(ctx)
				if err != nil {
					t.Errorf("Commit() error = %v", err)
					return
				}
			},
			beforeFunc: func() {
				sqlDB, mock, err := sqlmock.New()
				if err != nil {
					panic(err)
				}
				mock.ExpectBegin()
				mock.ExpectCommit()
				dialector := postgres.New(postgres.Config{
					Conn: sqlDB,
				})
				Invalidate()
				InvalidateDialector()
				InitializeDialector(&dialector)
			},
		},
		"happy_path/rollback": {
			sut: sut,
			action: func(t *testing.T, sut *Transaction) {
				ctx := context.Background()
				stmt := NewStatement(
					func(ctx context.Context, tx *gorm.DB, out db.StatementResult) error {
						out.Set("some string")
						return nil
					}, db.NewSingleStatementResult[string]())
				err := sut.ExecuteStatement(ctx, stmt)
				if err != nil {
					t.Errorf("ExecuteStatement() error = %v", err)
					return
				}
				err = sut.Rollback(ctx)
				if err != nil {
					t.Errorf("Commit() error = %v", err)
					return
				}
			},
			beforeFunc: func() {
				sqlDB, mock, err := sqlmock.New()
				if err != nil {
					panic(err)
				}

				mock.ExpectBegin()
				mock.ExpectRollback()
				dialector := postgres.New(postgres.Config{
					Conn: sqlDB,
				})
				Invalidate()
				InvalidateDialector()
				InitializeDialector(&dialector)
			},
		},
		"unhappy_path/gorm_is_not_initialized": {
			sut: sut,
			action: func(t *testing.T, sut *Transaction) {
				errQueue := sut.SubscribeError()
				err := <-errQueue
				if !errors.Is(err, ErrDialectorNotInitialized) {
					t.Errorf("Error() want = %v, error= %v", ErrDialectorNotInitialized, err)
					return
				}
			},
			beforeFunc: func() {
				Invalidate()
				InvalidateDialector()
			},
		},
		"unhappy_path/statement_returns_error": {
			sut: sut,
			action: func(t *testing.T, sut *Transaction) {
				ctx := context.Background()
				stmt := NewStatement(
					func(ctx context.Context, tx *gorm.DB, out db.StatementResult) error {
						out.Set("")
						return errTestManagerGetAndStart
					}, db.NewSingleStatementResult[string]())
				errQueue := sut.SubscribeError()
				err := sut.ExecuteStatement(ctx, stmt)
				if !errors.Is(err, errTestManagerGetAndStart) {
					t.Errorf("ExecuteStatement() want = %v, error = %v", errTestManagerGetAndStart, err)
					return
				}
				err = <-errQueue
				if !errors.Is(err, errTestManagerGetAndStart) {
					t.Errorf("SubscribeError() want = %v, error = %v", errTestManagerGetAndStart, err)
					return
				}
			},
			beforeFunc: func() {
				sqlDB, mock, err := sqlmock.New()
				if err != nil {
					panic(err)
				}
				mock.ExpectBegin()
				dialector := postgres.New(postgres.Config{
					Conn: sqlDB,
				})
				Invalidate()
				InvalidateDialector()
				InitializeDialector(&dialector)
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if tt.beforeFunc != nil {
				tt.beforeFunc()
			}
			tx := tt.sut()
			tt.action(t, tx)
		})
	}
}
