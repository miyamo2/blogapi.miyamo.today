package usecase

import (
	"context"
	"log/slog"

	"github.com/miyamo2/altnrslog"
	"github.com/miyamo2/api.miyamo.today/core/log"
	"github.com/miyamo2/api.miyamo.today/federator/internal/utils"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/cockroachdb/errors"
	"github.com/miyamo2/api.miyamo.today/core/util/duration"
	"github.com/miyamo2/api.miyamo.today/federator/internal/app/usecase/dto"
	"github.com/miyamo2/api.miyamo.today/protogen/tag/client/pb"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// Tags is a use-case of getting tags.
type Tags struct {
	// tSvcClt is a client of article service.
	tSvcClt pb.TagServiceClient
}

// Execute gets a tag by id.
func (u *Tags) Execute(ctx context.Context, in dto.TagsInDto) (dto.TagsOutDto, error) {
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
	out, err := func() (dto.TagsOutDto, error) {
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
				slog.Any("*dto.TagsOutDto", out),
				slog.Any("error", err)))
		return out, err
	}
	lgr.InfoContext(ctx, "END",
		slog.String("duration", dw.SDuration()),
		slog.Group("return",
			slog.Any("*dto.TagsOutDto", out),
			slog.Any("error", err)))
	return out, nil
}

// executeNextPaging
func (u *Tags) executeNextPaging(ctx context.Context, in dto.TagsInDto) (dto.TagsOutDto, error) {
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

	response, err := u.tSvcClt.GetNextTags(ctx, &pb.GetNextTagsRequest{
		First: int32(in.First()),
		After: utils.PtrFromString(in.After()),
	})
	if err != nil {
		err = errors.WithStack(err)
		lgr.WarnContext(ctx, "END",
			slog.String("duration", dw.SDuration()),
			slog.Group("return",
				slog.Any("*dto.TagsOutDto", nil),
				slog.Any("error", err)))
		return dto.TagsOutDto{}, err
	}
	pts := response.Tags
	dts := make([]dto.TagArticle, 0, len(pts))
	for _, pt := range pts {
		pas := pt.GetArticles()
		das := make([]dto.Article, 0, len(pas))
		for _, pa := range pt.Articles {
			das = append(das, dto.NewArticle(
				pa.Id,
				pa.Title,
				"",
				pa.ThumbnailUrl,
				pa.CreatedAt,
				pa.UpdatedAt))
		}
		dts = append(dts, dto.NewTagArticle(
			pt.Id,
			pt.Name,
			das))
	}
	out := dto.NewTagsOutDto(dts, dto.TagsOutDtoWithHasNext(response.StillExists))
	lgr.InfoContext(ctx, "END",
		slog.String("duration", dw.SDuration()),
		slog.Group("return",
			slog.Any("*dto.TagsOutDto", out),
			slog.Any("error", nil)))
	return out, nil
}

// executePrevPaging
func (u *Tags) executePrevPaging(ctx context.Context, in dto.TagsInDto) (dto.TagsOutDto, error) {
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

	response, err := u.tSvcClt.GetPrevTags(ctx, &pb.GetPrevTagsRequest{
		Last:   int32(in.Last()),
		Before: utils.PtrFromString(in.Before()),
	})
	if err != nil {
		err = errors.WithStack(err)
		lgr.WarnContext(ctx, "END",
			slog.String("duration", dw.SDuration()),
			slog.Group("return",
				slog.Any("*dto.TagsOutDto", nil),
				slog.Any("error", err)))
		return dto.TagsOutDto{}, err
	}
	pts := response.Tags
	dts := make([]dto.TagArticle, 0, len(pts))
	for _, pt := range pts {
		pas := pt.GetArticles()
		das := make([]dto.Article, 0, len(pas))
		for _, pa := range pt.Articles {
			das = append(das, dto.NewArticle(
				pa.Id,
				pa.Title,
				"",
				pa.ThumbnailUrl,
				pa.CreatedAt,
				pa.UpdatedAt))
		}
		dts = append(dts, dto.NewTagArticle(
			pt.Id,
			pt.Name,
			das))
	}
	out := dto.NewTagsOutDto(dts, dto.TagsOutDtoWithHasPrev(response.StillExists))
	lgr.InfoContext(ctx, "END",
		slog.String("duration", dw.SDuration()),
		slog.Group("return",
			slog.Any("*dto.TagsOutDto", out),
			slog.Any("error", nil)))
	return out, nil
}

// execute
func (u *Tags) execute(ctx context.Context) (dto.TagsOutDto, error) {
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
	response, err := u.tSvcClt.GetAllTags(ctx, &emptypb.Empty{})
	if err != nil {
		err = errors.WithStack(err)
		lgr.WarnContext(ctx, "END",
			slog.String("duration", dw.SDuration()),
			slog.Group("return",
				slog.Any("*dto.TagsOutDto", nil),
				slog.Any("error", err)))
		return dto.TagsOutDto{}, err
	}
	pts := response.Tags
	dts := make([]dto.TagArticle, 0, len(pts))
	for _, pt := range pts {
		pas := pt.GetArticles()
		das := make([]dto.Article, 0, len(pas))
		for _, pa := range pt.Articles {
			das = append(das, dto.NewArticle(
				pa.Id,
				pa.Title,
				"",
				pa.ThumbnailUrl,
				pa.CreatedAt,
				pa.UpdatedAt))
		}
		dts = append(dts, dto.NewTagArticle(
			pt.Id,
			pt.Name,
			das))
	}
	out := dto.NewTagsOutDto(dts)
	lgr.InfoContext(ctx, "END",
		slog.String("duration", dw.SDuration()),
		slog.Group("return",
			slog.Any("*dto.TagsOutDto", out),
			slog.Any("error", nil)))
	return out, nil
}

// NewTag is a constructor of Tag.
func NewTags(tSvcClt pb.TagServiceClient) *Tags {
	return &Tags{
		tSvcClt: tSvcClt,
	}
}
