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

// Tag is a DTO for tag
type Tag[A Article] interface {
	OutDto
	// Id returns the id of the tag
	Id() string
	// Name returns the name of the tag
	Name() string
	// Articles returns the articles of the tag
	Articles() []A
}

// Article is an Output DTO for GetById use-case.
type Article interface {
	OutDto
	// Id returns the id of the article
	Id() string
	// Title returns the title of the article
	Title() string
	// ThumbnailUrl returns the thumbnail url of the article
	ThumbnailUrl() string
	// CreatedAt returns date the article was created
	CreatedAt() string
	// UpdatedAt returns the date the article was last updated.
	UpdatedAt() string
}

// GetAllOutDto is an Output DTO for GetAll use-case.
type GetAllOutDto[A Article, T Tag[A]] interface {
	OutDto
	// Tags returns the tags of the GetAllOutDto
	Tags() []T
}

// GetNextInDto is an Input DTO for GetNext use-case
type GetNextInDto interface {
	InDto
	// First returns the number of tags to get
	First() int
	// Cursor returns the cursor that indicates the position to get tags from here.
	Cursor() *string
}

// GetNextOutDto is an Output DTO for GetNext use-case.
type GetNextOutDto[A Article, T Tag[A]] interface {
	OutDto
	// Tags returns the tags of the GetNextOutDto
	Tags() []T
	// HasNext returns whether there are more tags to get
	HasNext() bool
}

// GetPrevInDto is an Input DTO for GetPrev use-case
type GetPrevInDto interface {
	InDto
	// Last returns the number of tags to get
	Last() int
	// Cursor returns the cursor that indicates the position to get tags from here.
	Cursor() *string
}

// GetPrevOutDto is an Output DTO for GetPrev use-case.
type GetPrevOutDto[A Article, T Tag[A]] interface {
	OutDto
	// Tags returns the tags of the GetPrevOutDto
	Tags() []T
	// HasPrevious returns whether there are more tags to get
	HasPrevious() bool
}
