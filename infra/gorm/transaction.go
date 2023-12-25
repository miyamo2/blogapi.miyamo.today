package gorm

import (
	"context"
	"log/slog"

	"github.com/cockroachdb/errors"
	"github.com/miyamo2/blogapi-core/infra"
	"github.com/miyamo2/blogapi-core/infra/internal"
	"github.com/miyamo2/blogapi-core/log"
	"gorm.io/gorm"
)

var _ infra.Transaction = (*Transaction)(nil)

type Transaction struct {
	stmtQueue chan *internal.StatementRequest
	commit    chan struct{}
	rollback  chan struct{}
	errQueue  chan error
}

func (t *Transaction) start(ctx context.Context) {
	log.DefaultLogger().InfoContext(ctx, "start", slog.Group("parameters", slog.Any("ctx", ctx)))
	defer log.DefaultLogger().InfoContext(ctx, "end")

	defer close(t.stmtQueue)
	defer close(t.commit)
	defer close(t.rollback)

	db, err := Get()
	if err != nil {
		err = errors.Wrap(err, "failed to get gorm connection")
		t.errQueue <- err
		return
	}
	err = db.Transaction(func(tx *gorm.DB) error {
		for {
			select {
			case nx := <-t.stmtQueue:
				err = nx.Statement.Execute(WithTransaction(db), WithContext(ctx))
				if err != nil {
					nx.ErrCh <- err
					t.errQueue <- err
					return err
				}
				nx.ErrCh <- nil
			case <-t.commit:
				db.Commit()
				return nil
			case <-t.rollback:
				db.Rollback()
				return nil
			}
		}
	})
	if err != nil {
		err = errors.Wrap(err, "failed to run gorm transaction")
		t.errQueue <- err
	}
	t.errQueue <- nil
}

func (t *Transaction) SubscribeError() <-chan error {
	return t.errQueue
}

func (t *Transaction) ExecuteStatement(ctx context.Context, statement infra.Statement) error {
	log.DefaultLogger().InfoContext(ctx, "start", slog.Group("parameters", slog.Any("ctx", ctx), slog.Any("statement", statement)))
	// error will always be nil.
	defer log.DefaultLogger().InfoContext(ctx, "end", slog.Group("returns", slog.Any("error", nil)))
	errCh := make(chan error, 1)
	defer close(errCh)
	t.stmtQueue <- &internal.StatementRequest{
		Statement: statement,
		ErrCh:     errCh,
	}
	if err := <-errCh; err != nil {
		return err
	}

	return nil
}

func (t *Transaction) Commit(ctx context.Context) error {
	log.DefaultLogger().InfoContext(ctx, "start", slog.Group("parameters", slog.Any("ctx", ctx)))
	// error will always be nil.
	defer log.DefaultLogger().InfoContext(ctx, "end", slog.Group("returns", slog.Any("error", nil)))
	t.commit <- struct{}{}

	return nil
}

func (t *Transaction) Rollback(ctx context.Context) error {
	log.DefaultLogger().InfoContext(ctx, "start", slog.Group("parameters", slog.Any("ctx", ctx)))
	// error will always be nil.
	defer log.DefaultLogger().InfoContext(ctx, "end", slog.Group("returns", slog.Any("error", nil)))
	t.rollback <- struct{}{}

	return nil
}

var _ infra.TransactionManager = (*manager)(nil)

type manager struct {
}

func (m manager) GetAndStart(ctx context.Context) (infra.Transaction, error) {
	log.DefaultLogger().InfoContext(ctx, "start", slog.Group("parameters", slog.Any("ctx", ctx)))
	stmtQueue := make(chan *internal.StatementRequest)
	t := &Transaction{
		stmtQueue: stmtQueue,
		commit:    make(chan struct{}, 1),
		rollback:  make(chan struct{}, 1),
		errQueue:  make(chan error),
	}
	// error will always be nil.
	defer log.DefaultLogger().InfoContext(ctx, "end", slog.Group("returns", slog.Any("infra.Transaction", *t), slog.Any("error", nil)))
	go t.start(ctx)
	return t, nil
}

var m = manager{}

func Manager() manager {
	return m
}
