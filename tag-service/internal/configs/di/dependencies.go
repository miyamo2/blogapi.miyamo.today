package di

import (
	"fmt"
	"log/slog"

	blogapictx "blogapi.miyamo.today/core/context"
	"blogapi.miyamo.today/core/echo/middlewares"
	"connectrpc.com/grpcreflect"
	"github.com/google/wire"
	"github.com/miyamo2/altnrslog"

	"blogapi.miyamo.today/core/echo/s11n"
	"blogapi.miyamo.today/tag-service/internal/infra/grpc/grpcconnect"
	"connectrpc.com/grpchealth"
	"github.com/goccy/go-json"
	"github.com/labstack/echo/v4"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func Echo(service grpcconnect.TagServiceHandler, nr *newrelic.Application) *echo.Echo {
	slog.Info("creating echo server")
	e := echo.New()

	path, handler := grpcconnect.NewTagServiceHandler(service)
	e.POST(
		fmt.Sprintf("%s*", path),
		echo.WrapHandler(handler),
		nrecho.Middleware(nr),
		middlewares.NRConnect(nr),
		middlewares.SetBlogAPIContextToContext(blogapictx.RequestTypeGRPC),
		middlewares.SetLoggerToContext(nr),
	)

	gRPCReflector := grpcreflect.NewStaticReflector(grpcconnect.TagServiceName)

	reflectV1Path, reflectV1Handler := grpcreflect.NewHandlerV1(gRPCReflector)
	e.POST(fmt.Sprintf("%s*", reflectV1Path), echo.WrapHandler(reflectV1Handler))

	reflectV1AlphaPath, reflectV1AlphaHandler := grpcreflect.NewHandlerV1Alpha(gRPCReflector)
	e.POST(fmt.Sprintf("%s*", reflectV1AlphaPath), echo.WrapHandler(reflectV1AlphaHandler))

	healthPath, healthHandler := grpchealth.NewHandler(grpchealth.NewStaticChecker(grpcconnect.TagServiceName))
	e.POST(fmt.Sprintf("%s*", healthPath), echo.WrapHandler(healthHandler))

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		req := c.Request()
		ctx := req.Context()
		nrtx := newrelic.FromContext(ctx)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger, _ := altnrslog.FromContext(ctx)
		if logger == nil {
			logger = slog.Default()
		}
		logger.ErrorContext(
			ctx,
			fmt.Sprintf("request: %v %v", req.Method, req.URL),
			slog.String("error", err.Error()),
		)
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
