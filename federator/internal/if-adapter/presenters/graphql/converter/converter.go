package converter

import (
	"context"
	"log/slog"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/miyamo2/altnrslog"
	"github.com/miyamo2/api.miyamo.today/core/log"
	"github.com/miyamo2/api.miyamo.today/core/util/duration"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"

	"github.com/miyamo2/api.miyamo.today/federator/internal/app/usecase/dto"
	"github.com/miyamo2/api.miyamo.today/federator/internal/if-adapter/presenters/graphql/model"
	"github.com/newrelic/go-agent/v3/newrelic"
)

const (
	timeLayout = "2006-01-02T15:04:05Z"
)

var (
	ErrParseTime                = errors.New("failed to parse time")
	ErrFailedToConvertToTagNode = errors.New("failed to convert to tag node")
)

type Converter struct{}

// ToArticle converts dto.ArticleOutDto to model.ArticleNode.
func (c Converter) ToArticle(ctx context.Context, from dto.ArticleOutDto) (*model.ArticleNode, bool) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToArticle").End()
	dw := duration.Start()
	lgr, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		lgr = log.DefaultLogger()
	}
	lgr.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.Any("from", from)))
	n, err := c.articleNodeFromArticleTagDto(ctx, from.Article())
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		lgr.WarnContext(ctx, "END",
			slog.String("duration", dw.SDuration()),
			slog.Group("returns",
				slog.Any("*model.ArticleNode", nil),
				slog.Any("error", err)))
		return nil, false
	}
	return n, true
}

// articleNodeFromArticleTagDto converts dto.ArticleTag to model.ArticleNode.
func (c Converter) articleNodeFromArticleTagDto(ctx context.Context, from dto.ArticleTag) (*model.ArticleNode, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("articleNodeFromArticleTagDto").End()
	dw := duration.Start()
	lgr, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		lgr = log.DefaultLogger()
	}
	lgr.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.Any("from", from)))
	tegs := make([]*model.ArticleTagEdge, 0, len(from.Tags()))
	for _, tag := range from.Tags() {
		tegs = append(tegs, &model.ArticleTagEdge{
			Cursor: tag.Id(),
			Node: &model.ArticleTagNode{
				ID:   tag.Id(),
				Name: tag.Name(),
			},
		})
	}
	// TODO: check HasNextPage, HasPreviousPage
	tpg := func() model.PageInfo {
		if len(tegs) == 0 {
			return model.PageInfo{}
		}
		return model.PageInfo{
			StartCursor: tegs[0].Cursor,
			EndCursor:   tegs[len(tegs)-1].Cursor,
		}
	}()
	tcnn := model.ArticleTagConnection{
		Edges:      tegs,
		PageInfo:   &tpg,
		TotalCount: len(tegs),
	}
	crtd, err := time.Parse(timeLayout, from.CreatedAt())
	if err != nil {
		err = errors.Join(err, ErrParseTime)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		lgr.WarnContext(ctx, "END",
			slog.String("duration", dw.SDuration()),
			slog.Group("parameters",
				slog.Any("*model.ArticleNode", nil),
				slog.Any("error", err)))
		return nil, err
	}
	updtd, err := time.Parse(timeLayout, from.UpdatedAt())
	if err != nil {
		err = errors.Join(err, ErrParseTime)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		lgr.WarnContext(ctx, "END",
			slog.String("duration", dw.SDuration()),
			slog.Group("parameters",
				slog.Any("*model.ArticleNode", nil),
				slog.Any("error", err)))
		return nil, err
	}
	an := model.ArticleNode{
		ID:           from.Id(),
		Title:        from.Title(),
		Content:      from.Body(),
		ThumbnailURL: from.ThumbnailUrl(),
		CreatedAt:    crtd,
		UpdatedAt:    updtd,
		Tags:         &tcnn,
	}
	lgr.InfoContext(ctx, "END",
		slog.String("duration", dw.SDuration()),
		slog.Group("parameters",
			slog.Any("*model.ArticleNode", an),
			slog.Any("error", nil)))
	return &an, nil
}

