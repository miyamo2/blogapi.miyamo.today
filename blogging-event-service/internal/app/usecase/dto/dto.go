package dto

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
