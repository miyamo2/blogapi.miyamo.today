package pb

import (
	"context"
	"github.com/cockroachdb/errors"
	"github.com/miyamo2/altnrslog"
	"github.com/miyamo2/blogapi-article-service/internal/app/usecase/dto"
	"github.com/miyamo2/blogapi-article-service/internal/if-adapter/controller/pb/presenter"
	"github.com/miyamo2/blogapi-article-service/internal/if-adapter/controller/pb/usecase"
	"github.com/miyamo2/blogapi-core/log"
	"github.com/miyamo2/blogapi-core/util/duration"
	"github.com/miyamo2/blogproto-gen/article/server/pb"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"
	"google.golang.org/protobuf/types/known/emptypb"
	"log/slog"
)

// ArticleServiceServer is implementation of pb.ArticleServiceServer
type ArticleServiceServer struct {
	pb.UnimplementedArticleServiceServer
	getByIdUsecase usecase.GetById[dto.GetByIdInDto, dto.Tag, *dto.GetByIdOutDto]
	getAllUsecase  usecase.GetAll[dto.Tag, dto.Article, *dto.GetAllOutDto]
	getNextUsecase usecase.GetNext[dto.GetNextInDto, dto.Tag, dto.Article, *dto.GetNextOutDto]
	getPrevUsecase usecase.GetPrev[dto.GetPrevInDto, dto.Tag, dto.Article, *dto.GetPrevOutDto]
	getNextConv    presenter.ToGetNextConverter[dto.Tag, dto.Article, *dto.GetNextOutDto]
	getAllConv     presenter.ToGetAllConverter[dto.Tag, dto.Article, *dto.GetAllOutDto]
	getByIdConv    presenter.ToGetByIdConverter[dto.Tag, *dto.GetByIdOutDto]
	getPrevConv    presenter.ToGetPrevConverter[dto.Tag, dto.Article, *dto.GetPrevOutDto]
}

var (
	ErrConversionToGetNextArticlesFailed = errors.New("conversion to get_next_articles_response failed")
	ErrConversionToGetAllArticlesFailed  = errors.New("conversion to get_all_articles_response failed")
	ErrConversionToGetArticleByIdFailed  = errors.New("conversion to get_article_by_id_response failed")
	ErrConversionToGetPrevArticlesFailed = errors.New("conversion to get_prev_articles_response failed")
)

// GetAllArticles is implementation of pb.ArticleServiceServer.GetAllArticles
func (s *ArticleServiceServer) GetAllArticles(ctx context.Context, in *emptypb.Empty) (*pb.GetAllArticlesResponse, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("GetAllArticles").End()
	dw := duration.Start()
	lgr, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		lgr = log.DefaultLogger()
	}
	lgr.InfoContext(ctx, "BEGIN",
		slog.Group("parameters",
			slog.String("in", in.String())))
	oDto, err := s.getAllUsecase.Execute(ctx)
	if err != nil {
		err = errors.WithStack(err)
		lgr.InfoContext(ctx, "END",
			slog.String("duration", dw.SDuration()),
			slog.Group("return",
				slog.Any("pb.GetAllArticlesResponse", nil),
				slog.Any("error", err)))
		return nil, err
	}
	res, ok := s.getAllConv.ToGetAllArticlesResponse(ctx, oDto)
	if !ok {
		nrtx.NoticeError(nrpkgerrors.Wrap(ErrConversionToGetAllArticlesFailed))
		return nil, ErrConversionToGetAllArticlesFailed
	}
	lgr.InfoContext(ctx, "END",
		slog.String("duration", dw.SDuration()),
		slog.Group("return",
			slog.Any("error", nil)))
	return res, nil
}

