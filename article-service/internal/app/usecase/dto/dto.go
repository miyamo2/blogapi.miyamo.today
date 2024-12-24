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

// Article is an Output DTO for GetById use-case.
type Article struct {
	id           string
	title        string
	body         string
	thumbnailUrl string
	createdAt    synchro.Time[tz.UTC]
	updatedAt    synchro.Time[tz.UTC]
	tags         []Tag
}

// IsOutDto is a marker method for OutDTO.
func (a Article) IsOutDto() {}

// Id returns the id of the article
func (a Article) Id() string { return a.id }

// Title returns the title of the article
func (a Article) Title() string { return a.title }

// Body returns the body of the article
func (a Article) Body() string { return a.body }

// ThumbnailUrl returns the thumbnail url of the article
func (a Article) ThumbnailUrl() string { return a.thumbnailUrl }

// CreatedAt returns date the article was created
func (a Article) CreatedAt() synchro.Time[tz.UTC] { return a.createdAt }

// UpdatedAt returns the date the article was last updated.
func (a Article) UpdatedAt() synchro.Time[tz.UTC] { return a.updatedAt }

// Tags return the tags attached to the article
func (a Article) Tags() []Tag { return a.tags }

// NewArticle is constructor of Article
func NewArticle(
	id string,
	title string,
	body string,
	thumbnailUrl string,
	createdAt synchro.Time[tz.UTC],
	updatedAt synchro.Time[tz.UTC],
	tags []Tag,
) Article {
	return Article{
		id:           id,
		title:        title,
		body:         body,
		thumbnailUrl: thumbnailUrl,
		createdAt:    createdAt,
		updatedAt:    updatedAt,
		tags:         tags,
	}
}

// Tag is a DTO for tag
type Tag struct {
	id   string
	name string
}

// IsOutDto is a marker method for OutDTO.
func (t Tag) IsOutDto() {}

// Id returns the id of the tag
func (t Tag) Id() string { return t.id }

// Name returns the name of the tag
func (t Tag) Name() string { return t.name }

// NewTag is constructor of Tag
func NewTag(id string, name string) Tag {
	return Tag{id: id, name: name}
}

// GetByIdOutDto is an Output DTO for GetById use-case.
type GetByIdOutDto struct {
	Article
}

// NewGetByIdOutDto is constructor of GetByIdOutDto
func NewGetByIdOutDto(
	id string,
	title string,
	body string,
	thumbnailUrl string,
	createdAt synchro.Time[tz.UTC],
	updatedAt synchro.Time[tz.UTC],
	tags []Tag,
) GetByIdOutDto {
	return GetByIdOutDto{
		Article{
			id:           id,
			title:        title,
			body:         body,
			thumbnailUrl: thumbnailUrl,
			createdAt:    createdAt,
			updatedAt:    updatedAt,
			tags:         tags,
		},
	}
}

// GetAllOutDto is an Output DTO for GetAll use-case.
type GetAllOutDto struct {
	articles []Article
}

// NewGetAllOutDto is constructor of GetAllOutDto.
func NewGetAllOutDto(articles []Article) GetAllOutDto {
	return GetAllOutDto{articles: articles}
}

// IsOutDto is a marker method for OutDTO.
func (o GetAllOutDto) IsOutDto() {}

// Articles returns the articles.
func (o *GetAllOutDto) Articles() []Article { return o.articles }

// GetNextInDto is an Input DTO for GetNext use-case.
type GetNextInDto struct {
	first  int
	cursor *string
}

// IsInDto is a marker method for InDTO.
func (i GetNextInDto) IsInDto() {}

// First returns the first.
func (i GetNextInDto) First() int { return i.first }

// Cursor returns the cursor.
func (i GetNextInDto) Cursor() *string { return i.cursor }

// NewGetNextInDto is constructor of GetNextInDto.
func NewGetNextInDto(first int, cursor *string) GetNextInDto {
	return GetNextInDto{first: first, cursor: cursor}
}

// GetNextOutDto is an Output DTO for GetNext use-case.
type GetNextOutDto struct {
	articles []Article
	hasNext  bool
}

// NewGetNextOutDto is constructor of GetNextOutDto.
func NewGetNextOutDto(articles []Article, hasNext bool) GetNextOutDto {
	return GetNextOutDto{articles: articles, hasNext: hasNext}
}

// IsOutDto is a marker method for OutDTO.
func (o GetNextOutDto) IsOutDto() {}

// Articles returns the articles.
func (o *GetNextOutDto) Articles() []Article { return o.articles }

// HasNext returns whether there is still next items.
func (o *GetNextOutDto) HasNext() bool { return o.hasNext }

// GetPrevInDto is an Input DTO for GetPrev use-case.
type GetPrevInDto struct {
	last   int
	cursor *string
}

// IsInDto is a marker method for InDTO.
func (i GetPrevInDto) IsInDto() {}

// Last returns the last.
func (i GetPrevInDto) Last() int { return i.last }

// Cursor returns the cursor.
func (i GetPrevInDto) Cursor() *string { return i.cursor }

// NewGetPrevInDto is constructor of GetPrevInDto.
func NewGetPrevInDto(last int, cursor *string) GetPrevInDto {
	return GetPrevInDto{last: last, cursor: cursor}
}

// GetPrevOutDto is an Output DTO for GetPrev use-case.
type GetPrevOutDto struct {
	articles []Article
	hasPrev  bool
}

// NewGetPrevOutDto is constructor of GetPrevOutDto.
func NewGetPrevOutDto(articles []Article, hasPrev bool) GetPrevOutDto {
	return GetPrevOutDto{articles: articles, hasPrev: hasPrev}
}

// IsOutDto is a marker method for OutDTO.
func (o GetPrevOutDto) IsOutDto() {}

// Articles returns the articles.
func (o *GetPrevOutDto) Articles() []Article { return o.articles }

// HasPrevious returns whether there is still precious items.
func (o *GetPrevOutDto) HasPrevious() bool { return o.hasPrev }
