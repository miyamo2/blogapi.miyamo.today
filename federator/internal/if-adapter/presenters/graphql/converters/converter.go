package converters

import (
	"blogapi.miyamo.today/core/log"
	"blogapi.miyamo.today/federator/internal/pkg/gqlscalar"
	"context"
	"github.com/cockroachdb/errors"
	"github.com/miyamo2/altnrslog"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"log/slog"

	"blogapi.miyamo.today/federator/internal/app/usecase/dto"
	"blogapi.miyamo.today/federator/internal/if-adapter/presenters/graphql/model"
	"github.com/newrelic/go-agent/v3/newrelic"
)

var (
	ErrFailedToConvertToTagNode = errors.New("failed to convert to tag node")
)

type Converter struct{}

func (c Converter) ToCreateArticle(ctx context.Context, from dto.CreateArticleOutDTO) (*model.CreateArticlePayload, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToCreateArticle").End()

	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.Any("from", from)))

	var clientMutationID *string
	if v := from.ClientMutationID(); len(v) > 0 {
		clientMutationID = &v
	}
	payload := model.CreateArticlePayload{
		ClientMutationID: clientMutationID,
		EventID:          from.EventID(),
		ArticleID:        from.ArticleID(),
	}
	logger.InfoContext(ctx, "END",
		slog.Group("returns",
			slog.Any("*model.CreateArticlePayload", payload),
			slog.Any("error", nil)))
	return &payload, nil
}

// ToArticle converts dto.ArticleOutDTO to model.ArticleNode.
func (c Converter) ToArticle(ctx context.Context, from dto.ArticleOutDTO) (*model.ArticleNode, bool) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToArticle").End()

	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.Any("from", from)))
	node, err := c.articleNodeFromArticleTagDTO(ctx, from.Article())
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger.WarnContext(ctx, "END",
			slog.Group("returns",
				slog.Any("*model.ArticleNode", nil),
				slog.Any("error", err)))
		return nil, false
	}
	return node, true
}

// articleNodeFromArticleTagDTO converts dto.ArticleTag to model.ArticleNode.
func (c Converter) articleNodeFromArticleTagDTO(ctx context.Context, from dto.ArticleTag) (*model.ArticleNode, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("articleNodeFromArticleTagDTO").End()

	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.Any("from", from)))
	tegs := make([]*model.ArticleTagEdge, 0, len(from.Tags()))
	for _, tag := range from.Tags() {
		tegs = append(tegs, &model.ArticleTagEdge{
			Cursor: tag.ID(),
			Node: &model.ArticleTagNode{
				ID:   tag.ID(),
				Name: tag.Name(),
			},
		})
	}
	// TODO: check HasNextPage, HasPreviousPage
	tagPageInfo := func() model.PageInfo {
		if len(tegs) == 0 {
			return model.PageInfo{}
		}
		return model.PageInfo{
			StartCursor: tegs[0].Cursor,
			EndCursor:   tegs[len(tegs)-1].Cursor,
		}
	}()
	tagConnection := model.ArticleTagConnection{
		Edges:      tegs,
		PageInfo:   &tagPageInfo,
		TotalCount: len(tegs),
	}
	articleNode := model.ArticleNode{
		ID:           from.ID(),
		Title:        from.Title(),
		Content:      from.Body(),
		ThumbnailURL: gqlscalar.URL(from.ThumbnailURL()),
		CreatedAt:    gqlscalar.UTC(from.CreatedAt()),
		UpdatedAt:    gqlscalar.UTC(from.UpdatedAt()),
		Tags:         &tagConnection,
	}
	logger.InfoContext(ctx, "END",
		slog.Group("parameters",
			slog.Any("*model.ArticleNode", articleNode),
			slog.Any("error", nil)))
	return &articleNode, nil
}

func (c Converter) ToArticles(ctx context.Context, from dto.ArticlesOutDTO) (*model.ArticleConnection, bool) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToArticles").End()

	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.Any("from", from)))
	articleEdges := make([]*model.ArticleEdge, 0, len(from.Articles()))
	for _, article := range from.Articles() {
		node, err := c.articleNodeFromArticleTagDTO(ctx, article)
		if err != nil {
			err = errors.WithStack(err)
			nrtx.NoticeError(nrpkgerrors.Wrap(err))
			logger.WarnContext(ctx, "END",
				slog.Group("returns",
					slog.Any("*model.ArticleConnection", nil),
					slog.Any("error", err)))
			return nil, false
		}
		articleEdges = append(articleEdges, &model.ArticleEdge{
			Cursor: article.ID(),
			Node:   node,
		})
	}
	pageInfo := func() model.PageInfo {
		if len(articleEdges) == 0 {
			return model.PageInfo{}
		}
		if from.ByForward() {
			hasNext := from.HasNext()
			return model.PageInfo{
				StartCursor: articleEdges[0].Cursor,
				EndCursor:   articleEdges[len(articleEdges)-1].Cursor,
				HasNextPage: &hasNext,
			}
		}
		if from.ByBackward() {
			hasPrevious := from.HasPrev()
			return model.PageInfo{
				StartCursor:     articleEdges[0].Cursor,
				EndCursor:       articleEdges[len(articleEdges)-1].Cursor,
				HasPreviousPage: &hasPrevious,
			}
		}
		return model.PageInfo{
			StartCursor: articleEdges[0].Cursor,
			EndCursor:   articleEdges[len(articleEdges)-1].Cursor,
		}
	}()
	connection := model.ArticleConnection{
		Edges:      articleEdges,
		PageInfo:   &pageInfo,
		TotalCount: len(articleEdges),
	}
	logger.InfoContext(ctx, "END",
		slog.Group("returns",
			slog.Any("*model.ArticleConnection", connection),
			slog.Any("bool", true)))
	return &connection, true
}

