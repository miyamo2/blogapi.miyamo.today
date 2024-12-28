package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/miyamo2/altnrslog"
	blogapicontext "github.com/miyamo2/blogapi.miyamo.today/core/context"
	"github.com/miyamo2/blogapi.miyamo.today/core/log"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/oklog/ulid/v2"
)

func SetBlogAPIContextToContext(requestType blogapicontext.RequestType) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			seg := newrelic.FromContext(c.Request().Context()).StartSegment("BlogAPICore: Set BlogAPIContext to Context")
			headers := c.Request().Header
			rid := func() string {
				v := headers.Get("x-request-id")
				if len(v) > 0 {
					return v
				}
				return ulid.Make().String()
			}()
			ctx = blogapicontext.StoreToContext(ctx, blogapicontext.New(rid, c.Path(), requestType, headers, nil))
			c.SetRequest(c.Request().WithContext(ctx))
			seg.End()
			return next(c)
		}
	}
}

func SetLoggerToContext(app *newrelic.Application) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			nrtx := newrelic.FromContext(c.Request().Context())
			seg := nrtx.StartSegment("BlogAPICore: Set Transactional Logger to Context")
			logger := log.New(log.WithAltNRSlogTransactionalHandler(app, nrtx))
			ctx, err := altnrslog.StoreToContext(c.Request().Context(), logger)
			if err != nil {
				return err
			}
			c.SetRequest(c.Request().WithContext(ctx))
			seg.End()
			return next(c)
		}
	}
}