func (c Converter) ToArticles(ctx context.Context, from dto.ArticlesOutDto) (*model.ArticleConnection, bool) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToArticles").End()
	dw := duration.Start()
	lgr, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		lgr = log.DefaultLogger()
	}
	lgr.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.Any("from", from)))
	aegs := make([]*model.ArticleEdge, 0, len(from.Articles()))
	for _, article := range from.Articles() {
		n, err := c.articleNodeFromArticleTagDto(ctx, article)
		if err != nil {
			err = errors.WithStack(err)
			nrtx.NoticeError(nrpkgerrors.Wrap(err))
			lgr.WarnContext(ctx, "END",
				slog.String("duration", dw.SDuration()),
				slog.Group("returns",
					slog.Any("*model.ArticleConnection", nil),
					slog.Any("error", err)))
			return nil, false
		}
		aegs = append(aegs, &model.ArticleEdge{
			Cursor: article.Id(),
			Node:   n,
		})
	}
	pg := func() model.PageInfo {
		if len(aegs) == 0 {
			return model.PageInfo{}
		}
		if from.ByForward() {
			hn := from.HasNext()
			return model.PageInfo{
				StartCursor: aegs[0].Cursor,
				EndCursor:   aegs[len(aegs)-1].Cursor,
				HasNextPage: &hn,
			}
		}
		if from.ByBackward() {
			hp := from.HasPrev()
			return model.PageInfo{
				StartCursor:     aegs[0].Cursor,
				EndCursor:       aegs[len(aegs)-1].Cursor,
				HasPreviousPage: &hp,
			}
		}
		return model.PageInfo{
			StartCursor: aegs[0].Cursor,
			EndCursor:   aegs[len(aegs)-1].Cursor,
		}
	}()
	cnctn := model.ArticleConnection{
		Edges:      aegs,
		PageInfo:   &pg,
		TotalCount: len(aegs),
	}
	lgr.InfoContext(ctx, "END",
		slog.String("duration", dw.SDuration()),
		slog.Group("returns",
			slog.Any("*model.ArticleConnection", cnctn),
			slog.Any("bool", true)))
	return &cnctn, true
}

// ToTag converts dto.TagOutDto to model.TagNode.
func (c Converter) ToTag(ctx context.Context, from dto.TagOutDto) (*model.TagNode, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToTag").End()
	dw := duration.Start()
	lgr, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		lgr = log.DefaultLogger()
	}
	lgr.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.Any("from", from)))
	n, err := c.tagNodeFromTagArticleDto(ctx, from.Tag())
	if err != nil {
		err = errors.Join(err, ErrFailedToConvertToTagNode)
		lgr.WarnContext(ctx, "END",
			slog.String("duration", dw.SDuration()),
			slog.Group("returns",
				slog.Any("*model.TagNode", nil),
				slog.Any("error", err)))
		return nil, err
	}
	lgr.InfoContext(ctx, "END",
		slog.String("duration", dw.SDuration()),
		slog.Group("returns",
			slog.Any("*model.TagNode", n),
			slog.Any("bool", true)))
	return n, nil
}

