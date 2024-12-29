// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"blogapi.miyamo.today/federator/internal/pkg/gqlscalar"
	"github.com/99designs/gqlgen/graphql"
)

type Node interface {
	IsNode()
	GetID() string
}

type ArticleConnection struct {
	Edges      []*ArticleEdge `json:"edges"`
	PageInfo   *PageInfo      `json:"pageInfo"`
	TotalCount int            `json:"totalCount"`
}

type ArticleEdge struct {
	Cursor string       `json:"cursor"`
	Node   *ArticleNode `json:"node"`
}

type ArticleNode struct {
	ID           string                `json:"id"`
	Title        string                `json:"title"`
	Content      string                `json:"content"`
	ThumbnailURL gqlscalar.URL         `json:"thumbnailUrl"`
	CreatedAt    gqlscalar.UTC         `json:"createdAt"`
	UpdatedAt    gqlscalar.UTC         `json:"updatedAt"`
	Tags         *ArticleTagConnection `json:"tags"`
}

func (ArticleNode) IsNode()            {}
func (this ArticleNode) GetID() string { return this.ID }

type ArticleTagConnection struct {
	Edges      []*ArticleTagEdge `json:"edges"`
	PageInfo   *PageInfo         `json:"pageInfo"`
	TotalCount int               `json:"totalCount"`
}

type ArticleTagEdge struct {
	Cursor string          `json:"cursor"`
	Node   *ArticleTagNode `json:"node"`
}

type ArticleTagNode struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type AttachTagsInput struct {
	ArticleID        string   `json:"articleId"`
	TagNames         []string `json:"tagNames"`
	ClientMutationID *string  `json:"clientMutationId,omitempty"`
}

type AttachTagsPayload struct {
	ArticleID        string  `json:"articleId"`
	EventID          string  `json:"eventID"`
	ClientMutationID *string `json:"clientMutationId,omitempty"`
}

type CreateArticleInput struct {
	Title            string        `json:"title"`
	Content          string        `json:"content"`
	ThumbnailURL     gqlscalar.URL `json:"thumbnailURL"`
	TagNames         []string      `json:"tagNames"`
	ClientMutationID *string       `json:"clientMutationId,omitempty"`
}

type CreateArticlePayload struct {
	ArticleID        string  `json:"articleId"`
	EventID          string  `json:"eventID"`
	ClientMutationID *string `json:"clientMutationId,omitempty"`
}

type DetachTagsInput struct {
	ArticleID        string   `json:"articleId"`
	TagNames         []string `json:"tagNames"`
	ClientMutationID *string  `json:"clientMutationId,omitempty"`
}

type DetachTagsPayload struct {
	ArticleID        string  `json:"articleId"`
	EventID          string  `json:"eventID"`
	ClientMutationID *string `json:"clientMutationId,omitempty"`
}

type Mutation struct {
}

type NoopInput struct {
	ClientMutationID *string `json:"clientMutationId,omitempty"`
}

type NoopPayload struct {
	ClientMutationID *string `json:"clientMutationId,omitempty"`
}

type PageInfo struct {
	HasNextPage     *bool  `json:"hasNextPage,omitempty"`
	HasPreviousPage *bool  `json:"hasPreviousPage,omitempty"`
	StartCursor     string `json:"startCursor"`
	EndCursor       string `json:"endCursor"`
}

type Query struct {
}

type TagArticleConnection struct {
	Edges      []*TagArticleEdge `json:"edges"`
	PageInfo   *PageInfo         `json:"pageInfo"`
	TotalCount int               `json:"totalCount"`
}

type TagArticleEdge struct {
	Cursor string          `json:"cursor"`
	Node   *TagArticleNode `json:"node"`
}

type TagArticleNode struct {
	ID           string        `json:"id"`
	Title        string        `json:"title"`
	ThumbnailURL gqlscalar.URL `json:"thumbnailUrl"`
	CreatedAt    gqlscalar.UTC `json:"createdAt"`
	UpdatedAt    gqlscalar.UTC `json:"updatedAt"`
}

func (TagArticleNode) IsNode()            {}
func (this TagArticleNode) GetID() string { return this.ID }

type TagConnection struct {
	Edges      []*TagEdge `json:"edges"`
	PageInfo   *PageInfo  `json:"pageInfo"`
	TotalCount int        `json:"totalCount"`
}

type TagEdge struct {
	Cursor string   `json:"cursor"`
	Node   *TagNode `json:"node"`
}

type TagNode struct {
	ID       string                `json:"id"`
	Name     string                `json:"name"`
	Articles *TagArticleConnection `json:"articles"`
}

type UpdateArticleBodyInput struct {
	ArticleID        string  `json:"articleId"`
	Content          string  `json:"content"`
	ClientMutationID *string `json:"clientMutationId,omitempty"`
}

type UpdateArticleBodyPayload struct {
	ArticleID        string  `json:"articleId"`
	EventID          string  `json:"eventID"`
	ClientMutationID *string `json:"clientMutationId,omitempty"`
}

type UpdateArticleThumbnailInput struct {
	ArticleID        string        `json:"articleId"`
	ThumbnailURL     gqlscalar.URL `json:"thumbnailURL"`
	ClientMutationID *string       `json:"clientMutationId,omitempty"`
}

type UpdateArticleThumbnailPayload struct {
	ArticleID        string  `json:"articleId"`
	EventID          string  `json:"eventID"`
	ClientMutationID *string `json:"clientMutationId,omitempty"`
}

type UpdateArticleTitleInput struct {
	ArticleID        string  `json:"articleId"`
	Title            string  `json:"title"`
	ClientMutationID *string `json:"clientMutationId,omitempty"`
}

type UpdateArticleTitlePayload struct {
	ArticleID        string  `json:"articleId"`
	EventID          string  `json:"eventID"`
	ClientMutationID *string `json:"clientMutationId,omitempty"`
}

type UploadImageInput struct {
	Image            graphql.Upload `json:"image"`
	ClientMutationID *string        `json:"clientMutationId,omitempty"`
}

type UploadImagePayload struct {
	ImageURL         gqlscalar.URL `json:"imageURL"`
	ClientMutationID *string       `json:"clientMutationId,omitempty"`
}
