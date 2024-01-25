package provider

import (
	"os"

	"github.com/newrelic/go-agent/v3/newrelic"
	"go.uber.org/fx"
)

var NewRelic = fx.Options(
	fx.Provide(func() *newrelic.Application {
		app, err := newrelic.NewApplication(
			newrelic.ConfigAppName(os.Getenv("NEW_RELIC_CONFIG_APP_NAME")),
			newrelic.ConfigLicense(os.Getenv("NEW_RELIC_CONFIG_LICENSE")),
			newrelic.ConfigAppLogForwardingEnabled(true),
		)
		if err != nil {
			panic(err)
		}
		return app
	}),
)
