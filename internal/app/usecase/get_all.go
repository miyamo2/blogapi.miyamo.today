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

// GetAll is an implementation of github.com/miyamo2/blogapi-article-service/internal/if-adapter/controller/pb/usecase.GetAll
type GetAll struct {
	txmn db.TransactionManager
	qs   iquery.ArticleService[query.Tag, *query.Article]
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

	out := db.NewMultipleStatementResult[*query.Article]()
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
	// convert model.Tag to dto.Tag
	convTagModels := func(tms []query.Tag) []dto.Tag {
		t := make([]dto.Tag, 0, len(tms))
		for _, tm := range tms {
			t = append(t, dto.NewTag(tm.ID(), tm.Name()))
		}
		return t
	}
	// convert model.Article to dto.Article
	convArticleModels := func() []dto.Article {
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
		return a
	}
	result := dto.NewGetAllOutDto(convArticleModels())

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

// NewGetAll is constructor of GetAll.
func NewGetAll(txmn db.TransactionManager, qs iquery.ArticleService[query.Tag, *query.Article]) *GetAll {
	return &GetAll{txmn: txmn, qs: qs}
}
