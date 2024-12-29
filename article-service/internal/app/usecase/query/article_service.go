//go:generate mockgen -source=article_service.go -destination=../../../mock/app/usecase/query/mock_article_service.go -package=mock_query
package query

import (
	"blogapi.miyamo.today/article-service/internal/infra/rdb/query"
	"context"

	"blogapi.miyamo.today/core/db"
)

// ArticleService is a query service interface.
type ArticleService interface {
	// GetById returns a single article with tags.
	GetById(ctx context.Context, id string, out *db.SingleStatementResult[query.Article]) db.Statement
	// GetAll returns all articles with tags.
	//
	// If PaginationOption is specified, paging is performed.
	// multiple PaginationOption is specified, the last one is used.
	GetAll(ctx context.Context, out *db.MultipleStatementResult[query.Article], paginationOption ...db.PaginationOption) db.Statement
}
