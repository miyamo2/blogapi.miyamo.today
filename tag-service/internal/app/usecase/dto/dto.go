package dto

import (
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
)

// GetByIdInput is an Input DTO for GetById use-case
type GetByIdInput struct {
	id string
}

// Id returns the ID of the article to be got
func (i GetByIdInput) Id() string { return i.id }

// NewGetByIdInput constructs GetByIdInputDta.
func NewGetByIdInput(id string) GetByIdInput {
	return GetByIdInput{id: id}
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

// Id returns the id of the tag
func (t Tag) Id() string { return t.id }

// Name returns the name of the tag
func (t Tag) Name() string { return t.name }

// NewTag constructs Tag
func NewTag(id string, name string, articles ...Article) Tag {
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

// NewArticle constructs Article
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

// GetByIdOutput is an Output DTO for GetById use-case.
type GetByIdOutput = Tag

// ListAllOutput is an Output DTO for ListAll use-case.
type ListAllOutput struct {
	tags []Tag
}

// NewListAllOutput constructs ListAllOutput
func NewListAllOutput(tags ...Tag) ListAllOutput {
	return ListAllOutput{
		tags: tags,
	}
}

// Tags returns the tags of the ListAllOutput
func (g ListAllOutput) Tags() []Tag {
	return g.tags
}

// ListAfterInput is an Input DTO for ListAfter use-case.
type ListAfterInput struct {
	first  int
	cursor *string
}

// First returns the number of tags to get
func (i ListAfterInput) First() int { return i.first }

// Cursor returns the cursor that indicates the position to get tags from here.
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
	tags    []Tag
	hasNext bool
}

// NewListAfterOutput constructs ListAfterOutput.
func NewListAfterOutput(hasNext bool, tags ...Tag) ListAfterOutput {
	return ListAfterOutput{tags: tags, hasNext: hasNext}
}

// Tags returns the tags.
func (o ListAfterOutput) Tags() []Tag { return o.tags }

// HasNext returns whether there is still next items.
func (o ListAfterOutput) HasNext() bool { return o.hasNext }

// ListBeforeInput is an Input DTO for ListBefore use-case.
type ListBeforeInput struct {
	last   int
	cursor *string
}

// Last returns the number of tags to get
func (i ListBeforeInput) Last() int { return i.last }

// Cursor returns the cursor that indicates the position to get tags from here.
func (i ListBeforeInput) Cursor() *string { return i.cursor }

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

// ListBeforeOutput is an Output DTO for ListBefore use-case.
type ListBeforeOutput struct {
	tags        []Tag
	hasPrevious bool
}

// NewListBeforeOutput constructs ListBeforeOutput.
func NewListBeforeOutput(hasPrevious bool, tags ...Tag) ListBeforeOutput {
	return ListBeforeOutput{tags: tags, hasPrevious: hasPrevious}
}

// Tags returns the tags.
func (o ListBeforeOutput) Tags() []Tag { return o.tags }

// HasPrevious returns whether there is still previous items.
func (o ListBeforeOutput) HasPrevious() bool { return o.hasPrevious }
