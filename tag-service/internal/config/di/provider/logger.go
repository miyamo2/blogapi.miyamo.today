package provider

import (
	"log/slog"
	"os"

	"github.com/miyamo2/blogapi.miyamo.today/core/log"
	"github.com/newrelic/go-agent/v3/integrations/logcontext-v2/nrslog"
	"github.com/newrelic/go-agent/v3/newrelic"
	"go.uber.org/fx"
)

var Logger = fx.Options(
	fx.Invoke(func(app *newrelic.Application) {
		slog.SetDefault(log.New(log.WithInnerHandler(nrslog.JSONHandler(app, os.Stdout, &slog.HandlerOptions{}))))
	}))
