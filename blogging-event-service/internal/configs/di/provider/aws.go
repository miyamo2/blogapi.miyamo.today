package provider

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	awss3 "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/wire"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/infra/s3"
)

func AWSConfig() *aws.Config {
	ctx := context.Background()
	awsConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic(err)
	}
	return &awsConfig
}

func S3Client(awsConfig *aws.Config) *awss3.Client {
	return awss3.NewFromConfig(*awsConfig)
}

var AWSSet = wire.NewSet(
	AWSConfig,
	S3Client,
	wire.Bind(new(s3.Client), new(*awss3.Client)),
)
