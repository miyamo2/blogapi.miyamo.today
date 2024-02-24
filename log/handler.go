package log

import (
	"context"
	"log/slog"
	"os"
)

type Logger struct {
	*slog.Logger
}

type PreHandle func(ctx context.Context, r *slog.Record) error

// BlogAPILogHandler is an implementation of slog.Handler for blogapi.
type BlogAPILogHandler struct {
	*slog.JSONHandler
	preHandle PreHandle
}

// Handle add trace id to slog.Attrs from context before output.
// See: https://pkg.go.dev/log/slog#JSONHandler.Handle
func (h *BlogAPILogHandler) Handle(ctx context.Context, r slog.Record) error {
	if h.preHandle != nil {
		err := h.preHandle(ctx, &r)
		if err != nil {
			return err
		}
	}
	return h.JSONHandler.Handle(ctx, r)
}

// NewBlogAPILogHandler is constructor of BlogAPILogHandler.
func NewBlogAPILogHandler(handlerOption *slog.HandlerOptions, preHandle PreHandle) *BlogAPILogHandler {
	h := &BlogAPILogHandler{
		JSONHandler: slog.NewJSONHandler(os.Stdout, handlerOption),
		preHandle:   preHandle,
	}
	return h
}
