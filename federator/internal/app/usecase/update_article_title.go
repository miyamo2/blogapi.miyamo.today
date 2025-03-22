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

// UpdateArticleTitle is a use-case for updating an article title.
type UpdateArticleTitle struct {
	// bloggingEventServiceClient is a client of article service.
	bloggingEventServiceClient blogging_eventconnect.BloggingEventServiceClient
}

// Execute updates an article title.
func (u *UpdateArticleTitle) Execute(ctx context.Context, in dto.UpdateArticleTitleInDTO) (dto.UpdateArticleTitleOutDTO, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("UpdateArticleTitle#Execute").End()

	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.Any("in", in)))

	response, err := u.bloggingEventServiceClient.UpdateArticleTitle(
		newrelic.NewContext(ctx, nrtx),
		connect.NewRequest(&grpc.UpdateArticleTitleRequest{
			Id:    in.ID(),
			Title: in.Title(),
		}))
	if err != nil {
		err = errors.WithStack(err)
		logger.WarnContext(ctx, "END",
			slog.Group("return",
				slog.Any("*dto.UpdateArticleTitleOutDTO", nil),
				slog.Any("error", err)))
		return dto.UpdateArticleTitleOutDTO{}, err
	}

	message := response.Msg
	out := dto.NewUpdateArticleTitleOutDTO(message.EventId, message.ArticleId, in.ClientMutationID())
	logger.InfoContext(ctx, "END",
		slog.Group("return",
			slog.Any("*dto.UpdateArticleTitleOutDTO", out),
			slog.Any("error", nil)))
	return out, nil
}

// NewUpdateArticleTitle is a constructor of UpdateArticleTitle.
func NewUpdateArticleTitle(bloggingEventServiceClient blogging_eventconnect.BloggingEventServiceClient) *UpdateArticleTitle {
	return &UpdateArticleTitle{
		bloggingEventServiceClient: bloggingEventServiceClient,
	}
}
