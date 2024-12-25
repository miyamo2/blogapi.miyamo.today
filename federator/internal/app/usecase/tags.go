package usecase

import (
	"context"
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	"github.com/cockroachdb/errors"
	"github.com/miyamo2/altnrslog"
	"github.com/miyamo2/blogapi.miyamo.today/core/log"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/app/usecase/dto"
	grpc "github.com/miyamo2/blogapi.miyamo.today/federator/internal/infra/grpc/tag"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/utils"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"
	"google.golang.org/protobuf/types/known/emptypb"
	"log/slog"
	"net/url"
)

// Tags is a use-case of getting tags.
type Tags struct {
	// tagServiceClient is a client of article service.
	tagServiceClient grpc.TagServiceClient
}

// Execute gets a tag by id.
func (u *Tags) Execute(ctx context.Context, in dto.TagsInDTO) (dto.TagsOutDTO, error) {
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
	out, err := func() (dto.TagsOutDTO, error) {
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
				slog.Any("*dto.TagsOutDTO", out),
				slog.Any("error", err)))
		return out, err
	}
	logger.InfoContext(ctx, "END",
		slog.Group("return",
			slog.Any("*dto.TagsOutDTO", out),
			slog.Any("error", err)))
	return out, nil
}

// executeNextPaging
func (u *Tags) executeNextPaging(ctx context.Context, in dto.TagsInDTO) (dto.TagsOutDTO, error) {
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
				slog.Any("*dto.TagsOutDTO", nil),
				slog.Any("error", err)))
		return dto.TagsOutDTO{}, err
	}
	tagPBs := response.Tags
	tagDTOs := make([]dto.TagArticle, 0, len(tagPBs))
	for _, tag := range tagPBs {
		articlePBs := tag.GetArticles()
		articleDTOs := make([]dto.Article, 0, len(articlePBs))
		for _, article := range articlePBs {
			createdAt := synchro.In[tz.UTC](article.CreatedAt.AsTime())
			updatedAt := synchro.In[tz.UTC](article.UpdatedAt.AsTime())

			thumbnailURL, err := url.Parse(article.ThumbnailUrl)
			if err != nil {
				err = errors.WithStack(err)
				logger.WarnContext(ctx, "END",
					slog.Group("return",
						slog.Any("*dto.ArticleOutDTO", nil),
						slog.Any("error", err)))
				return dto.TagsOutDTO{}, err
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
	out := dto.NewTagsOutDTO(tagDTOs, dto.TagsOutDTOWithHasNext(response.StillExists))
	logger.InfoContext(ctx, "END",
		slog.Group("return",
			slog.Any("*dto.TagsOutDTO", out),
			slog.Any("error", nil)))
	return out, nil
}

// executePrevPaging
func (u *Tags) executePrevPaging(ctx context.Context, in dto.TagsInDTO) (dto.TagsOutDTO, error) {
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
				slog.Any("*dto.TagsOutDTO", nil),
				slog.Any("error", err)))
		return dto.TagsOutDTO{}, err
	}
	tagPBs := response.Tags
	tagDTO := make([]dto.TagArticle, 0, len(tagPBs))
	for _, tag := range tagPBs {
		articlePBs := tag.GetArticles()
		articleDTOs := make([]dto.Article, 0, len(articlePBs))
		for _, article := range articlePBs {
			createdAt := synchro.In[tz.UTC](article.CreatedAt.AsTime())
			updatedAt := synchro.In[tz.UTC](article.UpdatedAt.AsTime())

			thumbnailURL, err := url.Parse(article.ThumbnailUrl)
			if err != nil {
				err = errors.WithStack(err)
				logger.WarnContext(ctx, "END",
					slog.Group("return",
						slog.Any("*dto.ArticleOutDTO", nil),
						slog.Any("error", err)))
				return dto.TagsOutDTO{}, err
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
	out := dto.NewTagsOutDTO(tagDTO, dto.TagsOutDTOWithHasPrev(response.StillExists))
	logger.InfoContext(ctx, "END",
		slog.Group("return",
			slog.Any("*dto.TagsOutDTO", out),
			slog.Any("error", nil)))
	return out, nil
}

// execute
func (u *Tags) execute(ctx context.Context) (dto.TagsOutDTO, error) {
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
				slog.Any("*dto.TagsOutDTO", nil),
				slog.Any("error", err)))
		return dto.TagsOutDTO{}, err
	}
	tagPBs := response.Tags
	tagDTOs := make([]dto.TagArticle, 0, len(tagPBs))
	for _, tag := range tagPBs {
		articlePBs := tag.GetArticles()
		articleDTOs := make([]dto.Article, 0, len(articlePBs))
		for _, article := range tag.Articles {
			createdAt := synchro.In[tz.UTC](article.CreatedAt.AsTime())
			updatedAt := synchro.In[tz.UTC](article.UpdatedAt.AsTime())

			thumbnailURL, err := url.Parse(article.ThumbnailUrl)
			if err != nil {
				err = errors.WithStack(err)
				logger.WarnContext(ctx, "END",
					slog.Group("return",
						slog.Any("*dto.ArticleOutDTO", nil),
						slog.Any("error", err)))
				return dto.TagsOutDTO{}, err
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
	out := dto.NewTagsOutDTO(tagDTOs)
	logger.InfoContext(ctx, "END",
		slog.Group("return",
			slog.Any("*dto.TagsOutDTO", out),
			slog.Any("error", nil)))
	return out, nil
}

// NewTag is a constructor of Tag.
func NewTags(tagServiceClient grpc.TagServiceClient) *Tags {
	return &Tags{
		tagServiceClient: tagServiceClient,
	}
}
