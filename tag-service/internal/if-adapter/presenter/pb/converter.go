package pb

import (
	"connectrpc.com/connect"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
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
func (c Converter) ToGetByIdTagResponse(ctx context.Context, from *dto.GetByIdOutDto) (response *connect.Response[grpc.GetTagByIdResponse], ok bool) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToGetByIdTagResponse").End()
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN")
	defer func() { logger.InfoContext(ctx, "END", slog.Group("response", slog.Bool("ok", ok))) }()
	articleDTOs := from.Articles()
	articlePBs := make([]*grpc.Article, 0, len(articleDTOs))
	for _, a := range articleDTOs {
		articlePBs = append(articlePBs, &grpc.Article{
			Id:           a.Id(),
			Title:        a.Title(),
			ThumbnailUrl: a.ThumbnailUrl(),
			CreatedAt:    timestamppb.New(a.CreatedAt().StdTime()),
			UpdatedAt:    timestamppb.New(a.UpdatedAt().StdTime()),
		})
	}
	tagPB := &grpc.Tag{
		Id:       from.Id(),
		Name:     from.Name(),
		Articles: articlePBs,
	}
	rawResponse := &grpc.GetTagByIdResponse{
		Tag: tagPB,
	}
	response = connect.NewResponse(rawResponse)
	ok = true
	return
}

// ToGetAllTagsResponse is an implementation of presenter.ToGetAllConverter#ToGetAllTagsResponse
func (c Converter) ToGetAllTagsResponse(ctx context.Context, from *dto.GetAllOutDto) (response *connect.Response[grpc.GetAllTagsResponse], ok bool) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToGetAllTagsResponse").End()
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN")
	defer func() { logger.InfoContext(ctx, "END", slog.Group("response", slog.Bool("ok", ok))) }()
	tagDTOs := from.Tags()
	tagPBs := make([]*grpc.Tag, 0, len(tagDTOs))
	for _, t := range tagDTOs {
		articleDTOs := t.Articles()
		articlePBs := make([]*grpc.Article, 0, len(articleDTOs))
		for _, a := range articleDTOs {
			articlePBs = append(articlePBs, &grpc.Article{
				Id:           a.Id(),
				Title:        a.Title(),
				ThumbnailUrl: a.ThumbnailUrl(),
				CreatedAt:    timestamppb.New(a.CreatedAt().StdTime()),
				UpdatedAt:    timestamppb.New(a.UpdatedAt().StdTime()),
			})
		}
		tagPBs = append(tagPBs, &grpc.Tag{
			Id:       t.Id(),
			Name:     t.Name(),
			Articles: articlePBs,
		})
	}
	rawResponse := &grpc.GetAllTagsResponse{
		Tags: tagPBs,
	}
	response = connect.NewResponse(rawResponse)
	ok = true
	return
}

// ToGetNextTagsResponse is an implementation of presenter.ToGetNextConverter#ToGetNextTagsResponse
func (c Converter) ToGetNextTagsResponse(ctx context.Context, from *dto.GetNextOutDto) (response *connect.Response[grpc.GetNextTagResponse], ok bool) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToGetNextTagsResponse").End()
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN")
	defer func() { logger.InfoContext(ctx, "END", slog.Group("response", slog.Bool("ok", ok))) }()
	tagDTOs := from.Tags()
	tagPBs := make([]*grpc.Tag, 0, len(tagDTOs))
	for _, t := range tagDTOs {
		articleDTOs := t.Articles()
		articlePBs := make([]*grpc.Article, 0, len(articleDTOs))
		for _, a := range articleDTOs {
			articlePBs = append(articlePBs, &grpc.Article{
				Id:           a.Id(),
				Title:        a.Title(),
				ThumbnailUrl: a.ThumbnailUrl(),
				CreatedAt:    timestamppb.New(a.CreatedAt().StdTime()),
				UpdatedAt:    timestamppb.New(a.UpdatedAt().StdTime()),
			})
		}
		tagPBs = append(tagPBs, &grpc.Tag{
			Id:       t.Id(),
			Name:     t.Name(),
			Articles: articlePBs,
		})
	}
	rawResponse := &grpc.GetNextTagResponse{
		Tags:        tagPBs,
		StillExists: from.HasNext(),
	}
	response = connect.NewResponse(rawResponse)
	ok = true
	return
}

// ToGetPrevTagsResponse is an implementation of presenter.ToGetPrevConverter#ToGetPrevTagsResponse
func (c Converter) ToGetPrevTagsResponse(ctx context.Context, from *dto.GetPrevOutDto) (response *connect.Response[grpc.GetPrevTagResponse], ok bool) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToGetPrevTagsResponse").End()
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN")
	defer func() { logger.InfoContext(ctx, "END", slog.Group("response", slog.Bool("ok", ok))) }()
	tagDTOs := from.Tags()
	tagPBs := make([]*grpc.Tag, 0, len(tagDTOs))
	for _, t := range tagDTOs {
		articleDTOs := t.Articles()
		articlePBs := make([]*grpc.Article, 0, len(articleDTOs))
		for _, a := range articleDTOs {
			articlePBs = append(articlePBs, &grpc.Article{
				Id:           a.Id(),
				Title:        a.Title(),
				ThumbnailUrl: a.ThumbnailUrl(),
				CreatedAt:    timestamppb.New(a.CreatedAt().StdTime()),
				UpdatedAt:    timestamppb.New(a.UpdatedAt().StdTime()),
			})
		}
		tagPBs = append(tagPBs, &grpc.Tag{
			Id:       t.Id(),
			Name:     t.Name(),
			Articles: articlePBs,
		})
	}
	rawResponse := &grpc.GetPrevTagResponse{
		Tags:        tagPBs,
		StillExists: from.HasPrevious(),
	}
	response = connect.NewResponse(rawResponse)
	ok = true
	return
}

func NewConverter() *Converter {
	return &Converter{}
}
