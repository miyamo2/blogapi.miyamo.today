package usecase

import (
	"context"
	"log/slog"

	"github.com/miyamo2/altnrslog"
	"github.com/miyamo2/blogapi.miyamo.today/core/log"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/utils"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/cockroachdb/errors"
	"github.com/miyamo2/blogapi.miyamo.today/core/util/duration"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/app/usecase/dto"
	"github.com/miyamo2/blogapi.miyamo.today/protogen/article/client/pb"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// Articles is a use-case of getting an articles.
type Articles struct {
	// aSvcClt is a client of article service.
	aSvcClt pb.ArticleServiceClient
}

// Execute gets an article by id.
func (u *Articles) Execute(ctx context.Context, in dto.ArticlesInDto) (dto.ArticlesOutDto, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("Execute").End()
	dw := duration.Start()
	lgr, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		lgr = log.DefaultLogger()
	}
	lgr.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.Any("in", in)))
	out, err := func() (dto.ArticlesOutDto, error) {
		if in.First() != 0 {
			return u.executeNextPaging(ctx, in)
		} else if in.Last() != 0 {
			return u.executePrevPaging(ctx, in)
		}
		return u.execute(ctx)
	}()
	if err != nil {
		lgr.WarnContext(ctx, "END",
			slog.String("duration", dw.SDuration()),
			slog.Group("return",
				slog.Any("*dto.ArticleOutDto", out),
				slog.Any("error", err)))
		return out, err
	}
	lgr.InfoContext(ctx, "END",
		slog.String("duration", dw.SDuration()),
		slog.Group("return",
			slog.Any("*dto.ArticleOutDto", out),
			slog.Any("error", err)))
	return out, nil
}

// executeNextPaging
func (u *Articles) executeNextPaging(ctx context.Context, in dto.ArticlesInDto) (dto.ArticlesOutDto, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("executeNextPaging").End()
	dw := duration.Start()
	lgr, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		lgr = log.DefaultLogger()
	}
	lgr.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.Any("in", in)))
	response, err := u.aSvcClt.GetNextArticles(ctx, &pb.GetNextArticlesRequest{
		First: int32(in.First()),
		After: utils.PtrFromString(in.After()),
	})
	if err != nil {
		err = errors.WithStack(err)
		lgr.WarnContext(ctx, "END",
			slog.String("duration", dw.SDuration()),
			slog.Group("return",
				slog.Any("*dto.ArticleOutDto", nil),
				slog.Any("error", err)))
		return dto.ArticlesOutDto{}, err
	}
	pas := response.Articles
	das := make([]dto.ArticleTag, 0, len(pas))
	for _, pa := range pas {
		pts := pa.GetTags()
		ts := make([]dto.Tag, 0, len(pts))
		for _, pt := range pa.Tags {
			ts = append(ts, dto.NewTag(
				pt.Id,
				pt.Name))
		}
		das = append(das, dto.NewArticleTag(
			pa.Id,
			pa.Title,
			pa.Body,
			pa.ThumbnailUrl,
			pa.CreatedAt,
			pa.UpdatedAt,
			ts))
	}
	out := dto.NewArticlesOutDto(das, dto.ArticlesOutDtoWithHasNext(response.StillExists))
	lgr.InfoContext(ctx, "END",
		slog.String("duration", dw.SDuration()),
		slog.Group("return",
			slog.Any("*dto.ArticleOutDto", out),
			slog.Any("error", nil)))
	return out, nil
}

// executePrevPaging
func (u *Articles) executePrevPaging(ctx context.Context, in dto.ArticlesInDto) (dto.ArticlesOutDto, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("executePrevPaging").End()
	dw := duration.Start()
	lgr, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		lgr = log.DefaultLogger()
	}
	lgr.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.Any("in", in)))

	response, err := u.aSvcClt.GetPrevArticles(ctx, &pb.GetPrevArticlesRequest{
		Last:   int32(in.Last()),
		Before: utils.PtrFromString(in.Before()),
	})
	if err != nil {
		err = errors.WithStack(err)
		lgr.WarnContext(ctx, "END",
			slog.String("duration", dw.SDuration()),
			slog.Group("return",
				slog.Any("*dto.ArticleOutDto", nil),
				slog.Any("error", err)))
		return dto.ArticlesOutDto{}, err
	}
	pas := response.Articles
	das := make([]dto.ArticleTag, 0, len(pas))
	for _, pa := range pas {
		pts := pa.GetTags()
		ts := make([]dto.Tag, 0, len(pts))
		for _, pt := range pa.Tags {
			ts = append(ts, dto.NewTag(
				pt.Id,
				pt.Name))
		}
		das = append(das, dto.NewArticleTag(
			pa.Id,
			pa.Title,
			pa.Body,
			pa.ThumbnailUrl,
			pa.CreatedAt,
			pa.UpdatedAt,
			ts))
	}
	out := dto.NewArticlesOutDto(das, dto.ArticlesOutDtoWithHasPrev(response.StillExists))
	lgr.InfoContext(ctx, "END",
		slog.String("duration", dw.SDuration()),
		slog.Group("return",
			slog.Any("*dto.ArticleOutDto", out),
			slog.Any("error", nil)))
	return out, nil
}

// execute
func (u *Articles) execute(ctx context.Context) (dto.ArticlesOutDto, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("execute").End()
	dw := duration.Start()
	lgr, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		lgr = log.DefaultLogger()
	}
	lgr.InfoContext(ctx, "BEGIN")
	response, err := u.aSvcClt.GetAllArticles(ctx, &emptypb.Empty{})
	if err != nil {
		err = errors.WithStack(err)
		lgr.WarnContext(ctx, "END",
			slog.String("duration", dw.SDuration()),
			slog.Group("return",
				slog.Any("*dto.ArticleOutDto", nil),
				slog.Any("error", err)))
		return dto.ArticlesOutDto{}, err
	}
	pas := response.Articles
	das := make([]dto.ArticleTag, 0, len(pas))
	for _, pa := range pas {
		pts := pa.GetTags()
		ts := make([]dto.Tag, 0, len(pts))
		for _, pt := range pa.Tags {
			ts = append(ts, dto.NewTag(
				pt.Id,
				pt.Name))
		}
		das = append(das, dto.NewArticleTag(
			pa.Id,
			pa.Title,
			pa.Body,
			pa.ThumbnailUrl,
			pa.CreatedAt,
			pa.UpdatedAt,
			ts))
	}
	out := dto.NewArticlesOutDto(das)
	lgr.InfoContext(ctx, "END",
		slog.String("duration", dw.SDuration()),
		slog.Group("return",
			slog.Any("*dto.ArticleOutDto", out),
			slog.Any("error", nil)))
	return out, nil
}

// NewArticles is a constructor of Articles.
func NewArticles(aSvcClt pb.ArticleServiceClient) *Articles {
	return &Articles{
		aSvcClt: aSvcClt,
	}
}