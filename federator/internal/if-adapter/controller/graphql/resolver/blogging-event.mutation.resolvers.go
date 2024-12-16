package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.60

import (
	"context"
	"fmt"
	"log/slog"
	"net/url"

	"github.com/miyamo2/altnrslog"
	"github.com/miyamo2/blogapi.miyamo.today/core/log"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/app/usecase/dto"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/if-adapter/presenters/graphql/model"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// CreateArticle is the resolver for the createArticle field.
func (r *mutationResolver) CreateArticle(ctx context.Context, input model.CreateArticleInput) (*model.CreateArticlePayload, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("CreateArticle").End()

	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = ErrorWithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.String("input", fmt.Sprintf("%+v", input))))

	var clientMutationID string
	if input.ClientMutationID != nil {
		clientMutationID = *input.ClientMutationID
	}
	outDTO, err := r.usecases.createArticle.Execute(ctx, dto.NewCreateArticleInDTO(input.Title, input.Content, url.URL(input.ThumbnailURL), input.TagNames, clientMutationID))
	if err != nil {
		return nil, err
	}

	payload, err := r.converters.createArticle.ToCreateArticle(ctx, outDTO)
	if err != nil {
		return nil, err
	}
	return payload, err
}
