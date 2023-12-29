package internal

import (
	"github.com/miyamo2/blogapi-core/db"
)

type StatementRequest struct {
	Statement db.Statement
	ErrCh     chan error
}
