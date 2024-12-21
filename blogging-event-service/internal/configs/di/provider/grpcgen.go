package provider

import (
	"github.com/google/wire"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/if-adapter/controller/pb"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/if-adapter/controller/pb/presenters"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/if-adapter/controller/pb/usecase"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/infra/grpc"
)

// compatibility check
var _ grpc.BloggingEventServiceServer = (*pb.BloggingEventServiceServer)(nil)

func NewBloggingEventServiceServer(
	createArticleUsecase usecase.CreateArticle,
	createArticleConverter presenters.ToCreateArticleResponse,
	updateArticleTitleUsecase usecase.UpdateArticleTitle,
	updateArticleTitleConverter presenters.ToUpdateArticleTitleResponse,
	updateArticleBodyUsecase usecase.UpdateArticleBody,
	updateArticleBodyConverter presenters.ToUpdateArticleBodyResponse,
	updateArticleThumbnailUsecase usecase.UpdateArticleThumbnail,
	updateArticleThumbnailConverter presenters.ToUpdateArticleThumbnailResponse,
	attachTagsUsecase usecase.AttachTags,
	attachTagsConverter presenters.ToAttachTagsResponse,
	detachTagsUsecase usecase.DetachTags,
	detachTagsConverter presenters.ToDetachTagsResponse,
) *pb.BloggingEventServiceServer {
	return pb.NewBloggingEventServiceServer(
		pb.WithCreateArticleUsecase(createArticleUsecase),
		pb.WithCreateArticleConverter(createArticleConverter),
		pb.WithUpdateArticleTitleUsecase(updateArticleTitleUsecase),
		pb.WithUpdateArticleTitleConverter(updateArticleTitleConverter),
		pb.WithUpdateArticleBodyUsecase(updateArticleBodyUsecase),
		pb.WithUpdateArticleBodyConverter(updateArticleBodyConverter),
		pb.WithUpdateArticleThumbnailUsecase(updateArticleThumbnailUsecase),
		pb.WithUpdateArticleThumbnailConverter(updateArticleThumbnailConverter),
		pb.WithAttachTagsUsecase(attachTagsUsecase),
		pb.WithAttachTagsConverter(attachTagsConverter),
		pb.WithDetachTagsUsecase(detachTagsUsecase),
		pb.WithDetachTagsConverter(detachTagsConverter))
}

var BloggingEventServiceServerSet = wire.NewSet(
	NewBloggingEventServiceServer,
	wire.Bind(new(grpc.BloggingEventServiceServer), new(*pb.BloggingEventServiceServer)),
)
