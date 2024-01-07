package pb

import (
	"context"
	"fmt"
	"github.com/cockroachdb/errors"
	"github.com/miyamo2/blogapi-article-service/internal/app/usecase/dto"
	"github.com/miyamo2/blogapi-article-service/internal/if-adapter/controller/pb/presenter"
	"github.com/miyamo2/blogapi-article-service/internal/if-adapter/controller/pb/usecase"
	"github.com/miyamo2/blogapi-core/util/duration"
	"github.com/miyamo2/blogproto-gen/article/server/pb"
	"google.golang.org/protobuf/types/known/emptypb"
	"log/slog"
)

// ArticleServiceServer is implementation of pb.ArticleServiceServer
type ArticleServiceServer struct {
	pb.UnimplementedArticleServiceServer
	getByIdUsecase usecase.GetById[dto.GetByIdInDto, dto.Tag, *dto.GetByIdOutDto]
	getAllUsecase  usecase.GetAll[dto.Tag, dto.Article, *dto.GetAllOutDto]
	getNextUsecase usecase.GetNext[dto.GetNextInDto, dto.Tag, dto.Article, *dto.GetNextOutDto]
	getNextConv    presenter.ToGetNextConverter[dto.Tag, dto.Article, *dto.GetNextOutDto]
	getAllConv     presenter.ToGetAllConverter[dto.Tag, dto.Article, *dto.GetAllOutDto]
	getByIdConv    presenter.ToGetByIdConverter[dto.Tag, *dto.GetByIdOutDto]
}

var (
	ErrConversionToGetNextArticlesFailed = errors.New("conversion to get_next_articles_response failed")
	ErrConversionToGetAllArticlesFailed  = errors.New("conversion to get_all_articles_response failed")
	ErrConversionToGetArticleByIdFailed  = errors.New("conversion to get_article_by_id_response failed")
)

// GetAllArticles is implementation of pb.ArticleServiceServer.GetAllArticles
func (s *ArticleServiceServer) GetAllArticles(ctx context.Context, in *emptypb.Empty) (*pb.GetAllArticlesResponse, error) {
	dw := duration.Start()
	slog.InfoContext(ctx, "BEGIN",
		slog.Group("parameters",
			slog.String("in", in.String())))
	oDto, err := s.getAllUsecase.Execute(ctx)
	if err != nil {
		err = errors.WithStack(err)
		slog.InfoContext(ctx, "END",
			slog.String("duration", dw.SDuration()),
			slog.Group("return",
				slog.Any("*pb.GetAllArticlesResponse", nil),
				slog.String("error", fmt.Sprintf("%+v", err))))
		return nil, err
	}
	res, ok := s.getAllConv.ToGetAllArticlesResponse(oDto)
	if !ok {
		return nil, ErrConversionToGetAllArticlesFailed
	}
	slog.InfoContext(ctx, "END",
		slog.String("duration", dw.SDuration()),
		slog.Group("return",
			slog.String("*pb.GetAllArticlesResponse", res.String()),
			slog.Any("error", nil)))
	return res, nil
}

// GetNextArticles is implementation of pb.ArticleServiceServer.GetNextArticles
func (s *ArticleServiceServer) GetNextArticles(ctx context.Context, in *pb.GetNextArticlesRequest) (*pb.GetNextArticlesResponse, error) {
	dw := duration.Start()
	slog.InfoContext(ctx, "BEGIN",
		slog.Group("parameters",
			slog.String("in", in.String())))
	oDto, err := s.getNextUsecase.Execute(ctx, dto.NewGetNextInDto(int(in.First), in.After))
	if err != nil {
		err = errors.WithStack(err)
		slog.InfoContext(ctx, "END",
			slog.String("duration", dw.SDuration()),
			slog.Group("return",
				slog.Any("*pb.GetNextArticlesResponse", nil),
				slog.String("error", fmt.Sprintf("%+v", err))))
		return nil, err
	}
	res, ok := s.getNextConv.ToGetNextArticlesResponse(oDto)
	if !ok {
		return nil, ErrConversionToGetNextArticlesFailed
	}
	slog.InfoContext(ctx, "END",
		slog.String("duration", dw.SDuration()),
		slog.Group("return",
			slog.String("*pb.GetNextArticlesResponse", res.String()),
			slog.Any("error", nil)))
	return res, nil
}

// GetArticleById is implementation of pb.ArticleServiceServer.GetArticleById
func (s *ArticleServiceServer) GetArticleById(ctx context.Context, in *pb.GetArticleByIdRequest) (*pb.GetArticleByIdResponse, error) {
	dw := duration.Start()
	slog.InfoContext(ctx, "BEGIN",
		slog.Group("parameters",
			slog.String("in", in.String())))
	oDto, err := s.getByIdUsecase.Execute(ctx, dto.NewGetByIdInDto(in.GetId()))
	if err != nil {
		err = errors.WithStack(err)
		slog.InfoContext(ctx, "END",
			slog.String("duration", dw.SDuration()),
			slog.Group("return",
				slog.Any("*pb.GetArticleByIdResponse", nil),
				slog.String("error", fmt.Sprintf("%+v", err))))
		return nil, err
	}
	dTags := oDto.Tags()
	pTags := make([]*pb.Tag, 0, len(dTags))
	for _, t := range dTags {
		pTags = append(pTags, &pb.Tag{
			Id:   t.Id(),
			Name: t.Name(),
		})
	}
	res, ok := s.getByIdConv.ToGetByIdArticlesResponse(oDto)
	if !ok {
		return nil, ErrConversionToGetArticleByIdFailed
	}
	slog.InfoContext(ctx, "END",
		slog.String("duration", dw.SDuration()),
		slog.Group("return",
			slog.String("*pb.GetArticleByIdResponse", res.String()),
			slog.Any("error", nil)))
	return res, nil
}

// NewArticleServiceServer is constructor of ArticleServiceServer
func NewArticleServiceServer(
	getByIdUsecase usecase.GetById[dto.GetByIdInDto, dto.Tag, *dto.GetByIdOutDto],
	getAllUsecase usecase.GetAll[dto.Tag, dto.Article, *dto.GetAllOutDto],
	getNextUsecase usecase.GetNext[dto.GetNextInDto, dto.Tag, dto.Article, *dto.GetNextOutDto],
	getNextConv presenter.ToGetNextConverter[dto.Tag, dto.Article, *dto.GetNextOutDto],
	getAllConv presenter.ToGetAllConverter[dto.Tag, dto.Article, *dto.GetAllOutDto],
	getByIdConv presenter.ToGetByIdConverter[dto.Tag, *dto.GetByIdOutDto],
) *ArticleServiceServer {
	return &ArticleServiceServer{
		getByIdUsecase: getByIdUsecase,
		getAllUsecase:  getAllUsecase,
		getNextUsecase: getNextUsecase,
		getNextConv:    getNextConv,
		getAllConv:     getAllConv,
		getByIdConv:    getByIdConv,
	}
}
