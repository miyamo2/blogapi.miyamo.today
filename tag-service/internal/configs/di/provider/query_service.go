package provider

import (
	"context"
	"time"

	"blogapi.miyamo.today/tag-service/internal/app/usecase/query"
	"blogapi.miyamo.today/tag-service/internal/infra/rdb"
	"github.com/google/wire"
)

// compatibility check
var _ query.Queries = (*rdb.Queries)(nil)

func QueryService(tx rdb.DBTX) *rdb.Queries {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	q, err := rdb.Prepare(ctx, tx)
	if err != nil {
		panic(err)
	}
	return q
}

var QueryServiceSet = wire.NewSet(
	QueryService,
	wire.Bind(new(query.Queries), new(*rdb.Queries)),
)
