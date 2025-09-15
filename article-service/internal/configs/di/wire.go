//go:build wireinject

package di

import (
	"blogapi.miyamo.today/article-service/internal/configs/di/provider"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

func GetEchoApp() *echo.Echo {
	wire.Build(
		provider.NewRelicSet,
		provider.SQLDBSet,
		provider.QuerySet,
		provider.PresenterSet,
		provider.UsecaseSet,
		provider.ArticleServiceSet,
		EchoSet,
	)
	return nil
}
