package dynamodb

import (
	"context"
	"fmt"
	"github.com/miyamo2/altnrslog"
	"log/slog"
	"sync"

	"github.com/miyamo2/blogapi-core/util/duration"
	"github.com/newrelic/go-agent/v3/newrelic"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/cockroachdb/errors"
	"github.com/miyamo2/blogapi-core/db"
	"github.com/miyamo2/blogapi-core/db/internal"
	"github.com/miyamo2/blogapi-core/log"
)

var _ db.Transaction = (*Transaction)(nil)

type Transaction struct {
	stmtQueue chan *internal.StatementRequest
	commit    chan struct{}
	rollback  chan struct{}
	errQueue  chan error
}

func (t *Transaction) start(ctx context.Context) {
	nrtx := newrelic.FromContext(ctx).NewGoroutine()
	defer nrtx.StartSegment("BlogAPICore: DynamoDB Transaction Start").End()
	ctx = newrelic.NewContext(ctx, nrtx)
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

	clt, err := Get()
	mu := &sync.Mutex{}
	stmts := make([]*Statement, 0)
	tx := &dynamodb.ExecuteTransactionInput{}
	if err != nil {
		err = errors.Wrap(err, "failed to get gorm connection")
		t.errQueue <- err
		return
	}
	for {
		select {
		case nx := <-t.stmtQueue:
			stmt := nx.Statement.(*Statement)
			go func(stmt *Statement) {
				mu.Lock()
				stmts = append(stmts, stmt)
				tx.TransactStatements = append(tx.TransactStatements, stmt.partiQLStmt...)
				mu.Unlock()
			}(stmt)
		case <-t.commit:
			_, err := clt.ExecuteTransaction(ctx, tx)
			slog.Info("A")
			t.errQueue <- err
			return
		case <-t.rollback:
			t.errQueue <- nil
			return
		}
	}
}

func (t *Transaction) SubscribeError() <-chan error {
	return t.errQueue
}

func (t *Transaction) ExecuteStatement(ctx context.Context, statement db.Statement) error {
	defer newrelic.FromContext(ctx).StartSegment("BlogAPICore: DynamoDB Transaction Execute Statement").End()
	dw := duration.Start()
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters",
			slog.String("statement", fmt.Sprintf("%+v", statement))))
	// error will always be nil.
	defer logger.InfoContext(ctx, "END",
		slog.String("duration", dw.SDuration()),
		slog.Group("returns",
			slog.Any("error", nil)))
	t.stmtQueue <- &internal.StatementRequest{
		Statement: statement,
	}

	return nil
}

func (t *Transaction) Commit(ctx context.Context) error {
	defer newrelic.FromContext(ctx).StartSegment("BlogAPICore: DynamoDB Transaction Commit").End()
	dw := duration.Start()
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN")
	// error will always be nil.
	defer logger.InfoContext(ctx, "END",
		slog.String("duration", dw.SDuration()),
		slog.Group("returns", slog.Any("error", nil)))
	t.commit <- struct{}{}

	return nil
}

// Deprecated: DynamoDB Not Support Rollback.
//
// DynamoDB's Transaction are submit as a single all-or-nothing.
func (t *Transaction) Rollback(ctx context.Context) error {
	defer newrelic.FromContext(ctx).StartSegment("BlogAPICore: DynamoDB Transaction Rollback").End()
	dw := duration.Start()
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN")
	// error will always be nil.
	defer logger.InfoContext(ctx, "END",
		slog.String("duration", dw.SDuration()),
		slog.Group("returns",
			slog.Any("error", nil)))
	t.rollback <- struct{}{}

	return nil
}

var _ db.TransactionManager = (*manager)(nil)

type manager struct {
}

func (m manager) GetAndStart(ctx context.Context) (db.Transaction, error) {
	defer newrelic.FromContext(ctx).StartSegment("BlogAPICore: DynamoDB Get And Start Transaction").End()
	dw := duration.Start()
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
		slog.String("duration", dw.SDuration()),
		slog.Group("returns",
			slog.String("db.Transaction", fmt.Sprintf("%+v", *t)),
			slog.Any("error", nil)))

	go t.start(ctx)
	return t, nil
}

var m = manager{}

func Manager() manager {
	return m
}