// ToTag converts dto.TagOutDTO to model.TagNode.
func (c Converter) ToTag(ctx context.Context, from dto.TagOutDTO) (*model.TagNode, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToTag").End()

	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.Any("from", from)))
	node, err := c.tagNodeFromTagArticleDTO(ctx, from.Tag())
	if err != nil {
		err = errors.Join(err, ErrFailedToConvertToTagNode)
		logger.WarnContext(ctx, "END",
			slog.Group("returns",
				slog.Any("*model.TagNode", nil),
				slog.Any("error", err)))
		return nil, err
	}
	logger.InfoContext(ctx, "END",
		slog.Group("returns",
			slog.Any("*model.TagNode", node),
			slog.Any("bool", true)))
	return node, nil
}

// tagNodeFromTagArticleDTO converts dto.TagArticle to model.TagNode.
func (c Converter) tagNodeFromTagArticleDTO(ctx context.Context, from dto.TagArticle) (*model.TagNode, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("tagNodeFromTagArticleDTO").End()

	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.Any("from", from)))
	articleEdges := make([]*model.TagArticleEdge, 0, len(from.Articles()))
	for _, article := range from.Articles() {
		articleEdges = append(articleEdges, &model.TagArticleEdge{
			Cursor: article.ID(),
			Node: &model.TagArticleNode{
				ID:           article.ID(),
				Title:        article.Title(),
				ThumbnailURL: gqlscalar.URL(article.ThumbnailURL()),
				CreatedAt:    gqlscalar.UTC(article.CreatedAt()),
				UpdatedAt:    gqlscalar.UTC(article.UpdatedAt()),
			},
		})
	}
	pageInfo := func() model.PageInfo {
		if len(articleEdges) == 0 {
			return model.PageInfo{}
		}
		// TODO: check HasNextPage, HasPreviousPage
		return model.PageInfo{
			StartCursor: articleEdges[0].Cursor,
			EndCursor:   articleEdges[len(articleEdges)-1].Cursor,
		}
	}()
	articleConnection := model.TagArticleConnection{
		Edges:      articleEdges,
		PageInfo:   &pageInfo,
		TotalCount: len(articleEdges),
	}
	node := model.TagNode{
		ID:       from.ID(),
		Name:     from.Name(),
		Articles: &articleConnection,
	}
	logger.InfoContext(ctx, "END",
		slog.Group("parameters",
			slog.Any("*model.TagNode", node),
			slog.Any("error", nil)))
	return &node, nil
}

func (c Converter) ToTags(ctx context.Context, from dto.TagsOutDTO) (*model.TagConnection, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToTags").End()

	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.Any("from", from)))
	tegs := make([]*model.TagEdge, 0, len(from.Tags()))
	for _, tag := range from.Tags() {
		node, err := c.tagNodeFromTagArticleDTO(ctx, tag)
		if err != nil {
			err = errors.WithStack(err)
			nrtx.NoticeError(nrpkgerrors.Wrap(err))
			logger.WarnContext(ctx, "END",
				slog.Group("returns",
					slog.Any("*model.TagConnection", nil),
					slog.Any("error", err)))
			return nil, err
		}
		tegs = append(tegs, &model.TagEdge{
			Cursor: tag.ID(),
			Node:   node,
		})
	}
	pageInfo := func() model.PageInfo {
		if len(tegs) == 0 {
			return model.PageInfo{}
		}
		if from.ByForward() {
			hasNext := from.HasNext()
			return model.PageInfo{
				StartCursor: tegs[0].Cursor,
				EndCursor:   tegs[len(tegs)-1].Cursor,
				HasNextPage: &hasNext,
			}
		}
		if from.ByBackward() {
			hasPrevious := from.HasPrev()
			return model.PageInfo{
				StartCursor:     tegs[0].Cursor,
				EndCursor:       tegs[len(tegs)-1].Cursor,
				HasPreviousPage: &hasPrevious,
			}
		}
		return model.PageInfo{
			StartCursor: tegs[0].Cursor,
			EndCursor:   tegs[len(tegs)-1].Cursor,
		}
	}()
	connection := model.TagConnection{
		Edges:      tegs,
		PageInfo:   &pageInfo,
		TotalCount: len(tegs),
	}
	logger.InfoContext(ctx, "END",
		slog.Group("returns",
			slog.Any("*model.TagConnection", connection),
			slog.Any("error", nil)))
	return &connection, nil
}

