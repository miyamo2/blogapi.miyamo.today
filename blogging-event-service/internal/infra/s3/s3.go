package s3

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/cockroachdb/errors"
	"github.com/miyamo2/altnrslog"
	"github.com/miyamo2/blogapi.miyamo.today/core/log"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"
	"log/slog"
	"net/url"
	"os"

	s3sdk "github.com/aws/aws-sdk-go-v2/service/s3"
)

type Uploader struct {
	client Client
}

func (s *Uploader) Upload(ctx context.Context, name string, data []byte, contentType string) (url *url.URL, err error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("Upload").End()

	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN")
	defer func() {
		if err != nil {
			logger.WarnContext(ctx, "END",
				slog.Group("return",
					slog.Any("*url.URL", nil),
					slog.Any("error", err)))
			return
		}
		logger.InfoContext(ctx, "END", slog.Group("return", slog.Any("*url.URL", fmt.Sprintf("%+v", *url))))
	}()

	_, err = s.client.PutObject(ctx, &s3sdk.PutObjectInput{
		Bucket:      aws.String(os.Getenv("S3_BUCKET")),
		Key:         aws.String(name),
		Body:        bytes.NewBuffer(data),
		ContentType: aws.String(contentType),
	})
	if err != nil {
		err = errors.WithStack(err)
		return nil, err
	}

	uri, err := url.Parse(fmt.Sprintf("%s/%s", os.Getenv("CDN_HOST"), name))
	return uri, nil
}

// NewUploader creates a new Uploader
func NewUploader(client Client) *Uploader {
	return &Uploader{
		client: client,
	}
}
