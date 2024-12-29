package usecase

import (
	"blogapi.miyamo.today/core/log"
	grpc "blogapi.miyamo.today/federator/internal/infra/grpc/tag"
	"context"
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	"github.com/miyamo2/altnrslog"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"log/slog"
	"net/url"

	"blogapi.miyamo.today/federator/internal/app/usecase/dto"
	"github.com/cockroachdb/errors"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// Tag is a use-case of getting a tag by id.
type Tag struct {
	// tagServiceClient is a client of article service.
	tagServiceClient grpc.TagServiceClient
}

// Execute gets a tag by id.
func (u *Tag) Execute(ctx context.Context, in dto.TagInDTO) (dto.TagOutDTO, error) {
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
	response, err := u.tagServiceClient.GetTagById(
		newrelic.NewContext(ctx, nrtx),
		&grpc.GetTagByIdRequest{
			Id: in.ID(),
		})
	if err != nil {
		err = errors.WithStack(err)
		logger.WarnContext(ctx, "END",
			slog.Group("return",
				slog.Any("*dto.ArticleOutDTO", nil),
				slog.Any("error", err)))
		return dto.TagOutDTO{}, err
	}
	tagPB := response.Tag
	articlePBs := tagPB.Articles
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
			return dto.TagOutDTO{}, err
		}

		articleDTOs = append(articleDTOs, dto.NewArticle(
			article.Id,
			article.Title,
			"",
			*thumbnailURL,
			createdAt,
			updatedAt))
	}
	tagDTO := dto.NewTagArticle(
		tagPB.Id,
		tagPB.Name,
		articleDTOs)
	out := dto.NewTagOutDTO(tagDTO)
	logger.InfoContext(ctx, "END",

		slog.Group("return",
			slog.Any("*dto.TagOutDTO", out),
			slog.Any("error", nil)))
	return out, nil
}

// NewTag is a constructor of Tag.
func NewTag(tagServiceClient grpc.TagServiceClient) *Tag {
	return &Tag{
		tagServiceClient: tagServiceClient,
	}
}