func (c Converter) ToUpdateArticleTitle(ctx context.Context, from dto.UpdateArticleTitleOutDTO) (*model.UpdateArticleTitlePayload, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToUpdateArticleTitle").End()

	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.Any("from", from)))
	var clientMutationID *string
	if v := from.ClientMutationID(); len(v) > 0 {
		clientMutationID = &v
	}
	payload := model.UpdateArticleTitlePayload{
		ClientMutationID: clientMutationID,
		EventID:          from.EventID(),
		ArticleID:        from.ArticleID(),
	}
	logger.InfoContext(ctx, "END",
		slog.Group("returns",
			slog.Any("*model.UpdateArticleTitlePayload", payload),
			slog.Any("error", nil)))
	return &payload, nil
}

func (c Converter) ToUpdateArticleBody(ctx context.Context, from dto.UpdateArticleBodyOutDTO) (*model.UpdateArticleBodyPayload, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToUpdateArticleBody").End()

	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.Any("from", from)))
	var clientMutationID *string
	if v := from.ClientMutationID(); len(v) > 0 {
		clientMutationID = &v
	}
	payload := model.UpdateArticleBodyPayload{
		ClientMutationID: clientMutationID,
		EventID:          from.EventID(),
		ArticleID:        from.ArticleID(),
	}
	logger.InfoContext(ctx, "END",
		slog.Group("returns",
			slog.Any("*model.UpdateArticleBodyPayload", payload),
			slog.Any("error", nil)))
	return &payload, nil
}

func (c Converter) ToUpdateArticleThumbnail(ctx context.Context, from dto.UpdateArticleThumbnailOutDTO) (*model.UpdateArticleThumbnailPayload, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToUpdateArticleThumbnail").End()

	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.Any("from", from)))
	var clientMutationID *string
	if v := from.ClientMutationID(); len(v) > 0 {
		clientMutationID = &v
	}
	payload := model.UpdateArticleThumbnailPayload{
		ClientMutationID: clientMutationID,
		EventID:          from.EventID(),
		ArticleID:        from.ArticleID(),
	}
	logger.InfoContext(ctx, "END",
		slog.Group("returns",
			slog.Any("*model.UpdateArticleThumbnailPayload", payload),
			slog.Any("error", nil)))
	return &payload, nil
}

func (c Converter) ToAttachTags(ctx context.Context, from dto.AttachTagsOutDTO) (*model.AttachTagsPayload, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToAttachTags").End()

	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.Any("from", from)))
	var clientMutationID *string
	if v := from.ClientMutationID(); len(v) > 0 {
		clientMutationID = &v
	}
	payload := model.AttachTagsPayload{
		ClientMutationID: clientMutationID,
		EventID:          from.EventID(),
		ArticleID:        from.ArticleID(),
	}
	logger.InfoContext(ctx, "END",
		slog.Group("returns",
			slog.Any("*model.AttachTagsPayload", payload),
			slog.Any("error", nil)))
	return &payload, nil
}

func (c Converter) ToDetachTags(ctx context.Context, from dto.DetachTagsOutDTO) (*model.DetachTagsPayload, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToDetachTags").End()

	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.Any("from", from)))
	var clientMutationID *string
	if v := from.ClientMutationID(); len(v) > 0 {
		clientMutationID = &v
	}
	payload := model.DetachTagsPayload{
		ClientMutationID: clientMutationID,
		EventID:          from.EventID(),
		ArticleID:        from.ArticleID(),
	}
	logger.InfoContext(ctx, "END",
		slog.Group("returns",
			slog.Any("*model.DetachTagsPayload", payload),
			slog.Any("error", nil)))
	return &payload, nil
}

func (c Converter) ToUploadImage(ctx context.Context, from dto.UploadImageOutDTO) (*model.UploadImagePayload, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToUploadImage").End()

	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.Any("from", from)))
	var clientMutationID *string
	if v := from.ClientMutationID(); len(v) > 0 {
		clientMutationID = &v
	}
	payload := model.UploadImagePayload{
		ImageURL:         gqlscalar.URL(from.ImageURL()),
		ClientMutationID: clientMutationID,
	}
	logger.InfoContext(ctx, "END",
		slog.Group("returns",
			slog.Any("*model.UploadImagePayload", payload),
			slog.Any("error", nil)))
	return &payload, nil
}

func NewConverter() *Converter {
	return &Converter{}
}
