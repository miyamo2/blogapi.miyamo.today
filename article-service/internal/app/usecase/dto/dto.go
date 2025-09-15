package dto

import (
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
)

// GetByIDInput is an Input DTO for GetById use-case
type GetByIDInput struct {
	id string
}

// ID returns the ID of the article to be got
func (i GetByIDInput) ID() string { return i.id }

// NewGetByIDInput constructs GetByIdInDta.
func NewGetByIDInput(id string) GetByIDInput {
	return GetByIDInput{id: id}
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

// ID returns the id of the article
func (a Article) ID() string { return a.id }

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

// NewArticle constructs Article
func NewArticle(
	id string,
	title string,
	body string,
	thumbnailUrl string,
	createdAt synchro.Time[tz.UTC],
	updatedAt synchro.Time[tz.UTC],
	tags ...Tag,
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

// ID returns the id of the tag
func (t Tag) ID() string { return t.id }

// Name returns the name of the tag
func (t Tag) Name() string { return t.name }

// NewTag constructs Tag
func NewTag(id string, name string) Tag {
	return Tag{id: id, name: name}
}

// GetByIDOutput is an Output DTO for GetById use-case.
type GetByIDOutput = Article

// NewGetByIDOutput constructs GetByIDOutput
func NewGetByIDOutput(
	id string,
	title string,
	body string,
	thumbnailUrl string,
	createdAt synchro.Time[tz.UTC],
	updatedAt synchro.Time[tz.UTC],
	tags ...Tag,
) GetByIDOutput {
	return NewArticle(id, title, body, thumbnailUrl, createdAt, updatedAt, tags...)
}

// ListAllOutput is an Output DTO for ListAll use-case.
type ListAllOutput struct {
	articles []Article
}

// NewListAllOutput constructs ListAllOutput.
func NewListAllOutput(articles ...Article) ListAllOutput {
	return ListAllOutput{articles: articles}
}

// Articles returns the articles.
func (o *ListAllOutput) Articles() []Article { return o.articles }

// ListAfterInput is an Input DTO for ListAfter use-case.
type ListAfterInput struct {
	first  int
	cursor *string
}

// First returns the first.
func (i ListAfterInput) First() int { return i.first }

// Cursor returns the cursor.
func (i ListAfterInput) Cursor() *string { return i.cursor }

// NewListAfterInputOption is an option for NewListAfterInput
type NewListAfterInputOption func(*ListAfterInput)

// ListAfterInputWithCursor sets the cursor option for NewListAfterInput
func ListAfterInputWithCursor[T string | *string](cursor T) NewListAfterInputOption {
	return func(i *ListAfterInput) {
		switch v := any(cursor).(type) {
		case string:
			i.cursor = &v
		case *string:
			i.cursor = v
		}
	}
}

// NewListAfterInput constructs ListAfterInput.
func NewListAfterInput(first int, options ...NewListAfterInputOption) ListAfterInput {
	input := ListAfterInput{first: first}
	for _, opt := range options {
		opt(&input)
	}
	return input
}

// ListAfterOutput is an Output DTO for ListAfter use-case.
type ListAfterOutput struct {
	articles []Article
	hasNext  bool
}

// NewListAfterOutput constructs ListAfterOutput.
func NewListAfterOutput(hasNext bool, articles ...Article) ListAfterOutput {
	return ListAfterOutput{articles: articles, hasNext: hasNext}
}

// Articles returns the articles.
func (o *ListAfterOutput) Articles() []Article { return o.articles }

// HasNext returns whether there is still next items.
func (o *ListAfterOutput) HasNext() bool { return o.hasNext }

// ListBeforeInput is an Input DTO for ListBefore use-case.
type ListBeforeInput struct {
	last   int
	cursor *string
}

// Last returns the last.
func (i ListBeforeInput) Last() int { return i.last }

// Cursor returns the cursor.
func (i ListBeforeInput) Cursor() *string { return i.cursor }

// ListBeforeOutput is an Output DTO for ListBefore use-case.
type ListBeforeOutput struct {
	articles []Article
	hasPrev  bool
}

// NewListBeforeInputOption is an option for NewListBeforeInput
type NewListBeforeInputOption func(*ListBeforeInput)

// ListBeforeInputWithCursor sets the cursor option for NewListBeforeInput
func ListBeforeInputWithCursor[T string | *string](cursor T) NewListBeforeInputOption {
	return func(i *ListBeforeInput) {
		switch v := any(cursor).(type) {
		case string:
			i.cursor = &v
		case *string:
			i.cursor = v
		}
	}
}

// NewListBeforeInput constructs ListBeforeInput.
func NewListBeforeInput(last int, options ...NewListBeforeInputOption) ListBeforeInput {
	input := ListBeforeInput{last: last}
	for _, opt := range options {
		opt(&input)
	}
	return input
}

// Articles returns the articles.
func (o *ListBeforeOutput) Articles() []Article { return o.articles }

// HasPrevious returns whether there is still precious items.
func (o *ListBeforeOutput) HasPrevious() bool { return o.hasPrev }

// NewListBeforeOutput constructs ListBeforeOutput.
func NewListBeforeOutput(hasPrev bool, articles ...Article) ListBeforeOutput {
	return ListBeforeOutput{articles: articles, hasPrev: hasPrev}
}
