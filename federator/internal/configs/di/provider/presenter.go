package provider

import (
	"github.com/google/wire"
	abstract "github.com/miyamo2/blogapi.miyamo.today/federator/internal/if-adapter/controller/graphql/resolver/presenter/converter"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/if-adapter/presenters/graphql/converter"
)

// compatibility check
var (
	_ abstract.ArticleConverter  = (*converter.Converter)(nil)
	_ abstract.ArticlesConverter = (*converter.Converter)(nil)
	_ abstract.TagConverter      = (*converter.Converter)(nil)
	_ abstract.TagsConverter     = (*converter.Converter)(nil)
)

var PresenterSet = wire.NewSet(
	converter.NewConverter,
	wire.Bind(new(abstract.ArticleConverter), new(*converter.Converter)),
	wire.Bind(new(abstract.ArticlesConverter), new(*converter.Converter)),
	wire.Bind(new(abstract.TagConverter), new(*converter.Converter)),
	wire.Bind(new(abstract.TagsConverter), new(*converter.Converter)),
)
