package provider

import (
	"github.com/miyamo2/blogapi-core/log"
	"github.com/newrelic/go-agent/v3/integrations/logcontext-v2/nrslog"
	"github.com/newrelic/go-agent/v3/newrelic"
	"go.uber.org/fx"
	"log/slog"
	"os"
)

var Logger = fx.Options(
	fx.Invoke(func(app *newrelic.Application) {
		slog.SetDefault(log.New(log.WithInnerHandler(nrslog.JSONHandler(app, os.Stdout, &slog.HandlerOptions{}))))
	}))
