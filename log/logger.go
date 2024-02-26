package log

import (
	"context"
	"github.com/miyamo2/altnrslog"
	"github.com/miyamo2/blogapi-core/internal"
	"io"
	"log/slog"
	"os"

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

// WithAltNRSlogTransactionalHandler returns a slog.Handler wrapped altnrslog,TransactionalHandler.
func WithAltNRSlogTransactionalHandler(app *newrelic.Application, nrtx *newrelic.Transaction) HandlerWrapOption {
	return func(h slog.Handler) slog.Handler {
		return WithInnerHandler(altnrslog.NewTransactionalHandler(app,
			nrtx, altnrslog.WithSlogHandlerSpecify(true, internal.JSONHandlerOption)))(h)
	}
}

func WithInnerHandler(innerHandler slog.Handler) HandlerWrapOption {
	return func(h slog.Handler) slog.Handler {
		switch handler := h.(type) {
		case *BlogAPILogHandler:
			handler.handler = innerHandler
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
