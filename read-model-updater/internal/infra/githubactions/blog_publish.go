package githubactions

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/cockroachdb/errors"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/oklog/ulid/v2"
)

const payload = `{"event_type":"sync-read-model", "client_payload": { "event_id": "%s" }}`

type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

type BlogPublisher struct {
	endpoint string
	token    string
	client   Client
}

func (b *BlogPublisher) Publish(ctx context.Context) error {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("BlogPublisher#Publish").End()

	req, err := http.NewRequest(http.MethodPost, b.endpoint, bytes.NewBuffer([]byte(fmt.Sprintf(payload, ulid.Make()))))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", b.token))
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	seg := newrelic.StartExternalSegment(nrtx, req)

	res, err := b.client.Do(req)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}

	seg.Response = res
	seg.End()

	return nil
}

// NewBlogPublisher returns new *BlogPublisher
func NewBlogPublisher(endpoint, token string, client Client) *BlogPublisher {
	return &BlogPublisher{
		endpoint: endpoint,
		token:    token,
		client:   client,
	}
}
