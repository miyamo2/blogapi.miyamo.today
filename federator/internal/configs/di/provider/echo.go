package provider

import (
	_ "embed"
	"fmt"
	"log/slog"

	"net/http"

	"blogapi.miyamo.today/core/echo/middlewares"
	"blogapi.miyamo.today/core/echo/s11n"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/goccy/go-json"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/miyamo2/altnrslog"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"
)

//go:embed remote_import_paths.html
var remoteImportPaths string

func Echo(srv *handler.Server, nr *newrelic.Application, verifier middlewares.Verifier) *echo.Echo {
	slog.Info("creating echo server")
	e := echo.New()

	authMiddleware := middlewares.Auth(verifier)

	queryGroup := e.Group("/query")
	queryGroup.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
			echo.HeaderAccessControlRequestHeaders,
			echo.HeaderAccessControlRequestHeaders,
		},
		AllowMethods: []string{
			http.MethodOptions,
			http.MethodPost,
		},
	}))
	queryGroup.POST("/",
		echo.WrapHandler(srv),
		nrecho.Middleware(nr),
		middlewares.SetLoggerToContext(nr),
		middlewares.RequestLog(),
		authMiddleware,
	)
	e.GET("/playground",
		echo.WrapHandler(playground.Handler("GraphQL playground", "/query")),
		authMiddleware,
	)
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})
	e.Any("/", func(c echo.Context) error {
		switch c.Request().Method {
		case http.MethodGet, http.MethodHead:
			return c.HTML(http.StatusOK, remoteImportPaths)
		default:
			return echo.NewHTTPError(http.StatusMethodNotAllowed, fmt.Sprintf("%s / unsupported", c.Request().Method))
		}
	})
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		req := c.Request()
		ctx := req.Context()
		nrtx := newrelic.FromContext(ctx)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger, _ := altnrslog.FromContext(ctx)
		if logger == nil {
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
