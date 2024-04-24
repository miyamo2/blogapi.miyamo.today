package provider

import (
	"github.com/miyamo2/blogapi-core/log"
	"github.com/newrelic/go-agent/v3/newrelic"
	"go.uber.org/fx"
	"log/slog"
)

var Logger = fx.Options(
	fx.Invoke(func(app *newrelic.Application) {
		l := log.New()
		slog.SetDefault(l)
	}))
