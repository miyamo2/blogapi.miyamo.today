package internal

import (
	"context"
	"github.com/miyamo2/api.miyamo.today/core/db"
)

type StatementRequest struct {
	Statement db.Statement
	ErrCh     chan error
	Ctx       context.Context
}
