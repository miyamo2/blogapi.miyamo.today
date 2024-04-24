//go:generate mockgen -source=tag_service.go -destination=../../../mock/app/usecase/query/mock_tag_service.go -package=mock_query
package query

import (
	"context"
	"github.com/miyamo2/api.miyamo.today/core/db"
	"github.com/miyamo2/api.miyamo.today/tag-service/internal/app/usecase/query/model"
)

// TagService is a query service interface.
type TagService[A model.Article, T model.Tag[A]] interface {
	// GetById returns a single tag with articles.
	GetById(ctx context.Context, id string, out *db.SingleStatementResult[T]) db.Statement
	// GetAll returns multiple tag with articles.
	GetAll(ctx context.Context, out *db.MultipleStatementResult[T], paginationOption ...db.PaginationOption) db.Statement
}
