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
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/utils"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/cockroachdb/errors"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/app/usecase/dto"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// Articles is a use-case of getting an articles.
type Articles struct {
	// articleServiceClient is a client of article service.
	articleServiceClient grpc.ArticleServiceClient
}

// Execute gets an article by id.
func (u *Articles) Execute(ctx context.Context, in dto.ArticlesInDto) (dto.ArticlesOutDto, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("Execute").End()

	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("articlerameters", slog.Any("in", in)))
	out, err := func() (dto.ArticlesOutDto, error) {
		if in.First() != 0 {
			return u.executeNextPaging(ctx, in)
		} else if in.Last() != 0 {
			return u.executePrevPaging(ctx, in)
		}
		return u.execute(ctx)
	}()
	if err != nil {
		logger.WarnContext(ctx, "END",
			slog.Group("return",
				slog.Any("*dto.ArticleOutDto", out),
				slog.Any("error", err)))
		return out, err
	}
	logger.InfoContext(ctx, "END",
		slog.Group("return",
			slog.Any("*dto.ArticleOutDto", out),
			slog.Any("error", err)))
	return out, nil
}

// executeNextPaging
func (u *Articles) executeNextPaging(ctx context.Context, in dto.ArticlesInDto) (dto.ArticlesOutDto, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("executeNextPaging").End()

	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("articlerameters", slog.Any("in", in)))
	response, err := u.articleServiceClient.GetNextArticles(ctx, &grpc.GetNextArticlesRequest{
		First: int32(in.First()),
		After: utils.PtrFromString(in.After()),
	})
	if err != nil {
		err = errors.WithStack(err)
		logger.WarnContext(ctx, "END",
			slog.Group("return",
				slog.Any("*dto.ArticleOutDto", nil),
				slog.Any("error", err)))
		return dto.ArticlesOutDto{}, err
	}
	articlePBs := response.Articles
	articleDTOs := make([]dto.ArticleTag, 0, len(articlePBs))
	for _, article := range articlePBs {
		tagPBs := article.GetTags()
		tagDTOs := make([]dto.Tag, 0, len(tagPBs))
		for _, tag := range tagPBs {
			tagDTOs = append(tagDTOs, dto.NewTag(
				tag.Id,
				tag.Name))
		}
		createdAt, err := synchro.Parse[tz.UTC](time.RFC3339Nano, article.CreatedAt)
		if err != nil {
			err = errors.WithStack(err)
			logger.WarnContext(ctx, "END",
				slog.Group("return",
					slog.Any("*dto.ArticleOutDto", nil),
					slog.Any("error", err)))
			return dto.ArticlesOutDto{}, err
		}

		updatedAt, err := synchro.Parse[tz.UTC](time.RFC3339Nano, article.UpdatedAt)
		if err != nil {
			err = errors.WithStack(err)
			logger.WarnContext(ctx, "END",
				slog.Group("return",
					slog.Any("*dto.ArticleOutDto", nil),
					slog.Any("error", err)))
			return dto.ArticlesOutDto{}, err
		}

		thumbnailURL, err := url.Parse(article.ThumbnailUrl)
		if err != nil {
			err = errors.WithStack(err)
			logger.WarnContext(ctx, "END",
				slog.Group("return",
					slog.Any("*dto.ArticleOutDto", nil),
					slog.Any("error", err)))
			return dto.ArticlesOutDto{}, err
		}

		articleDTOs = append(articleDTOs, dto.NewArticleTag(
			article.Id,
			article.Title,
			article.Body,
			*thumbnailURL,
			createdAt,
			updatedAt,
			tagDTOs))
	}
	out := dto.NewArticlesOutDto(articleDTOs, dto.ArticlesOutDtoWithHasNext(response.StillExists))
	logger.InfoContext(ctx, "END",
		slog.Group("return",
			slog.Any("*dto.ArticleOutDto", out),
			slog.Any("error", nil)))
	return out, nil
}

