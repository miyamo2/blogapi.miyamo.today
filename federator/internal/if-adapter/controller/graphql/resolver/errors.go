package resolver

import "github.com/cockroachdb/errors"

var (
	ErrFailedToConvertToArticleNode       = errors.New("failed to convert to article node")
	ErrFailedToConvertToArticleConnection = errors.New("failed to convert to article connection")
)

func ErrorWithStack(err error) error {
	return errors.WithStack(err)
}
