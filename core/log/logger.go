package log

import (
	"io"
	"log/slog"
	"os"

	"blogapi.miyamo.today/core/internal"
	"github.com/miyamo2/altnrslog"

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
		switch handler := h.(type) {
		case *BlogAPILogHandler:
			return altnrslog.NewTransactionalHandler(app, nrtx, altnrslog.WithInnerHandlerProvider(
				func(w io.Writer) slog.Handler {
					handler.handler = slog.NewJSONHandler(w, internal.JSONHandlerOption)
					return handler
				}))
		default:
			return handler
		}
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
