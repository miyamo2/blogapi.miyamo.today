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
	"github.com/miyamo2/blogapi.miyamo.today/protogen/article/client/pb"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// Article is a use-case of getting an article by id.
type Article struct {
	// aSvcClt is a client of article service.
	aSvcClt pb.ArticleServiceClient
}

// Execute gets an article by id.
func (u *Article) Execute(ctx context.Context, in dto.ArticleInDto) (dto.ArticleOutDto, error) {
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
	response, err := u.aSvcClt.GetArticleById(
		newrelic.NewContext(ctx, nrtx),
		&pb.GetArticleByIdRequest{
			Id: in.Id(),
		})
	if err != nil {
		err = errors.WithStack(err)
		lgr.WarnContext(ctx, "END",
			slog.String("duration", dw.SDuration()),
			slog.Group("return",
				slog.Any("*dto.ArticleOutDto", nil),
				slog.Any("error", err)))
		return dto.ArticleOutDto{}, err
	}
	pa := response.Article
	pts := pa.GetTags()
	ts := make([]dto.Tag, 0, len(pts))
	for _, pt := range pts {
		ts = append(ts, dto.NewTag(
			pt.Id,
			pt.Name))
	}
	a := dto.NewArticleTag(
		pa.Id,
		pa.Title,
		pa.Body,
		pa.ThumbnailUrl,
		pa.CreatedAt,
		pa.UpdatedAt,
		ts)
	out := dto.NewArticleOutDto(a)
	lgr.InfoContext(ctx, "END",
		slog.String("duration", dw.SDuration()),
		slog.Group("return",
			slog.Any("*dto.ArticleOutDto", out),
			slog.Any("error", nil)))
	return out, nil
}

// NewArticle is a constructor of Article.
func NewArticle(aSvcClt pb.ArticleServiceClient) *Article {
	return &Article{
		aSvcClt: aSvcClt,
	}
}