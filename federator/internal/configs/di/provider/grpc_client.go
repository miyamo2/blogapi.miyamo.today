package provider

import (
	"blogapi.miyamo.today/federator/internal/infra/grpc/article/articleconnect"
	"blogapi.miyamo.today/federator/internal/infra/grpc/blogging_event/blogging_eventconnect"
	"blogapi.miyamo.today/federator/internal/infra/grpc/tag/tagconnect"
	"connectrpc.com/connect"
	"fmt"
	"github.com/google/wire"
	"net/http"
	"os"
)

func ArticleClient(httpClient *http.Client) articleconnect.ArticleServiceClient {
	address := os.Getenv("ARTICLE_SERVICE_ADDRESS")
	return articleconnect.NewArticleServiceClient(httpClient, fmt.Sprintf("http://%s", address), connect.WithGRPC())
}

func TagClient(httpClient *http.Client) tagconnect.TagServiceClient {
	address := os.Getenv("TAG_SERVICE_ADDRESS")
	return tagconnect.NewTagServiceClient(httpClient, fmt.Sprintf("http://%s", address), connect.WithGRPC())
}

func BloggingEventClient(httpClient *http.Client) blogging_eventconnect.BloggingEventServiceClient {
	address := os.Getenv("BLOGGING_EVENT_SERVICE_ADDRESS")
	return blogging_eventconnect.NewBloggingEventServiceClient(httpClient, fmt.Sprintf("http://%s", address), connect.WithGRPC())
}

var GRPCClientSet = wire.NewSet(ArticleClient, TagClient, BloggingEventClient)
