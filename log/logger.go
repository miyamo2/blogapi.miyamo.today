package log

import (
	"context"
	"io"
	"log/slog"
	"os"

	"github.com/miyamo2/blogapi-core/log/internal"
	"github.com/newrelic/go-agent/v3/integrations/logcontext-v2/nrslog"
	"github.com/newrelic/go-agent/v3/newrelic"
)

var (
	defaultLogger *slog.Logger
)

func init() {
	defaultLogger = New()
	slog.SetDefault(defaultLogger)
}

func DefaultLogger() *slog.Logger {
	return defaultLogger
}

// HandlerWrapOption is an option for slog.Handler.
type HandlerWrapOption func(slog.Handler) slog.Handler

// WrapNRHandler returns a slog.Handler wrapped in nrslog.
func WrapNRHandler(app *newrelic.Application, nrtx *newrelic.Transaction) HandlerWrapOption {
	return func(h slog.Handler) slog.Handler {
		switch handler := h.(type) {
		case *BlogAPILogHandler:
			innerHandler := nrslog.JSONHandler(app, os.Stdout, internal.JSONHandlerOption)
			handler.handler = innerHandler.WithTransaction(nrtx)
			return handler
		default:
			return handler
		}
	}
}

// WithWriter returns a slog.Handler with a modified log.writer if the handler type is BlogAPILogHandle otherwise unmodified slog.Handler.
func WithWriter(w io.Writer) HandlerWrapOption {
	return func(h slog.Handler) slog.Handler {
		switch handler := h.(type) {
		case *BlogAPILogHandler:
			handler.handler = slog.NewJSONHandler(w, internal.JSONHandlerOption)
			return handler
		default:
			return handler
		}
	}
}

// New is wrapped constructor of log.slog
func New(options ...HandlerWrapOption) *slog.Logger {
	var h slog.Handler = NewBlogAPILogHandler(os.Stdout, internal.JSONHandlerOption)
	for _, o := range options {
		h = o(h)
	}
	return slog.New(h)
}

type loggerKey struct{}

// FromContext returns the *slog.Logger stored in context.Context.
func FromContext(ctx context.Context) *slog.Logger {
	lgr, ok := ctx.Value(loggerKey{}).(*slog.Logger)
	if !ok {
		return nil
	}
	return lgr
}

// StoreToContext stores the *slog.Logger in context.Context.
func StoreToContext(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}
