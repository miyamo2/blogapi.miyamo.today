package model

// Tag is a query service model for tag.
type Tag struct {
	id       string
	name     string
	articles []Article
}

// IsQueryServiceModel is a marker method for QueryServiceModel.
func (t *Tag) IsQueryServiceModel() {}

// ID returns the id of the tag
func (t *Tag) ID() string { return t.id }

// Name returns the name of the tag
func (t *Tag) Name() string { return t.name }

// Articles returns the articles of the tag
func (t *Tag) Articles() []Article { return t.articles }

func (t *Tag) AddArticle(article Article) {
	t.articles = append(t.articles, article)
}

type NewTagOption func(*Tag)

func WithTagsSize(size int) NewTagOption {
	return func(t *Tag) {
		t.articles = make([]Article, 0, size)
	}
}

// NewTag constructor of Tag.
func NewTag(id, name string, opt ...NewTagOption) Tag {
	t := Tag{
		id:       id,
		name:     name,
		articles: make([]Article, 0),
	}
	for _, o := range opt {
		o(&t)
	}
	return t
}

// Article is a query service model for article.
type Article struct {
	id        string
	title     string
	thumbnail string
	createdAt string
	updatedAt string
}

// IsQueryServiceModel is a marker method for QueryServiceModel.
func (a Article) IsQueryServiceModel() {}

// ID returns the id of the article
func (a Article) ID() string { return a.id }

// Title returns the title of the article
func (a Article) Title() string { return a.title }

// Thumbnail returns the thumbnail url of the article
func (a Article) Thumbnail() string { return a.thumbnail }

// CreatedAt returns date the article was created
func (a Article) CreatedAt() string { return a.createdAt }

// UpdatedAt returns the date the article was last updated.
func (a Article) UpdatedAt() string { return a.updatedAt }

// NewArticle constructor of Article.
func NewArticle(id, title, thumbnail, createdAt, updatedAt string) Article {
	a := Article{
		id:        id,
		title:     title,
		thumbnail: thumbnail,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
	return a
}
