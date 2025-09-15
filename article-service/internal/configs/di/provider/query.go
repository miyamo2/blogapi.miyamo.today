package provider

import (
	"context"
	"time"

	"blogapi.miyamo.today/article-service/internal/app/usecase/query"
	"blogapi.miyamo.today/article-service/internal/infra/rdb/sqlc"
	"github.com/google/wire"
)

// compatibility check
var _ query.Queries = (*sqlc.Queries)(nil)

func QueryService(tx sqlc.DBTX) *sqlc.Queries {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	q, err := sqlc.Prepare(ctx, tx)
	if err != nil {
		panic(err)
	}
	return q
}

var QuerySet = wire.NewSet(
	QueryService,
	wire.Bind(new(query.Queries), new(*sqlc.Queries)),
)
