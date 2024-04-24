package dynamodb

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	dynamotypes "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/cockroachdb/errors"
	"github.com/miyamo2/api.miyamo.today/core/db"
	"github.com/miyamo2/api.miyamo.today/core/db/dynamodb/client"
	mclient "github.com/miyamo2/api.miyamo.today/core/internal/mock/infra/dynamodb/client"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestStatement_Execute(t *testing.T) {
	errStatementTest := errors.New("error statement test")

	type args struct {
		opts []db.ExecuteOption
	}

	type want struct {
		err error
	}

	type testCase struct {
		client     func(ctrl *gomock.Controller) client.Client
		statement  func() db.Statement
		args       args
		want       want
		wantErr    bool
		beforeFunc func()
	}

	tests := map[string]testCase{
		"happy_path/with_transaction": {
			client: func(ctrl *gomock.Controller) client.Client {
				clt := mclient.NewMockClient(ctrl)
				clt.EXPECT().ExecuteTransaction(gomock.Any(), gomock.Any()).Return(nil, nil).Times(0)
				return clt
			},
			statement: func() db.Statement {
				stmt := NewStatement(make([]dynamotypes.ParameterizedStatement, 0))
				return stmt
			},
			args: args{
				opts: []db.ExecuteOption{
					WithTransaction(&dynamodb.ExecuteTransactionInput{}),
				},
			},
			want: want{err: nil},
		},
		"happy_path/without_option": {
			client: func(ctrl *gomock.Controller) client.Client {
				clt := mclient.NewMockClient(ctrl)
				clt.EXPECT().ExecuteTransaction(gomock.Any(), gomock.Any()).Return(nil, nil).Times(1)
				return clt
			},
			statement: func() db.Statement {
				stmt := NewStatement(make([]dynamotypes.ParameterizedStatement, 0))
				return stmt
			},
			args: args{
				opts: []db.ExecuteOption{},
			},
			want: want{err: nil},
		},
		"unhappy_path/client_is_nil": {
			client: func(ctrl *gomock.Controller) client.Client {
				return nil
			},
			statement: func() db.Statement {
				stmt := NewStatement(make([]dynamotypes.ParameterizedStatement, 0))
				return stmt
			},
			args: args{
				opts: []db.ExecuteOption{},
			},
			want:    want{err: ErrClientNotInitialized},
			wantErr: true,
			beforeFunc: func() {
				Invalidate()
			},
		},
		"unhappy_path/statement_is_already_executed": {
			client: func(ctrl *gomock.Controller) client.Client {
				clt := mclient.NewMockClient(ctrl)
				clt.EXPECT().ExecuteTransaction(gomock.Any(), gomock.Any()).Return(nil, nil).Times(0)
				return clt
			},
			statement: func() db.Statement {
				stmt := NewStatement(make([]dynamotypes.ParameterizedStatement, 0))
				stmt.(*Statement).executed = true
				return stmt
			},
			args: args{
				opts: []db.ExecuteOption{},
			},
			want:    want{err: ErrAlreadyExecuted},
			wantErr: true,
		},
		"unhappy_path/transaction_returns_error": {
			client: func(ctrl *gomock.Controller) client.Client {
				clt := mclient.NewMockClient(ctrl)
				clt.EXPECT().ExecuteTransaction(gomock.Any(), gomock.Any()).Return(nil, errStatementTest).Times(1)
				return clt
			},
			statement: func() db.Statement {
				stmt := NewStatement(make([]dynamotypes.ParameterizedStatement, 0))
				return stmt
			},
			args: args{
				opts: []db.ExecuteOption{},
			},
			want:    want{err: errStatementTest},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			clt := tt.client(mockCtrl)
			Invalidate()
			Initialize(clt)
			if beforeFunc := tt.beforeFunc; beforeFunc != nil {
				beforeFunc()
			}
			stmt := tt.statement()
			err := stmt.Execute(context.Background(), tt.args.opts...)
			if tt.wantErr {
				if !errors.Is(err, tt.want.err) {
					t.Errorf("Execute() error = %+v, want %+v", err, tt.want.err)
					return
				}
			} else if err != nil {
				t.Errorf("Execute() returned error: %+v", err)
				return
			}
			res := stmt.Result()
			switch tp := res.(type) {
			case zeroValueResult:
				// Do Nothing
			default:
				t.Errorf("StatementResult.type = %v", tp)
				return
			}
			if res.Get() != nil {
				t.Errorf("zeroValueResult Get() returns = %v, want = %v", res.Get(), nil)
				return
			}
		})
	}
}
