package usecase

import (
	"context"
	"log/slog"

	"github.com/miyamo2/altnrslog"
	"github.com/miyamo2/blogapi.miyamo.today/core/log"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"

	"github.com/cockroachdb/errors"
	"github.com/miyamo2/blogapi.miyamo.today/core/util/duration"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/app/usecase/dto"
	"github.com/miyamo2/blogapi.miyamo.today/protogen/tag/client/pb"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// Tag is a use-case of getting a tag by id.
type Tag struct {
	// tSvcClt is a client of article service.
	tSvcClt pb.TagServiceClient
}

// Execute gets a tag by id.
func (u *Tag) Execute(ctx context.Context, in dto.TagInDto) (dto.TagOutDto, error) {
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
	response, err := u.tSvcClt.GetTagById(
		newrelic.NewContext(ctx, nrtx),
		&pb.GetTagByIdRequest{
			Id: in.Id(),
		})
	if err != nil {
		err = errors.WithStack(err)
		lgr.WarnContext(ctx, "END",
			slog.String("duration", dw.SDuration()),
			slog.Group("return",
				slog.Any("*dto.ArticleOutDto", nil),
				slog.Any("error", err)))
		return dto.TagOutDto{}, err
	}
	pt := response.Tag
	pas := pt.Articles
	atcls := make([]dto.Article, 0, len(pas))
	for _, pa := range pas {
		atcls = append(atcls, dto.NewArticle(
			pa.Id,
			pa.Title,
			"",
			pa.ThumbnailUrl,
			pa.CreatedAt,
			pa.UpdatedAt))
	}
	t := dto.NewTagArticle(
		pt.Id,
		pt.Name,
		atcls)
	out := dto.NewTagOutDto(t)
	lgr.InfoContext(ctx, "END",
		slog.String("duration", dw.SDuration()),
		slog.Group("return",
			slog.Any("*dto.TagOutDto", out),
			slog.Any("error", nil)))
	return out, nil
}

// NewTag is a constructor of Tag.
func NewTag(tSvcClt pb.TagServiceClient) *Tag {
	return &Tag{
		tSvcClt: tSvcClt,
	}
}