package usecase

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/miyamo2/altnrslog"
	"github.com/miyamo2/blogapi.miyamo.today/core/log"

	"github.com/miyamo2/blogapi.miyamo.today/tag-service/internal/infra/rdb/query/model"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"

	"github.com/cockroachdb/errors"
	"github.com/miyamo2/blogapi.miyamo.today/core/db"
	"github.com/miyamo2/blogapi.miyamo.today/tag-service/internal/app/usecase/dto"
	iquery "github.com/miyamo2/blogapi.miyamo.today/tag-service/internal/app/usecase/query"
)

// GetAll is an implementation of usecase.GetAll
type GetAll struct {
	transactionManager db.TransactionManager
	queryService       iquery.TagService
}

func (u *GetAll) Execute(ctx context.Context) (*dto.GetAllOutDto, error) {
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
				slog.Any("*dto.GetAllOutDto", nil),
				slog.String("error", fmt.Sprintf("%+v", err))))
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		return nil, err
	}
	errCh := tx.SubscribeError()

	out := db.NewMultipleStatementResult[model.Tag]()
	stmt := u.queryService.GetAll(ctx, out)
	err = tx.ExecuteStatement(ctx, stmt)
	if err != nil {
		err := errors.WithStack(err)
		logger.WarnContext(ctx, "END",
			slog.Group("return",
				slog.Any("*dto.GetAllOutDto", nil),
				slog.String("error", fmt.Sprintf("%+v", err))))
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		return nil, err
	}
	queryResult := out.StrictGet()
	result := dto.NewGetAllOutDto()

	for _, tag := range queryResult {
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
			// slog.String("*dto.GetAllOutDto", fmt.Sprintf("%v", result)),
			slog.Any("error", nil)))
	return &result, nil
}

// NewGetAll is constructor of GetAll
func NewGetAll(transactionManager db.TransactionManager, queryService iquery.TagService) *GetAll {
	return &GetAll{transactionManager: transactionManager, queryService: queryService}
}
