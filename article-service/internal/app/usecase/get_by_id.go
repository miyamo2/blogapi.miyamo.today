package usecase

import (
	"context"
	"fmt"
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

// GetById is an implementation of blogapi.miyamo.today/article-service/internal/if-adapter/controller/pb/usecase.GetById
type GetById struct {
	transactionManager db.TransactionManager
	queryService       iquery.ArticleService
}

func (u *GetById) Execute(ctx context.Context, in dto.GetByIdInDto) (*dto.GetByIdOutDto, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("Execute").End()
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters",
			slog.String("in", fmt.Sprintf("%v", in))))
	tx, err := u.transactionManager.GetAndStart(ctx)
	if err != nil {
		err := errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger.WarnContext(ctx, "END",
			slog.Group("return",
				slog.Any("dto.GetByIdOutDto", nil),
				slog.Any("error", err)))
		return nil, err
	}
	errCh := tx.SubscribeError()

	out := db.NewSingleStatementResult[query.Article]()
	stmt := u.queryService.GetById(ctx, in.Id(), out)
	err = tx.ExecuteStatement(ctx, stmt)
	if err != nil {
		err := errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger.WarnContext(ctx, "END",
			slog.Group("return",
				slog.Any("dto.GetByIdOutDto", nil),
				slog.Any("error", err)))
		return nil, err
	}
	queryResult := out.StrictGet()
	result := dto.NewGetByIdOutDto(
		queryResult.ID(),
		queryResult.Title(),
		queryResult.Body(),
		queryResult.Thumbnail(),
		queryResult.CreatedAt(),
		queryResult.UpdatedAt(),
		func() []dto.Tag {
			tagDto := make([]dto.Tag, 0, len(queryResult.Tags()))
			for _, tag := range queryResult.Tags() {
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
			logger.WarnContext(ctx, "transaction has error.", slog.String("err", err.Error()))
		}
	}
	logger.InfoContext(ctx, "END",
		slog.Group("return",
			slog.Any("dto.GetByIdOutDto", result),
			slog.Any("error", nil)))
	return &result, nil
}

// NewGetById is constructor of GetById
func NewGetById(transactionManager db.TransactionManager, queryService iquery.ArticleService) *GetById {
	return &GetById{transactionManager: transactionManager, queryService: queryService}
}
