package convert

import (
	"context"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/types/known/timestamppb"

	"blogapi.miyamo.today/tag-service/internal/app/usecase/dto"
	"blogapi.miyamo.today/tag-service/internal/infra/grpc"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type GetByIdTag struct{}

func (c *GetByIdTag) ToResponse(
	ctx context.Context, from *dto.GetByIdOutput,
) (response *connect.Response[grpc.GetTagByIdResponse], ok bool) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToResponse").End()

	articleDTOs := from.Articles()
	articlePBs := make([]*grpc.Article, 0, len(articleDTOs))
	for _, a := range articleDTOs {
		articlePBs = append(
			articlePBs, &grpc.Article{
				Id:           a.Id(),
				Title:        a.Title(),
				ThumbnailUrl: a.ThumbnailUrl(),
				CreatedAt:    timestamppb.New(a.CreatedAt().StdTime()),
				UpdatedAt:    timestamppb.New(a.UpdatedAt().StdTime()),
			},
		)
	}
	tagPB := &grpc.Tag{
		Id:       from.Id(),
		Name:     from.Name(),
		Articles: articlePBs,
	}
	rawResponse := &grpc.GetTagByIdResponse{
		Tag: tagPB,
	}
	response = connect.NewResponse(rawResponse)
	ok = true
	return
}

func NewGetByIdTag() *GetByIdTag {
	return &GetByIdTag{}
}

type GetAllTags struct{}

func (c *GetAllTags) ToResponse(
	ctx context.Context, from *dto.ListAllOutput,
) (response *connect.Response[grpc.GetAllTagsResponse], ok bool) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToResponse").End()

	tagDTOs := from.Tags()
	tagPBs := make([]*grpc.Tag, 0, len(tagDTOs))
	for _, t := range tagDTOs {
		articleDTOs := t.Articles()
		articlePBs := make([]*grpc.Article, 0, len(articleDTOs))
		for _, a := range articleDTOs {
			articlePBs = append(
				articlePBs, &grpc.Article{
					Id:           a.Id(),
					Title:        a.Title(),
					ThumbnailUrl: a.ThumbnailUrl(),
					CreatedAt:    timestamppb.New(a.CreatedAt().StdTime()),
					UpdatedAt:    timestamppb.New(a.UpdatedAt().StdTime()),
				},
			)
		}
		tagPBs = append(
			tagPBs, &grpc.Tag{
				Id:       t.Id(),
				Name:     t.Name(),
				Articles: articlePBs,
			},
		)
	}
	rawResponse := &grpc.GetAllTagsResponse{
		Tags: tagPBs,
	}
	response = connect.NewResponse(rawResponse)
	ok = true
	return
}

func NewGetAllTags() *GetAllTags {
	return &GetAllTags{}
}

type GetNextTags struct{}

func (c *GetNextTags) ToResponse(
	ctx context.Context, from *dto.ListAfterOutput,
) (response *connect.Response[grpc.GetNextTagResponse], ok bool) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToResponse").End()

	tagDTOs := from.Tags()
	tagPBs := make([]*grpc.Tag, 0, len(tagDTOs))
	for _, t := range tagDTOs {
		articleDTOs := t.Articles()
		articlePBs := make([]*grpc.Article, 0, len(articleDTOs))
		for _, a := range articleDTOs {
			articlePBs = append(
				articlePBs, &grpc.Article{
					Id:           a.Id(),
					Title:        a.Title(),
					ThumbnailUrl: a.ThumbnailUrl(),
					CreatedAt:    timestamppb.New(a.CreatedAt().StdTime()),
					UpdatedAt:    timestamppb.New(a.UpdatedAt().StdTime()),
				},
			)
		}
		tagPBs = append(
			tagPBs, &grpc.Tag{
				Id:       t.Id(),
				Name:     t.Name(),
				Articles: articlePBs,
			},
		)
	}
	rawResponse := &grpc.GetNextTagResponse{
		Tags:        tagPBs,
		StillExists: from.HasNext(),
	}
	response = connect.NewResponse(rawResponse)
	ok = true
	return
}

func NewGetNextTags() *GetNextTags {
	return &GetNextTags{}
}

type GetPrevTags struct{}

func (c *GetPrevTags) ToResponse(
	ctx context.Context, from *dto.ListBeforeOutput,
) (response *connect.Response[grpc.GetPrevTagResponse], ok bool) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToResponse").End()

	tagDTOs := from.Tags()
	tagPBs := make([]*grpc.Tag, 0, len(tagDTOs))
	for _, t := range tagDTOs {
		articleDTOs := t.Articles()
		articlePBs := make([]*grpc.Article, 0, len(articleDTOs))
		for _, a := range articleDTOs {
			articlePBs = append(
				articlePBs, &grpc.Article{
					Id:           a.Id(),
					Title:        a.Title(),
					ThumbnailUrl: a.ThumbnailUrl(),
					CreatedAt:    timestamppb.New(a.CreatedAt().StdTime()),
					UpdatedAt:    timestamppb.New(a.UpdatedAt().StdTime()),
				},
			)
		}
		tagPBs = append(
			tagPBs, &grpc.Tag{
				Id:       t.Id(),
				Name:     t.Name(),
				Articles: articlePBs,
			},
		)
	}
	rawResponse := &grpc.GetPrevTagResponse{
		Tags:        tagPBs,
		StillExists: from.HasPrevious(),
	}
	response = connect.NewResponse(rawResponse)
	ok = true
	return
}

func NewGetPrevTags() *GetPrevTags {
	return &GetPrevTags{}
}
