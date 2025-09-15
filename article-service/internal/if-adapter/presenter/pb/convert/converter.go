package convert

import (
	"context"

	"blogapi.miyamo.today/article-service/internal/app/usecase/dto"
	"blogapi.miyamo.today/article-service/internal/infra/grpc"
	"github.com/newrelic/go-agent/v3/newrelic"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ListAfter struct{}

func (c *ListAfter) ToResponse(
	ctx context.Context, from *dto.ListAfterOutput,
) (response *grpc.GetNextArticlesResponse, ok bool) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToListAfterArticlesResponse").End()

	articleDTOs := from.Articles()
	articlePBs := make([]*grpc.Article, 0, len(articleDTOs))
	for _, a := range articleDTOs {
		tagDTOs := a.Tags()
		tagPBs := make([]*grpc.Tag, 0, len(tagDTOs))
		for _, t := range tagDTOs {
			tagPBs = append(
				tagPBs, &grpc.Tag{
					Id:   t.ID(),
					Name: t.Name(),
				},
			)
		}
		articlePBs = append(
			articlePBs, &grpc.Article{
				Id:           a.ID(),
				Title:        a.Title(),
				Body:         a.Body(),
				ThumbnailUrl: a.ThumbnailUrl(),
				CreatedAt:    timestamppb.New(a.CreatedAt().StdTime()),
				UpdatedAt:    timestamppb.New(a.UpdatedAt().StdTime()),
				Tags:         tagPBs,
			},
		)
	}
	response = &grpc.GetNextArticlesResponse{
		Articles:    articlePBs,
		StillExists: from.HasNext(),
	}
	ok = true
	return
}

func NewListAfter() *ListAfter {
	return &ListAfter{}
}

type ListAll struct{}

func (c *ListAll) ToResponse(
	ctx context.Context, from *dto.ListAllOutput,
) (response *grpc.GetAllArticlesResponse, ok bool) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToGetAllArticlesResponse").End()

	articleDTOs := from.Articles()
	articlePBs := make([]*grpc.Article, 0, len(articleDTOs))
	for _, a := range articleDTOs {
		tagDTOs := a.Tags()
		tagPBs := make([]*grpc.Tag, 0, len(tagDTOs))
		for _, t := range tagDTOs {
			tagPBs = append(
				tagPBs, &grpc.Tag{
					Id:   t.ID(),
					Name: t.Name(),
				},
			)
		}
		articlePBs = append(
			articlePBs, &grpc.Article{
				Id:           a.ID(),
				Title:        a.Title(),
				Body:         a.Body(),
				ThumbnailUrl: a.ThumbnailUrl(),
				CreatedAt:    timestamppb.New(a.CreatedAt().StdTime()),
				UpdatedAt:    timestamppb.New(a.UpdatedAt().StdTime()),
				Tags:         tagPBs,
			},
		)
	}
	response = &grpc.GetAllArticlesResponse{
		Articles: articlePBs,
	}
	ok = true
	return
}

func NewListAll() *ListAll {
	return &ListAll{}
}

type GetByID struct{}

func (c *GetByID) ToResponse(
	ctx context.Context, from *dto.GetByIDOutput,
) (response *grpc.GetArticleByIdResponse, ok bool) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToGetByIdArticlesResponse").End()

	tagDTOs := from.Tags()
	tagPBs := make([]*grpc.Tag, 0, len(tagDTOs))
	for _, t := range tagDTOs {
		tagPBs = append(
			tagPBs, &grpc.Tag{
				Id:   t.ID(),
				Name: t.Name(),
			},
		)
	}
	articlePB := &grpc.Article{
		Id:           from.ID(),
		Title:        from.Title(),
		Body:         from.Body(),
		ThumbnailUrl: from.ThumbnailUrl(),
		CreatedAt:    timestamppb.New(from.CreatedAt().StdTime()),
		UpdatedAt:    timestamppb.New(from.UpdatedAt().StdTime()),
		Tags:         tagPBs,
	}
	response = &grpc.GetArticleByIdResponse{
		Article: articlePB,
	}
	ok = true
	return
}

func NewGetByID() *GetByID {
	return &GetByID{}
}

type ListBefore struct{}

func (c *ListBefore) ToResponse(
	ctx context.Context, from *dto.ListBeforeOutput,
) (response *grpc.GetPrevArticlesResponse, ok bool) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToListBeforeArticlesResponse").End()

	articleDTOs := from.Articles()
	articlePBs := make([]*grpc.Article, 0, len(articleDTOs))
	for _, a := range articleDTOs {
		tagDTOs := a.Tags()
		tagPBs := make([]*grpc.Tag, 0, len(tagDTOs))
		for _, t := range tagDTOs {
			tagPBs = append(
				tagPBs, &grpc.Tag{
					Id:   t.ID(),
					Name: t.Name(),
				},
			)
		}
		articlePBs = append(
			articlePBs, &grpc.Article{
				Id:           a.ID(),
				Title:        a.Title(),
				Body:         a.Body(),
				ThumbnailUrl: a.ThumbnailUrl(),
				CreatedAt:    timestamppb.New(a.CreatedAt().StdTime()),
				UpdatedAt:    timestamppb.New(a.UpdatedAt().StdTime()),
				Tags:         tagPBs,
			},
		)
	}
	response = &grpc.GetPrevArticlesResponse{
		Articles:    articlePBs,
		StillExists: from.HasPrevious(),
	}
	ok = true
	return
}

func NewListBefore() *ListBefore {
	return &ListBefore{}
}
