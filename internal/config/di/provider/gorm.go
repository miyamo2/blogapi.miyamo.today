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

type Dns string

var Gorm = fx.Options(
	fx.Provide(func() Dns {
		return Dns(os.Getenv("COCKROACHDB_DNS"))
	}),
	fx.Provide(func(dns Dns) *gorm.Dialector {
		db, err := sql.Open("nrpgx", string(dns))
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
