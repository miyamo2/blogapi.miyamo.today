package gorm

import (
	"context"
	"github.com/cockroachdb/errors"
	"github.com/miyamo2/blogapi-core/infra"
	"github.com/miyamo2/blogapi-core/log"
	"gorm.io/gorm"
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
func WithTransaction(tx *gorm.DB) infra.ExecuteOption {
	return func(s infra.Statement) {
		switch v := s.(type) {
		case *Statement:
			v.tx = tx
		}
	}
}

// Statement is a implementation of infra.Statement for gorm.
type Statement struct {
	ctx      context.Context
	tx       *gorm.DB
	out      infra.StatementResult
	function func(ctx context.Context, db *gorm.DB, out infra.StatementResult) error
	executed bool
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
		db, err := Get()
		if err != nil {
			return errors.Wrap(err, "failed to get gorm db connection")
		}
		err = db.Transaction(func(tx *gorm.DB) error {
			err := s.function(ctx, tx, s.out)
			if err != nil {
				return errors.Wrap(err, "failed to execute stmt")
			}
			return nil
		})
		return err
	}
	return s.function(ctx, tx, s.out)
}

func (s *Statement) Result() infra.StatementResult {
	return s.out
}

func NewStatement(fn func(ctx context.Context, tx *gorm.DB, out infra.StatementResult) error, out infra.StatementResult) infra.Statement {
	return &Statement{function: fn, out: out}
}
