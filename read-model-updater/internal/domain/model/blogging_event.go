package model

type BloggingEvent struct {
	eventID    string
	articleID  string
	title      *string
	content    *string
	thumbnail  *string
	tags       []string
	attachTags []string
	detachTags []string
	invisible  *bool
}

func (b BloggingEvent) EventID() string {
	return b.eventID
}

func (b BloggingEvent) ArticleID() string {
	return b.articleID
}

func (b BloggingEvent) Title() *string {
	return b.title
}

func (b BloggingEvent) Content() *string {
	return b.content
}

func (b BloggingEvent) Thumbnail() *string {
	return b.thumbnail
}

func (b BloggingEvent) AttachTags() []string {
	return b.attachTags
}

func (b BloggingEvent) Tags() []string {
	return b.tags
}

func (b BloggingEvent) DetachTags() []string {
	return b.detachTags
}

func (b BloggingEvent) Invisible() *bool {
	return b.invisible
}

func NewBloggingEvent(
	eventID, articleID string, title, content, thumbnail *string, tags, attachTag, detacheTag []string, invisible *bool,
) BloggingEvent {
	return BloggingEvent{
		eventID:    eventID,
		articleID:  articleID,
		title:      title,
		content:    content,
		thumbnail:  thumbnail,
		tags:       tags,
		attachTags: attachTag,
		detachTags: detacheTag,
		invisible:  invisible,
	}
}
