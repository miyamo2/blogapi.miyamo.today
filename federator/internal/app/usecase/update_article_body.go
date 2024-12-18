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

// UpdateArticleBody is a use-case for updating an article body.
type UpdateArticleBody struct {
	// bloggingEventServiceClient is a client of article service.
	bloggingEventServiceClient grpc.BloggingEventServiceClient
}

// Execute updates an article body.
func (u *UpdateArticleBody) Execute(ctx context.Context, in dto.UpdateArticleBodyInDTO) (dto.UpdateArticleBodyOutDTO, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("UpdateArticleBody#Execute").End()

	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.Any("in", in)))

	response, err := u.bloggingEventServiceClient.UpdateArticleBody(
		newrelic.NewContext(ctx, nrtx),
		&grpc.UpdateArticleBodyRequest{
			Id:   in.ID(),
			Body: in.Content(),
		})
	if err != nil {
		err = errors.WithStack(err)
		logger.WarnContext(ctx, "END",
			slog.Group("return",
				slog.Any("*dto.UpdateArticleBodyOutDTO", nil),
				slog.Any("error", err)))
		return dto.UpdateArticleBodyOutDTO{}, err
	}

	out := dto.NewUpdateArticleBodyOutDTO(response.EventId, response.ArticleId, in.ClientMutationID())
	logger.InfoContext(ctx, "END",
		slog.Group("return",
			slog.Any("*dto.UpdateArticleBodyOutDTO", out),
			slog.Any("error", nil)))
	return out, nil
}

// NewUpdateArticleBody is a constructor of UpdateArticleBody.
func NewUpdateArticleBody(bloggingEventServiceClient grpc.BloggingEventServiceClient) *UpdateArticleBody {
	return &UpdateArticleBody{
		bloggingEventServiceClient: bloggingEventServiceClient,
	}
}
