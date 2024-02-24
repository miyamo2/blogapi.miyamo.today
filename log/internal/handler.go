package internal

import (
	"log/slog"
	"path/filepath"

	blogapictx "github.com/miyamo2/blogapi-core/context"
)

var JSONHandlerOption = &slog.HandlerOptions{
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

func ParseRequest(request blogapictx.Request) map[string]interface{} {
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
