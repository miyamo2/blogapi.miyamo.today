package usecase

import (
	"context"
	"fmt"
	"log/slog"

	"blogapi.miyamo.today/core/log"
	"github.com/miyamo2/altnrslog"

	"blogapi.miyamo.today/tag-service/internal/infra/rdb/query/model"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"

	"blogapi.miyamo.today/core/db"
	"blogapi.miyamo.today/tag-service/internal/app/usecase/dto"
	iquery "blogapi.miyamo.today/tag-service/internal/app/usecase/query"
	"github.com/cockroachdb/errors"
)

// GetPrev is an implementation of usecase.GetPrev
type GetPrev struct {
	transactionManager db.TransactionManager
	queryService       iquery.TagService
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
		logger.WarnContext(ctx, "END",
			slog.Group("return",
				slog.Any("*dto.GetPrevOutDto", nil),
				slog.String("error", fmt.Sprintf("%+v", err))))
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		return nil, err
	}
	errCh := tx.SubscribeError()

	out := db.NewMultipleStatementResult[model.Tag]()
	last := in.Last()
	stmt := u.queryService.GetAll(ctx, out, db.WithPreviousPaging(last, in.Cursor()))
	err = tx.ExecuteStatement(ctx, stmt)
	if err != nil {
		err := errors.WithStack(err)
		logger.WarnContext(ctx, "END",
			slog.Group("return",
				slog.Any("*dto.GetPrevOutDto", nil),
				slog.String("error", fmt.Sprintf("%+v", err))))
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		return nil, err
	}
	queryResult := out.StrictGet()

	result := dto.NewGetPrevOutDto(len(queryResult) > last)
	for i, tag := range queryResult {
		if i >= last {
			break
		}
		result = result.WithTagDto(tagDtoFromQueryModel(tag))
	}

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
			// slog.String("*dto.GetPrevOutDto", fmt.Sprintf("%v", result)),
			slog.Any("error", nil)))
	return &result, nil
}

// NewGetPrev is constructor of GetPrev
func NewGetPrev(transactionManager db.TransactionManager, queryService iquery.TagService) *GetPrev {
	return &GetPrev{transactionManager: transactionManager, queryService: queryService}
}
