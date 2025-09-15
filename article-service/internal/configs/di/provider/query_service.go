package provider

import (
	"blogapi.miyamo.today/article-service/internal/app/usecase/query"
	impl "blogapi.miyamo.today/article-service/internal/infra/rdb/query"
	"github.com/google/wire"
)

// compatibility check
var _ query.ArticleService = (*impl.ArticleService)(nil)

var QueryServiceSet = wire.NewSet(
	impl.NewArticleService,
	wire.Bind(new(query.ArticleService), new(*impl.ArticleService)),
)