// tagNodeFromTagArticleDto converts dto.TagArticle to model.TagNode.
func (c Converter) tagNodeFromTagArticleDto(ctx context.Context, from dto.TagArticle) (*model.TagNode, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("tagNodeFromTagArticleDto").End()
	dw := duration.Start()
	lgr, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		lgr = log.DefaultLogger()
	}
	lgr.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.Any("from", from)))
	aegs := make([]*model.TagArticleEdge, 0, len(from.Articles()))
	for _, article := range from.Articles() {
		crtd, err := time.Parse(timeLayout, article.CreatedAt())
		if err != nil {
			err = errors.Join(err, ErrParseTime)
			nrtx.NoticeError(nrpkgerrors.Wrap(err))
			lgr.WarnContext(ctx, "END",
				slog.String("duration", dw.SDuration()),
				slog.Group("parameters",
					slog.Any("*model.ArticleNode", nil),
					slog.Any("error", err)))
			return nil, err
		}
		updtd, err := time.Parse(timeLayout, article.UpdatedAt())
		if err != nil {
			err = errors.Join(err, ErrParseTime)
			nrtx.NoticeError(nrpkgerrors.Wrap(err))
			lgr.WarnContext(ctx, "END",
				slog.String("duration", dw.SDuration()),
				slog.Group("parameters",
					slog.Any("*model.ArticleNode", nil),
					slog.Any("error", err)))
			return nil, err
		}
		aegs = append(aegs, &model.TagArticleEdge{
			Cursor: article.Id(),
			Node: &model.TagArticleNode{
				ID:           article.Id(),
				Title:        article.Title(),
				ThumbnailURL: article.ThumbnailUrl(),
				CreatedAt:    crtd,
				UpdatedAt:    updtd,
			},
		})
	}
	pg := func() model.PageInfo {
		if len(aegs) == 0 {
			return model.PageInfo{}
		}
		// TODO: check HasNextPage, HasPreviousPage
		return model.PageInfo{
			StartCursor: aegs[0].Cursor,
			EndCursor:   aegs[len(aegs)-1].Cursor,
		}
	}()
	acnctn := model.TagArticleConnection{
		Edges:      aegs,
		PageInfo:   &pg,
		TotalCount: len(aegs),
	}
	n := model.TagNode{
		ID:       from.Id(),
		Name:     from.Name(),
		Articles: &acnctn,
	}
	lgr.InfoContext(ctx, "END",
		slog.String("duration", dw.SDuration()),
		slog.Group("parameters",
			slog.Any("*model.TagNode", n),
			slog.Any("error", nil)))
	return &n, nil
}

func (c Converter) ToTags(ctx context.Context, from dto.TagsOutDto) (*model.TagConnection, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToTags").End()
	dw := duration.Start()
	lgr, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		lgr = log.DefaultLogger()
	}
	lgr.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.Any("from", from)))
	tegs := make([]*model.TagEdge, 0, len(from.Tags()))
	for _, tag := range from.Tags() {
		n, err := c.tagNodeFromTagArticleDto(ctx, tag)
		if err != nil {
			err = errors.WithStack(err)
			nrtx.NoticeError(nrpkgerrors.Wrap(err))
			lgr.WarnContext(ctx, "END",
				slog.String("duration", dw.SDuration()),
				slog.Group("returns",
					slog.Any("*model.TagConnection", nil),
					slog.Any("error", err)))
			return nil, err
		}
		tegs = append(tegs, &model.TagEdge{
			Cursor: tag.Id(),
			Node:   n,
		})
	}
	pg := func() model.PageInfo {
		if len(tegs) == 0 {
			return model.PageInfo{}
		}
		if from.ByForward() {
			hn := from.HasNext()
			return model.PageInfo{
				StartCursor: tegs[0].Cursor,
				EndCursor:   tegs[len(tegs)-1].Cursor,
				HasNextPage: &hn,
			}
		}
		if from.ByBackward() {
			hp := from.HasPrev()
			return model.PageInfo{
				StartCursor:     tegs[0].Cursor,
				EndCursor:       tegs[len(tegs)-1].Cursor,
				HasPreviousPage: &hp,
			}
		}
		return model.PageInfo{
			StartCursor: tegs[0].Cursor,
			EndCursor:   tegs[len(tegs)-1].Cursor,
		}
	}()
	cnctn := model.TagConnection{
		Edges:      tegs,
		PageInfo:   &pg,
		TotalCount: len(tegs),
	}
	lgr.InfoContext(ctx, "END",
		slog.String("duration", dw.SDuration()),
		slog.Group("returns",
			slog.Any("*model.TagConnection", cnctn),
			slog.Any("error", nil)))
	return &cnctn, nil
}

func NewConverter() *Converter {
	return &Converter{}
}
