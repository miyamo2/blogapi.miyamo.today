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

// GetById is an implementation of usecase.GetById
type GetById struct {
	transactionManager db.TransactionManager
	queryService       iquery.TagService
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
		logger.WarnContext(ctx, "END",
			slog.Group("return",
				slog.Any("*dto.GetByIdOutDto", nil),
				slog.String("error", fmt.Sprintf("%+v", err))))
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		return nil, err
	}
	errCh := tx.SubscribeError()

	out := db.NewSingleStatementResult[model.Tag]()
	stmt := u.queryService.GetById(ctx, in.Id(), out)
	err = tx.ExecuteStatement(ctx, stmt)
	if err != nil {
		err := errors.WithStack(err)
		logger.WarnContext(ctx, "END",
			slog.Group("return",
				slog.Any("*dto.GetByIdOutDto", nil),
				slog.String("error", fmt.Sprintf("%+v", err))))
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		return nil, err
	}
	queryResult := out.StrictGet()
	result := dto.NewGetByIdOutDto(
		queryResult.ID(),
		queryResult.Name(),
		func() []dto.Article {
			articleDto := make([]dto.Article, 0, len(queryResult.Articles()))
			for _, a := range queryResult.Articles() {
				articleDto = append(articleDto, dto.NewArticle(
					a.ID(),
					a.Title(),
					a.Thumbnail(),
					a.CreatedAt(),
					a.UpdatedAt()))
			}
			return articleDto
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
			// slog.String("*dto.GetByIdOutDto", fmt.Sprintf("%v", result)),
			slog.Any("error", nil)))
	return &result, nil
}

// NewGetById is constructor of GetById
func NewGetById(transactionManager db.TransactionManager, queryService iquery.TagService) *GetById {
	return &GetById{transactionManager: transactionManager, queryService: queryService}
}
