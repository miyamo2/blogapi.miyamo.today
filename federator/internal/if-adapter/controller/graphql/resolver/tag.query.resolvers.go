package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"
	"log/slog"

	"github.com/cockroachdb/errors"
	"github.com/miyamo2/altnrslog"
	"github.com/miyamo2/api.miyamo.today/core/log"

	"github.com/miyamo2/api.miyamo.today/core/util/duration"
	"github.com/miyamo2/api.miyamo.today/federator/internal/app/usecase/dto"
	"github.com/miyamo2/api.miyamo.today/federator/internal/if-adapter/presenters/graphql/model"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// Tags is the resolver for the tags field.
func (r *queryResolver) Tags(ctx context.Context, first *int, last *int, after *string, before *string) (*model.TagConnection, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("Tags").End()
	dw := duration.Start()
	lgr, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		lgr = log.DefaultLogger()
	}
	lgr.InfoContext(ctx, "BEGIN",
		slog.Group("parameters",
			slog.Any("first", first),
			slog.Any("last", last),
			slog.Any("after", after),
			slog.Any("before", before)))
	opts := make([]dto.TagsInDtoOption, 0, 4)
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
	in, err := dto.NewTagsInDto(opts...)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		lgr.WarnContext(ctx, "END",
			slog.String("duration", dw.SDuration()),
			slog.Group("returns",
				slog.Any("*model.TagConnection", nil),
				slog.Any("error", err)))
		return nil, err
	}
	oDto, err := r.usecases.tags.Execute(ctx, in)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		lgr.WarnContext(ctx, "END",
			slog.String("duration", dw.SDuration()),
			slog.Group("returns",
				slog.Any("*model.TagConnection", nil),
				slog.Any("error", err)))
		return nil, err
	}
	cnctn, err := r.converters.tags.ToTags(ctx, oDto)
	if err != nil {
		lgr.InfoContext(ctx, "END",
			slog.String("duration", dw.SDuration()),
			slog.Group("returns",
				slog.Any("*model.TagConnection", nil),
				slog.Any("error", err)))
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		return nil, err
	}
	lgr.InfoContext(ctx, "END",
		slog.String("duration", dw.SDuration()),
		slog.Group("returns",
			slog.Any("*model.TagConnection", &cnctn),
			slog.Any("error", nil)))
	return cnctn, nil
}

// Tag is the resolver for the tag field.
func (r *queryResolver) Tag(ctx context.Context, id string) (*model.TagNode, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("Tag").End()
	dw := duration.Start()
	lgr, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		lgr = log.DefaultLogger()
	}
	lgr.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.String("id", id)))
	oDto, err := r.usecases.tag.Execute(ctx, dto.NewTagInDto(id))
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		lgr.WarnContext(ctx, "END",
			slog.String("duration", dw.SDuration()),
			slog.Group("returns",
				slog.Any("*model.TagNode", nil),
				slog.Any("error", err)))
		return nil, err
	}
	nd, err := r.converters.tag.ToTag(ctx, oDto)
	if err != nil {
		lgr.InfoContext(ctx, "END",
			slog.String("duration", dw.SDuration()),
			slog.Group("returns",
				slog.Any("*model.TagNode", nil),
				slog.Any("error", err)))
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		return nil, err
	}
	lgr.InfoContext(ctx, "END",
		slog.String("duration", dw.SDuration()),
		slog.Group("returns",
			slog.Any("*model.TageNode", &nd),
			slog.Any("error", nil)))
	return nd, nil
}
