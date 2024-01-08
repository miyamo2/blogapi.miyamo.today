package pb

import (
	"github.com/miyamo2/blogapi-article-service/internal/app/usecase/dto"
	"github.com/miyamo2/blogproto-gen/article/server/pb"
)

type Converter struct{}

func (c Converter) ToGetNextArticlesResponse(from *dto.GetNextOutDto) (response *pb.GetNextArticlesResponse, ok bool) {
	fa := from.Articles()
	pa := make([]*pb.Article, 0, len(fa))
	for _, a := range fa {
		ft := a.Tags()
		pt := make([]*pb.Tag, 0, len(ft))
		for _, t := range ft {
			pt = append(pt, &pb.Tag{
				Id:   t.Id(),
				Name: t.Name(),
			})
		}
		pa = append(pa, &pb.Article{
			Id:           a.Id(),
			Title:        a.Title(),
			Body:         a.Body(),
			ThumbnailUrl: a.ThumbnailUrl(),
			CreatedAt:    a.CreatedAt(),
			UpdatedAt:    a.UpdatedAt(),
			Tags:         pt,
		})
	}
	response = &pb.GetNextArticlesResponse{
		Articles:    pa,
		StillExists: from.HasNext(),
	}
	ok = true
	return
}

func (c Converter) ToGetAllArticlesResponse(from *dto.GetAllOutDto) (response *pb.GetAllArticlesResponse, ok bool) {
	fa := from.Articles()
	pa := make([]*pb.Article, 0, len(fa))
	for _, a := range fa {
		ft := a.Tags()
		pt := make([]*pb.Tag, 0, len(ft))
		for _, t := range ft {
			pt = append(pt, &pb.Tag{
				Id:   t.Id(),
				Name: t.Name(),
			})
		}
		pa = append(pa, &pb.Article{
			Id:           a.Id(),
			Title:        a.Title(),
			Body:         a.Body(),
			ThumbnailUrl: a.ThumbnailUrl(),
			CreatedAt:    a.CreatedAt(),
			UpdatedAt:    a.UpdatedAt(),
			Tags:         pt,
		})
	}
	response = &pb.GetAllArticlesResponse{
		Articles: pa,
	}
	ok = true
	return
}

func (c Converter) ToGetByIdArticlesResponse(from *dto.GetByIdOutDto) (response *pb.GetArticleByIdResponse, ok bool) {
	ft := from.Tags()
	pt := make([]*pb.Tag, 0, len(ft))
	for _, t := range ft {
		pt = append(pt, &pb.Tag{
			Id:   t.Id(),
			Name: t.Name()})
	}
	a := &pb.Article{
		Id:           from.Id(),
		Title:        from.Title(),
		Body:         from.Body(),
		ThumbnailUrl: from.ThumbnailUrl(),
		CreatedAt:    from.CreatedAt(),
		UpdatedAt:    from.UpdatedAt(),
		Tags:         pt,
	}
	response = &pb.GetArticleByIdResponse{
		Article: a,
	}
	ok = true
	return
}

func (c Converter) ToGetPrevArticlesResponse(from *dto.GetPrevOutDto) (response *pb.GetPrevArticlesResponse, ok bool) {
	fa := from.Articles()
	pa := make([]*pb.Article, 0, len(fa))
	for _, a := range fa {
		ft := a.Tags()
		pt := make([]*pb.Tag, 0, len(ft))
		for _, t := range ft {
			pt = append(pt, &pb.Tag{
				Id:   t.Id(),
				Name: t.Name(),
			})
		}
		pa = append(pa, &pb.Article{
			Id:           a.Id(),
			Title:        a.Title(),
			Body:         a.Body(),
			ThumbnailUrl: a.ThumbnailUrl(),
			CreatedAt:    a.CreatedAt(),
			UpdatedAt:    a.UpdatedAt(),
			Tags:         pt,
		})
	}
	response = &pb.GetPrevArticlesResponse{
		Articles:    pa,
		StillExists: from.HasPrevious(),
	}
	ok = true
	return
}

func NewConverter() *Converter {
	return &Converter{}
}
