package provider

import (
	"log/slog"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	nrecho "github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	"github.com/newrelic/go-agent/v3/newrelic"
	"go.uber.org/fx"
)

var Echo = fx.Options(
	fx.Provide(func(srv *handler.Server, nr *newrelic.Application) *echo.Echo {
		slog.Info("creating echo server")
		e := echo.New()
		e.POST("/query", echo.WrapHandler(srv))
		e.GET("/playground", echo.WrapHandler(playground.Handler("GraphQL playground", "/query")))
		e.Use(nrecho.Middleware(nr))
		slog.Info("echo server created")
		return e
	}),
)
