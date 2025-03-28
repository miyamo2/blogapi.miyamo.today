package provider

import (
	"blogapi.miyamo.today/core/echo/middlewares"
	"fmt"
	"github.com/google/wire"
	"github.com/miyamo2/altnrslog"
	"log/slog"
	"net/http"

	"blogapi.miyamo.today/core/echo/s11n"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/goccy/go-json"
	"github.com/labstack/echo/v4"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func Echo(srv *handler.Server, nr *newrelic.Application, verifier middlewares.Verifier) *echo.Echo {
	slog.Info("creating echo server")
	e := echo.New()

	authMiddleware := middlewares.Auth(verifier)
	e.Add(http.MethodPost, "/query", echo.WrapHandler(srv), nrecho.Middleware(nr), authMiddleware)
	e.GET("/playground", echo.WrapHandler(playground.Handler("GraphQL playground", "/query")), authMiddleware)
	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "ok")
	})
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		req := c.Request()
		ctx := req.Context()
		nrtx := newrelic.FromContext(ctx)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger, err := altnrslog.FromContext(ctx)
		if err != nil {
			logger = slog.Default()
		}
		logger.ErrorContext(ctx,
			fmt.Sprintf("request: %v %v", req.Method, req.URL),
			slog.String("error", err.Error()))
		e.DefaultHTTPErrorHandler(err, c)
	}
	e.JSONSerializer = &s11n.JSONSerializer[*json.Encoder, *json.Decoder]{
		Encoder: json.NewEncoder,
		Decoder: json.NewDecoder,
	}
	slog.Info("echo server created")
	return e
}

var EchoSet = wire.NewSet(Echo)
