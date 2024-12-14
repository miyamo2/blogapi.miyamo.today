package usecase

import (
	"context"
	"github.com/cockroachdb/errors"
	"github.com/miyamo2/altnrslog"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/app/usecase/command"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/app/usecase/dto"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/domain/model"
	"github.com/miyamo2/blogapi.miyamo.today/core/db"
	"github.com/miyamo2/blogapi.miyamo.today/core/log"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"
	"log/slog"
)

// CreateArticle is a use-case for creating an article.
type CreateArticle struct {
	bloggingEventCommand command.BloggingEventService
}

// Execute executes the CreateArticle use-case.
func (u *CreateArticle) Execute(ctx context.Context, in *dto.CreateArticleInDto) (*dto.CreateArticleOutDto, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("Execute").End()

	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN")

	command := model.NewCreateArticleEvent(in.Title(), in.Body(), in.ThumbnailUrl(), in.TagNames())
	commandOut := db.NewSingleStatementResult[*model.BloggingEventKey]()
	err = u.bloggingEventCommand.CreateArticle(ctx, command, commandOut).Execute(ctx)
	if err != nil {
		err := errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger.WarnContext(ctx, "END",
			slog.Group("return",
				slog.Any("dto.GetAllOutDto", nil),
				slog.Any("error", err)))
		return nil, err
	}

	key := commandOut.StrictGet()
	result := dto.NewCreateArticleOutDto(key.EventID(), key.ArticleID())
	return &result, nil
}

// NewCreateArticle is a constructor for CreateArticle use-case.
func NewCreateArticle(bloggingEventCommand command.BloggingEventService) *CreateArticle {
	return &CreateArticle{bloggingEventCommand: bloggingEventCommand}
}
