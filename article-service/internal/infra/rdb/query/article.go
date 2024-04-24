package query

// Article is a query service model for article.
type Article struct {
	id        string
	title     string
	body      string
	thumbnail string
	createdAt string
	updatedAt string
	tags      []Tag
}

// IsQueryServiceModel is a marker method for QueryServiceModel.
func (a Article) IsQueryServiceModel() {}

// ID returns the id of the article
func (a Article) ID() string { return a.id }

// Title returns the title of the article
func (a Article) Title() string { return a.title }

// Body returns the body of the article
func (a Article) Body() string { return a.body }

// Thumbnail returns the thumbnail url of the article
func (a Article) Thumbnail() string { return a.thumbnail }

// CreatedAt returns date the article was created
func (a Article) CreatedAt() string { return a.createdAt }

// UpdatedAt returns the date the article was last updated.
func (a Article) UpdatedAt() string { return a.updatedAt }

// Tags return the tags attached to the article
func (a Article) Tags() []Tag { return a.tags }

// Tag is a query service model for tag.
type Tag struct {
	id   string
	name string
}

// IsQueryServiceModel is a marker method for QueryServiceModel.
func (t Tag) IsQueryServiceModel() {}

// ID returns the id of the tag
func (t Tag) ID() string { return t.id }

// Name returns the name of the tag
func (t Tag) Name() string { return t.name }

type NewArticleOption func(*Article)

func WithTagsSize(size int) NewArticleOption {
	return func(a *Article) {
		a.tags = make([]Tag, 0, size)
	}
}

// NewArticle constructor of Article.
func NewArticle(id, title, body, thumbnail, createdAt, updatedAt string, opt ...NewArticleOption) Article {
	a := Article{
		id:        id,
		title:     title,
		body:      body,
		thumbnail: thumbnail,
		createdAt: createdAt,
		updatedAt: updatedAt,
		tags:      make([]Tag, 0),
	}
	for _, o := range opt {
		o(&a)
	}
	return a
}

func (a *Article) AddTag(tag Tag) {
	a.tags = append(a.tags, tag)
}

// NewTag constructor of Tag.
func NewTag(id, name string) Tag {
	return Tag{
		id:   id,
		name: name,
	}
}
