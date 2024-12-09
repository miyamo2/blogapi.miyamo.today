package model

type BloggingEvent struct {
	eventID     string
	articleID   string
	title       *string
	content     *string
	thumbnail   *string
	tags        []string
	attacheTags []string
	detachTags  []string
	invisible   *bool
}

// IsQueryModel is a marker method for the QueryModel.
func (b BloggingEvent) IsQueryModel() {}

// IsCommandModel is a marker method for the CommandModel.
func (b BloggingEvent) IsCommandModel() {}

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

func (b BloggingEvent) AttacheTags() []string {
	return b.attacheTags
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

func NewBloggingEvent(eventID, articleID string, title, content, thumbnail *string, tags, attacheTag, detacheTag []string, invisible *bool) BloggingEvent {
	return BloggingEvent{
		eventID:     eventID,
		articleID:   articleID,
		title:       title,
		content:     content,
		thumbnail:   thumbnail,
		tags:        tags,
		attacheTags: attacheTag,
		detachTags:  detacheTag,
		invisible:   invisible,
	}
}
