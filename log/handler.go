package log

import (
	"context"
	blogapictx "github.com/miyamo2/blogapi-core/context"
	"github.com/miyamo2/blogapi-core/log/internal"
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
	return h.JSONHandler.Handle(ctx, r)
}

// NewBlogAPILogHandler is constructor of BlogAPILogHandler.
func NewBlogAPILogHandler(handlerOption *slog.HandlerOptions) *BlogAPILogHandler {
	h := &BlogAPILogHandler{
		JSONHandler: slog.NewJSONHandler(os.Stdout, handlerOption),
	}
	return h
}
