package githubactions

import (
	"net/http"
)

type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

type BlogPublisher struct {
	endpoint string
	token    string
	client   Client
}

func (b *BlogPublisher) Publish() error {
	//req, err := http.NewRequest(http.MethodPost, b.endpoint, nil)
	//if err != nil {
	//	return err
	//}
	//req.Header.Set("Authorization", b.token)
	//if _, err = b.client.Do(req); err != nil {
	//	return err
	//}
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
