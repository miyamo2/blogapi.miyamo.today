//go:generate mockgen -source=converter.go -destination=../../../../mock/if-adapter/controller/pb/presenter/mock_converter.go -package=presenter
package presenter

import (
	"context"

	"github.com/miyamo2/api.miyamo.today/protogen/tag/server/pb"
	"github.com/miyamo2/api.miyamo.today/tag-service/internal/if-adapter/controller/pb/usecase"
)

// ToGetByIdConverter is a converter interface for converting from GetById use-case's dto to pb response.
type ToGetByIdConverter[A usecase.Article, T usecase.Tag[A]] interface {
	// ToGetByIdTagResponse converts from GetById use-case's dto to pb response.
	ToGetByIdTagResponse(ctx context.Context, from T) (response *pb.GetTagByIdResponse, ok bool)
}

// ToGetAllConverter is a converter interface for converting from GetAll use-case's dto to pb response.
type ToGetAllConverter[A usecase.Article, T usecase.Tag[A], O usecase.GetAllOutDto[A, T]] interface {
	// ToGetAllTagsResponse converts from GetAll use-case's dto to pb response.
	ToGetAllTagsResponse(ctx context.Context, from O) (response *pb.GetAllTagsResponse, ok bool)
}

// ToGetNextConverter is a converter interface for converting from GetNext use-case's dto to pb response.
type ToGetNextConverter[A usecase.Article, T usecase.Tag[A], O usecase.GetNextOutDto[A, T]] interface {
	// ToGetNextTagsResponse converts from GetNext use-case's dto to pb response.
	ToGetNextTagsResponse(ctx context.Context, from O) (response *pb.GetNextTagResponse, ok bool)
}

// ToGetPrevConverter is a converter interface for converting from GetPrev use-case's dto to pb response.
type ToGetPrevConverter[A usecase.Article, T usecase.Tag[A], O usecase.GetPrevOutDto[A, T]] interface {
	// ToGetPrevTagsResponse converts from GetPrev use-case's dto to pb response.
	ToGetPrevTagsResponse(ctx context.Context, from O) (response *pb.GetPrevTagResponse, ok bool)
}
