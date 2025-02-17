package gorm

import (
	"context"
	"fmt"
	"gorm.io/plugin/dbresolver"
	"log/slog"

	"github.com/miyamo2/altnrslog"

	"github.com/newrelic/go-agent/v3/newrelic"

	"blogapi.miyamo.today/core/db"
	"blogapi.miyamo.today/core/db/internal"
	"blogapi.miyamo.today/core/log"
	"github.com/cockroachdb/errors"
)

var _ db.Transaction = (*Transaction)(nil)

type Transaction struct {
	stmtQueue chan *internal.StatementRequest
	commit    chan struct{}
	rollback  chan struct{}
	errQueue  chan error
}

func (t *Transaction) process(ctx context.Context, dbSource string) {
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN")
	defer logger.InfoContext(ctx, "END")

	defer close(t.stmtQueue)
	defer close(t.commit)
	defer close(t.rollback)
	defer close(t.errQueue)

	conn, err := Get(ctx)
	if err != nil {
		err = errors.Wrap(err, "failed to get gorm connection")
		t.errQueue <- err
		return
	}
	if dbSource != "" {
		conn = conn.Clauses(dbresolver.Use(dbSource))
	}
	tx := conn.Begin()
TX:
	for {
		select {
		case nx := <-t.stmtQueue:
			err = nx.Statement.Execute(nx.Ctx, WithTransaction(tx))
			nx.ErrCh <- err
			if err != nil {
				t.errQueue <- err
			}
		case <-t.commit:
			tx.Commit()
			break TX
		case <-t.rollback:
			tx.Rollback()
			break TX
		}
	}
}

func (t *Transaction) SubscribeError() <-chan error {
	return t.errQueue
}

func (t *Transaction) ExecuteStatement(ctx context.Context, statement db.Statement) error {
	defer newrelic.FromContext(ctx).StartSegment("BlogAPICore: Gorm Transaction Execute Statement").End()
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN")
	// error will always be nil.
	defer logger.InfoContext(ctx, "END")
	errCh := make(chan error, 1)
	defer close(errCh)
	t.stmtQueue <- &internal.StatementRequest{
		Statement: statement,
		ErrCh:     errCh,
		Ctx:       ctx,
	}
	if err := <-errCh; err != nil {
		return err
	}

	return nil
}

func (t *Transaction) Commit(ctx context.Context) error {
	defer newrelic.FromContext(ctx).StartSegment("BlogAPICore: Gorm Transaction Commit").End()
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN")
	// error will always be nil.
	defer logger.InfoContext(ctx, "END",
		slog.Group("returns",
			slog.Any("error", nil)))
	t.commit <- struct{}{}

	return nil
}

func (t *Transaction) Rollback(ctx context.Context) error {
	defer newrelic.FromContext(ctx).StartSegment("BlogAPICore: Gorm Transaction Rollback").End()
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN")
	// error will always be nil.
	defer logger.InfoContext(ctx, "END",
		slog.Group("returns",
			slog.Any("error", nil)))
	t.rollback <- struct{}{}

	return nil
}

var _ db.TransactionManager = (*manager)(nil)

type manager struct {
}

func (m manager) GetAndStart(ctx context.Context, options ...db.GetAndStartOption) (db.Transaction, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("BlogAPICore: Gorm Get And Start Transaction").End()
	prop := db.GetAndStartProperty{}
	for _, opt := range options {
		opt(&prop)
	}
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN")
	stmtQueue := make(chan *internal.StatementRequest)
	t := &Transaction{
		stmtQueue: stmtQueue,
		commit:    make(chan struct{}, 1),
		rollback:  make(chan struct{}, 1),
		errQueue:  make(chan error),
	}
	// error will always be nil.
	defer logger.InfoContext(ctx, "END",
		slog.Group("returns",
			slog.String("conn.Transaction", fmt.Sprintf("%+v", *t)),
			slog.Any("error", nil)))

	pnrtx := nrtx.NewGoroutine()
	pctx, err := altnrslog.StoreToContext(
		newrelic.NewContext(ctx, pnrtx),
		log.New(log.WithAltNRSlogTransactionalHandler(nrtx.Application(), pnrtx)))
	if err != nil {
		pctx = ctx
	}
	go t.process(pctx, prop.Source)
	return t, nil
}

var m = manager{}

func Manager() manager {
	return m
}
