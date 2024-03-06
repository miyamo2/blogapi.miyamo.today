package provider

import (
	"database/sql"
	"os"

	gwrapper "github.com/miyamo2/blogapi-core/db/gorm"
	_ "github.com/newrelic/go-agent/v3/integrations/nrpgx"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Dsn string

var Gorm = fx.Options(
	fx.Provide(func() Dsn {
		return Dsn(os.Getenv("COCKROACHDB_DSN"))
	}),
	fx.Provide(func(dsn Dsn) *gorm.Dialector {
		db, err := sql.Open("nrpgx", string(dsn))
		if err != nil {
			panic(err)
		}
		dialector := postgres.New(postgres.Config{
			Conn: db,
		})
		return &dialector
	}),
	fx.Invoke(func(dialector *gorm.Dialector) {
		gwrapper.InitializeDialector(dialector)
	}),
)
