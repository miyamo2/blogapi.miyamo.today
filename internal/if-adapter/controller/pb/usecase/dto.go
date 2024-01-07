package usecase

type InDto interface {
	IsInDto()
}

type OutDto interface {
	IsOutDto()
}

// GetByIdInDto is an Input DTO for GetById use-case
type GetByIdInDto interface {
	InDto
	Id() string
}

// Article is an Output DTO for GetById use-case.
type Article[T Tag] interface {
	OutDto
	// Id returns the id of the article
	Id() string
	// Title returns the title of the article
	Title() string
	// Body returns the body of the article
	Body() string
	// ThumbnailUrl returns the thumbnail url of the article
	ThumbnailUrl() string
	// CreatedAt returns date the article was created
	CreatedAt() string
	// UpdatedAt returns the date the article was last updated.
	UpdatedAt() string
	// Tags return the tags attached to the article
	Tags() []T
}

// Tag is a DTO for tag
type Tag interface {
	OutDto
	// Id returns the id of the tag
	Id() string
	// Name returns the name of the tag
	Name() string
}

// GetAllOutDto is an Output DTO for GetAll use-case.
type GetAllOutDto[T Tag, A Article[T]] interface {
	OutDto
	Articles() []A
}

// GetNextInDto is an Input DTO for GetNext use-case
type GetNextInDto interface {
	InDto
	Limit() int
	Cursor() *string
}

// GetNextOutDto is an Output DTO for GetNext use-case.
type GetNextOutDto[T Tag, A Article[T]] interface {
	OutDto
	Articles() []A
	StillExists() bool
}
