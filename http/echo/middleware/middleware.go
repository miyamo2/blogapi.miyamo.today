package middleware

import (
	"context"
	"github.com/labstack/echo/v4"
	blogapicontext "github.com/miyamo2/blogapi-core/context"
	"github.com/oklog/ulid/v2"
)

func SetTraceIDAndRequestIDToContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		traceID := c.Request().Header.Get("X-Amzn-Trace-Id")
		if traceID == "" {
			traceID = ulid.Make().String()
		}
		requestID := c.Request().Header.Get("x-amzn-RequestId")
		if requestID == "" {
			requestID = ulid.Make().String()
		}
		ctx = context.WithValue(ctx, blogapicontext.TraceIDKey{}, traceID)
		ctx = context.WithValue(ctx, blogapicontext.RequestIDKey{}, requestID)
		c.SetRequest(c.Request().WithContext(ctx))
		err := next(c)
		return err
	}
}
