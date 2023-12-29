package gorm

import (
	"context"
	"fmt"
	"github.com/cockroachdb/errors"
	"github.com/miyamo2/blogapi-core/db"
	"github.com/miyamo2/blogapi-core/log"
	"gorm.io/gorm"
	"log/slog"
)

var ErrAlreadyExecuted = errors.New("statement is already executed.")

// WithTransaction is an option to set transaction to Statement.
func WithTransaction(tx *gorm.DB) db.ExecuteOption {
	return func(s db.Statement) {
		switch v := s.(type) {
		case *Statement:
			v.tx = tx
		}
	}
}

// Statement is a implementation of db.Statement for gorm.
type Statement struct {
	tx       *gorm.DB
	out      db.StatementResult
	function func(ctx context.Context, db *gorm.DB, out db.StatementResult) error
	executed bool
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

func (s *Statement) Result() db.StatementResult {
	return s.out
}

func NewStatement(fn func(ctx context.Context, tx *gorm.DB, out db.StatementResult) error, out db.StatementResult) db.Statement {
	return &Statement{function: fn, out: out}
}
