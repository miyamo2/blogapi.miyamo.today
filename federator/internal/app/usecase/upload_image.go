package usecase

import (
	"context"
	"github.com/miyamo2/altnrslog"
	"github.com/miyamo2/blogapi.miyamo.today/core/log"
	grpc "github.com/miyamo2/blogapi.miyamo.today/federator/internal/infra/grpc/bloggingevent"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"io"
	"log/slog"
	"net/url"

	"github.com/cockroachdb/errors"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/app/usecase/dto"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// UploadImage is a use-case of uploading image.
type UploadImage struct {
	// bloggingEventServiceClient is a client of article service.
	bloggingEventServiceClient grpc.BloggingEventServiceClient
}

// Execute uploads image
func (u *UploadImage) Execute(ctx context.Context, in dto.UploadImageInDTO) (dto.UploadImageOutDTO, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("UploadImage#Execute").End()

	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.Any("in", in)))

	stream, err := u.bloggingEventServiceClient.UploadImage(ctx)
	if err != nil {
		err = errors.WithStack(err)
		logger.WarnContext(ctx, "END",
			slog.Group("return",
				slog.Any("*dto.UploadImageOutDTO", nil),
				slog.Any("error", err)))
		return dto.UploadImageOutDTO{}, err
	}
	stream.Send(&grpc.UploadImageRequest{
		Value: &grpc.UploadImageRequest_Meta{
			Meta: &grpc.Meta{
				Name: in.Filename(),
			},
		},
	})

	data, err := io.ReadAll(in.Data())
	if err != nil {
		err = errors.WithStack(err)
		logger.WarnContext(ctx, "END",
			slog.Group("return",
				slog.Any("*dto.UploadImageOutDTO", nil),
				slog.Any("error", err)))
		return dto.UploadImageOutDTO{}, err
	}
	var read int
	reqOngoing := true
	for reqOngoing {
		end := read + 1024
		if end > len(data) {
			end = len(data)
			reqOngoing = false
		}
		err := stream.Send(&grpc.UploadImageRequest{
			Value: &grpc.UploadImageRequest_Data{
				Data: data[read:end],
			},
		})
		if err != nil {
			return dto.UploadImageOutDTO{}, err
		}
		read += 1024
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		return dto.UploadImageOutDTO{}, err
	}
	if !response.Success {
		return dto.UploadImageOutDTO{}, errors.WithStack(errors.New("failed to upload image"))
	}
	if response.Url == nil {
		return dto.UploadImageOutDTO{}, errors.WithStack(errors.New("failed to upload image"))
	}
	uri, err := url.Parse(*response.Url)
	if err != nil {
		return dto.UploadImageOutDTO{}, errors.WithStack(err)
	}

	out := dto.NewUploadImageOutDTO(*uri, in.ClientMutationID())
	logger.InfoContext(ctx, "END",
		slog.Group("return",
			slog.Any("*dto.UploadImageOutDTO", out),
			slog.Any("error", nil)))
	return out, nil
}

// NewUploadImage is a constructor of UploadImage.
func NewUploadImage(bloggingEventServiceClient grpc.BloggingEventServiceClient) *UploadImage {
	return &UploadImage{
		bloggingEventServiceClient: bloggingEventServiceClient,
	}
}
