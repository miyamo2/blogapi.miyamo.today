package usecase

import (
	"context"
	"github.com/miyamo2/altnrslog"
	"github.com/miyamo2/blogapi.miyamo.today/core/log"
	grpc "github.com/miyamo2/blogapi.miyamo.today/federator/internal/infra/grpc/bloggingevent"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"log/slog"

	"github.com/cockroachdb/errors"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/app/usecase/dto"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// AttachTags is a use-case for attaching tags to an article.
type AttachTags struct {
	// bloggingEventServiceClient is a client of article service.
	bloggingEventServiceClient grpc.BloggingEventServiceClient
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
		&grpc.AttachTagsRequest{
			Id:       in.ID(),
			TagNames: in.TagNames(),
		})
	if err != nil {
		err = errors.WithStack(err)
		logger.WarnContext(ctx, "END",
			slog.Group("return",
				slog.Any("*dto.AttachTagsOutDTO", nil),
				slog.Any("error", err)))
		return dto.AttachTagsOutDTO{}, err
	}

	out := dto.NewAttachTagsOutDTO(response.EventId, response.ArticleId, in.ClientMutationID())
	logger.InfoContext(ctx, "END",
		slog.Group("return",
			slog.Any("*dto.AttachTagsOutDTO", out),
			slog.Any("error", nil)))
	return out, nil
}

// NewAttachTags is a constructor of AttachTags.
func NewAttachTags(bloggingEventServiceClient grpc.BloggingEventServiceClient) *AttachTags {
	return &AttachTags{
		bloggingEventServiceClient: bloggingEventServiceClient,
	}
}
