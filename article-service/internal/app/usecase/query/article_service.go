//go:generate mockgen -source=article_service.go -destination=../../../mock/app/usecase/query/mock_article_service.go -package=mock_query
package query

import (
	"context"
	"github.com/miyamo2/api.miyamo.today/article-service/internal/app/usecase/query/model"
	"github.com/miyamo2/api.miyamo.today/core/db"
)

// ArticleService is a query service interface.
type ArticleService[T model.Tag, A model.Article[T]] interface {
	// GetById returns a single article with tags.
	GetById(ctx context.Context, id string, out *db.SingleStatementResult[A]) db.Statement
	// GetAll returns all articles with tags.
	//
	// If PaginationOption is specified, paging is performed.
	// multiple PaginationOption is specified, the last one is used.
	GetAll(ctx context.Context, out *db.MultipleStatementResult[A], paginationOption ...db.PaginationOption) db.Statement
}
