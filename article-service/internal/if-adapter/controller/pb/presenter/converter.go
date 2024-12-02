//go:generate mockgen -source=converter.go -destination=../../../../mock/if-adapter/controller/pb/presenter/mock_converter.go -package=presenter
package presenter

import (
	"context"
	"github.com/miyamo2/blogapi.miyamo.today/article-service/internal/app/usecase/dto"

	"github.com/miyamo2/blogapi.miyamo.today/protogen/article/server/pb"
)

type ToGetNextConverter interface {
	ToGetNextArticlesResponse(ctx context.Context, from *dto.GetNextOutDto) (response *pb.GetNextArticlesResponse, ok bool)
}

type ToGetAllConverter interface {
	ToGetAllArticlesResponse(ctx context.Context, from *dto.GetAllOutDto) (response *pb.GetAllArticlesResponse, ok bool)
}

type ToGetByIdConverter interface {
	ToGetByIdArticlesResponse(ctx context.Context, from *dto.GetByIdOutDto) (response *pb.GetArticleByIdResponse, ok bool)
}

type ToGetPrevConverter interface {
	ToGetPrevArticlesResponse(ctx context.Context, from *dto.GetPrevOutDto) (response *pb.GetPrevArticlesResponse, ok bool)
}
