package githubactions

import (
	"bytes"
	"fmt"
	"net/http"
)

const payload = `{"event_type":"sync-read-model"}}`

type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

type BlogPublisher struct {
	endpoint string
	token    string
	client   Client
}

func (b *BlogPublisher) Publish() error {
	req, err := http.NewRequest(http.MethodPost, b.endpoint, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", b.token))
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	if _, err = b.client.Do(req); err != nil {
		return err
	}
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
