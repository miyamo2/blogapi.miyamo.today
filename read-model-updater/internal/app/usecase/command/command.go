package command

import (
	"context"

	"blogapi.miyamo.today/core/db"
	"blogapi.miyamo.today/read-model-updater/internal/domain/model"
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
)

// ArticleService is a command service interface for the Article.
type ArticleService interface {
	ExecuteArticleCommand(ctx context.Context, article model.ArticleCommand, eventAt synchro.Time[tz.UTC]) db.Statement
}

type TagService interface {
	ExecuteTagCommand(ctx context.Context, article model.ArticleCommand, eventAt synchro.Time[tz.UTC]) db.Statement
}
