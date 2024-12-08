package model

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// EventStream is a stream of the Blogging event.
type EventStream struct {
	Records []Record `json:"Records"`
}

// Record is a record of the Blogging event.
type Record struct {
	EventID        string   `json:"eventID"`
	EventName      string   `json:"eventName"`
	EventSource    string   `json:"eventSource"`
	EventVersion   string   `json:"eventVersion"`
	AWSRegion      string   `json:"awsRegion"`
	DynamoDB       DynamoDB `json:"dynamodb"`
	EventSourceAtn string   `json:"eventSourceARN"`
}

type DynamoDB struct {
	NewImage map[string]types.AttributeValue `json:"NewImage"`
}

type Image struct {
	EventId     string   `json:"EventId" dynamodbav:"event_id"`
	ArticleId   string   `json:"ArticleId" dynamodbav:"article_id"`
	Title       *string  `json:"Title" dynamodbav:"title"`
	Content     *string  `json:"Content" dynamodbav:"content"`
	Thumbnail   *string  `json:"Thumbnail" dynamodbav:"thumbnail"`
	AttacheTags []string `json:"AttacheTags"  dynamodbav:"attache_tags"`
	DetachTags  []string `json:"DetachTags"  dynamodbav:"detach_tags"`
	Invisible   *bool    `json:"Invisible" dynamodbav:"invisible"`
}
