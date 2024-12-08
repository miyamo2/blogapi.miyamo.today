package command

import (
	"context"
	"github.com/miyamo2/blogapi.miyamo.today/core/db"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/domain/model"
)

// ArticleService is a command service interface for the Article.
type ArticleService interface {
	ExecuteArticleCommand(ctx context.Context, article model.ArticleCommand) db.Statement
}

type TagService interface {
	ExecuteTagCommand(ctx context.Context, article model.ArticleCommand) db.Statement
}
