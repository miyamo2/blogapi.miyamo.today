package log

import (
	"github.com/miyamo2/blogapi-core/internal/log"
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
