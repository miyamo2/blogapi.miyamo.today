package gorm

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/miyamo2/altnrslog"

	"github.com/cockroachdb/errors"
	"github.com/miyamo2/blogapi.miyamo.today/core/db"
	"github.com/miyamo2/blogapi.miyamo.today/core/log"
	"github.com/miyamo2/blogapi.miyamo.today/core/util/duration"
	"github.com/newrelic/go-agent/v3/newrelic"
	"gorm.io/gorm"
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
	defer newrelic.FromContext(ctx).StartSegment("BlogAPICore: Gorm Statement Execute").End()
	dw := duration.Start()
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		logger = log.DefaultLogger()
	}
	logger.Info("BEGIN",
		slog.Group("parameters",
			slog.String("opts", fmt.Sprintf("%+v", opts))))
	defer logger.Info("END", slog.String("duration", dw.SDuration()))
	if s.executed {
		return ErrAlreadyExecuted
	}
	defer func() { s.executed = true }()
	for _, opt := range opts {
		opt(s)
	}
	tx := s.tx

	if tx == nil {
		conn, err := Get(ctx)
		if err != nil {
			return errors.Wrap(err, "failed to get gorm conn connection")
		}
		tx = conn.Begin()
		err = s.function(ctx, tx, s.out)
		if err != nil {
			tx.Rollback()
			return errors.Wrap(err, "failed to execute stmt")
		}
		tx.Commit()
		return nil
	}
	return s.function(ctx, tx, s.out)
}

func (s *Statement) Result() db.StatementResult {
	return s.out
}

func NewStatement(fn func(ctx context.Context, tx *gorm.DB, out db.StatementResult) error, out db.StatementResult) db.Statement {
	return &Statement{function: fn, out: out}
}