// GetNextArticles is implementation of pb.ArticleServiceServer.GetNextArticles
func (s *ArticleServiceServer) GetNextArticles(ctx context.Context, in *pb.GetNextArticlesRequest) (*pb.GetNextArticlesResponse, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("GetNextArticles").End()
	dw := duration.Start()
	lgr, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		lgr = log.DefaultLogger()
	}
	lgr.InfoContext(ctx, "BEGIN",
		slog.Group("parameters",
			slog.String("in", in.String())))
	oDto, err := s.getNextUsecase.Execute(ctx, dto.NewGetNextInDto(int(in.First), in.After))
	if err != nil {
		err = errors.WithStack(err)
		lgr.InfoContext(ctx, "END",
			slog.String("duration", dw.SDuration()),
			slog.Group("return",
				slog.Any("pb.GetNextArticlesResponse", nil),
				slog.Any("error", err)))
		return nil, err
	}
	res, ok := s.getNextConv.ToGetNextArticlesResponse(ctx, oDto)
	if !ok {
		nrtx.NoticeError(nrpkgerrors.Wrap(ErrConversionToGetNextArticlesFailed))
		return nil, ErrConversionToGetNextArticlesFailed
	}
	lgr.InfoContext(ctx, "END",
		slog.String("duration", dw.SDuration()),
		slog.Group("return",
			slog.Any("error", nil)))
	return res, nil
}

// GetArticleById is implementation of pb.ArticleServiceServer.GetArticleById
func (s *ArticleServiceServer) GetArticleById(ctx context.Context, in *pb.GetArticleByIdRequest) (*pb.GetArticleByIdResponse, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("GetArticleById").End()
	dw := duration.Start()
	lgr, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		lgr = log.DefaultLogger()
	}
	lgr.InfoContext(ctx, "BEGIN",
		slog.Group("parameters",
			slog.String("in", in.String())))
	oDto, err := s.getByIdUsecase.Execute(ctx, dto.NewGetByIdInDto(in.GetId()))
	if err != nil {
		err = errors.WithStack(err)
		lgr.InfoContext(ctx, "END",
			slog.String("duration", dw.SDuration()),
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
	lgr.InfoContext(ctx, "END",
		slog.String("duration", dw.SDuration()),
		slog.Group("return",
			slog.Any("error", nil)))
	return res, nil
}

// GetPrevArticles is implementation of pb.ArticleServiceServer.GetPrevArticles
func (s *ArticleServiceServer) GetPrevArticles(ctx context.Context, in *pb.GetPrevArticlesRequest) (*pb.GetPrevArticlesResponse, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("GetPrevArticles").End()
	dw := duration.Start()
	lgr, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		lgr = log.DefaultLogger()
	}
	lgr.InfoContext(ctx, "BEGIN",
		slog.Group("parameters",
			slog.String("in", in.String())))
	oDto, err := s.getPrevUsecase.Execute(ctx, dto.NewGetPrevInDto(int(in.Last), in.Before))
	if err != nil {
		err = errors.WithStack(err)
		lgr.InfoContext(ctx, "END",
			slog.String("duration", dw.SDuration()),
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
	lgr.InfoContext(ctx, "END",
		slog.String("duration", dw.SDuration()),
		slog.Group("return",
			slog.Any("error", nil)))
	return res, nil
}

// NewArticleServiceServer is constructor of ArticleServiceServer
func NewArticleServiceServer(
	getByIdUsecase usecase.GetById[dto.GetByIdInDto, dto.Tag, *dto.GetByIdOutDto],
	getAllUsecase usecase.GetAll[dto.Tag, dto.Article, *dto.GetAllOutDto],
	getNextUsecase usecase.GetNext[dto.GetNextInDto, dto.Tag, dto.Article, *dto.GetNextOutDto],
	getPrevUsecase usecase.GetPrev[dto.GetPrevInDto, dto.Tag, dto.Article, *dto.GetPrevOutDto],
	getByIdConv presenter.ToGetByIdConverter[dto.Tag, *dto.GetByIdOutDto],
	getAllConv presenter.ToGetAllConverter[dto.Tag, dto.Article, *dto.GetAllOutDto],
	getNextConv presenter.ToGetNextConverter[dto.Tag, dto.Article, *dto.GetNextOutDto],
	getPrevConv presenter.ToGetPrevConverter[dto.Tag, dto.Article, *dto.GetPrevOutDto],
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
