//go:build wireinject

package di

import (
	"blogapi.miyamo.today/tag-service/internal/configs/di/provider"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

func GetEchoApp() *echo.Echo {
	wire.Build(
		provider.NewRelicSet,
		provider.SQLDBSet,
		provider.QueryServiceSet,
		provider.PresenterSet,
		provider.UsecaseSet,
		provider.TagServiceSet,
		EchoSet,
	)
	return nil
}
