package provider

import (
	"database/sql"
	"os"

	"blogapi.miyamo.today/tag-service/internal/infra/rdb"
	"github.com/google/wire"
	_ "github.com/newrelic/go-agent/v3/integrations/nrpgx"
)

func SQLDB() *sql.DB {
	db, err := sql.Open("nrpgx", os.Getenv("COCKROACHDB_DSN"))
	if err != nil {
		panic(err)
	}
	return db
}

var SQLDBSet = wire.NewSet(
	SQLDB,
	wire.Bind(new(rdb.DBTX), new(*sql.DB)),
)
