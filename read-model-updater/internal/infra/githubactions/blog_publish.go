package githubactions

import (
	"bytes"
	"context"
	"fmt"
	"github.com/cockroachdb/errors"
	"github.com/miyamo2/altnrslog"
	"github.com/newrelic/go-agent/v3/newrelic"
	"log/slog"
	"net/http"
)

const payload = `{"event_type":"sync-read-model", "client_payload": {}}`

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
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		logger = slog.Default()
	}
	logger.Info("START")
	defer logger.Info("END")

	req, err := http.NewRequest(http.MethodPost, b.endpoint, bytes.NewBuffer([]byte(payload)))
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
