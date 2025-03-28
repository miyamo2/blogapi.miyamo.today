package middlewares

import (
	blogapicontext "blogapi.miyamo.today/core/context"
	"blogapi.miyamo.today/core/log"
	"context"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/miyamo2/altnrslog"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/oklog/ulid/v2"
	"net/http"
	"net/url"
	"strings"
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
			ctx = blogapicontext.StoreToContext(ctx, blogapicontext.New(rid, c.Request().URL.Path, requestType, headers, nil))
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

func NRConnect(app *newrelic.Application) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			nrtx := newrelic.FromContext(ctx)
			if nrtx == nil {
				nrtx = app.StartTransaction(c.Request().URL.Path)
				defer nrtx.End()
			} else {
				nrtx.SetName(c.Request().URL.Path)
			}
			host := c.Request().Host
			nrtx.SetWebRequest(newrelic.WebRequest{
				Type:      "ConnectRPC",
				Host:      host,
				Header:    c.Request().Header,
				URL:       &url.URL{Scheme: "connectrpc", Host: host, Path: c.Request().URL.Path},
				Method:    "POST",
				Transport: newrelic.TransportHTTP})
			ctx = newrelic.NewContext(ctx, nrtx)
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}

// Verifier is an interface for verifying JWT tokens.
type Verifier interface {
	// Verify verifies the token and returns the claims.
	Verify(ctx context.Context, token string) (jwt.Token, error)
}

func Auth(verifier Verifier) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			defer newrelic.FromContext(ctx).StartSegment("BlogAPICore: Auth").End()
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "no authorization header provided"})
			}
			token := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))
			if token == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token"})
			}
			jwtToken, err := verifier.Verify(ctx, token)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "failed to verify token"})
			}
			ctx = context.WithValue(ctx, JWTContextKey{}, jwtToken)
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}

type JWTContextKey struct{}
