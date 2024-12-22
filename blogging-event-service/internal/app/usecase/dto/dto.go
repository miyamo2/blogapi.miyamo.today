package dto

import (
	"net/url"
)

// CreateArticleInDto is an Input DTO for CreateArticle use-case
type CreateArticleInDto struct {
	title        string
	body         string
	thumbnailUrl string
	tagNames     []string
}

// Title returns the title of the article to be created
func (i CreateArticleInDto) Title() string {
	return i.title
}

// Body returns the body of the article to be created
func (i CreateArticleInDto) Body() string {
	return i.body
}

// ThumbnailUrl returns the thumbnail URL of the article to be created
func (i CreateArticleInDto) ThumbnailUrl() string {
	return i.thumbnailUrl
}

// TagNames returns the tag names of the article to be created
func (i CreateArticleInDto) TagNames() []string {
	return i.tagNames
}

// NewCreateArticleInDto is constructor of CreateArticle.
func NewCreateArticleInDto(title, body, thumbnailUrl string, tagNames []string) CreateArticleInDto {
	return CreateArticleInDto{
		title:        title,
		body:         body,
		thumbnailUrl: thumbnailUrl,
		tagNames:     tagNames,
	}
}

// CreateArticleOutDto is an Output DTO for CreateArticle use-case
type CreateArticleOutDto struct {
	eventID   string
	articleID string
}

// EventID returns the ID of the event
func (o CreateArticleOutDto) EventID() string {
	return o.eventID
}

// ArticleID returns the ID of the article
func (o CreateArticleOutDto) ArticleID() string {
	return o.articleID
}

// NewCreateArticleOutDto is constructor of CreateArticleOutDto.
func NewCreateArticleOutDto(eventID, articleID string) CreateArticleOutDto {
	return CreateArticleOutDto{
		eventID:   eventID,
		articleID: articleID,
	}
}

// UpdateArticleTitleInDto is an Input DTO for UpdateArticleTitle use-case
type UpdateArticleTitleInDto struct {
	id    string
	title string
}

// ID returns the ID of the article to be updated
func (i UpdateArticleTitleInDto) ID() string {
	return i.id
}

// Title returns the title of the article to be updated
func (i UpdateArticleTitleInDto) Title() string {
	return i.title
}

// NewUpdateArticleTitleInDto is constructor of UpdateArticleTitleInDto.
func NewUpdateArticleTitleInDto(id string, title string) UpdateArticleTitleInDto {
	return UpdateArticleTitleInDto{
		id:    id,
		title: title,
	}
}

// UpdateArticleTitleOutDto is an Output DTO for UpdateArticleTitle use-case
type UpdateArticleTitleOutDto struct {
	eventID   string
	articleID string
}

// EventID returns the ID of the event
func (o UpdateArticleTitleOutDto) EventID() string {
	return o.eventID
}

// ArticleID returns the ID of the article
func (o UpdateArticleTitleOutDto) ArticleID() string {
	return o.articleID
}

// NewUpdateArticleTitleOutDto is constructor of UpdateArticleTitleOutDto.
func NewUpdateArticleTitleOutDto(eventID, articleID string) UpdateArticleTitleOutDto {
	return UpdateArticleTitleOutDto{
		eventID:   eventID,
		articleID: articleID,
	}
}

// UpdateArticleBodyInDto is an Input DTO for UpdateArticleBody use-case
type UpdateArticleBodyInDto struct {
	id   string
	body string
}

// ID returns the ID of the article to be updated
func (i UpdateArticleBodyInDto) ID() string {
	return i.id
}

// Body returns the body of the article to be updated
func (i UpdateArticleBodyInDto) Body() string {
	return i.body
}

// NewUpdateArticleBodyInDto is constructor of UpdateArticleBodyInDto.
func NewUpdateArticleBodyInDto(id string, body string) UpdateArticleBodyInDto {
	return UpdateArticleBodyInDto{
		id:   id,
		body: body,
	}
}

// UpdateArticleBodyOutDto is an Output DTO for UpdateArticleBody use-case
type UpdateArticleBodyOutDto struct {
	eventID   string
	articleID string
}

// EventID returns the ID of the event
func (o UpdateArticleBodyOutDto) EventID() string {
	return o.eventID
}

// ArticleID returns the ID of the article
func (o UpdateArticleBodyOutDto) ArticleID() string {
	return o.articleID
}

// NewUpdateArticleBodyOutDto is constructor of UpdateArticleBodyOutDto.
func NewUpdateArticleBodyOutDto(eventID, articleID string) UpdateArticleBodyOutDto {
	return UpdateArticleBodyOutDto{
		eventID:   eventID,
		articleID: articleID,
	}
}

// UpdateArticleThumbnailInDto is an Input DTO for UpdateArticleThumbnail use-case
type UpdateArticleThumbnailInDto struct {
	id           string
	thumbnailUrl url.URL
}

// ID returns the ID of the article to be updated
func (i UpdateArticleThumbnailInDto) ID() string {
	return i.id
}

