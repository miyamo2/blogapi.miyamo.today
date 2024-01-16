package pb

import (
	"github.com/miyamo2/blogapi-tag-service/internal/app/usecase/dto"
	"github.com/miyamo2/blogproto-gen/tag/server/pb"
)

// Converter is am implementation of presenter.ToGetByIdConverter
type Converter struct{}

// ToGetByIdTagResponse is an implementation of presenter.ToGetByIdConverter#ToGetByIdTagResponse
func (c Converter) ToGetByIdTagResponse(from *dto.GetByIdOutDto) (response *pb.GetTagByIdResponse, ok bool) {
	fa := from.Articles()
	pa := make([]*pb.Article, 0, len(fa))
	for _, a := range fa {
		pa = append(pa, &pb.Article{
			Id:           a.Id(),
			Title:        a.Title(),
			ThumbnailUrl: a.ThumbnailUrl(),
			CreatedAt:    a.CreatedAt(),
			UpdatedAt:    a.UpdatedAt(),
		})
	}
	t := &pb.Tag{
		Id:       from.Id(),
		Name:     from.Name(),
		Articles: pa,
	}
	response = &pb.GetTagByIdResponse{
		Tag: t,
	}
	ok = true
	return
}

// ToGetAllTagsResponse is an implementation of presenter.ToGetAllConverter#ToGetAllTagsResponse
func (c Converter) ToGetAllTagsResponse(from *dto.GetAllOutDto) (response *pb.GetAllTagsResponse, ok bool) {
	ft := from.Tags()
	pt := make([]*pb.Tag, 0, len(ft))
	for _, t := range ft {
		fa := t.Articles()
		pa := make([]*pb.Article, 0, len(fa))
		for _, a := range fa {
			pa = append(pa, &pb.Article{
				Id:           a.Id(),
				Title:        a.Title(),
				ThumbnailUrl: a.ThumbnailUrl(),
				CreatedAt:    a.CreatedAt(),
				UpdatedAt:    a.UpdatedAt(),
			})
		}
		pt = append(pt, &pb.Tag{
			Id:       t.Id(),
			Name:     t.Name(),
			Articles: pa,
		})
	}
	response = &pb.GetAllTagsResponse{
		Tags: pt,
	}
	ok = true
	return
}

// ToGetNextTagsResponse is an implementation of presenter.ToGetNextConverter#ToGetNextTagsResponse
func (c Converter) ToGetNextTagsResponse(from *dto.GetNextOutDto) (response *pb.GetNextTagResponse, ok bool) {
	ft := from.Tags()
	pt := make([]*pb.Tag, 0, len(ft))
	for _, t := range ft {
		fa := t.Articles()
		pa := make([]*pb.Article, 0, len(fa))
		for _, a := range fa {
			pa = append(pa, &pb.Article{
				Id:           a.Id(),
				Title:        a.Title(),
				ThumbnailUrl: a.ThumbnailUrl(),
				CreatedAt:    a.CreatedAt(),
				UpdatedAt:    a.UpdatedAt(),
			})
		}
		pt = append(pt, &pb.Tag{
			Id:       t.Id(),
			Name:     t.Name(),
			Articles: pa,
		})
	}
	response = &pb.GetNextTagResponse{
		Tags:        pt,
		StillExists: from.HasNext(),
	}
	ok = true
	return
}

// ToGetPrevTagsResponse is an implementation of presenter.ToGetPrevConverter#ToGetPrevTagsResponse
func (c Converter) ToGetPrevTagsResponse(from *dto.GetPrevOutDto) (response *pb.GetPrevTagResponse, ok bool) {
	ft := from.Tags()
	pt := make([]*pb.Tag, 0, len(ft))
	for _, t := range ft {
		fa := t.Articles()
		pa := make([]*pb.Article, 0, len(fa))
		for _, a := range fa {
			pa = append(pa, &pb.Article{
				Id:           a.Id(),
				Title:        a.Title(),
				ThumbnailUrl: a.ThumbnailUrl(),
				CreatedAt:    a.CreatedAt(),
				UpdatedAt:    a.UpdatedAt(),
			})
		}
		pt = append(pt, &pb.Tag{
			Id:       t.Id(),
			Name:     t.Name(),
			Articles: pa,
		})
	}
	response = &pb.GetPrevTagResponse{
		Tags:        pt,
		StillExists: from.HasPrevious(),
	}
	ok = true
	return
}

func NewConverter() *Converter {
	return &Converter{}
}
