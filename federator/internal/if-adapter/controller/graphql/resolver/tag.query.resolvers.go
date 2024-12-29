package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.60

import (
	"context"
	"log/slog"

	"blogapi.miyamo.today/core/log"
	"blogapi.miyamo.today/federator/internal/app/usecase/dto"
	"blogapi.miyamo.today/federator/internal/if-adapter/presenters/graphql/model"
	"github.com/miyamo2/altnrslog"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// Tags is the resolver for the tags field.
func (r *queryResolver) Tags(ctx context.Context, first *int, last *int, after *string, before *string) (*model.TagConnection, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("Tags").End()

	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = ErrorWithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters",
			slog.Any("first", first),
			slog.Any("last", last),
			slog.Any("after", after),
			slog.Any("before", before)))
	opts := make([]dto.TagsInDTOOption, 0, 4)
	if first != nil {
		opts = append(opts, dto.TagsInWithFirst(*first))
	}
	if last != nil {
		opts = append(opts, dto.TagsInWithLast(*last))
	}
	if after != nil {
		opts = append(opts, dto.TagsInWithAfter(*after))
	}
	if before != nil {
		opts = append(opts, dto.TagsInWithBefore(*before))
	}
	in, err := dto.NewTagsInDTO(opts...)
	if err != nil {
		err = ErrorWithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger.WarnContext(ctx, "END",
			slog.Group("returns",
				slog.Any("*model.TagConnection", nil),
				slog.Any("error", err)))
		return nil, err
	}
	oDTO, err := r.usecases.tags.Execute(ctx, in)
	if err != nil {
		err = ErrorWithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger.WarnContext(ctx, "END",
			slog.Group("returns",
				slog.Any("*model.TagConnection", nil),
				slog.Any("error", err)))
		return nil, err
	}
	connection, err := r.converters.tags.ToTags(ctx, oDTO)
	if err != nil {
		logger.InfoContext(ctx, "END",
			slog.Group("returns",
				slog.Any("*model.TagConnection", nil),
				slog.Any("error", err)))
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		return nil, err
	}
	logger.InfoContext(ctx, "END",
		slog.Group("returns",
			slog.Any("*model.TagConnection", &connection),
			slog.Any("error", nil)))
	return connection, nil
}

// Tag is the resolver for the tag field.
func (r *queryResolver) Tag(ctx context.Context, id string) (*model.TagNode, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("Tag").End()

	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = ErrorWithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.String("id", id)))
	oDTO, err := r.usecases.tag.Execute(ctx, dto.NewTagInDTO(id))
	if err != nil {
		err = ErrorWithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger.WarnContext(ctx, "END",
			slog.Group("returns",
				slog.Any("*model.TagNode", nil),
				slog.Any("error", err)))
		return nil, err
	}
	node, err := r.converters.tag.ToTag(ctx, oDTO)
	if err != nil {
		logger.InfoContext(ctx, "END",
			slog.Group("returns",
				slog.Any("*model.TagNode", nil),
				slog.Any("error", err)))
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		return nil, err
	}
	logger.InfoContext(ctx, "END",
		slog.Group("returns",
			slog.Any("*model.TageNode", &node),
			slog.Any("error", nil)))
	return node, nil
}
