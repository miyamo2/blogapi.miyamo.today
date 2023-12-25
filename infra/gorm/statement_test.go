package gorm

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cockroachdb/errors"
	"github.com/miyamo2/blogapi-core/infra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"regexp"
	"strconv"
	"testing"
)

func TestStatement_Execute(t *testing.T) {
	type Dummy struct {
		ID int `gorm:"primaryKey,autoincrement" json:"id,omitempty"`
	}

	errStatementTest := errors.New("error statement test")

	initializeConn := func() {
		sqlDB, mock, err := sqlmock.New()
		if err != nil {
			panic(err)
		}

		rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(
			`SELECT "id" FROM "dummies"`)).
			WillReturnRows(rows)
		mock.ExpectCommit()
		dialector := postgres.New(postgres.Config{
			Conn: sqlDB,
		})
		InvalidateDialector()
		InitializeDialector(&dialector)
	}

	type args struct {
		opts []infra.ExecuteOption
	}

	type want struct {
		err error
		out string
	}

	type testCase struct {
		statementResult func() *infra.SingleStatementResult[string]
		statement       func(result infra.StatementResult) infra.Statement
		args            args
		want            want
		wantErr         bool
		beforeFunc      func()
	}

	tests := map[string]testCase{
		"happy_path/with_context&transaction": {
			statementResult: func() *infra.SingleStatementResult[string] {
				return infra.NewSingleStatementResult[string]()
			},
			statement: func(result infra.StatementResult) infra.Statement {
				return NewStatement(
					func(ctx context.Context, tx *gorm.DB, out infra.StatementResult) error {
						out.Set("happy_path/with_context&transaction")
						return nil
					}, result)
			},
			args: args{
				opts: []infra.ExecuteOption{
					WithContext(context.Background()),
					WithTransaction(&gorm.DB{}),
				},
			},
			want: want{out: "happy_path/with_context&transaction", err: nil},
		},
		"happy_path/with_transaction": {
			statementResult: func() *infra.SingleStatementResult[string] {
				return infra.NewSingleStatementResult[string]()
			},
			statement: func(result infra.StatementResult) infra.Statement {
				return NewStatement(
					func(ctx context.Context, tx *gorm.DB, out infra.StatementResult) error {
						out.Set("happy_path/with_transaction")
						return nil
					}, result)
			},
			args: args{
				opts: []infra.ExecuteOption{
					WithTransaction(&gorm.DB{}),
				},
			},
			want: want{out: "happy_path/with_transaction", err: nil},
		},
		"happy_path/with_context": {
			statementResult: func() *infra.SingleStatementResult[string] {
				return infra.NewSingleStatementResult[string]()
			},
			statement: func(result infra.StatementResult) infra.Statement {
				return NewStatement(
					func(ctx context.Context, tx *gorm.DB, out infra.StatementResult) error {
						d := Dummy{}
						tx.Select("id").Find(&d)
						out.Set(strconv.Itoa(d.ID))
						return nil
					}, result)
			},
			args: args{
				opts: []infra.ExecuteOption{
					WithContext(context.Background()),
				},
			},
			want:       want{out: "1", err: nil},
			beforeFunc: initializeConn,
		},
		"happy_path/without_option": {
			statementResult: func() *infra.SingleStatementResult[string] {
				return infra.NewSingleStatementResult[string]()
			},
			statement: func(result infra.StatementResult) infra.Statement {
				return NewStatement(
					func(ctx context.Context, tx *gorm.DB, out infra.StatementResult) error {
						d := Dummy{}
						tx.Select("id").Find(&d)
						out.Set(strconv.Itoa(d.ID))
						return nil
					}, result)
			},
			args: args{
				opts: []infra.ExecuteOption{},
			},
			want:       want{out: "1", err: nil},
			beforeFunc: initializeConn,
		},
		"unhappy_path/statement_returned_error": {
			statementResult: func() *infra.SingleStatementResult[string] {
				return infra.NewSingleStatementResult[string]()
			},
			statement: func(result infra.StatementResult) infra.Statement {
				return NewStatement(
					func(ctx context.Context, tx *gorm.DB, out infra.StatementResult) error {
						d := Dummy{}
						tx.Select("id").Find(&d)
						return errStatementTest
					}, result)
			},
			args: args{
				opts: []infra.ExecuteOption{},
			},
			want:    want{out: "", err: errStatementTest},
			wantErr: true,
			beforeFunc: func() {
				sqlDB, mock, err := sqlmock.New()
				if err != nil {
					panic(err)
				}

				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT "id" FROM "dummies"`)).
					WillReturnRows(rows)
				mock.ExpectRollback()
				dialector := postgres.New(postgres.Config{
					Conn: sqlDB,
				})
				InvalidateDialector()
				InitializeDialector(&dialector)
			},
		},
		"unhappy_path/dialector_is_nil": {
			statementResult: func() *infra.SingleStatementResult[string] {
				return infra.NewSingleStatementResult[string]()
			},
			statement: func(result infra.StatementResult) infra.Statement {
				return NewStatement(
					func(ctx context.Context, tx *gorm.DB, out infra.StatementResult) error {
						out.Set("unhappy_path/client_is_nil")
						return nil
					}, result)
			},
			args: args{
				opts: []infra.ExecuteOption{},
			},
			want:    want{out: "", err: ErrDialectorNotInitialized},
			wantErr: true,
			beforeFunc: func() {
				InvalidateDialector()
			},
		},
		"unhappy_path/statement_is_already_executed": {
			statementResult: func() *infra.SingleStatementResult[string] {
				return infra.NewSingleStatementResult[string]()
			},
			statement: func(result infra.StatementResult) infra.Statement {
				stmt := NewStatement(
					func(ctx context.Context, tx *gorm.DB, out infra.StatementResult) error {
						out.Set("unhappy_path/statement_is_already_executed")
						return nil
					}, result)
				stmt.(*Statement).executed = true
				return stmt
			},
			args: args{
				opts: []infra.ExecuteOption{},
			},
			want:    want{out: "", err: ErrAlreadyExecuted},
			wantErr: true,
			beforeFunc: func() {
				sqlDB, mock, err := sqlmock.New()
				if err != nil {
					panic(err)
				}

				mock.ExpectBegin()
				dialector := postgres.New(postgres.Config{
					Conn: sqlDB,
				})
				InvalidateDialector()
				InitializeDialector(&dialector)
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if beforeFunc := tt.beforeFunc; beforeFunc != nil {
				beforeFunc()
			}
			result := tt.statementResult()
			stmt := tt.statement(result)
			err := stmt.Execute(tt.args.opts...)
			if tt.wantErr {
				if !errors.Is(err, tt.want.err) {
					t.Errorf("Execute() error = %+v, want %+v", err, tt.want.err)
					return
				}
			} else if err != nil {
				t.Errorf("Execute() returned error: %+v", err)
				return
			}
			if got := result.StrictGet(); got != tt.want.out {
				t.Errorf("Execute() got = %v, want %v", got, tt.want.out)
				return
			}
			if got := stmt.Result().Get(); got != tt.want.out {
				t.Errorf("Execute() got = %v, want %v", got, tt.want.out)
				return
			}
		})
	}
}
