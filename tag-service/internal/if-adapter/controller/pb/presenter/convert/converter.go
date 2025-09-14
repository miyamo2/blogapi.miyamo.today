package convert

import (
	"context"

	"blogapi.miyamo.today/tag-service/internal/app/usecase/dto"
	"connectrpc.com/connect"

	"blogapi.miyamo.today/tag-service/internal/infra/grpc"
)

// ToGetById provides conversion from GetById use-case's dto to connect.Response.
type ToGetById interface {
	// ToResponse converts from GetById use-case's dto to connect.Response.
	ToResponse(ctx context.Context, from *dto.GetByIdOutput) (
		response *connect.Response[grpc.GetTagByIdResponse], ok bool,
	)
}

// ToGetAll provides conversion from GetAll use-case's dto to connect.Response.
type ToGetAll interface {
	// ToResponse converts from GetAll use-case's dto to connect.Response.
	ToResponse(ctx context.Context, from *dto.ListAllOutput) (
		response *connect.Response[grpc.GetAllTagsResponse], ok bool,
	)
}

// ToGetNext provides conversion from GetNext use-case's dto to connect.Response.
type ToGetNext interface {
	// ToResponse converts from GetNext use-case's dto to connect.Response.
	ToResponse(ctx context.Context, from *dto.ListAfterOutput) (
		response *connect.Response[grpc.GetNextTagResponse], ok bool,
	)
}

// ToGetPrev provides conversion from GetPrev use-case's dto to connect.Response.
type ToGetPrev interface {
	// ToResponse converts from GetPrev use-case's dto to connect.Response.
	ToResponse(ctx context.Context, from *dto.ListBeforeOutput) (
		response *connect.Response[grpc.GetPrevTagResponse], ok bool,
	)
}
