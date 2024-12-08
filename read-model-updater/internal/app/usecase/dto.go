package usecase

// SyncUsecaseInDto is an in dto of the Sync.SyncBlogSnapshotWithEvents
type SyncUsecaseInDto struct {
	eventID    string
	articleID  string
	title      *string
	content    *string
	thumbnail  *string
	attachTags []string
	detachTags []string
	invisible  *bool
}

// IsInDto is a marker for in dto.
func (d SyncUsecaseInDto) IsInDto() {}

// EventId returns the event id.
func (d SyncUsecaseInDto) EventId() string {
	return d.eventID
}

// ArticleId returns the article id.
func (d SyncUsecaseInDto) ArticleId() string {
	return d.articleID
}

// Title returns the article title.
func (d SyncUsecaseInDto) Title() *string {
	return d.title
}

// Content returns the article content.
func (d SyncUsecaseInDto) Content() *string {
	return d.content
}

// Thumbnail returns the article thumbnail.
func (d SyncUsecaseInDto) Thumbnail() *string {
	return d.thumbnail
}

// AttacheTags returns the tags to attach.
func (d SyncUsecaseInDto) AttacheTags() []string {
	return d.attachTags
}

// DetachTags returns the tags to detach.
func (d SyncUsecaseInDto) DetachTags() []string {
	return d.detachTags
}

// Invisible returns the article visibility.
func (d SyncUsecaseInDto) Invisible() *bool {
	return d.invisible
}

// NewSyncUsecaseInDto creates a new SyncUsecaseInDto
func NewSyncUsecaseInDto(eventID, articleID string, title, content, thumbnail *string, attachTags, detachTags []string, invisible *bool) SyncUsecaseInDto {
	return SyncUsecaseInDto{
		eventID:    eventID,
		articleID:  articleID,
		title:      title,
		content:    content,
		thumbnail:  thumbnail,
		attachTags: attachTags,
		detachTags: detachTags,
		invisible:  invisible,
	}
}
