package model

type CreateArticleEvent struct {
	title     string
	content   string
	thumbnail string
	tags      []string
}

func (c CreateArticleEvent) Title() string {
	return c.title
}

func (c CreateArticleEvent) Content() string {
	return c.content
}

func (c CreateArticleEvent) Thumbnail() string {
	return c.thumbnail
}

func (c CreateArticleEvent) Tags() []string {
	return c.tags
}

func NewCreateArticleEvent(title, content, thumbnail string, tags []string) CreateArticleEvent {
	return CreateArticleEvent{
		title:     title,
		content:   content,
		thumbnail: thumbnail,
		tags:      tags,
	}
}

type UpdateArticleTitleEvent struct {
	articleID string
	title     string
}

func (u UpdateArticleTitleEvent) ArticleID() string {
	return u.articleID
}

func (u UpdateArticleTitleEvent) Title() string {
	return u.title
}

func NewUpdateArticleTitleEvent(articleID, title string) UpdateArticleTitleEvent {
	return UpdateArticleTitleEvent{
		articleID: articleID,
		title:     title,
	}
}

// UpdateArticleBodyEvent is an event to update the article body.
type UpdateArticleBodyEvent struct {
	articleID string
	body      string
}

// ArticleID returns the article id.
func (u UpdateArticleBodyEvent) ArticleID() string {
	return u.articleID
}

// Body returns the article body.
func (u UpdateArticleBodyEvent) Body() string {
	return u.body
}

// NewUpdateArticleBodyEvent creates a new UpdateArticleBodyEvent.
func NewUpdateArticleBodyEvent(id string, body string) UpdateArticleBodyEvent {
	return UpdateArticleBodyEvent{
		articleID: id,
		body:      body,
	}
}

type BloggingEventKey struct {
	eventID   string
	articleID string
}

func (b BloggingEventKey) EventID() string {
	return b.eventID
}

func (b BloggingEventKey) ArticleID() string {
	return b.articleID
}

func NewBloggingEventKey(eventID, articleID string) BloggingEventKey {
	return BloggingEventKey{
		eventID:   eventID,
		articleID: articleID,
	}
}
