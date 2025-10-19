package usecase

import (
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
)

// SyncUsecaseInDto is an in dto of the Sync.SyncBlogSnapshotWithEvents
type SyncUsecaseInDto struct {
	EventID    string               `dynamodbav:"event_id"`
	ArticleID  string               `dynamodbav:"article_id"`
	Title      *string              `dynamodbav:"title"`
	Content    *string              `dynamodbav:"content"`
	Thumbnail  *string              `dynamodbav:"thumbnail"`
	Tags       []string             `dynamodbav:"tags"`
	AttachTags []string             `dynamodbav:"attach_tags"`
	DetachTags []string             `dynamodbav:"detach_tags"`
	Invisible  *bool                `dynamodbav:"invisible"`
	EventAt    synchro.Time[tz.UTC] `dynamodbav:"-"`
}
