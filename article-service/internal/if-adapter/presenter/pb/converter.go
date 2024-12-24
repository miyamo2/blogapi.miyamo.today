package pb

import (
	"context"
	"log/slog"

	"github.com/cockroachdb/errors"
	"github.com/miyamo2/altnrslog"
	"github.com/miyamo2/blogapi.miyamo.today/article-service/internal/app/usecase/dto"
	"github.com/miyamo2/blogapi.miyamo.today/article-service/internal/infra/grpc"
	"github.com/miyamo2/blogapi.miyamo.today/core/log"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type Converter struct{}

func (c Converter) ToGetNextArticlesResponse(ctx context.Context, from *dto.GetNextOutDto) (response *grpc.GetNextArticlesResponse, ok bool) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToGetNextArticlesResponse").End()
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN", slog.Group("parameters", slog.Any("from", *from)))
	defer func() {
		logger.InfoContext(ctx, "END",
			slog.Group("return",
				slog.Bool("ok", ok)))
	}()
	articleDTOs := from.Articles()
	articlePBs := make([]*grpc.Article, 0, len(articleDTOs))
	for _, a := range articleDTOs {
		tagDTOs := a.Tags()
		tagPBs := make([]*grpc.Tag, 0, len(tagDTOs))
		for _, t := range tagDTOs {
			tagPBs = append(tagPBs, &grpc.Tag{
				Id:   t.Id(),
				Name: t.Name(),
			})
		}
		articlePBs = append(articlePBs, &grpc.Article{
			Id:           a.Id(),
			Title:        a.Title(),
			Body:         a.Body(),
			ThumbnailUrl: a.ThumbnailUrl(),
			CreatedAt:    a.CreatedAt(),
			UpdatedAt:    a.UpdatedAt(),
			Tags:         tagPBs,
		})
	}
	response = &grpc.GetNextArticlesResponse{
		Articles:    articlePBs,
		StillExists: from.HasNext(),
	}
	ok = true
	return
}

func (c Converter) ToGetAllArticlesResponse(ctx context.Context, from *dto.GetAllOutDto) (response *grpc.GetAllArticlesResponse, ok bool) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToGetAllArticlesResponse").End()
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN", slog.Group("parameters", slog.Any("from", *from)))
	defer func() {
		logger.InfoContext(ctx, "END",
			slog.Group("return",
				slog.Bool("ok", ok)))
	}()
	articleDTOs := from.Articles()
	articlePBs := make([]*grpc.Article, 0, len(articleDTOs))
	for _, a := range articleDTOs {
		tagDTOs := a.Tags()
		tagPBs := make([]*grpc.Tag, 0, len(tagDTOs))
		for _, t := range tagDTOs {
			tagPBs = append(tagPBs, &grpc.Tag{
				Id:   t.Id(),
				Name: t.Name(),
			})
		}
		articlePBs = append(articlePBs, &grpc.Article{
			Id:           a.Id(),
			Title:        a.Title(),
			Body:         a.Body(),
			ThumbnailUrl: a.ThumbnailUrl(),
			CreatedAt:    a.CreatedAt(),
			UpdatedAt:    a.UpdatedAt(),
			Tags:         tagPBs,
		})
	}
	response = &grpc.GetAllArticlesResponse{
		Articles: articlePBs,
	}
	ok = true
	return
}

func (c Converter) ToGetByIdArticlesResponse(ctx context.Context, from *dto.GetByIdOutDto) (response *grpc.GetArticleByIdResponse, ok bool) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToGetByIdArticlesResponse").End()
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN", slog.Group("parameters", slog.Any("from", *from)))
	defer func() {
		logger.InfoContext(ctx, "END",
			slog.Group("return",
				slog.Bool("ok", ok)))
	}()
	tagDTOs := from.Tags()
	tagPBs := make([]*grpc.Tag, 0, len(tagDTOs))
	for _, t := range tagDTOs {
		tagPBs = append(tagPBs, &grpc.Tag{
			Id:   t.Id(),
			Name: t.Name()})
	}
	articlePB := &grpc.Article{
		Id:           from.Id(),
		Title:        from.Title(),
		Body:         from.Body(),
		ThumbnailUrl: from.ThumbnailUrl(),
		CreatedAt:    from.CreatedAt(),
		UpdatedAt:    from.UpdatedAt(),
		Tags:         tagPBs,
	}
	response = &grpc.GetArticleByIdResponse{
		Article: articlePB,
	}
	ok = true
	return
}

func (c Converter) ToGetPrevArticlesResponse(ctx context.Context, from *dto.GetPrevOutDto) (response *grpc.GetPrevArticlesResponse, ok bool) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToGetPrevArticlesResponse").End()
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN", slog.Group("parameters", slog.Any("from", *from)))
	defer func() {
		logger.InfoContext(ctx, "END",
			slog.Group("return",
				slog.Bool("ok", ok)))
	}()
	articleDTOs := from.Articles()
	articlePBs := make([]*grpc.Article, 0, len(articleDTOs))
	for _, a := range articleDTOs {
		tagDTOs := a.Tags()
		tagPBs := make([]*grpc.Tag, 0, len(tagDTOs))
		for _, t := range tagDTOs {
			tagPBs = append(tagPBs, &grpc.Tag{
				Id:   t.Id(),
				Name: t.Name(),
			})
		}
		articlePBs = append(articlePBs, &grpc.Article{
			Id:           a.Id(),
			Title:        a.Title(),
			Body:         a.Body(),
			ThumbnailUrl: a.ThumbnailUrl(),
			CreatedAt:    a.CreatedAt(),
			UpdatedAt:    a.UpdatedAt(),
			Tags:         tagPBs,
		})
	}
	response = &grpc.GetPrevArticlesResponse{
		Articles:    articlePBs,
		StillExists: from.HasPrevious(),
	}
	ok = true
	return
}

func NewConverter() *Converter {
	return &Converter{}
}
