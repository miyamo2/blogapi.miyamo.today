package provider

import (
	"database/sql"
	"os"

	"blogapi.miyamo.today/article-service/internal/infra/rdb/sqlc"
	"github.com/google/wire"
	_ "github.com/newrelic/go-agent/v3/integrations/nrpgx"
)

func SQLDB() *sql.DB {
	db, err := sql.Open("nrpgx", os.Getenv("COCKROACHDB_DSN"))
	if err != nil {
		panic(err) // because they are critical errors
	}
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
