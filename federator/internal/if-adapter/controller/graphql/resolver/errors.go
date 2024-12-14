package resolver

import "github.com/cockroachdb/errors"

var (
	ErrFailedToConvertToArticleNode       = errors.New("failed to convert to article node")
	ErrFailedToConvertToArticleConnection = errors.New("failed to convert to article connection")
)
