package log

import (
	"github.com/miyamo2/blogapi-core/internal/log"
	"io"
	"log/slog"
)

var (
	defaultLogger *slog.Logger
)

func init() {
	defaultLogger = slog.New(log.NewBlogAPILogHandler(HandlerOption, PreHandle))
	slog.SetDefault(defaultLogger)
}

func DefaultLogger() *slog.Logger {
	return defaultLogger
}

// WithWriter returns a log.BlogAPILogHandlerOption that sets the writer to output.
func WithWriter(w io.Writer) log.BlogAPILogHandlerOption {
	return func(h *log.BlogAPILogHandler) {
		h.JSONHandler = slog.NewJSONHandler(w, HandlerOption)
	}
}

// New is constructor of log.NewBlogAPILogHandler
func New(options ...log.BlogAPILogHandlerOption) *slog.Logger {
	return slog.New(log.NewBlogAPILogHandler(HandlerOption, PreHandle, options...))
}
