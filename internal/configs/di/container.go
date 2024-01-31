package di

import (
	"github.com/miyamo2/blogapi/internal/configs/di/provider"
	"go.uber.org/fx"
)

var Container = fx.Options(
	provider.NewRelic,
	provider.Logger,
	provider.Grpc,
	provider.Presenter,
	provider.Usecase,
	provider.Gqlgen,
	provider.Echo,
)
