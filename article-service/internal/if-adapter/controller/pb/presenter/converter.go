//go:generate mockgen -source=converter.go -destination=../../../../mock/if-adapter/controller/pb/presenter/mock_converter.go -package=presenter
package presenter

import (
	"context"

	"github.com/miyamo2/blogapi.miyamo.today/article-service/internal/if-adapter/controller/pb/usecase"
	"github.com/miyamo2/blogapi.miyamo.today/protogen/article/server/pb"
)

type ToGetNextConverter[T usecase.Tag, A usecase.Article[T], O usecase.GetNextOutDto[T, A]] interface {
	ToGetNextArticlesResponse(ctx context.Context, from O) (response *pb.GetNextArticlesResponse, ok bool)
}

type ToGetAllConverter[T usecase.Tag, A usecase.Article[T], O usecase.GetAllOutDto[T, A]] interface {
	ToGetAllArticlesResponse(ctx context.Context, from O) (response *pb.GetAllArticlesResponse, ok bool)
}

type ToGetByIdConverter[T usecase.Tag, A usecase.Article[T]] interface {
	ToGetByIdArticlesResponse(ctx context.Context, from A) (response *pb.GetArticleByIdResponse, ok bool)
}

type ToGetPrevConverter[T usecase.Tag, A usecase.Article[T], O usecase.GetPrevOutDto[T, A]] interface {
	ToGetPrevArticlesResponse(ctx context.Context, from O) (response *pb.GetPrevArticlesResponse, ok bool)
}