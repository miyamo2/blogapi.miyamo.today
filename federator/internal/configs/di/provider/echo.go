package provider

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/goccy/go-json"
	"github.com/labstack/echo/v4"
	"github.com/miyamo2/blogapi.miyamo.today/core/echo/s11n"
	nrecho "github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"
	"go.uber.org/fx"
)

var Echo = fx.Options(
	fx.Provide(func(srv *handler.Server, nr *newrelic.Application) *echo.Echo {
		slog.Info("creating echo server")
		e := echo.New()
		e.Add(http.MethodPost, "/query", echo.WrapHandler(srv), nrecho.Middleware(nr))
		e.GET("/playground", echo.WrapHandler(playground.Handler("GraphQL playground", "/query")))
		e.GET("/health", func(c echo.Context) error {
			return c.String(200, "ok")
		})
		e.HTTPErrorHandler = func(err error, c echo.Context) {
			req := c.Request()
			ctx := req.Context()
			nrtx := newrelic.FromContext(ctx)
			if he, ok := err.(*echo.HTTPError); ok {
				code := he.Code
				switch code {
				case http.StatusNotFound:
					nrtx.NoticeError(nrpkgerrors.Wrap(err))
					slog.ErrorContext(c.Request().Context(),
						fmt.Sprintf("request: %v %v", req.Method, req.URL),
						slog.String("error", err.Error()))
				}
			}
			e.DefaultHTTPErrorHandler(err, c)
		}
		e.JSONSerializer = &s11n.JSONSerializer[*json.Encoder, *json.Decoder]{
			Encoder: json.NewEncoder,
			Decoder: json.NewDecoder,
		}
		slog.Info("echo server created")
		return e
	}),
)