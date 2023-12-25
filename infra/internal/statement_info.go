package internal

import (
	"github.com/miyamo2/blogapi-core/infra"
)

type StatementRequest struct {
	Statement infra.Statement
	ErrCh     chan error
}
