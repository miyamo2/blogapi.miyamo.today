package log

import (
	"context"
	blogapictx "github.com/miyamo2/blogapi-core/context"
	"log/slog"
	"path/filepath"
)

var HandlerOption = &slog.HandlerOptions{
	AddSource: true,
	ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
		switch {
		case a.Key == slog.SourceKey:
			source := a.Value.Any().(*slog.Source)
			source.File = filepath.Base(source.File)
			return slog.Any("source", source)
		}
		return a
	},
}

var PreHandle = func(ctx context.Context, r *slog.Record) error {
	bctx := blogapictx.FromContext(ctx)
	if bctx == nil {
		return nil
	}
	r.Add(slog.String("trace_id", bctx.TraceID))
	r.Add(slog.String("span_id", bctx.SpanID))
	r.Add(slog.Any("in_request", bctx.Incoming))
	if outgoing := bctx.Outgoing; outgoing == nil {
		r.Add(slog.Any("out_request", outgoing))
	}
	return nil
}
