package provider

import (
	"github.com/google/wire"
	"os"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func NewRelic() *newrelic.Application {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(os.Getenv("NEW_RELIC_CONFIG_APP_NAME")),
		newrelic.ConfigLicense(os.Getenv("NEW_RELIC_CONFIG_LICENSE")),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		panic(err)
	}
	return app
}

var NewRelicSet = wire.NewSet(NewRelic)
