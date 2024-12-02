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
	lgr, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		lgr = log.DefaultLogger()
	}
	lgr.InfoContext(ctx, "BEGIN", slog.Group("patameters", slog.Any("from", *from)))
	defer func() {
		lgr.InfoContext(ctx, "END",
			slog.Group("return",
				slog.Bool("ok", ok)))
	}()
	fa := from.Articles()
	pa := make([]*grpc.Article, 0, len(fa))
	for _, a := range fa {
		ft := a.Tags()
		pt := make([]*grpc.Tag, 0, len(ft))
		for _, t := range ft {
			pt = append(pt, &grpc.Tag{
				Id:   t.Id(),
				Name: t.Name(),
			})
		}
		pa = append(pa, &grpc.Article{
			Id:           a.Id(),
			Title:        a.Title(),
			Body:         a.Body(),
			ThumbnailUrl: a.ThumbnailUrl(),
			CreatedAt:    a.CreatedAt(),
			UpdatedAt:    a.UpdatedAt(),
			Tags:         pt,
		})
	}
	response = &grpc.GetNextArticlesResponse{
		Articles:    pa,
		StillExists: from.HasNext(),
	}
	ok = true
	return
}

func (c Converter) ToGetAllArticlesResponse(ctx context.Context, from *dto.GetAllOutDto) (response *grpc.GetAllArticlesResponse, ok bool) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToGetAllArticlesResponse").End()
	lgr, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		lgr = log.DefaultLogger()
	}
	lgr.InfoContext(ctx, "BEGIN", slog.Group("patameters", slog.Any("from", *from)))
	defer func() {
		lgr.InfoContext(ctx, "END",
			slog.Group("return",
				slog.Bool("ok", ok)))
	}()
	fa := from.Articles()
	pa := make([]*grpc.Article, 0, len(fa))
	for _, a := range fa {
		ft := a.Tags()
		pt := make([]*grpc.Tag, 0, len(ft))
		for _, t := range ft {
			pt = append(pt, &grpc.Tag{
				Id:   t.Id(),
				Name: t.Name(),
			})
		}
		pa = append(pa, &grpc.Article{
			Id:           a.Id(),
			Title:        a.Title(),
			Body:         a.Body(),
			ThumbnailUrl: a.ThumbnailUrl(),
			CreatedAt:    a.CreatedAt(),
			UpdatedAt:    a.UpdatedAt(),
			Tags:         pt,
		})
	}
	response = &grpc.GetAllArticlesResponse{
		Articles: pa,
	}
	ok = true
	return
}

func (c Converter) ToGetByIdArticlesResponse(ctx context.Context, from *dto.GetByIdOutDto) (response *grpc.GetArticleByIdResponse, ok bool) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToGetByIdArticlesResponse").End()
	lgr, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		lgr = log.DefaultLogger()
	}
	lgr.InfoContext(ctx, "BEGIN", slog.Group("patameters", slog.Any("from", *from)))
	defer func() {
		lgr.InfoContext(ctx, "END",
			slog.Group("return",
				slog.Bool("ok", ok)))
	}()
	ft := from.Tags()
	pt := make([]*grpc.Tag, 0, len(ft))
	for _, t := range ft {
		pt = append(pt, &grpc.Tag{
			Id:   t.Id(),
			Name: t.Name()})
	}
	a := &grpc.Article{
		Id:           from.Id(),
		Title:        from.Title(),
		Body:         from.Body(),
		ThumbnailUrl: from.ThumbnailUrl(),
		CreatedAt:    from.CreatedAt(),
		UpdatedAt:    from.UpdatedAt(),
		Tags:         pt,
	}
	response = &grpc.GetArticleByIdResponse{
		Article: a,
	}
	ok = true
	return
}

func (c Converter) ToGetPrevArticlesResponse(ctx context.Context, from *dto.GetPrevOutDto) (response *grpc.GetPrevArticlesResponse, ok bool) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToGetPrevArticlesResponse").End()
	lgr, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		lgr = log.DefaultLogger()
	}
	lgr.InfoContext(ctx, "BEGIN", slog.Group("patameters", slog.Any("from", *from)))
	defer func() {
		lgr.InfoContext(ctx, "END",
			slog.Group("return",
				slog.Bool("ok", ok)))
	}()
	fa := from.Articles()
	pa := make([]*grpc.Article, 0, len(fa))
	for _, a := range fa {
		ft := a.Tags()
		pt := make([]*grpc.Tag, 0, len(ft))
		for _, t := range ft {
			pt = append(pt, &grpc.Tag{
				Id:   t.Id(),
				Name: t.Name(),
			})
		}
		pa = append(pa, &grpc.Article{
			Id:           a.Id(),
			Title:        a.Title(),
			Body:         a.Body(),
			ThumbnailUrl: a.ThumbnailUrl(),
			CreatedAt:    a.CreatedAt(),
			UpdatedAt:    a.UpdatedAt(),
			Tags:         pt,
		})
	}
	response = &grpc.GetPrevArticlesResponse{
		Articles:    pa,
		StillExists: from.HasPrevious(),
	}
	ok = true
	return
}

func NewConverter() *Converter {
	return &Converter{}
}
