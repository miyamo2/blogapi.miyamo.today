package model

import (
	"encoding/base64"
	"fmt"
	"slices"

	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
)

type ArticleCommand struct {
	id        string
	title     string
	body      string
	thumbnail string
	tags      []ArticleTagCommand
	eventAt   synchro.Time[tz.UTC]
}

// IsCommandModel is a marker method for CommandModel.
func (a ArticleCommand) IsCommandModel() {}

func (a ArticleCommand) ID() string {
	return a.id
}

func (a ArticleCommand) Title() string {
	return a.title
}

func (a ArticleCommand) Body() string {
	return a.body
}

func (a ArticleCommand) Thumbnail() string {
	return a.thumbnail
}

func (a ArticleCommand) Tags() []ArticleTagCommand {
	return a.tags
}

func (a ArticleCommand) EventAt() synchro.Time[tz.UTC] {
	return a.eventAt
}

type ArticleTagCommand struct {
	id   string
	name string
}

// IsCommandModel is a marker method for CommandModel.
func (a ArticleTagCommand) IsCommandModel() {}

func (a ArticleTagCommand) ID() string {
	return a.id
}

func (a ArticleTagCommand) Name() string {
	return a.name
}

func newArticleTagCommands(names ...string) []ArticleTagCommand {
	result := make([]ArticleTagCommand, 0, len(names))
	for _, n := range names {
		result = append(
			result, ArticleTagCommand{
				id:   base64.URLEncoding.EncodeToString([]byte(fmt.Sprintf("tag:%s", n))),
				name: n,
			},
		)
	}
	return result
}

func ArticleCommandFromBloggingEvents(events []BloggingEvent) *ArticleCommand {
	if len(events) == 0 {
		return nil
	}
	result := ArticleCommand{
		id: events[0].ArticleID(),
	}
	var tagNames []string
	for _, e := range events {
		if e.title != nil {
			result.title = *e.title
		}
		if e.content != nil {
			result.body = *e.content
		}
		if e.thumbnail != nil {
			result.thumbnail = *e.thumbnail
		}
		tagNames = slices.DeleteFunc(
			append(tagNames, append(e.Tags(), e.AttachTags()...)...), func(v string) bool {
				return slices.Contains(e.DetachTags(), v)
			},
		)
	}
	tags := newArticleTagCommands(tagNames...)
	result.tags = tags
	return &result
}
