package dynamodb

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	dynamotypes "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/cockroachdb/errors"
	"github.com/miyamo2/blogapi-core/db"
	"github.com/miyamo2/blogapi-core/log"
	"log/slog"
)

var ErrAlreadyExecuted = errors.New("statement is already executed.")

// WithTransaction is an option to set transaction to Statement.
func WithTransaction(tx *dynamodb.ExecuteTransactionInput) db.ExecuteOption {
	return func(s db.Statement) {
		switch v := s.(type) {
		case *Statement:
			v.tx = tx
		}
	}
}

// Statement is a implementation of db.Statement for dynamodb.
//
// Supports only Insert, Update, and Delete.
type Statement struct {
	tx          *dynamodb.ExecuteTransactionInput
	partiQLStmt []dynamotypes.ParameterizedStatement
	executed    bool
}

// zeroValueResult is a implementation of db.StatementResult.
type zeroValueResult struct{}

func (r zeroValueResult) Get() interface{} {
	return nil
}

func (r zeroValueResult) Set(v interface{}) {
	return
}

func (s *Statement) Execute(ctx context.Context, opts ...db.ExecuteOption) error {
	log.DefaultLogger().Info("BEGIN",
		slog.Group("parameters",
			slog.String("ctx", fmt.Sprintf("%+v", ctx)),
			slog.String("opts", fmt.Sprintf("%+v", opts))))
	defer log.DefaultLogger().Info("END")
	if s.executed {
		return ErrAlreadyExecuted
	}
	defer func() { s.executed = true }()
	for _, opt := range opts {
		opt(s)
	}
	tx := s.tx

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
func (s *Statement) Result() db.StatementResult {
	return zeroValueResult{}
}

func NewStatement(
	partiQLStmt []dynamotypes.ParameterizedStatement,
) db.Statement {
	return &Statement{
		partiQLStmt: partiQLStmt,
	}
}
