package log

import (
	"context"
	"io"
	"log/slog"

	blogapictx "github.com/miyamo2/blogapi-core/context"
	"github.com/miyamo2/blogapi-core/log/internal"
)

type Logger struct {
	*slog.Logger
}

type PreHandle func(ctx context.Context, r *slog.Record) error

// BlogAPILogHandler is an implementation of slog.Handler for blogapi.
type BlogAPILogHandler struct {
	handler slog.Handler
}

func (h *BlogAPILogHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

func (h *BlogAPILogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h.handler.WithAttrs(attrs)
}

func (h *BlogAPILogHandler) WithGroup(name string) slog.Handler {
	return h.handler.WithGroup(name)
}

// Handle add trace id to slog.Attrs from context before output.
// See: https://pkg.go.dev/log/slog#JSONHandler.Handle
func (h *BlogAPILogHandler) Handle(ctx context.Context, r slog.Record) error {
	bctx := blogapictx.FromContext(ctx)
	if bctx == nil {
		return nil
	}
	r.Add(slog.String("request_id", bctx.RequestID))
	r.Add(slog.Any("in_request", internal.ParseRequest(bctx.Incoming)))
	if outgoing := bctx.Outgoing; outgoing != nil {
		r.Add(slog.Any("out_request", internal.ParseRequest(*outgoing)))
	}
	return h.handler.Handle(ctx, r)
}

// NewBlogAPILogHandler is constructor of BlogAPILogHandler.
func NewBlogAPILogHandler(w io.Writer, handlerOption *slog.HandlerOptions) *BlogAPILogHandler {
	h := &BlogAPILogHandler{
		handler: slog.NewJSONHandler(w, handlerOption),
	}
	return h
}