// executePrevPaging
func (u *Articles) executePrevPaging(ctx context.Context, in dto.ArticlesInDto) (dto.ArticlesOutDto, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("executePrevPaging").End()

	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("articlerameters", slog.Any("in", in)))

	response, err := u.articleServiceClient.GetPrevArticles(ctx, &grpc.GetPrevArticlesRequest{
		Last:   int32(in.Last()),
		Before: utils.PtrFromString(in.Before()),
	})
	if err != nil {
		err = errors.WithStack(err)
		logger.WarnContext(ctx, "END",
			slog.Group("return",
				slog.Any("*dto.ArticleOutDto", nil),
				slog.Any("error", err)))
		return dto.ArticlesOutDto{}, err
	}
	articlePBs := response.Articles
	articleDTOs := make([]dto.ArticleTag, 0, len(articlePBs))
	for _, article := range articlePBs {
		createdAt, err := synchro.Parse[tz.UTC](time.RFC3339Nano, article.CreatedAt)
		if err != nil {
			err = errors.WithStack(err)
			logger.WarnContext(ctx, "END",
				slog.Group("return",
					slog.Any("*dto.ArticleOutDto", nil),
					slog.Any("error", err)))
			return dto.ArticlesOutDto{}, err
		}

		updatedAt, err := synchro.Parse[tz.UTC](time.RFC3339Nano, article.UpdatedAt)
		if err != nil {
			err = errors.WithStack(err)
			logger.WarnContext(ctx, "END",
				slog.Group("return",
					slog.Any("*dto.ArticleOutDto", nil),
					slog.Any("error", err)))
			return dto.ArticlesOutDto{}, err
		}

		thumbnailURL, err := url.Parse(article.ThumbnailUrl)
		if err != nil {
			err = errors.WithStack(err)
			logger.WarnContext(ctx, "END",
				slog.Group("return",
					slog.Any("*dto.ArticleOutDto", nil),
					slog.Any("error", err)))
			return dto.ArticlesOutDto{}, err
		}

		tagPBs := article.GetTags()
		tagDTOs := make([]dto.Tag, 0, len(tagPBs))
		for _, pt := range article.Tags {
			tagDTOs = append(tagDTOs, dto.NewTag(
				pt.Id,
				pt.Name))
		}
		articleDTOs = append(articleDTOs, dto.NewArticleTag(
			article.Id,
			article.Title,
			article.Body,
			*thumbnailURL,
			createdAt,
			updatedAt,
			tagDTOs))
	}
	out := dto.NewArticlesOutDto(articleDTOs, dto.ArticlesOutDtoWithHasPrev(response.StillExists))
	logger.InfoContext(ctx, "END",
		slog.Group("return",
			slog.Any("*dto.ArticleOutDto", out),
			slog.Any("error", nil)))
	return out, nil
}

// execute
func (u *Articles) execute(ctx context.Context) (dto.ArticlesOutDto, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("execute").End()

	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN")
	response, err := u.articleServiceClient.GetAllArticles(ctx, &emptypb.Empty{})
	if err != nil {
		err = errors.WithStack(err)
		logger.WarnContext(ctx, "END",
			slog.Group("return",
				slog.Any("*dto.ArticleOutDto", nil),
				slog.Any("error", err)))
		return dto.ArticlesOutDto{}, err
	}
	articlePBs := response.Articles
	articleDTOs := make([]dto.ArticleTag, 0, len(articlePBs))

	for _, article := range articlePBs {
		createdAt, err := synchro.Parse[tz.UTC](time.RFC3339Nano, article.CreatedAt)
		if err != nil {
			err = errors.WithStack(err)
			logger.WarnContext(ctx, "END",
				slog.Group("return",
					slog.Any("*dto.ArticleOutDto", nil),
					slog.Any("error", err)))
			return dto.ArticlesOutDto{}, err
		}

		updatedAt, err := synchro.Parse[tz.UTC](time.RFC3339Nano, article.UpdatedAt)
		if err != nil {
			err = errors.WithStack(err)
			logger.WarnContext(ctx, "END",
				slog.Group("return",
					slog.Any("*dto.ArticleOutDto", nil),
					slog.Any("error", err)))
			return dto.ArticlesOutDto{}, err
		}

		thumbnailURL, err := url.Parse(article.ThumbnailUrl)
		if err != nil {
			err = errors.WithStack(err)
			logger.WarnContext(ctx, "END",
				slog.Group("return",
					slog.Any("*dto.ArticleOutDto", nil),
					slog.Any("error", err)))
			return dto.ArticlesOutDto{}, err
		}

		tagPBs := article.GetTags()
		tagDTOs := make([]dto.Tag, 0, len(tagPBs))
		for _, pt := range article.Tags {
			tagDTOs = append(tagDTOs, dto.NewTag(
				pt.Id,
				pt.Name))
		}
		articleDTOs = append(articleDTOs, dto.NewArticleTag(
			article.Id,
			article.Title,
			article.Body,
			*thumbnailURL,
			createdAt,
			updatedAt,
			tagDTOs))
	}
	out := dto.NewArticlesOutDto(articleDTOs)
	logger.InfoContext(ctx, "END",
		slog.Group("return",
			slog.Any("*dto.ArticleOutDto", out),
			slog.Any("error", nil)))
	return out, nil
}

// NewArticles is a constructor of Articles.
func NewArticles(articleServiceClient grpc.ArticleServiceClient) *Articles {
	return &Articles{
		articleServiceClient: articleServiceClient,
	}
}
