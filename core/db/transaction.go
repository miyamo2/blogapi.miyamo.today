//go:generate mockgen -source=$GOFILE -destination ../../tag-service/internal/mock/core/db/mock_db.go -package mock_db
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

// GetAndStartProperty is a property for GetAndStart.
type GetAndStartProperty struct {
	Source string
}

// GetAndStartOption is an option for GetAndStart.
type GetAndStartOption func(*GetAndStartProperty)

// GetAndStartWithDBSource specify db source for GetAndStart.
func GetAndStartWithDBSource(source string) GetAndStartOption {
	return func(p *GetAndStartProperty) {
		p.Source = source
	}
}

// TransactionManager manages transaction.
type TransactionManager interface {
	// GetAndStart gets transaction and starts it.
	GetAndStart(ctx context.Context, options ...GetAndStartOption) (Transaction, error)
}
