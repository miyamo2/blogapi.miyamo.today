package model

import "net/url"

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

// UpdateArticleThumbnailEvent is an event to update the article thumbnail.
type UpdateArticleThumbnailEvent struct {
	articleID string
	thumbnail url.URL
}

// ArticleID returns the article id.
func (u UpdateArticleThumbnailEvent) ArticleID() string {
	return u.articleID
}

// Thumbnail returns the article thumbnail.
func (u UpdateArticleThumbnailEvent) Thumbnail() url.URL {
	return u.thumbnail
}

// NewUpdateArticleThumbnailEvent creates a new UpdateArticleThumbnailEvent.
func NewUpdateArticleThumbnailEvent(id string, thumbnail url.URL) UpdateArticleThumbnailEvent {
	return UpdateArticleThumbnailEvent{
		articleID: id,
		thumbnail: thumbnail,
	}
}

// AttachTagsEvent is an event to attach tags to the article.
type AttachTagsEvent struct {
	articleID string
	tags      []string
}

// ArticleID returns the article id.
func (a AttachTagsEvent) ArticleID() string {
	return a.articleID
}

// Tags returns the tags.
func (a AttachTagsEvent) Tags() []string {
	return a.tags
}

// NewAttachTagsEvent creates a new AttachTagsEvent.
func NewAttachTagsEvent(articleID string, tags []string) AttachTagsEvent {
	return AttachTagsEvent{
		articleID: articleID,
		tags:      tags,
	}
}

// DetachTagsEvent is an event to detach tags from the article.
type DetachTagsEvent struct {
	articleID string
	tags      []string
}

// ArticleID returns the article id.
func (d DetachTagsEvent) ArticleID() string {
	return d.articleID
}

// Tags returns the tags.
func (d DetachTagsEvent) Tags() []string {
	return d.tags
}

// NewDetachTagsEvent creates a new DetachTagsEvent.
func NewDetachTagsEvent(articleID string, tags []string) DetachTagsEvent {
	return DetachTagsEvent{
		articleID: articleID,
		tags:      tags,
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
