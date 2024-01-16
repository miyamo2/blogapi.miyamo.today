package usecase

import (
	"context"
	"fmt"
	"github.com/miyamo2/blogapi-tag-service/internal/infra/rdb/query/model"
	"log/slog"

	"github.com/cockroachdb/errors"
	"github.com/miyamo2/blogapi-core/db"
	"github.com/miyamo2/blogapi-core/util/duration"
	"github.com/miyamo2/blogapi-tag-service/internal/app/usecase/dto"
	iquery "github.com/miyamo2/blogapi-tag-service/internal/app/usecase/query"
)

// GetAll is an implementation of usecase.GetAll
type GetAll struct {
	txmn db.TransactionManager
	qs   iquery.TagService[model.Article, *model.Tag]
}

func (u *GetAll) Execute(ctx context.Context) (*dto.GetAllOutDto, error) {
	dw := duration.Start()
	slog.InfoContext(ctx, "BEGIN")
	tx, err := u.txmn.GetAndStart(ctx)
	if err != nil {
		err := errors.WithStack(err)
		slog.WarnContext(ctx, "END",
			slog.String("duration", dw.SDuration()),
			slog.Group("return",
				slog.Any("*dto.GetAllOutDto", nil),
				slog.String("error", fmt.Sprintf("%+v", err))))
		return nil, err
	}
	errCh := tx.SubscribeError()

	out := db.NewMultipleStatementResult[*model.Tag]()
	stmt := u.qs.GetAll(ctx, out)
	err = tx.ExecuteStatement(ctx, stmt)
	if err != nil {
		err := errors.WithStack(err)
		slog.WarnContext(ctx, "END",
			slog.String("duration", dw.SDuration()),
			slog.Group("return",
				slog.Any("*dto.GetAllOutDto", nil),
				slog.String("error", fmt.Sprintf("%+v", err))))
		return nil, err
	}
	qres := out.StrictGet()
	result := dto.NewGetAllOutDto()

	for _, tag := range qres {
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
			slog.WarnContext(ctx, "transaction has error. err: %+v", err)
		}
	}
	slog.InfoContext(ctx, "END",
		slog.String("duration", dw.SDuration()),
		slog.Group("return",
			slog.String("*dto.GetAllOutDto", fmt.Sprintf("%v", result)),
			slog.Any("error", nil)))
	return &result, nil
}

// NewGetAll is constructor of GetAll
func NewGetAll(txmn db.TransactionManager, qs iquery.TagService[model.Article, *model.Tag]) *GetAll {
	return &GetAll{txmn: txmn, qs: qs}
}
