package provider

import (
	"github.com/google/wire"
	"github.com/miyamo2/blogapi.miyamo.today/article-service/internal/app/usecase/query"
	impl "github.com/miyamo2/blogapi.miyamo.today/article-service/internal/infra/rdb/query"
)

// compatibility check
var _ query.ArticleService = (*impl.ArticleService)(nil)

var QueryServiceSet = wire.NewSet(
	impl.NewArticleService,
	wire.Bind(new(query.ArticleService), new(*impl.ArticleService)),
)
