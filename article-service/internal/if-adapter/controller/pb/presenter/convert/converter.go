package convert

import (
	"context"

	"blogapi.miyamo.today/article-service/internal/app/usecase/dto"

	"blogapi.miyamo.today/article-service/internal/infra/grpc"
)

type GetByID interface {
	ToResponse(ctx context.Context, from *dto.GetByIDOutput) (
		response *grpc.GetArticleByIdResponse, ok bool,
	)
}

type ListAll interface {
	ToResponse(ctx context.Context, from *dto.ListAllOutput) (
		response *grpc.GetAllArticlesResponse, ok bool,
	)
}

type ListAfter interface {
	ToResponse(ctx context.Context, from *dto.ListAfterOutput) (
		response *grpc.GetNextArticlesResponse, ok bool,
	)
}

type ListBefore interface {
	ToResponse(ctx context.Context, from *dto.ListBeforeOutput) (
		response *grpc.GetPrevArticlesResponse, ok bool,
	)
}
