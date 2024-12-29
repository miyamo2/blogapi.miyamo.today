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

// CreateArticle is a use-case of create an article by id.
type CreateArticle struct {
	// bloggingEventServiceClient is a client of article service.
	bloggingEventServiceClient grpc.BloggingEventServiceClient
}

// Execute gets an article by id.
func (u *CreateArticle) Execute(ctx context.Context, in dto.CreateArticleInDTO) (dto.CreateArticleOutDTO, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("Execute").End()

	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.Any("in", in)))

	thumbnail := in.ThumbnailURL()
	response, err := u.bloggingEventServiceClient.CreateArticle(
		newrelic.NewContext(ctx, nrtx),
		&grpc.CreateArticleRequest{
			Title:        in.Title(),
			Body:         in.Body(),
			ThumbnailUrl: thumbnail.String(),
			TagNames:     in.TagNames(),
		})
	if err != nil {
		err = errors.WithStack(err)
		logger.WarnContext(ctx, "END",
			slog.Group("return",
				slog.Any("*dto.CreateArticleOutDTO", nil),
				slog.Any("error", err)))
		return dto.CreateArticleOutDTO{}, err
	}

	out := dto.NewCreateArticleOutDTO(response.EventId, response.ArticleId, in.ClientMutationID())
	logger.InfoContext(ctx, "END",
		slog.Group("return",
			slog.Any("*dto.CreateArticleOutDTO", out),
			slog.Any("error", nil)))
	return out, nil
}

// NewCreateArticle is a constructor of CreateArticle.
func NewCreateArticle(bloggingEventServiceClient grpc.BloggingEventServiceClient) *CreateArticle {
	return &CreateArticle{
		bloggingEventServiceClient: bloggingEventServiceClient,
	}
}
