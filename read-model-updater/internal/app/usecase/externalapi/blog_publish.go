package externalapi

import "context"

type BlogPublisher interface {
	Publish(ctx context.Context) error
}
