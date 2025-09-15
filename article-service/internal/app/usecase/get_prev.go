package usecase

import (
	"context"
	"log/slog"

	"blogapi.miyamo.today/core/log"
	"github.com/miyamo2/altnrslog"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"

	"blogapi.miyamo.today/article-service/internal/app/usecase/dto"
	iquery "blogapi.miyamo.today/article-service/internal/app/usecase/query"
	"blogapi.miyamo.today/article-service/internal/infra/rdb/query"
	"blogapi.miyamo.today/core/db"
	"github.com/cockroachdb/errors"
)

// GetPrev is an implementation of blogapi.miyamo.today/article-service/internal/if-adapter/controller/pb/usecase.GetPrev
type GetPrev struct {
	transactionManager db.TransactionManager
	queryService       iquery.ArticleService
}

func (u *GetPrev) Execute(ctx context.Context, in dto.GetPrevInDto) (*dto.GetPrevOutDto, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("Execute").End()
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN")
	tx, err := u.transactionManager.GetAndStart(ctx)
	if err != nil {
		err := errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger.WarnContext(ctx, "END",
			slog.Group("return",
				slog.Any("dto.GetPrevOutDto", nil),
				slog.Any("error", err)))
		return nil, err
	}
	errCh := tx.SubscribeError()

	out := db.NewMultipleStatementResult[query.Article]()
	limit := in.Last()
	stmt := u.queryService.GetAll(ctx, out, db.WithPreviousPaging(limit, in.Cursor()))
	err = tx.ExecuteStatement(ctx, stmt)
	if err != nil {
		err := errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger.WarnContext(ctx, "END",
			slog.Group("return",
				slog.Any("dto.GetPrevOutDto", nil),
				slog.Any("error", err)))
		return nil, err
	}
	queryResult := out.StrictGet()
	// convert model.Tag to dto.Tag
	convTagModels := func(tms []query.Tag) []dto.Tag {
		t := make([]dto.Tag, 0, len(tms))
		for _, tm := range tms {
			t = append(t, dto.NewTag(tm.ID(), tm.Name()))
		}
		return t
	}
	// convert model.Article to dto.Article
	convArticleModels := func() ([]dto.Article, bool) {
		a := make([]dto.Article, 0, len(queryResult))
		for _, am := range queryResult {
			// This would be a dead code.
			// if am == nil ...
			a = append(a, dto.NewArticle(
				am.ID(),
				am.Title(),
				am.Body(),
				am.Thumbnail(),
				am.CreatedAt(),
				am.UpdatedAt(),
				convTagModels(am.Tags())))
		}
		if len(a) > limit {
			return a[:limit], true
		}
		return a, false
	}
	result := dto.NewGetPrevOutDto(convArticleModels())

	// Never return an error
	_ = tx.Commit(ctx)
	for {
		err, alive := <-errCh
		if !alive {
			break
		}
		if err != nil {
			nrtx.NoticeError(nrpkgerrors.Wrap(err))
			logger.WarnContext(ctx, "transaction has error.", slog.String("err", err.Error()))
		}
	}
	logger.InfoContext(ctx, "END",
		slog.Group("return",
			slog.Any("dto.GetPrevOutDto", result),
			slog.Any("error", nil)))
	return &result, nil
}

// NewGetPrev is constructor of GetPrevPage.
func NewGetPrev(transactionManager db.TransactionManager, queryService iquery.ArticleService) *GetPrev {
	return &GetPrev{transactionManager: transactionManager, queryService: queryService}
}
