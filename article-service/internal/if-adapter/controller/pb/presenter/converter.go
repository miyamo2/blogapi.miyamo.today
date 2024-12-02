//go:generate mockgen -source=converter.go -destination=../../../../mock/if-adapter/controller/pb/presenter/mock_converter.go -package=presenter
package presenter

import (
	"context"
	"github.com/miyamo2/blogapi.miyamo.today/article-service/internal/app/usecase/dto"

	"github.com/miyamo2/blogapi.miyamo.today/article-service/internal/infra/grpc"
)

type ToGetNextConverter interface {
	ToGetNextArticlesResponse(ctx context.Context, from *dto.GetNextOutDto) (response *grpc.GetNextArticlesResponse, ok bool)
}

type ToGetAllConverter interface {
	ToGetAllArticlesResponse(ctx context.Context, from *dto.GetAllOutDto) (response *grpc.GetAllArticlesResponse, ok bool)
}

type ToGetByIdConverter interface {
	ToGetByIdArticlesResponse(ctx context.Context, from *dto.GetByIdOutDto) (response *grpc.GetArticleByIdResponse, ok bool)
}

type ToGetPrevConverter interface {
	ToGetPrevArticlesResponse(ctx context.Context, from *dto.GetPrevOutDto) (response *grpc.GetPrevArticlesResponse, ok bool)
}
