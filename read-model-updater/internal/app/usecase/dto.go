package usecase

import (
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
)

// SyncUsecaseInDto is an in dto of the Sync.SyncBlogSnapshotWithEvents
type SyncUsecaseInDto struct {
	EventID    string               `json:"event_id"`
	ArticleID  string               `json:"article_id"`
	Title      *string              `json:"title"`
	Content    *string              `json:"content"`
	Thumbnail  *string              `json:"thumbnail"`
	Tags       []string             `json:"tags"`
	AttachTags []string             `json:"attach_tags"`
	DetachTags []string             `json:"detach_tags"`
	Invisible  *bool                `json:"invisible"`
	EventAt    synchro.Time[tz.UTC] `json:"-"`
}
