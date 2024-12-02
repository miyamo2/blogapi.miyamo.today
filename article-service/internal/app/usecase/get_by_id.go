package usecase

import (
	"context"
	"fmt"
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
	"github.com/miyamo2/blogapi.miyamo.today/core/util/duration"
)

// GetById is an implementation of github.com/miyamo2/blogapi.miyamo.today/article-service/internal/if-adapter/controller/pb/usecase.GetById
type GetById struct {
	txmn db.TransactionManager
	qs   iquery.ArticleService
}

func (u *GetById) Execute(ctx context.Context, in dto.GetByIdInDto) (*dto.GetByIdOutDto, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("Execute").End()
	dw := duration.Start()
	lgr, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		lgr = log.DefaultLogger()
	}
	lgr.InfoContext(ctx, "BEGIN",
		slog.Group("parameters",
			slog.String("in", fmt.Sprintf("%v", in))))
	tx, err := u.txmn.GetAndStart(ctx)
	if err != nil {
		err := errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		lgr.WarnContext(ctx, "END",
			slog.String("duration", dw.SDuration()),
			slog.Group("return",
				slog.Any("dto.GetByIdOutDto", nil),
				slog.Any("error", err)))
		return nil, err
	}
	errCh := tx.SubscribeError()

	out := db.NewSingleStatementResult[*query.Article]()
	stmt := u.qs.GetById(ctx, in.Id(), out)
	err = tx.ExecuteStatement(ctx, stmt)
	if err != nil {
		err := errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		lgr.WarnContext(ctx, "END",
			slog.String("duration", dw.SDuration()),
			slog.Group("return",
				slog.Any("dto.GetByIdOutDto", nil),
				slog.Any("error", err)))
		return nil, err
	}
	qres := out.StrictGet()
	result := dto.NewGetByIdOutDto(
		qres.ID(),
		qres.Title(),
		qres.Body(),
		qres.Thumbnail(),
		qres.CreatedAt(),
		qres.UpdatedAt(),
		func() []dto.Tag {
			tagDto := make([]dto.Tag, 0, len(qres.Tags()))
			for _, tag := range qres.Tags() {
				tagDto = append(tagDto, dto.NewTag(tag.ID(), tag.Name()))
			}
			return tagDto
		}())

	// Never return an error
	_ = tx.Commit(ctx)
	for {
		err, alive := <-errCh
		if !alive {
			break
		}
		if err != nil {
			nrtx.NoticeError(nrpkgerrors.Wrap(err))
			lgr.WarnContext(ctx, "transaction has error.", slog.String("err", err.Error()))
		}
	}
	lgr.InfoContext(ctx, "END",
		slog.String("duration", dw.SDuration()),
		slog.Group("return",
			slog.Any("dto.GetByIdOutDto", result),
			slog.Any("error", nil)))
	return &result, nil
}

// NewGetById is constructor of GetById
func NewGetById(txmn db.TransactionManager, qs iquery.ArticleService) *GetById {
	return &GetById{txmn: txmn, qs: qs}
}
