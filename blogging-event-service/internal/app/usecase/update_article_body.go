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

// UpdateArticleBody is a use-case for updating an article body.
type UpdateArticleBody struct {
	bloggingEventCommand command.BloggingEventService
}

// Execute executes the UpdateArticleBody use-case.
func (u *UpdateArticleBody) Execute(ctx context.Context, in *dto.UpdateArticleBodyInDto) (*dto.UpdateArticleBodyOutDto, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("Execute").End()

	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN")
	defer func() {
		if err != nil {
			logger.WarnContext(ctx, "END",
				slog.Group("return",
					slog.Any("dto.UpdateArticleBody", nil),
					slog.Any("error", err)))
		}
		logger.InfoContext(ctx, "END")
	}()

	command := model.NewUpdateArticleBodyEvent(in.ID(), in.Body())
	commandOut := db.NewSingleStatementResult[*model.BloggingEventKey]()
	err = u.bloggingEventCommand.UpdateArticleBody(ctx, command, commandOut).Execute(ctx)
	if err != nil {
		err = errors.WithStack(err)
		return nil, err
	}

	key := commandOut.StrictGet()
	result := dto.NewUpdateArticleBodyOutDto(key.EventID(), key.ArticleID())
	return &result, nil
}

// NewUpdateArticleBody is a constructor for UpdateArticleBody use-case.
func NewUpdateArticleBody(bloggingEventCommand command.BloggingEventService) *UpdateArticleBody {
	return &UpdateArticleBody{bloggingEventCommand: bloggingEventCommand}
}
