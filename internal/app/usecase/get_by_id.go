package usecase

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/cockroachdb/errors"
	"github.com/miyamo2/blogapi-article-service/internal/app/usecase/dto"
	iquery "github.com/miyamo2/blogapi-article-service/internal/app/usecase/query"
	"github.com/miyamo2/blogapi-article-service/internal/infra/rdb/query"
	"github.com/miyamo2/blogapi-core/db"
	"github.com/miyamo2/blogapi-core/util/duration"
)

// GetById is an implementation of github.com/miyamo2/blogapi-article-service/internal/if-adapter/controller/pb/usecase.GetById
type GetById struct {
	txmn db.TransactionManager
	qs   iquery.ArticleService[query.Tag, *query.Article]
}

func (u *GetById) Execute(ctx context.Context, in dto.GetByIdInDto) (*dto.GetByIdOutDto, error) {
	dw := duration.Start()
	slog.InfoContext(ctx, "BEGIN",
		slog.Group("parameters",
			slog.String("in", fmt.Sprintf("%v", in))))
	tx, err := u.txmn.GetAndStart(ctx)
	if err != nil {
		err := errors.WithStack(err)
		slog.WarnContext(ctx, "END",
			slog.String("duration", dw.SDuration()),
			slog.Group("return",
				slog.Any("*dto.GetByIdOutDto", nil),
				slog.String("error", fmt.Sprintf("%+v", err))))
		return nil, err
	}
	errCh := tx.SubscribeError()

	out := db.NewSingleStatementResult[*query.Article]()
	stmt := u.qs.GetById(ctx, in.Id(), out)
	err = tx.ExecuteStatement(ctx, stmt)
	if err != nil {
		err := errors.WithStack(err)
		slog.WarnContext(ctx, "END",
			slog.String("duration", dw.SDuration()),
			slog.Group("return",
				slog.Any("*dto.GetByIdOutDto", nil),
				slog.String("error", fmt.Sprintf("%+v", err))))
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
			slog.WarnContext(ctx, "transaction has error. err: %+v", err)
		}
	}
	slog.InfoContext(ctx, "END",
		slog.String("duration", dw.SDuration()),
		slog.Group("return",
			slog.String("*dto.GetByIdOutDto", fmt.Sprintf("%v", result)),
			slog.Any("error", nil)))
	return &result, nil
}

// NewGetById is constructor of GetById
func NewGetById(txmn db.TransactionManager, qs iquery.ArticleService[query.Tag, *query.Article]) *GetById {
	return &GetById{txmn: txmn, qs: qs}
}
