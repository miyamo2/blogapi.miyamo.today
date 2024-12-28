package pb

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	"github.com/miyamo2/blogapi.miyamo.today/tag-service/internal/infra/grpc"
	"github.com/miyamo2/blogapi.miyamo.today/tag-service/internal/infra/grpc/grpcconnect"
	"log/slog"

	"github.com/miyamo2/altnrslog"

	"github.com/newrelic/go-agent/v3/newrelic"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/cockroachdb/errors"
	"github.com/miyamo2/blogapi.miyamo.today/tag-service/internal/app/usecase/dto"
	"github.com/miyamo2/blogapi.miyamo.today/tag-service/internal/if-adapter/controller/pb/presenter"
	"github.com/miyamo2/blogapi.miyamo.today/tag-service/internal/if-adapter/controller/pb/usecase"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
)

// TagServiceServer is implementation of grpcconnect.TagServiceServer
type TagServiceServer struct {
	grpcconnect.UnimplementedTagServiceHandler
	getByIdUsecase usecase.GetById
	getByIdConv    presenter.ToGetByIdConverter
	getAllUsecase  usecase.GetAll
	getAllConv     presenter.ToGetAllConverter
	getNextUsecase usecase.GetNext
	getNextConv    presenter.ToGetNextConverter
	getPrevUsecase usecase.GetPrev
	getPrevConv    presenter.ToGetPrevConverter
}

var (
	ErrConversionToGetTagByIdFailed  = errors.New("conversion to get_tag_by_id_response failed")
	ErrConversionToGetAllTagsFailed  = errors.New("conversion to get_all_tags_response failed")
	ErrConversionToGetNextTagsFailed = errors.New("conversion to get_next_tags_response failed")
	ErrConversionToGetPrevTagsFailed = errors.New("conversion to get_prev_tags_response failed")
)

// GetTagById is implementation of grpcconnect.TagServiceServer#GetTagById
func (s *TagServiceServer) GetTagById(ctx context.Context, in *connect.Request[grpc.GetTagByIdRequest]) (*connect.Response[grpc.GetTagByIdResponse], error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("GetTagById").End()
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = slog.Default()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters",
			slog.Any("in", in)))
	oDto, err := s.getByIdUsecase.Execute(ctx, dto.NewGetByIdInDto(in.Msg.GetId()))
	if err != nil {
		err = errors.WithStack(err)
		logger.InfoContext(ctx, "END",
			slog.Group("return",
				slog.Any("*grpcconnect.GetTagByIdResponse", nil),
				slog.String("error", fmt.Sprintf("%+v", err))))
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		return nil, err
	}
	res, ok := s.getByIdConv.ToGetByIdTagResponse(ctx, oDto)
	if !ok {
		err = ErrConversionToGetTagByIdFailed
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		return nil, err
	}
	logger.InfoContext(ctx, "END",
		slog.Group("return",
			// slog.Any("*grpcconnect.GetTagByIdResponse", res),
			slog.Any("error", nil)))
	return res, nil
}

// GetAllTags is implementation of grpcconnect.TagServiceServer#GetTagById
func (s *TagServiceServer) GetAllTags(ctx context.Context, _ *connect.Request[emptypb.Empty]) (*connect.Response[grpc.GetAllTagsResponse], error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("GetAllTags").End()
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = slog.Default()
	}
	logger.InfoContext(ctx, "BEGIN")
	oDto, err := s.getAllUsecase.Execute(ctx)
	if err != nil {
		err = errors.WithStack(err)
		logger.InfoContext(ctx, "END",
			slog.Group("return",
				slog.Any("*grpcconnect.GetAllTagsResponse", nil),
				slog.String("error", fmt.Sprintf("%+v", err))))
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		return nil, err
	}
	res, ok := s.getAllConv.ToGetAllTagsResponse(ctx, oDto)
	if !ok {
		err = ErrConversionToGetAllTagsFailed
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		return nil, err
	}
	logger.InfoContext(ctx, "END",
		slog.Group("return",
			// slog.Any("*grpcconnect.GetAllTagsResponse", res),
			slog.Any("error", nil)))
	return res, nil
}

// GetNextTags is implementation of grpcconnect.TagServiceServer#GetNextTags
func (s *TagServiceServer) GetNextTags(ctx context.Context, in *connect.Request[grpc.GetNextTagsRequest]) (*connect.Response[grpc.GetNextTagResponse], error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("GetTagById").End()
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = slog.Default()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters",
			slog.Any("in", in)))
	oDto, err := s.getNextUsecase.Execute(ctx, dto.NewGetNextInDto(int(in.Msg.GetFirst()), in.Msg.After))
	if err != nil {
		err = errors.WithStack(err)
		logger.InfoContext(ctx, "END",
			slog.Group("return",
				slog.Any("*grpcconnect.GetNextTagResponse", nil),
				slog.String("error", fmt.Sprintf("%+v", err))))
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		return nil, err
	}
	res, ok := s.getNextConv.ToGetNextTagsResponse(ctx, oDto)
	if !ok {
		err = ErrConversionToGetNextTagsFailed
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		return nil, err
	}
	logger.InfoContext(ctx, "END",
		slog.Group("return",
			// slog.Any("*grpcconnect.GetNextTagResponse", res),
			slog.Any("error", nil)))
	return res, nil
}

// GetPrevTags is implementation of grpcconnect.TagServiceServer#GetPrevTags
func (s *TagServiceServer) GetPrevTags(ctx context.Context, in *connect.Request[grpc.GetPrevTagsRequest]) (*connect.Response[grpc.GetPrevTagResponse], error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("GetPrevTags").End()
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = slog.Default()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters",
			slog.Any("in", in)))
	oDto, err := s.getPrevUsecase.Execute(ctx, dto.NewGetPrevInDto(int(in.Msg.GetLast()), in.Msg.Before))
	if err != nil {
		err = errors.WithStack(err)
		logger.InfoContext(ctx, "END",
			slog.Group("return",
				slog.Any("*grpcconnect.GetPrevTagResponse", nil),
				slog.String("error", fmt.Sprintf("%+v", err))))
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		return nil, err
	}
	res, ok := s.getPrevConv.ToGetPrevTagsResponse(ctx, oDto)
	if !ok {
		err = ErrConversionToGetPrevTagsFailed
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		return nil, err
	}
	logger.InfoContext(ctx, "END",
		slog.Group("return",
			// slog.Any("*grpcconnect.GetPrevTagResponse", res),
			slog.Any("error", nil)))
	return res, nil
}

// NewTagServiceServer is constructor of TagServiceServer
func NewTagServiceServer(
	getByIdUsecase usecase.GetById,
	getByIdConv presenter.ToGetByIdConverter,
	getAllUsecase usecase.GetAll,
	getAllConv presenter.ToGetAllConverter,
	getNextUsecase usecase.GetNext,
	getNextConv presenter.ToGetNextConverter,
	getPrevUsecase usecase.GetPrev,
	getPrevConv presenter.ToGetPrevConverter,
) *TagServiceServer {
	return &TagServiceServer{
		getByIdUsecase: getByIdUsecase,
		getByIdConv:    getByIdConv,
		getAllUsecase:  getAllUsecase,
		getAllConv:     getAllConv,
		getNextUsecase: getNextUsecase,
		getNextConv:    getNextConv,
		getPrevUsecase: getPrevUsecase,
		getPrevConv:    getPrevConv,
	}
}
