package pb

import (
	"context"

	"blogapi.miyamo.today/article-service/internal/if-adapter/controller/pb/presenter/convert"
	"blogapi.miyamo.today/article-service/internal/infra/grpc/grpcconnect"
	"connectrpc.com/connect"

	"blogapi.miyamo.today/article-service/internal/app/usecase/dto"
	"blogapi.miyamo.today/article-service/internal/if-adapter/controller/pb/usecase"
	"blogapi.miyamo.today/article-service/internal/infra/grpc"
	"github.com/cockroachdb/errors"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"
	"google.golang.org/protobuf/types/known/emptypb"
)

// compatibility check
var _ grpcconnect.ArticleServiceHandler = (*ArticleServiceServer)(nil)

// ArticleServiceServer implements grpc.ArticleServiceServer
type ArticleServiceServer struct {
	getByIDUsecase      usecase.GetByID
	listAllUsecase      usecase.ListAll
	listAfterUsecase    usecase.ListAfter
	listBeforeUsecase   usecase.ListBefore
	listAfterConverter  convert.ListAfter
	listAllConverter    convert.ListAll
	getByIDConverter    convert.GetByID
	listBeforeConverter convert.ListBefore
}

var (
	ErrConversionToListNextFailed = errors.New("conversion to get_next_articles_response failed")
	ErrConversionToListAllFailed  = errors.New("conversion to get_all_articles_response failed")
	ErrConversionToGetByIDFailed  = errors.New("conversion to get_article_by_id_response failed")
	ErrConversionToListPrevFailed = errors.New("conversion to get_prev_articles_response failed")
)

// GetAllArticles implements grpc.ArticleServiceServer.GetAllArticles
func (s *ArticleServiceServer) GetAllArticles(
	ctx context.Context, in *connect.Request[emptypb.Empty],
) (*connect.Response[grpc.GetAllArticlesResponse], error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("GetAllArticles").End()

	oDto, err := s.listAllUsecase.Execute(ctx)
	if err != nil {
		err = nrpkgerrors.Wrap(errors.WithStack(err))
		nrtx.NoticeError(err)
		return nil, err
	}
	res, ok := s.listAllConverter.ToResponse(ctx, oDto)
	if !ok {
		err := nrpkgerrors.Wrap(ErrConversionToListAllFailed)
		nrtx.NoticeError(err)
		return nil, err
	}

	return connect.NewResponse(res), nil
}

// GetNextArticles implements grpc.ArticleServiceServer.GetNextArticles
func (s *ArticleServiceServer) GetNextArticles(
	ctx context.Context, in *connect.Request[grpc.GetNextArticlesRequest],
) (*connect.Response[grpc.GetNextArticlesResponse], error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("GetNextArticles").End()

	oDto, err := s.listAfterUsecase.Execute(
		ctx, dto.NewListAfterInput(
			int(in.Msg.First),
			dto.ListAfterInputWithCursor(in.Msg.After),
		),
	)
	if err != nil {
		err = nrpkgerrors.Wrap(errors.WithStack(err))
		nrtx.NoticeError(err)
		return nil, err
	}
	res, ok := s.listAfterConverter.ToResponse(ctx, oDto)
	if !ok {
		err := nrpkgerrors.Wrap(ErrConversionToListNextFailed)
		nrtx.NoticeError(err)
		return nil, err
	}

	return connect.NewResponse(res), nil
}

// GetArticleById implements grpc.ArticleServiceServer.GetArticleById
func (s *ArticleServiceServer) GetArticleById(
	ctx context.Context, in *connect.Request[grpc.GetArticleByIdRequest],
) (*connect.Response[grpc.GetArticleByIdResponse], error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("GetArticleById").End()

	oDto, err := s.getByIDUsecase.Execute(ctx, dto.NewGetByIDInput(in.Msg.GetId()))
	if err != nil {
		err = nrpkgerrors.Wrap(errors.WithStack(err))
		nrtx.NoticeError(err)
		return nil, err
	}
	res, ok := s.getByIDConverter.ToResponse(ctx, oDto)
	if !ok {
		err := nrpkgerrors.Wrap(ErrConversionToGetByIDFailed)
		nrtx.NoticeError(err)
		return nil, err
	}
	return connect.NewResponse(res), nil
}

// GetPrevArticles implements grpc.ArticleServiceServer.GetPrevArticles
func (s *ArticleServiceServer) GetPrevArticles(
	ctx context.Context, in *connect.Request[grpc.GetPrevArticlesRequest],
) (
	*connect.Response[grpc.GetPrevArticlesResponse], error,
) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("GetPrevArticles").End()

	oDto, err := s.listBeforeUsecase.Execute(
		ctx,
		dto.NewListBeforeInput(int(in.Msg.Last), dto.ListBeforeInputWithCursor(in.Msg.Before)),
	)
	if err != nil {
		err = nrpkgerrors.Wrap(errors.WithStack(err))
		nrtx.NoticeError(err)
		return nil, err
	}
	res, ok := s.listBeforeConverter.ToResponse(ctx, oDto)
	if !ok {
		nrtx.NoticeError(nrpkgerrors.Wrap(ErrConversionToListPrevFailed))
		return nil, ErrConversionToListPrevFailed
	}
	return connect.NewResponse(res), nil
}

// NewArticleServiceServerOption sets options for NewArticleServiceServer
type NewArticleServiceServerOption func(server *ArticleServiceServer)

// WithGetByID sets GetById usecase and converter
func WithGetByID(u usecase.GetByID, conv convert.GetByID) NewArticleServiceServerOption {
	return func(s *ArticleServiceServer) {
		s.getByIDUsecase = u
		s.getByIDConverter = conv
	}
}

// WithListAll sets ListAll usecase and converter
func WithListAll(u usecase.ListAll, conv convert.ListAll) NewArticleServiceServerOption {
	return func(s *ArticleServiceServer) {
		s.listAllUsecase = u
		s.listAllConverter = conv
	}
}

// WithListAfter sets ListAfter usecase and converter
func WithListAfter(u usecase.ListAfter, conv convert.ListAfter) NewArticleServiceServerOption {
	return func(s *ArticleServiceServer) {
		s.listAfterUsecase = u
		s.listAfterConverter = conv
	}
}

// WithListBefore sets ListBefore usecase and converter
func WithListBefore(u usecase.ListBefore, conv convert.ListBefore) NewArticleServiceServerOption {
	return func(s *ArticleServiceServer) {
		s.listBeforeUsecase = u
		s.listBeforeConverter = conv
	}
}

// NewArticleServiceServer constructs ArticleServiceServer
func NewArticleServiceServer(options ...NewArticleServiceServerOption) *ArticleServiceServer {
	var s ArticleServiceServer
	for _, opt := range options {
		opt(&s)
	}
	return &s
}
