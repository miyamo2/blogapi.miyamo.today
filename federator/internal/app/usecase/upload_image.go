package usecase

import (
	"blogapi.miyamo.today/core/log"
	grpc "blogapi.miyamo.today/federator/internal/infra/grpc/blogging_event"
	"blogapi.miyamo.today/federator/internal/infra/grpc/blogging_event/blogging_eventconnect"
	"context"
	"github.com/miyamo2/altnrslog"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"io"
	"log/slog"
	"net/url"

	"blogapi.miyamo.today/federator/internal/app/usecase/dto"
	"github.com/cockroachdb/errors"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// UploadImage is a use-case of uploading image.
type UploadImage struct {
	// bloggingEventServiceClient is a client of article service.
	bloggingEventServiceClient blogging_eventconnect.BloggingEventServiceClient
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

	stream := u.bloggingEventServiceClient.UploadImage(ctx)
	stream.Send(&grpc.UploadImageRequest{
		Value: &grpc.UploadImageRequest_Meta{
			Meta: &grpc.Meta{
				Name:        in.Filename(),
				ContentType: in.ContentType(),
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
		end := read + chunkSize
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
		read += chunkSize
	}

	response, err := stream.CloseAndReceive()
	if err != nil {
		return dto.UploadImageOutDTO{}, err
	}

	message := response.Msg
	if !message.Success {
		return dto.UploadImageOutDTO{}, errors.WithStack(errors.New("failed to upload image"))
	}
	if message.Url == nil {
		return dto.UploadImageOutDTO{}, errors.WithStack(errors.New("failed to upload image"))
	}
	uri, err := url.Parse(*message.Url)
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
func NewUploadImage(bloggingEventServiceClient blogging_eventconnect.BloggingEventServiceClient) *UploadImage {
	return &UploadImage{
		bloggingEventServiceClient: bloggingEventServiceClient,
	}
}

const chunkSize = 1022976
