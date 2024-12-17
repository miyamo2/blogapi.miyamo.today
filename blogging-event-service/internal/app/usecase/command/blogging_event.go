//go:generate mockgen -source=$GOFILE -destination=../../../mock/app/usecase/command/$GOFILE -package=$GOPACKAGE
package command

import (
	"context"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/domain/model"
	"github.com/miyamo2/blogapi.miyamo.today/core/db"
)

type BloggingEventService interface {
	CreateArticle(ctx context.Context, in model.CreateArticleEvent, out *db.SingleStatementResult[*model.BloggingEventKey]) db.Statement
	UpdateArticleTitle(ctx context.Context, in model.UpdateArticleTitleEvent, out *db.SingleStatementResult[*model.BloggingEventKey]) db.Statement
	UpdateArticleBody(ctx context.Context, command model.UpdateArticleBodyEvent, out *db.SingleStatementResult[*model.BloggingEventKey]) db.Statement
}
