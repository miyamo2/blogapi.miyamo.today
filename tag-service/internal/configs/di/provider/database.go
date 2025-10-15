package provider

import (
	"context"
	"database/sql"
	"os"

	"blogapi.miyamo.today/tag-service/internal/infra/rdb/sqlc"
	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/newrelic/go-agent/v3/integrations/nrpgx5"
)

func SQLDB() *sql.DB {
	cfg, err := pgxpool.ParseConfig(os.Getenv("COCKROACHDB_DSN"))
	if err != nil {
		panic(err) // because they are critical errors
	}

	cfg.ConnConfig.Tracer = nrpgx5.NewTracer()
	pool, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		panic(err) // because they are critical errors
	}
	db := stdlib.OpenDBFromPool(pool)
	err = db.Ping()
	if err != nil {
		panic(err) // because they are critical errors
	}
	return db
}

var SQLDBSet = wire.NewSet(
	SQLDB,
	wire.Bind(new(sqlc.DBTX), new(*sql.DB)),
)
