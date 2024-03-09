package provider

import (
	"github.com/newrelic/go-agent/v3/integrations/logcontext-v2/nrslog"
	"github.com/newrelic/go-agent/v3/newrelic"
	"go.uber.org/fx"
	"log/slog"
	"os"
)

var Logger = fx.Options(
	fx.Provide(func(nr *newrelic.Application) *slog.Logger {
		logger := slog.New(nrslog.JSONHandler(nr, os.Stdout, nil))
		slog.SetDefault(logger)
		return logger
	}))
