package db

import (
	"context"
)

// Transaction is a database transaction.
type Transaction interface {
	// ExecuteStatement executes statement in transaction.
	ExecuteStatement(ctx context.Context, statement Statement) error
	// Commit commits transaction.
	Commit(ctx context.Context) error
	// Rollback rollbacks transaction.
	Rollback(ctx context.Context) error
	// SubscribeError returns error channel.
	SubscribeError() <-chan error
}

// TransactionManager manages transaction.
type TransactionManager interface {
	// GetAndStart gets transaction and starts it.
	GetAndStart(ctx context.Context) (Transaction, error)
}
