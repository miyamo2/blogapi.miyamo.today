package internal

import (
	"context"

	"blogapi.miyamo.today/core/db"
)

type StatementRequest struct {
	Statement db.Statement
	ErrCh     chan error
	Ctx       context.Context
}
