package model

type QueryServiceModel interface {
	IsQueryServiceModel()
}

// Article is a query service model for article.
type Article[T Tag] interface {
	QueryServiceModel
	ID() string
	Title() string
	Body() string
	Thumbnail() string
	CreatedAt() string
	UpdatedAt() string
	Tags() []T
}

// Tag is a query service model for tag.
type Tag interface {
	QueryServiceModel
	ID() string
	Name() string
}
