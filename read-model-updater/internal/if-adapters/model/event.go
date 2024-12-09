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
	EventID    string   `json:"event_id" dynamodbav:"event_id"`
	ArticleID  string   `json:"article_id" dynamodbav:"article_id"`
	Title      *string  `json:"title" dynamodbav:"title"`
	Content    *string  `json:"content" dynamodbav:"content"`
	Thumbnail  *string  `json:"thumbnail" dynamodbav:"thumbnail"`
	AttachTags []string `json:"attach_tags"  dynamodbav:"attach_tags"`
	DetachTags []string `json:"detach_tags"  dynamodbav:"detach_tags"`
	Invisible  *bool    `json:"invisible" dynamodbav:"invisible"`
}