// ThumbnailUrl returns the thumbnail URL of the article to be updated
func (i UpdateArticleThumbnailInDto) ThumbnailUrl() url.URL {
	return i.thumbnailUrl
}

// NewUpdateArticleThumbnailInDto is constructor of UpdateArticleThumbnailInDto.
func NewUpdateArticleThumbnailInDto(id string, thumbnailUrl url.URL) UpdateArticleThumbnailInDto {
	return UpdateArticleThumbnailInDto{
		id:           id,
		thumbnailUrl: thumbnailUrl,
	}
}

// UpdateArticleThumbnailOutDto is an Output DTO for UpdateArticleThumbnail use-case
type UpdateArticleThumbnailOutDto struct {
	eventID   string
	articleID string
}

// EventID returns the ID of the event
func (o UpdateArticleThumbnailOutDto) EventID() string {
	return o.eventID
}

// ArticleID returns the ID of the article
func (o UpdateArticleThumbnailOutDto) ArticleID() string {
	return o.articleID
}

// NewUpdateArticleThumbnailOutDto is constructor of UpdateArticleThumbnailOutDto.
func NewUpdateArticleThumbnailOutDto(eventID, articleID string) UpdateArticleThumbnailOutDto {
	return UpdateArticleThumbnailOutDto{
		eventID:   eventID,
		articleID: articleID,
	}
}

// AttachTagsInDto is an Input DTO for AttachTag use-case
type AttachTagsInDto struct {
	id       string
	tagNames []string
}

// ID returns the ID of the article to attach tags
func (i AttachTagsInDto) ID() string {
	return i.id
}

// TagNames returns the tag names to attach
func (i AttachTagsInDto) TagNames() []string {
	return i.tagNames
}

// NewAttachTagsInDto is constructor of AttachTagsInDto.
func NewAttachTagsInDto(id string, names []string) AttachTagsInDto {
	return AttachTagsInDto{
		id:       id,
		tagNames: names,
	}
}

// AttachTagsOutDto is an Output DTO for AttachTags use-case
type AttachTagsOutDto struct {
	eventID   string
	articleID string
}

// EventID returns the ID of the event
func (o AttachTagsOutDto) EventID() string {
	return o.eventID
}

// ArticleID returns the ID of the article
func (o AttachTagsOutDto) ArticleID() string {
	return o.articleID
}

// NewAttachTagsOutDto is constructor of AttachTagsOutDto.
func NewAttachTagsOutDto(eventID, articleID string) AttachTagsOutDto {
	return AttachTagsOutDto{
		eventID:   eventID,
		articleID: articleID,
	}
}

// DetachTagsInDto is an Input DTO for DetachTags use-case
type DetachTagsInDto struct {
	id       string
	tagNames []string
}

// ID returns the ID of the article to detach tags
func (i DetachTagsInDto) ID() string {
	return i.id
}

// TagNames returns the tag names to detach
func (i DetachTagsInDto) TagNames() []string {
	return i.tagNames
}

// NewDetachTagsInDto is constructor of DetachTagsInDto.
func NewDetachTagsInDto(id string, names []string) DetachTagsInDto {
	return DetachTagsInDto{
		id:       id,
		tagNames: names,
	}
}

// DetachTagsOutDto is an Output DTO for DetachTags use-case
type DetachTagsOutDto struct {
	eventID   string
	articleID string
}

// EventID returns the ID of the event
func (o DetachTagsOutDto) EventID() string {
	return o.eventID
}

// ArticleID returns the ID of the article
func (o DetachTagsOutDto) ArticleID() string {
	return o.articleID
}

// NewDetachTagsOutDto is constructor of DetachTagsOutDto.
func NewDetachTagsOutDto(eventID, articleID string) DetachTagsOutDto {
	return DetachTagsOutDto{
		eventID:   eventID,
		articleID: articleID,
	}
}

// UploadImageInDto is an Input DTO for UploadImage use-case
type UploadImageInDto struct {
	name  string
	bytes []byte
}

// Name returns the name of the image to be uploaded
func (i UploadImageInDto) Name() string {
	return i.name
}

// Bytes returns the bytes of the image to be uploaded
func (i UploadImageInDto) Bytes() []byte {
	return i.bytes
}

// NewUploadImageInDto is constructor of UploadImageInDto.
func NewUploadImageInDto(name string, bytes []byte) UploadImageInDto {
	return UploadImageInDto{
		name:  name,
		bytes: bytes,
	}
}

// UploadImageOutDto is an Output DTO for UploadImage use-case
type UploadImageOutDto struct {
	uri url.URL
}

// URL returns the URL of the uploaded image
func (o UploadImageOutDto) URL() url.URL {
	return o.uri
}

// NewUploadImageOutDto is constructor of UploadImageOutDto.
func NewUploadImageOutDto(uri url.URL) UploadImageOutDto {
	return UploadImageOutDto{
		uri: uri,
	}
}
