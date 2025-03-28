//go:generate mockgen -source=converter.go -destination=../../../../mock/if-adapter/controller/pb/presenter/mock_converter.go -package=presenter
package presenter

import (
	"blogapi.miyamo.today/tag-service/internal/app/usecase/dto"
	"connectrpc.com/connect"
	"context"

	"blogapi.miyamo.today/tag-service/internal/infra/grpc"
)

// ToGetByIdConverter is a converter interface for converting from GetById use-case's dto to pb response.
type ToGetByIdConverter interface {
	// ToGetByIdTagResponse converts from GetById use-case's dto to pb response.
	ToGetByIdTagResponse(ctx context.Context, from *dto.GetByIdOutDto) (response *connect.Response[grpc.GetTagByIdResponse], ok bool)
}

// ToGetAllConverter is a converter interface for converting from GetAll use-case's dto to pb response.
type ToGetAllConverter interface {
	// ToGetAllTagsResponse converts from GetAll use-case's dto to pb response.
	ToGetAllTagsResponse(ctx context.Context, from *dto.GetAllOutDto) (response *connect.Response[grpc.GetAllTagsResponse], ok bool)
}

// ToGetNextConverter is a converter interface for converting from GetNext use-case's dto to pb response.
type ToGetNextConverter interface {
	// ToGetNextTagsResponse converts from GetNext use-case's dto to pb response.
	ToGetNextTagsResponse(ctx context.Context, from *dto.GetNextOutDto) (response *connect.Response[grpc.GetNextTagResponse], ok bool)
}

// ToGetPrevConverter is a converter interface for converting from GetPrev use-case's dto to pb response.
type ToGetPrevConverter interface {
	// ToGetPrevTagsResponse converts from GetPrev use-case's dto to pb response.
	ToGetPrevTagsResponse(ctx context.Context, from *dto.GetPrevOutDto) (response *connect.Response[grpc.GetPrevTagResponse], ok bool)
}
