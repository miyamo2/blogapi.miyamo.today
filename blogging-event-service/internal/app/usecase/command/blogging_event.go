//go:generate mockgen -source=$GOFILE -destination=../../../mock/app/usecase/command/$GOFILE -package=$GOPACKAGE
package command

import (
	"blogapi.miyamo.today/blogging-event-service/internal/domain/model"
	"blogapi.miyamo.today/core/db"
	"context"
)

// BloggingEventService is a command service for blogging events.
type BloggingEventService interface {
	// CreateArticle creates a new article.
	CreateArticle(ctx context.Context, in model.CreateArticleEvent, out *db.SingleStatementResult[*model.BloggingEventKey]) db.Statement
	// UpdateArticleTitle updates the title of the article.
	UpdateArticleTitle(ctx context.Context, in model.UpdateArticleTitleEvent, out *db.SingleStatementResult[*model.BloggingEventKey]) db.Statement
	// UpdateArticleBody updates the body of the article.
	UpdateArticleBody(ctx context.Context, command model.UpdateArticleBodyEvent, out *db.SingleStatementResult[*model.BloggingEventKey]) db.Statement
	// UpdateArticleThumbnail updates the thumbnail of the article.
	UpdateArticleThumbnail(ctx context.Context, command model.UpdateArticleThumbnailEvent, out *db.SingleStatementResult[*model.BloggingEventKey]) db.Statement
	// AttachTags attaches tags to the article.
	AttachTags(ctx context.Context, command model.AttachTagsEvent, out *db.SingleStatementResult[*model.BloggingEventKey]) db.Statement
	// DetachTags detaches tags from the article.
	DetachTags(ctx context.Context, command model.DetachTagsEvent, out *db.SingleStatementResult[*model.BloggingEventKey]) db.Statement
}
