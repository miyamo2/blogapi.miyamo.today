package usecase

import (
	"blogapi.miyamo.today/core/log"
	grpc "blogapi.miyamo.today/federator/internal/infra/grpc/blogging_event"
	"blogapi.miyamo.today/federator/internal/infra/grpc/blogging_event/blogging_eventconnect"
	"connectrpc.com/connect"
	"context"
	"github.com/miyamo2/altnrslog"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"log/slog"

	"blogapi.miyamo.today/federator/internal/app/usecase/dto"
	"github.com/cockroachdb/errors"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// AttachTags is a use-case for attaching tags to an article.
type AttachTags struct {
	// bloggingEventServiceClient is a client of article service.
	bloggingEventServiceClient blogging_eventconnect.BloggingEventServiceClient
}

// Execute attaches tags to an article.
func (u *AttachTags) Execute(ctx context.Context, in dto.AttachTagsInDTO) (dto.AttachTagsOutDTO, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("AttachTags#Execute").End()

	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.Any("in", in)))

	response, err := u.bloggingEventServiceClient.AttachTags(
		newrelic.NewContext(ctx, nrtx),
		connect.NewRequest(&grpc.AttachTagsRequest{
			Id:       in.ID(),
			TagNames: in.TagNames(),
		}))
	if err != nil {
		err = errors.WithStack(err)
		logger.WarnContext(ctx, "END",
			slog.Group("return",
				slog.Any("*dto.AttachTagsOutDTO", nil),
				slog.Any("error", err)))
		return dto.AttachTagsOutDTO{}, err
	}

	message := response.Msg
	out := dto.NewAttachTagsOutDTO(message.EventId, message.ArticleId, in.ClientMutationID())
	logger.InfoContext(ctx, "END",
		slog.Group("return",
			slog.Any("*dto.AttachTagsOutDTO", out),
			slog.Any("error", nil)))
	return out, nil
}

// NewAttachTags is a constructor of AttachTags.
func NewAttachTags(bloggingEventServiceClient blogging_eventconnect.BloggingEventServiceClient) *AttachTags {
	return &AttachTags{
		bloggingEventServiceClient: bloggingEventServiceClient,
	}
}
