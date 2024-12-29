package usecase

import (
	"blogapi.miyamo.today/core/log"
	grpc "blogapi.miyamo.today/federator/internal/infra/grpc/bloggingevent"
	"context"
	"github.com/miyamo2/altnrslog"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"log/slog"

	"blogapi.miyamo.today/federator/internal/app/usecase/dto"
	"github.com/cockroachdb/errors"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// DetachTags is a use-case for detaching tags from an article.
type DetachTags struct {
	// bloggingEventServiceClient is a client of article service.
	bloggingEventServiceClient grpc.BloggingEventServiceClient
}

// Execute detaches tags from an article.
func (u *DetachTags) Execute(ctx context.Context, in dto.DetachTagsInDTO) (dto.DetachTagsOutDTO, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("DetachTags#Execute").End()

	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.Any("in", in)))

	response, err := u.bloggingEventServiceClient.DetachTags(
		newrelic.NewContext(ctx, nrtx),
		&grpc.DetachTagsRequest{
			Id:       in.ID(),
			TagNames: in.TagNames(),
		})
	if err != nil {
		err = errors.WithStack(err)
		logger.WarnContext(ctx, "END",
			slog.Group("return",
				slog.Any("*dto.DetachTagsOutDTO", nil),
				slog.Any("error", err)))
		return dto.DetachTagsOutDTO{}, err
	}

	out := dto.NewDetachTagsOutDTO(response.EventId, response.ArticleId, in.ClientMutationID())
	logger.InfoContext(ctx, "END",
		slog.Group("return",
			slog.Any("*dto.DetachTagsOutDTO", out),
			slog.Any("error", nil)))
	return out, nil
}

// NewDetachTags is a constructor of DetachTags.
func NewDetachTags(bloggingEventServiceClient grpc.BloggingEventServiceClient) *DetachTags {
	return &DetachTags{
		bloggingEventServiceClient: bloggingEventServiceClient,
	}
}
