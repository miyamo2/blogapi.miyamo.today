package pb

import (
	"context"
	"log/slog"

	"github.com/cockroachdb/errors"
	"github.com/miyamo2/altnrslog"
	"github.com/miyamo2/blogapi.miyamo.today/core/log"
	"github.com/miyamo2/blogapi.miyamo.today/tag-service/internal/app/usecase/dto"
	"github.com/miyamo2/blogapi.miyamo.today/tag-service/internal/infra/grpc"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// Converter is am implementation of presenter.ToGetByIdConverter
type Converter struct{}

// ToGetByIdTagResponse is an implementation of presenter.ToGetByIdConverter#ToGetByIdTagResponse
func (c Converter) ToGetByIdTagResponse(ctx context.Context, from *dto.GetByIdOutDto) (response *grpc.GetTagByIdResponse, ok bool) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToGetByIdTagResponse").End()
	lgr, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		lgr = log.DefaultLogger()
	}
	lgr.InfoContext(ctx, "BEGIN")
	defer func() { lgr.InfoContext(ctx, "END", slog.Group("response", slog.Bool("ok", ok))) }()
	fa := from.Articles()
	pa := make([]*grpc.Article, 0, len(fa))
	for _, a := range fa {
		pa = append(pa, &grpc.Article{
			Id:           a.Id(),
			Title:        a.Title(),
			ThumbnailUrl: a.ThumbnailUrl(),
			CreatedAt:    a.CreatedAt(),
			UpdatedAt:    a.UpdatedAt(),
		})
	}
	t := &grpc.Tag{
		Id:       from.Id(),
		Name:     from.Name(),
		Articles: pa,
	}
	response = &grpc.GetTagByIdResponse{
		Tag: t,
	}
	ok = true
	return
}

// ToGetAllTagsResponse is an implementation of presenter.ToGetAllConverter#ToGetAllTagsResponse
func (c Converter) ToGetAllTagsResponse(ctx context.Context, from *dto.GetAllOutDto) (response *grpc.GetAllTagsResponse, ok bool) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToGetAllTagsResponse").End()
	lgr, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		lgr = log.DefaultLogger()
	}
	lgr.InfoContext(ctx, "BEGIN")
	defer func() { lgr.InfoContext(ctx, "END", slog.Group("response", slog.Bool("ok", ok))) }()
	ft := from.Tags()
	pt := make([]*grpc.Tag, 0, len(ft))
	for _, t := range ft {
		fa := t.Articles()
		pa := make([]*grpc.Article, 0, len(fa))
		for _, a := range fa {
			pa = append(pa, &grpc.Article{
				Id:           a.Id(),
				Title:        a.Title(),
				ThumbnailUrl: a.ThumbnailUrl(),
				CreatedAt:    a.CreatedAt(),
				UpdatedAt:    a.UpdatedAt(),
			})
		}
		pt = append(pt, &grpc.Tag{
			Id:       t.Id(),
			Name:     t.Name(),
			Articles: pa,
		})
	}
	response = &grpc.GetAllTagsResponse{
		Tags: pt,
	}
	ok = true
	return
}

// ToGetNextTagsResponse is an implementation of presenter.ToGetNextConverter#ToGetNextTagsResponse
func (c Converter) ToGetNextTagsResponse(ctx context.Context, from *dto.GetNextOutDto) (response *grpc.GetNextTagResponse, ok bool) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToGetNextTagsResponse").End()
	lgr, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		lgr = log.DefaultLogger()
	}
	lgr.InfoContext(ctx, "BEGIN")
	defer func() { lgr.InfoContext(ctx, "END", slog.Group("response", slog.Bool("ok", ok))) }()
	ft := from.Tags()
	pt := make([]*grpc.Tag, 0, len(ft))
	for _, t := range ft {
		fa := t.Articles()
		pa := make([]*grpc.Article, 0, len(fa))
		for _, a := range fa {
			pa = append(pa, &grpc.Article{
				Id:           a.Id(),
				Title:        a.Title(),
				ThumbnailUrl: a.ThumbnailUrl(),
				CreatedAt:    a.CreatedAt(),
				UpdatedAt:    a.UpdatedAt(),
			})
		}
		pt = append(pt, &grpc.Tag{
			Id:       t.Id(),
			Name:     t.Name(),
			Articles: pa,
		})
	}
	response = &grpc.GetNextTagResponse{
		Tags:        pt,
		StillExists: from.HasNext(),
	}
	ok = true
	return
}

// ToGetPrevTagsResponse is an implementation of presenter.ToGetPrevConverter#ToGetPrevTagsResponse
func (c Converter) ToGetPrevTagsResponse(ctx context.Context, from *dto.GetPrevOutDto) (response *grpc.GetPrevTagResponse, ok bool) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToGetPrevTagsResponse").End()
	lgr, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		lgr = log.DefaultLogger()
	}
	lgr.InfoContext(ctx, "BEGIN")
	defer func() { lgr.InfoContext(ctx, "END", slog.Group("response", slog.Bool("ok", ok))) }()
	ft := from.Tags()
	pt := make([]*grpc.Tag, 0, len(ft))
	for _, t := range ft {
		fa := t.Articles()
		pa := make([]*grpc.Article, 0, len(fa))
		for _, a := range fa {
			pa = append(pa, &grpc.Article{
				Id:           a.Id(),
				Title:        a.Title(),
				ThumbnailUrl: a.ThumbnailUrl(),
				CreatedAt:    a.CreatedAt(),
				UpdatedAt:    a.UpdatedAt(),
			})
		}
		pt = append(pt, &grpc.Tag{
			Id:       t.Id(),
			Name:     t.Name(),
			Articles: pa,
		})
	}
	response = &grpc.GetPrevTagResponse{
		Tags:        pt,
		StillExists: from.HasPrevious(),
	}
	ok = true
	return
}

func NewConverter() *Converter {
	return &Converter{}
}
