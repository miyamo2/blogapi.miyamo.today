//go:generate mockgen -source=$GOFILE -destination=../../mock/infra/s3/$GOFILE -package=$GOPACKAGE
package s3

import (
	"context"
	awss3 "github.com/aws/aws-sdk-go-v2/service/s3"
)

// Client is a subset of the Uploader client methods used by the application.
type Client interface {
	// PutObject uploads an object to an Uploader bucket.
	// See https://docs.aws.amazon.com/sdk-for-go/api/service/s3/#S3.PutObject
	PutObject(ctx context.Context, params *awss3.PutObjectInput, optFns ...func(*awss3.Options)) (*awss3.PutObjectOutput, error)
}
