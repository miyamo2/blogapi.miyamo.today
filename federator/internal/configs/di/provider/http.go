package provider

import (
	"github.com/google/wire"
	"github.com/newrelic/go-agent/v3/newrelic"
	"net/http"
)

func HTTPClient() *http.Client {
	client := &http.Client{
		Transport: newrelic.NewRoundTripper(nil),
	}
	return client
}

var HTTPSet = wire.NewSet(HTTPClient)
