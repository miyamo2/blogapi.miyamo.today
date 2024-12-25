package dto

import (
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
)

// GetByIdInDto is an Input DTO for GetById use-case
type GetByIdInDto struct {
	id string
}

// IsInDto is a marker method for InDTO.
func (i GetByIdInDto) IsInDto() {}

// Id returns the ID of the article to be got
func (i GetByIdInDto) Id() string { return i.id }

// NewGetByIdInDto is constructor of GetByIdInDta.
func NewGetByIdInDto(id string) GetByIdInDto {
	return GetByIdInDto{id: id}
}

// Tag is a DTO for tag
type Tag struct {
	id       string
	name     string
	articles []Article
}

// Articles returns the articles of the tag
func (t Tag) Articles() []Article {
	return t.articles
}

// IsOutDto is a marker method for OutDTO.
func (t Tag) IsOutDto() {}

// Id returns the id of the tag
func (t Tag) Id() string { return t.id }

// Name returns the name of the tag
func (t Tag) Name() string { return t.name }

// NewTag is constructor of Tag
func NewTag(id string, name string, articles []Article) Tag {
	return Tag{id: id, name: name, articles: articles}
}

// Article is an Output DTO for GetById use-case.
type Article struct {
	id           string
	title        string
	thumbnailUrl string
	createdAt    synchro.Time[tz.UTC]
	updatedAt    synchro.Time[tz.UTC]
}

// IsOutDto is a marker method for OutDTO.
func (a Article) IsOutDto() {}

// Id returns the id of the article
func (a Article) Id() string { return a.id }

// Title returns the title of the article
func (a Article) Title() string { return a.title }

// ThumbnailUrl returns the thumbnail url of the article
func (a Article) ThumbnailUrl() string { return a.thumbnailUrl }

// CreatedAt returns date the article was created
func (a Article) CreatedAt() synchro.Time[tz.UTC] { return a.createdAt }

// UpdatedAt returns the date the article was last updated.
func (a Article) UpdatedAt() synchro.Time[tz.UTC] { return a.updatedAt }

// NewArticle is constructor of Article
func NewArticle(
	id string,
	title string,
	thumbnail string,
	createdAt synchro.Time[tz.UTC],
	updatedAt synchro.Time[tz.UTC],
) Article {
	return Article{
		id:           id,
		title:        title,
		thumbnailUrl: thumbnail,
		createdAt:    createdAt,
		updatedAt:    updatedAt,
	}
}

// GetByIdOutDto is an Output DTO for GetById use-case.
type GetByIdOutDto struct {
	Tag
}

// NewGetByIdOutDto is constructor of GetByIdOutDto
func NewGetByIdOutDto(id string, name string, articles []Article) GetByIdOutDto {
	return GetByIdOutDto{
		Tag{
			id:       id,
			name:     name,
			articles: articles,
		},
	}
}

// GetAllOutDto is an Output DTO for GetAll use-case.
type GetAllOutDto struct {
	tags []Tag
}

// NewGetAllOutDto is constructor of GetAllOutDto
func NewGetAllOutDto() GetAllOutDto {
	return GetAllOutDto{
		make([]Tag, 0),
	}
}

// WithTagDto returns a new GetAllOutDto with the given tag.
func (g GetAllOutDto) WithTagDto(tag Tag) GetAllOutDto {
	return GetAllOutDto{
		append(g.tags, tag),
	}
}

// Tags returns the tags of the GetAllOutDto
func (g GetAllOutDto) Tags() []Tag {
	return g.tags
}

// IsOutDto is a marker method for OutDTO.
func (g GetAllOutDto) IsOutDto() {}

// GetNextInDto is an Input DTO for GetNext use-case.
type GetNextInDto struct {
	first  int
	cursor *string
}

// IsInDto is a marker method for InDTO.
func (i GetNextInDto) IsInDto() {}

// First returns the number of tags to get
func (i GetNextInDto) First() int { return i.first }

// Cursor returns the cursor that indicates the position to get tags from here.
func (i GetNextInDto) Cursor() *string { return i.cursor }

// NewGetNextInDto is constructor of GetNextInDto.
func NewGetNextInDto(first int, cursor *string) GetNextInDto {
	return GetNextInDto{first: first, cursor: cursor}
}

// GetNextOutDto is an Output DTO for GetNext use-case.
type GetNextOutDto struct {
	tags    []Tag
	hasNext bool
}

// NewGetNextOutDto is constructor of GetNextOutDto.
func NewGetNextOutDto(hasNext bool) GetNextOutDto {
	return GetNextOutDto{tags: make([]Tag, 0), hasNext: hasNext}
}

// WithTagDto returns a new GetNextOutDto with the given tag.
func (o GetNextOutDto) WithTagDto(tag Tag) GetNextOutDto {
	return GetNextOutDto{
		append(o.tags, tag),
		o.hasNext,
	}
}

// IsOutDto is a marker method for OutDTO.
func (o GetNextOutDto) IsOutDto() {}

// Tags returns the tags.
func (o GetNextOutDto) Tags() []Tag { return o.tags }

// HasNext returns whether there is still next items.
func (o GetNextOutDto) HasNext() bool { return o.hasNext }

// GetPrevInDto is an Input DTO for GetPrev use-case.
type GetPrevInDto struct {
	last   int
	cursor *string
}

// IsInDto is a marker method for InDTO.
func (i GetPrevInDto) IsInDto() {}

// Last returns the number of tags to get
func (i GetPrevInDto) Last() int { return i.last }

// Cursor returns the cursor that indicates the position to get tags from here.
func (i GetPrevInDto) Cursor() *string { return i.cursor }

// NewGetPrevInDto is constructor of GetPrevInDto.
func NewGetPrevInDto(last int, cursor *string) GetPrevInDto {
	return GetPrevInDto{last: last, cursor: cursor}
}

// GetPrevOutDto is an Output DTO for GetPrev use-case.
type GetPrevOutDto struct {
	tags        []Tag
	hasPrevious bool
}

// NewGetPrevOutDto is constructor of GetPrevOutDto.
func NewGetPrevOutDto(hasPrevious bool) GetPrevOutDto {
	return GetPrevOutDto{tags: make([]Tag, 0), hasPrevious: hasPrevious}
}

// WithTagDto returns a new GetPrevOutDto with the given tag.
func (o GetPrevOutDto) WithTagDto(tag Tag) GetPrevOutDto {
	return GetPrevOutDto{
		append(o.tags, tag),
		o.hasPrevious,
	}
}

// IsOutDto is a marker method for OutDTO.
func (o GetPrevOutDto) IsOutDto() {}

// Tags returns the tags.
func (o GetPrevOutDto) Tags() []Tag { return o.tags }

// HasPrevious returns whether there is still previous items.
func (o GetPrevOutDto) HasPrevious() bool { return o.hasPrevious }
