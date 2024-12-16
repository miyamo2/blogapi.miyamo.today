package provider

import (
	"github.com/google/wire"
	abstract "github.com/miyamo2/blogapi.miyamo.today/federator/internal/if-adapter/controller/graphql/resolver/presenter/converters"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/if-adapter/presenters/graphql/converters"
)

// compatibility check
var (
	_ abstract.ArticleConverter       = (*converters.Converter)(nil)
	_ abstract.ArticlesConverter      = (*converters.Converter)(nil)
	_ abstract.TagConverter           = (*converters.Converter)(nil)
	_ abstract.TagsConverter          = (*converters.Converter)(nil)
	_ abstract.CreateArticleConverter = (*converters.Converter)(nil)
)

var PresenterSet = wire.NewSet(
	converters.NewConverter,
	wire.Bind(new(abstract.ArticleConverter), new(*converters.Converter)),
	wire.Bind(new(abstract.ArticlesConverter), new(*converters.Converter)),
	wire.Bind(new(abstract.TagConverter), new(*converters.Converter)),
	wire.Bind(new(abstract.TagsConverter), new(*converters.Converter)),
	wire.Bind(new(abstract.CreateArticleConverter), new(*converters.Converter)),
)
