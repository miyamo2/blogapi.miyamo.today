package internal

import (
	"context"

	"github.com/miyamo2/blogapi.miyamo.today/core/db"
)

type StatementRequest struct {
	Statement db.Statement
	ErrCh     chan error
	Ctx       context.Context
}
