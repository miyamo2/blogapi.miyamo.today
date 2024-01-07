package provider

import (
	gwrapper "github.com/miyamo2/blogapi-core/db/gorm"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type Dsn string

var Gorm = fx.Options(
	fx.Provide(func() Dsn {
		return Dsn(os.Getenv("COCKROACHDB_DSN"))
	}),
	fx.Provide(func(dsn Dsn) *gorm.Dialector {
		dialector := postgres.Open(string(dsn))
		return &dialector
	}),
	fx.Invoke(func(dialector *gorm.Dialector) {
		gwrapper.InitializeDialector(dialector)
	}),
)
