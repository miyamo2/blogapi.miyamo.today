package di

import (
	"github.com/miyamo2/blogapi.miyamo.today/tag-service/internal/configs/di/provider"
	"go.uber.org/fx"
)

var Container = fx.Options(
	provider.NewRelic,
	provider.Logger,
	provider.Gorm,
	provider.TxManager,
	provider.QueryService,
	provider.Usecase,
	provider.Presenter,
	provider.Controller,
	provider.Tcp,
)
