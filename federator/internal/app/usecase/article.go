package usecase

import (
	"context"
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	grpc "github.com/miyamo2/blogapi.miyamo.today/federator/internal/infra/grpc/article"
	"log/slog"
	"net/url"
	"time"

	"github.com/miyamo2/altnrslog"
	"github.com/miyamo2/blogapi.miyamo.today/core/log"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"

	"github.com/cockroachdb/errors"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/app/usecase/dto"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// Article is a use-case of getting an article by id.
type Article struct {
	// articleServiceClient is a client of article service.
	articleServiceClient grpc.ArticleServiceClient
}

// Execute gets an article by id.
func (u *Article) Execute(ctx context.Context, in dto.ArticleInDto) (dto.ArticleOutDto, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("Execute").End()

	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.Any("in", in)))
	response, err := u.articleServiceClient.GetArticleById(
		newrelic.NewContext(ctx, nrtx),
		&grpc.GetArticleByIdRequest{
			Id: in.Id(),
		})
	if err != nil {
		err = errors.WithStack(err)
		logger.WarnContext(ctx, "END",
			slog.Group("return",
				slog.Any("*dto.ArticleOutDto", nil),
				slog.Any("error", err)))
		return dto.ArticleOutDto{}, err
	}
	articlePB := response.Article
	tagPBs := articlePB.GetTags()
	tagDTOs := make([]dto.Tag, 0, len(tagPBs))
	for _, tag := range tagPBs {
		tagDTOs = append(tagDTOs, dto.NewTag(
			tag.Id,
			tag.Name))
	}
	createdAt, err := synchro.Parse[tz.UTC](time.RFC3339Nano, articlePB.CreatedAt)
	if err != nil {
		err = errors.WithStack(err)
		logger.WarnContext(ctx, "END",
			slog.Group("return",
				slog.Any("*dto.ArticleOutDto", nil),
				slog.Any("error", err)))
		return dto.ArticleOutDto{}, err
	}
	updatedAt, err := synchro.Parse[tz.UTC](time.RFC3339Nano, articlePB.UpdatedAt)
	if err != nil {
		err = errors.WithStack(err)
		logger.WarnContext(ctx, "END",
			slog.Group("return",
				slog.Any("*dto.ArticleOutDto", nil),
				slog.Any("error", err)))
		return dto.ArticleOutDto{}, err
	}
	thumbnailURL, err := url.Parse(articlePB.ThumbnailUrl)
	if err != nil {
		err = errors.WithStack(err)
		logger.WarnContext(ctx, "END",
			slog.Group("return",
				slog.Any("*dto.ArticleOutDto", nil),
				slog.Any("error", err)))
		return dto.ArticleOutDto{}, err
	}
	articleDTO := dto.NewArticleTag(
		articlePB.Id,
		articlePB.Title,
		articlePB.Body,
		*thumbnailURL,
		createdAt,
		updatedAt,
		tagDTOs)
	out := dto.NewArticleOutDto(articleDTO)
	logger.InfoContext(ctx, "END",
		slog.Group("return",
			slog.Any("*dto.ArticleOutDto", out),
			slog.Any("error", nil)))
	return out, nil
}

// NewArticle is a constructor of Article.
func NewArticle(articleServiceClient grpc.ArticleServiceClient) *Article {
	return &Article{
		articleServiceClient: articleServiceClient,
	}
}
