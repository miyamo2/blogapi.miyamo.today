package usecase

import (
	"context"
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	"log/slog"
	"net/url"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/miyamo2/altnrslog"
	"github.com/miyamo2/blogapi.miyamo.today/core/log"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/app/usecase/dto"
	grpc "github.com/miyamo2/blogapi.miyamo.today/federator/internal/infra/grpc/tag"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/utils"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Tags is a use-case of getting tags.
type Tags struct {
	// tagServiceClient is a client of article service.
	tagServiceClient grpc.TagServiceClient
}

// Execute gets a tag by id.
func (u *Tags) Execute(ctx context.Context, in dto.TagsInDto) (dto.TagsOutDto, error) {
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
	out, err := func() (dto.TagsOutDto, error) {
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
				slog.Any("*dto.TagsOutDto", out),
				slog.Any("error", err)))
		return out, err
	}
	logger.InfoContext(ctx, "END",
		slog.Group("return",
			slog.Any("*dto.TagsOutDto", out),
			slog.Any("error", err)))
	return out, nil
}

// executeNextPaging
func (u *Tags) executeNextPaging(ctx context.Context, in dto.TagsInDto) (dto.TagsOutDto, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("executeNextPaging").End()

	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.Any("in", in)))

	response, err := u.tagServiceClient.GetNextTags(ctx, &grpc.GetNextTagsRequest{
		First: int32(in.First()),
		After: utils.PtrFromString(in.After()),
	})
	if err != nil {
		err = errors.WithStack(err)
		logger.WarnContext(ctx, "END",
			slog.Group("return",
				slog.Any("*dto.TagsOutDto", nil),
				slog.Any("error", err)))
		return dto.TagsOutDto{}, err
	}
	tagPBs := response.Tags
	tagDTOs := make([]dto.TagArticle, 0, len(tagPBs))
	for _, tag := range tagPBs {
		articlePBs := tag.GetArticles()
		articleDTOs := make([]dto.Article, 0, len(articlePBs))
		for _, article := range articlePBs {
			createdAt, err := synchro.Parse[tz.UTC](time.RFC3339Nano, article.CreatedAt)
			if err != nil {
				err = errors.WithStack(err)
				logger.WarnContext(ctx, "END",
					slog.Group("return",
						slog.Any("*dto.ArticleOutDto", nil),
						slog.Any("error", err)))
				return dto.TagsOutDto{}, err
			}

			updatedAt, err := synchro.Parse[tz.UTC](time.RFC3339Nano, article.UpdatedAt)
			if err != nil {
				err = errors.WithStack(err)
				logger.WarnContext(ctx, "END",
					slog.Group("return",
						slog.Any("*dto.ArticleOutDto", nil),
						slog.Any("error", err)))
				return dto.TagsOutDto{}, err
			}

			thumbnailURL, err := url.Parse(article.ThumbnailUrl)
			if err != nil {
				err = errors.WithStack(err)
				logger.WarnContext(ctx, "END",
					slog.Group("return",
						slog.Any("*dto.ArticleOutDto", nil),
						slog.Any("error", err)))
				return dto.TagsOutDto{}, err
			}

			articleDTOs = append(articleDTOs, dto.NewArticle(
				article.Id,
				article.Title,
				"",
				*thumbnailURL,
				createdAt,
				updatedAt))
		}
		tagDTOs = append(tagDTOs, dto.NewTagArticle(
			tag.Id,
			tag.Name,
			articleDTOs))
	}
	out := dto.NewTagsOutDto(tagDTOs, dto.TagsOutDtoWithHasNext(response.StillExists))
	logger.InfoContext(ctx, "END",
		slog.Group("return",
			slog.Any("*dto.TagsOutDto", out),
			slog.Any("error", nil)))
	return out, nil
}

