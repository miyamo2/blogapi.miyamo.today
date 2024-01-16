package model

type QueryServiceModel interface {
	IsQueryServiceModel()
}

// Tag is a query service model for tag.
type Tag[A Article] interface {
	QueryServiceModel
	ID() string
	Name() string
	Articles() []A
}

// Article is a query service model for article.
type Article interface {
	QueryServiceModel
	ID() string
	Title() string
	Thumbnail() string
	CreatedAt() string
	UpdatedAt() string
}
