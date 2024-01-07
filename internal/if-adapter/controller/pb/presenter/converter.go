//go:generate mockgen -source=converter.go -destination=../../../../mock/if-adapter/controller/pb/presenter/mock_converter.go -package=presenter
package presenter

import (
	"github.com/miyamo2/blogapi-article-service/internal/if-adapter/controller/pb/usecase"
	"github.com/miyamo2/blogproto-gen/article/server/pb"
)

type ToGetNextConverter[T usecase.Tag, A usecase.Article[T], O usecase.GetNextOutDto[T, A]] interface {
	ToGetNextArticlesResponse(from O) (response *pb.GetNextArticlesResponse, ok bool)
}

type ToGetAllConverter[T usecase.Tag, A usecase.Article[T], O usecase.GetAllOutDto[T, A]] interface {
	ToGetAllArticlesResponse(from O) (response *pb.GetAllArticlesResponse, ok bool)
}

type ToGetByIdConverter[T usecase.Tag, A usecase.Article[T]] interface {
	ToGetByIdArticlesResponse(from A) (response *pb.GetArticleByIdResponse, ok bool)
}
