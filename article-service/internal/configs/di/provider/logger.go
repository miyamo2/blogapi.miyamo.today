package provider

import (
	"log/slog"

	"github.com/miyamo2/blogapi.miyamo.today/core/log"
	"github.com/newrelic/go-agent/v3/newrelic"
	"go.uber.org/fx"
)

var Logger = fx.Options(
	fx.Invoke(func(app *newrelic.Application) {
		l := log.New()
		slog.SetDefault(l)
	}))
