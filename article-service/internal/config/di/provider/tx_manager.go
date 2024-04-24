package provider

import (
	"github.com/miyamo2/api.miyamo.today/core/db"
	"github.com/miyamo2/api.miyamo.today/core/db/gorm"
	"go.uber.org/fx"
)

var TxManager = fx.Options(
	fx.Provide(
		fx.Annotate(
			gorm.Manager,
			fx.As(new(db.TransactionManager)),
		),
	),
)
