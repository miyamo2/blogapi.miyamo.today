package provider

import (
	"database/sql"
	"github.com/google/wire"
	"os"

	_ "github.com/newrelic/go-agent/v3/integrations/nrpgx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GormDialector() *gorm.Dialector {
	db, err := sql.Open("nrpgx", os.Getenv("COCKROACHDB_DSN"))
	if err != nil {
		panic(err)
	}
	dialector := postgres.New(postgres.Config{
		Conn: db,
	})
	return &dialector
}

var GormSet = wire.NewSet(GormDialector)