// executePrevPaging
func (u *Tags) executePrevPaging(ctx context.Context, in dto.TagsInDto) (dto.TagsOutDto, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("executePrevPaging").End()

	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.Any("in", in)))

	response, err := u.tagServiceClient.GetPrevTags(ctx, &grpc.GetPrevTagsRequest{
		Last:   int32(in.Last()),
		Before: utils.PtrFromString(in.Before()),
	})
	if err != nil {
		err = errors.WithStack(err)
		logger.WarnContext(ctx, "END",
			slog.Group("return",
				slog.Any("*dto.TagsOutDto", nil),
				slog.Any("error", err)))
		return dto.TagsOutDto{}, err
	}
	tagPBs := response.Tags
	tagDTO := make([]dto.TagArticle, 0, len(tagPBs))
	for _, tag := range tagPBs {
		articlePBs := tag.GetArticles()
		articleDTOs := make([]dto.Article, 0, len(articlePBs))
		for _, article := range articlePBs {
			createdAt, err := synchro.Parse[tz.UTC](time.RFC3339Nano, article.CreatedAt)
			if err != nil {
				err = errors.WithStack(err)
				logger.WarnContext(ctx, "END",
					slog.Group("return",
						slog.Any("*dto.ArticleOutDto", nil),
						slog.Any("error", err)))
				return dto.TagsOutDto{}, err
			}

			updatedAt, err := synchro.Parse[tz.UTC](time.RFC3339Nano, article.UpdatedAt)
			if err != nil {
				err = errors.WithStack(err)
				logger.WarnContext(ctx, "END",
					slog.Group("return",
						slog.Any("*dto.ArticleOutDto", nil),
						slog.Any("error", err)))
				return dto.TagsOutDto{}, err
			}

			thumbnailURL, err := url.Parse(article.ThumbnailUrl)
			if err != nil {
				err = errors.WithStack(err)
				logger.WarnContext(ctx, "END",
					slog.Group("return",
						slog.Any("*dto.ArticleOutDto", nil),
						slog.Any("error", err)))
				return dto.TagsOutDto{}, err
			}

			articleDTOs = append(articleDTOs, dto.NewArticle(
				article.Id,
				article.Title,
				"",
				*thumbnailURL,
				createdAt,
				updatedAt))
		}
		tagDTO = append(tagDTO, dto.NewTagArticle(
			tag.Id,
			tag.Name,
			articleDTOs))
	}
	out := dto.NewTagsOutDto(tagDTO, dto.TagsOutDtoWithHasPrev(response.StillExists))
	logger.InfoContext(ctx, "END",
		slog.Group("return",
			slog.Any("*dto.TagsOutDto", out),
			slog.Any("error", nil)))
	return out, nil
}

// execute
func (u *Tags) execute(ctx context.Context) (dto.TagsOutDto, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("execute").End()

	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN")
	response, err := u.tagServiceClient.GetAllTags(ctx, &emptypb.Empty{})
	if err != nil {
		err = errors.WithStack(err)
		logger.WarnContext(ctx, "END",
			slog.Group("return",
				slog.Any("*dto.TagsOutDto", nil),
				slog.Any("error", err)))
		return dto.TagsOutDto{}, err
	}
	tagPBs := response.Tags
	tagDTOs := make([]dto.TagArticle, 0, len(tagPBs))
	for _, tag := range tagPBs {
		articlePBs := tag.GetArticles()
		articleDTOs := make([]dto.Article, 0, len(articlePBs))
		for _, article := range tag.Articles {
			createdAt, err := synchro.Parse[tz.UTC](time.RFC3339Nano, article.CreatedAt)
			if err != nil {
				err = errors.WithStack(err)
				logger.WarnContext(ctx, "END",
					slog.Group("return",
						slog.Any("*dto.ArticleOutDto", nil),
						slog.Any("error", err)))
				return dto.TagsOutDto{}, err
			}

			updatedAt, err := synchro.Parse[tz.UTC](time.RFC3339Nano, article.UpdatedAt)
			if err != nil {
				err = errors.WithStack(err)
				logger.WarnContext(ctx, "END",
					slog.Group("return",
						slog.Any("*dto.ArticleOutDto", nil),
						slog.Any("error", err)))
				return dto.TagsOutDto{}, err
			}

			thumbnailURL, err := url.Parse(article.ThumbnailUrl)
			if err != nil {
				err = errors.WithStack(err)
				logger.WarnContext(ctx, "END",
					slog.Group("return",
						slog.Any("*dto.ArticleOutDto", nil),
						slog.Any("error", err)))
				return dto.TagsOutDto{}, err
			}

			articleDTOs = append(articleDTOs, dto.NewArticle(
				article.Id,
				article.Title,
				"",
				*thumbnailURL,
				createdAt,
				updatedAt))
		}
		tagDTOs = append(tagDTOs, dto.NewTagArticle(
			tag.Id,
			tag.Name,
			articleDTOs))
	}
	out := dto.NewTagsOutDto(tagDTOs)
	logger.InfoContext(ctx, "END",
		slog.Group("return",
			slog.Any("*dto.TagsOutDto", out),
			slog.Any("error", nil)))
	return out, nil
}

// NewTag is a constructor of Tag.
func NewTags(tagServiceClient grpc.TagServiceClient) *Tags {
	return &Tags{
		tagServiceClient: tagServiceClient,
	}
}
