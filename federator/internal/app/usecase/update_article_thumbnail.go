package usecase

import (
	"blogapi.miyamo.today/core/log"
	"blogapi.miyamo.today/federator/internal/app/usecase/dto"
	grpc "blogapi.miyamo.today/federator/internal/infra/grpc/blogging_event"
	"blogapi.miyamo.today/federator/internal/infra/grpc/blogging_event/blogging_eventconnect"
	"connectrpc.com/connect"
	"context"
	"github.com/cockroachdb/errors"
	"github.com/miyamo2/altnrslog"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"
	"log/slog"
)

// UpdateArticleThumbnail is a use-case for updating an article body.
type UpdateArticleThumbnail struct {
	// bloggingEventServiceClient is a client of article service.
	bloggingEventServiceClient blogging_eventconnect.BloggingEventServiceClient
}

// Execute updates an article body.
func (u *UpdateArticleThumbnail) Execute(ctx context.Context, in dto.UpdateArticleThumbnailInDTO) (dto.UpdateArticleThumbnailOutDTO, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("UpdateArticleThumbnail#Execute").End()

	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.Any("in", in)))

	thumbnail := in.Thumbnail()
	response, err := u.bloggingEventServiceClient.UpdateArticleThumbnail(
		newrelic.NewContext(ctx, nrtx),
		connect.NewRequest(&grpc.UpdateArticleThumbnailRequest{
			Id:           in.ID(),
			ThumbnailUrl: thumbnail.String(),
		}))
	if err != nil {
		err = errors.WithStack(err)
		logger.WarnContext(ctx, "END",
			slog.Group("return",
				slog.Any("*dto.UpdateArticleThumbnailOutDTO", nil),
				slog.Any("error", err)))
		return dto.UpdateArticleThumbnailOutDTO{}, err
	}

	message := response.Msg
	out := dto.NewUpdateArticleThumbnailOutDTO(message.EventId, message.ArticleId, in.ClientMutationID())
	logger.InfoContext(ctx, "END",
		slog.Group("return",
			slog.Any("*dto.UpdateArticleThumbnailOutDTO", out),
			slog.Any("error", nil)))
	return out, nil
}

// NewUpdateArticleThumbnail is a constructor of UpdateArticleThumbnail.
func NewUpdateArticleThumbnail(bloggingEventServiceClient blogging_eventconnect.BloggingEventServiceClient) *UpdateArticleThumbnail {
	return &UpdateArticleThumbnail{
		bloggingEventServiceClient: bloggingEventServiceClient,
	}
}
