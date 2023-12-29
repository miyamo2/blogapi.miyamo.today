package middleware

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	blogapicontext "github.com/miyamo2/blogapi-core/context"
	"github.com/oklog/ulid/v2"
	"strings"
)

func SetTraceIDAndRequestIDToContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		traceID := c.Request().Header.Get("X-Amzn-Trace-Id")
		if traceID == "" {
			// UUID: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
			// X-Ray Trace ID: 1-xxxxxxxx-xxxxxxxxxxxxxxxxxxxxxxxx
			suuid := strings.ReplaceAll(uuid.New().String(), "-", "")
			traceID = fmt.Sprintf("1-%v-%v", suuid[0:8], suuid[8:])
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
