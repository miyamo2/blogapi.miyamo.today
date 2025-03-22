package provider

import (
	"blogapi.miyamo.today/blogging-event-service/internal/if-adapter/controller/pb"
	"blogapi.miyamo.today/blogging-event-service/internal/if-adapter/controller/pb/presenters"
	"blogapi.miyamo.today/blogging-event-service/internal/if-adapter/controller/pb/usecase"
	_ "blogapi.miyamo.today/blogging-event-service/internal/infra/grpc"
	"blogapi.miyamo.today/blogging-event-service/internal/infra/grpc/grpcconnect"
	"github.com/google/wire"
)

// compatibility check
var _ grpcconnect.BloggingEventServiceHandler = (*pb.BloggingEventServiceServer)(nil)

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
	uploadImageUsecase usecase.UploadImage,
	uploadImageConverter presenters.ToUploadImageResponse,
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
		pb.WithDetachTagsConverter(detachTagsConverter),
		pb.WithUploadImageUsecase(uploadImageUsecase),
		pb.WithUploadImageConverter(uploadImageConverter))
}

var BloggingEventServiceServerSet = wire.NewSet(
	NewBloggingEventServiceServer,
	wire.Bind(new(grpcconnect.BloggingEventServiceHandler), new(*pb.BloggingEventServiceServer)),
)
