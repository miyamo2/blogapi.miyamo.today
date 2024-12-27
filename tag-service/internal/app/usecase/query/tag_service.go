//go:generate mockgen -source=tag_service.go -destination=../../../mock/app/usecase/query/mock_tag_service.go -package=mock_query
package query

import (
	"context"
	"github.com/miyamo2/blogapi.miyamo.today/tag-service/internal/infra/rdb/query/model"

	"github.com/miyamo2/blogapi.miyamo.today/core/db"
)

// TagService is a query service interface.
type TagService interface {
	// GetById returns a single tag with articles.
	GetById(ctx context.Context, id string, out *db.SingleStatementResult[model.Tag]) db.Statement
	// GetAll returns multiple tag with articles.
	GetAll(ctx context.Context, out *db.MultipleStatementResult[model.Tag], paginationOption ...db.PaginationOption) db.Statement
}
