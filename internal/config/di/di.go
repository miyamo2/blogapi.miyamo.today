package di

import (
	"github.com/miyamo2/blogapi-tag-service/internal/config/di/provider"
	"go.uber.org/fx"
)

var Container = fx.Options(
	provider.Gorm,
	provider.TxManager,
	provider.QueryService,
	provider.Usecase,
	provider.Presenter,
	provider.Controller,
	provider.Tcp,
)
