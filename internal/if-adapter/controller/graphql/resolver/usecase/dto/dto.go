package dto

type InDto interface {
	// IsInDto is a marker for in dto.
	IsInDto()
}

type OutDto interface {
	// IsOutDto is a marker for out dto.
	IsOutDto()
}

type ArticleInDto interface {
	InDto
	// Id returns id.
	Id() string
}

type Tag interface {
	OutDto
	Id() string
	Name() string
}

type ArticleTag[T Tag] interface {
	OutDto
	Id() string
	Title() string
	ThumbnailUrl() string
	CreatedAt() string
	UpdatedAt() string
	Tags() []T
}

type ArticleOutDto[T Tag, AT ArticleTag[T]] interface {
	OutDto
	Article() AT
}

type ArticlesInDto interface {
	InDto
	First() int
	After() string
	Last() int
	Before() string
}

type ArticlesOutDto[T Tag, AT ArticleTag[T]] interface {
	OutDto
	Articles() []AT
	HasNext() bool
	HasPrev() bool
}

type TagInDto interface {
	InDto
	Id() string
}

type Article interface {
	OutDto
	Id() string
	Title() string
	ThumbnailUrl() string
	CreatedAt() string
	UpdatedAt() string
}

type TagArticle[A Article] interface {
	OutDto
	Id() string
	Name() string
	Articles() []A
}

type TagOutDto[A Article, T TagArticle[A]] interface {
	IsOutDto()
	Tag() T
}

type TagsInDto interface {
	InDto
	First() int
	After() string
	Last() int
	Before() string
}

type TagsOutDto[A Article, T TagArticle[A]] interface {
	OutDto
	Tags() []T
	HasNext() bool
	HasPrev() bool
}
