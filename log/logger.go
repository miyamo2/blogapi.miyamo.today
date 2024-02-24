package log

import (
	"io"
	"log/slog"

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
func WrapNRHandler(app *newrelic.Application) HandlerWrapOption {
	return func(h slog.Handler) slog.Handler {
		return nrslog.WrapHandler(app, h)
	}
}

// WithWriter returns a slog.Handler with a modified log.writer if the handler type is BlogAPILogHandle otherwise unmodified slog.Handler.
func WithWriter(w io.Writer) HandlerWrapOption {
	return func(h slog.Handler) slog.Handler {
		switch handler := h.(type) {
		case *BlogAPILogHandler:
			handler.JSONHandler = slog.NewJSONHandler(w, internal.JSONHandlerOption)
			return handler
		default:
			return handler
		}
	}
}

// New is wrapped constructor of log.slog
func New(options ...HandlerWrapOption) *slog.Logger {
	var h slog.Handler = NewBlogAPILogHandler(internal.JSONHandlerOption, internal.PreHandle)
	for _, o := range options {
		h = o(h)
	}
	return slog.New(h)
}
