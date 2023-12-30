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
	r.Add(slog.Any("in_request", parseRequest(bctx.Incoming)))
	if outgoing := bctx.Outgoing; outgoing != nil {
		r.Add(slog.Any("out_request", parseRequest(*outgoing)))
	}
	return nil
}

func parseRequest(request blogapictx.Request) map[string]interface{} {
	parsedReq := map[string]interface{}{}
	if request.Service != "" {
		parsedReq["service"] = request.Service
	}
	if request.Path != "" {
		parsedReq["path"] = request.Path
	}
	if len(request.Headers) > 0 {
		parsedReq["headers"] = request.Headers
	}
	if request.Duration != nil {
		parsedReq["duration_ms"] = request.Duration
	}
	if request.Status != nil {
		parsedReq["status"] = request.Status
	}
	if request.Body != nil {
		parsedReq["body"] = request.Body
	}
	return parsedReq
}
