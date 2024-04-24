package usecase

import (
	"context"
	"github.com/miyamo2/altnrslog"
	"github.com/miyamo2/api.miyamo.today/core/log"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"
	"log/slog"

	"github.com/cockroachdb/errors"
	"github.com/miyamo2/api.miyamo.today/article-service/internal/app/usecase/dto"
	iquery "github.com/miyamo2/api.miyamo.today/article-service/internal/app/usecase/query"
	"github.com/miyamo2/api.miyamo.today/article-service/internal/infra/rdb/query"
	"github.com/miyamo2/api.miyamo.today/core/db"
	"github.com/miyamo2/api.miyamo.today/core/util/duration"
)

// GetNext is an implementation of github.com/miyamo2/api.miyamo.today/article-service/internal/if-adapter/controller/pb/usecase.GetNext
type GetNext struct {
	txmn db.TransactionManager
	qs   iquery.ArticleService[query.Tag, *query.Article]
}

func (u *GetNext) Execute(ctx context.Context, in dto.GetNextInDto) (*dto.GetNextOutDto, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("Execute").End()
	dw := duration.Start()
	lgr, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		lgr = log.DefaultLogger()
	}
	lgr.InfoContext(ctx, "BEGIN")
	tx, err := u.txmn.GetAndStart(ctx)
	if err != nil {
		err := errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		lgr.WarnContext(ctx, "END",
			slog.String("duration", dw.SDuration()),
			slog.Group("return",
				slog.Any("dto.GetAllOutDto", nil),
				slog.Any("error", err)))
		return nil, err
	}
	errCh := tx.SubscribeError()

	out := db.NewMultipleStatementResult[*query.Article]()
	limit := in.First()
	stmt := u.qs.GetAll(ctx, out, db.WithNextPaging(limit, in.Cursor()))
	err = tx.ExecuteStatement(ctx, stmt)
	if err != nil {
		err := errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		lgr.WarnContext(ctx, "END",
			slog.String("duration", dw.SDuration()),
			slog.Group("return",
				slog.Any("dto.GetAllOutDto", nil),
				slog.Any("error", err)))
		return nil, err
	}
	qres := out.StrictGet()
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
		a := make([]dto.Article, 0, len(qres))
		for _, am := range qres {
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
			lgr.WarnContext(ctx, "transaction has error. err: %+v", err)
		}
	}
	lgr.InfoContext(ctx, "END",
		slog.String("duration", dw.SDuration()),
		slog.Group("return",
			slog.Any("dto.GetAllOutDto", result),
			slog.Any("error", nil)))
	return &result, nil
}

// NewGetNext is constructor of GetNextPage.
func NewGetNext(txmn db.TransactionManager, qs iquery.ArticleService[query.Tag, *query.Article]) *GetNext {
	return &GetNext{txmn: txmn, qs: qs}
}
