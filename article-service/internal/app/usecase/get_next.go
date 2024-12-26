package usecase

import (
	"context"
	"log/slog"

	"github.com/miyamo2/altnrslog"
	"github.com/miyamo2/blogapi.miyamo.today/core/log"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"

	"github.com/cockroachdb/errors"
	"github.com/miyamo2/blogapi.miyamo.today/article-service/internal/app/usecase/dto"
	iquery "github.com/miyamo2/blogapi.miyamo.today/article-service/internal/app/usecase/query"
	"github.com/miyamo2/blogapi.miyamo.today/article-service/internal/infra/rdb/query"
	"github.com/miyamo2/blogapi.miyamo.today/core/db"
)

// GetNext is an implementation of github.com/miyamo2/blogapi.miyamo.today/article-service/internal/if-adapter/controller/pb/usecase.GetNext
type GetNext struct {
	transactionManager db.TransactionManager
	queryService       iquery.ArticleService
}

func (u *GetNext) Execute(ctx context.Context, in dto.GetNextInDto) (*dto.GetNextOutDto, error) {
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
				slog.Any("dto.GetAllOutDto", nil),
				slog.Any("error", err)))
		return nil, err
	}
	errCh := tx.SubscribeError()

	out := db.NewMultipleStatementResult[query.Article]()
	limit := in.First()
	stmt := u.queryService.GetAll(ctx, out, db.WithNextPaging(limit, in.Cursor()))
	err = tx.ExecuteStatement(ctx, stmt)
	if err != nil {
		err := errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger.WarnContext(ctx, "END",
			slog.Group("return",
				slog.Any("dto.GetAllOutDto", nil),
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
	result := dto.NewGetNextOutDto(convArticleModels())

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
			slog.Any("dto.GetAllOutDto", result),
			slog.Any("error", nil)))
	return &result, nil
}

// NewGetNext is constructor of GetNextPage.
func NewGetNext(transactionManager db.TransactionManager, queryService iquery.ArticleService) *GetNext {
	return &GetNext{transactionManager: transactionManager, queryService: queryService}
}
