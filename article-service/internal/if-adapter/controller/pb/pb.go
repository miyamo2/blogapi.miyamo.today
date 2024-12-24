package pb

import (
	"context"
	"log/slog"

	"github.com/cockroachdb/errors"
	"github.com/miyamo2/altnrslog"
	"github.com/miyamo2/blogapi.miyamo.today/article-service/internal/app/usecase/dto"
	"github.com/miyamo2/blogapi.miyamo.today/article-service/internal/if-adapter/controller/pb/presenter"
	"github.com/miyamo2/blogapi.miyamo.today/article-service/internal/if-adapter/controller/pb/usecase"
	"github.com/miyamo2/blogapi.miyamo.today/article-service/internal/infra/grpc"
	"github.com/miyamo2/blogapi.miyamo.today/core/log"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"
	"google.golang.org/protobuf/types/known/emptypb"
)

// ArticleServiceServer is implementation of grpc.ArticleServiceServer
type ArticleServiceServer struct {
	grpc.UnimplementedArticleServiceServer
	getByIdUsecase usecase.GetById
	getAllUsecase  usecase.GetAll
	getNextUsecase usecase.GetNext
	getPrevUsecase usecase.GetPrev
	getNextConv    presenter.ToGetNextConverter
	getAllConv     presenter.ToGetAllConverter
	getByIdConv    presenter.ToGetByIdConverter
	getPrevConv    presenter.ToGetPrevConverter
}

var (
	ErrConversionToGetNextArticlesFailed = errors.New("conversion to get_next_articles_response failed")
	ErrConversionToGetAllArticlesFailed  = errors.New("conversion to get_all_articles_response failed")
	ErrConversionToGetArticleByIdFailed  = errors.New("conversion to get_article_by_id_response failed")
	ErrConversionToGetPrevArticlesFailed = errors.New("conversion to get_prev_articles_response failed")
)

// GetAllArticles is implementation of grpc.ArticleServiceServer.GetAllArticles
func (s *ArticleServiceServer) GetAllArticles(ctx context.Context, in *emptypb.Empty) (*grpc.GetAllArticlesResponse, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("GetAllArticles").End()
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters",
			slog.String("in", in.String())))
	oDto, err := s.getAllUsecase.Execute(ctx)
	if err != nil {
		err = errors.WithStack(err)
		logger.InfoContext(ctx, "END",
			slog.Group("return",
				slog.Any("grpc.GetAllArticlesResponse", nil),
				slog.Any("error", err)))
		return nil, err
	}
	res, ok := s.getAllConv.ToGetAllArticlesResponse(ctx, oDto)
	if !ok {
		nrtx.NoticeError(nrpkgerrors.Wrap(ErrConversionToGetAllArticlesFailed))
		return nil, ErrConversionToGetAllArticlesFailed
	}
	logger.InfoContext(ctx, "END",
		slog.Group("return",
			slog.Any("error", nil)))
	return res, nil
}

// GetNextArticles is implementation of grpc.ArticleServiceServer.GetNextArticles
func (s *ArticleServiceServer) GetNextArticles(ctx context.Context, in *grpc.GetNextArticlesRequest) (*grpc.GetNextArticlesResponse, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("GetNextArticles").End()
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters",
			slog.String("in", in.String())))
	oDto, err := s.getNextUsecase.Execute(ctx, dto.NewGetNextInDto(int(in.First), in.After))
	if err != nil {
		err = errors.WithStack(err)
		logger.InfoContext(ctx, "END",
			slog.Group("return",
				slog.Any("grpc.GetNextArticlesResponse", nil),
				slog.Any("error", err)))
		return nil, err
	}
	res, ok := s.getNextConv.ToGetNextArticlesResponse(ctx, oDto)
	if !ok {
		nrtx.NoticeError(nrpkgerrors.Wrap(ErrConversionToGetNextArticlesFailed))
		return nil, ErrConversionToGetNextArticlesFailed
	}
	logger.InfoContext(ctx, "END",
		slog.Group("return",
			slog.Any("error", nil)))
	return res, nil
}

// GetArticleById is implementation of grpc.ArticleServiceServer.GetArticleById
func (s *ArticleServiceServer) GetArticleById(ctx context.Context, in *grpc.GetArticleByIdRequest) (*grpc.GetArticleByIdResponse, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("GetArticleById").End()
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters",
			slog.String("in", in.String())))
	oDto, err := s.getByIdUsecase.Execute(ctx, dto.NewGetByIdInDto(in.GetId()))
	if err != nil {
		err = errors.WithStack(err)
		logger.InfoContext(ctx, "END",
			slog.Group("return",
				slog.Any("pb.GetArticleByIdResponse", nil),
				slog.Any("error", err)))
		return nil, err
	}
	res, ok := s.getByIdConv.ToGetByIdArticlesResponse(ctx, oDto)
	if !ok {
		nrtx.NoticeError(nrpkgerrors.Wrap(ErrConversionToGetArticleByIdFailed))
		return nil, ErrConversionToGetArticleByIdFailed
	}
	logger.InfoContext(ctx, "END",
		slog.Group("return",
			slog.Any("error", nil)))
	return res, nil
}

// GetPrevArticles is implementation of grpc.ArticleServiceServer.GetPrevArticles
func (s *ArticleServiceServer) GetPrevArticles(ctx context.Context, in *grpc.GetPrevArticlesRequest) (*grpc.GetPrevArticlesResponse, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("GetPrevArticles").End()
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters",
			slog.String("in", in.String())))
	oDto, err := s.getPrevUsecase.Execute(ctx, dto.NewGetPrevInDto(int(in.Last), in.Before))
	if err != nil {
		err = errors.WithStack(err)
		logger.InfoContext(ctx, "END",
			slog.Group("return",
				slog.Any("pb.GetPrevArticlesResponse", nil),
				slog.Any("error", err)))
		return nil, err
	}
	res, ok := s.getPrevConv.ToGetPrevArticlesResponse(ctx, oDto)
	if !ok {
		nrtx.NoticeError(nrpkgerrors.Wrap(ErrConversionToGetPrevArticlesFailed))
		return nil, ErrConversionToGetPrevArticlesFailed
	}
	logger.InfoContext(ctx, "END",
		slog.Group("return",
			slog.Any("error", nil)))
	return res, nil
}

// NewArticleServiceServer is constructor of ArticleServiceServer
func NewArticleServiceServer(
	getByIdUsecase usecase.GetById,
	getAllUsecase usecase.GetAll,
	getNextUsecase usecase.GetNext,
	getPrevUsecase usecase.GetPrev,
	getByIdConv presenter.ToGetByIdConverter,
	getAllConv presenter.ToGetAllConverter,
	getNextConv presenter.ToGetNextConverter,
	getPrevConv presenter.ToGetPrevConverter,
) *ArticleServiceServer {
	return &ArticleServiceServer{
		getByIdUsecase: getByIdUsecase,
		getAllUsecase:  getAllUsecase,
		getNextUsecase: getNextUsecase,
		getNextConv:    getNextConv,
		getAllConv:     getAllConv,
		getByIdConv:    getByIdConv,
		getPrevUsecase: getPrevUsecase,
		getPrevConv:    getPrevConv,
	}
}
