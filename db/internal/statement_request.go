package internal

import (
	"context"
	"github.com/miyamo2/blogapi-core/db"
)

type StatementRequest struct {
	Statement db.Statement
	Ctx       context.Context
	ErrCh     chan error
}
