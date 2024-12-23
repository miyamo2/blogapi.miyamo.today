//go:generate mockgen -source=$GOFILE -destination=../../../mock/app/usecase/storage/$GOFILE -package=storage
package storage

import (
	"context"
	"net/url"
)

// Uploader is an interface for uploading files.
type Uploader interface {
	// Upload uploads a file.
	Upload(ctx context.Context, name string, bytes []byte, contentType string) (*url.URL, error)
}
