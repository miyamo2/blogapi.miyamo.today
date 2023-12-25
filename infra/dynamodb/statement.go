package dynamodb

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	dynamotypes "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/cockroachdb/errors"
	"github.com/miyamo2/blogapi-core/infra"
	"github.com/miyamo2/blogapi-core/log"
)

var ErrAlreadyExecuted = errors.New("statement is already executed.")

// WithContext is an option to set context to Statement.
func WithContext(ctx context.Context) infra.ExecuteOption {
	return func(s infra.Statement) {
		switch v := s.(type) {
		case *Statement:
			v.ctx = ctx
		}
	}
}

// WithTransaction is an option to set transaction to Statement.
func WithTransaction(tx *dynamodb.ExecuteTransactionInput) infra.ExecuteOption {
	return func(s infra.Statement) {
		switch v := s.(type) {
		case *Statement:
			v.tx = tx
		}
	}
}

// Statement is a implementation of infra.Statement for dynamodb.
//
// Supports only Insert, Update, and Delete.
type Statement struct {
	ctx         context.Context
	tx          *dynamodb.ExecuteTransactionInput
	partiQLStmt []dynamotypes.ParameterizedStatement
	executed    bool
}

// zeroValueResult is a implementation of infra.StatementResult.
type zeroValueResult struct{}

func (r zeroValueResult) Get() interface{} {
	return nil
}

func (r zeroValueResult) Set(v interface{}) {
	return
}

func (s *Statement) Execute(opts ...infra.ExecuteOption) error {
	log.DefaultLogger().Info("start")
	defer log.DefaultLogger().Info("end")
	if s.executed {
		return ErrAlreadyExecuted
	}
	defer func() { s.executed = true }()
	for _, opt := range opts {
		opt(s)
	}
	ctx := s.ctx
	tx := s.tx
	if ctx == nil {
		ctx = context.Background()
	}

	if tx == nil {
		clt, err := Get()
		if err != nil {
			return errors.Wrap(err, "failed to get dynamodb connection.")
		}

		if _, err = clt.ExecuteTransaction(ctx, &dynamodb.ExecuteTransactionInput{
			TransactStatements: s.partiQLStmt,
		}); err != nil {
			return errors.Wrap(err, "failed to execute transaction.")
		}
	}
	return nil
}

// Result always returns zero value.
func (s *Statement) Result() infra.StatementResult {
	return zeroValueResult{}
}

func NewStatement(
	partiQLStmt []dynamotypes.ParameterizedStatement,
) infra.Statement {
	return &Statement{
		partiQLStmt: partiQLStmt,
	}
}
