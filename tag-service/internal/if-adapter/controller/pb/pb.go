package pb

import (
	"context"
	"fmt"
	"github.com/miyamo2/blogapi.miyamo.today/tag-service/internal/infra/grpc"
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

// TagServiceServer is implementation of grpc.TagServiceServer
type TagServiceServer struct {
	grpc.UnimplementedTagServiceServer
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

// GetTagById is implementation of grpc.TagServiceServer#GetTagById
func (s *TagServiceServer) GetTagById(ctx context.Context, in *grpc.GetTagByIdRequest) (*grpc.GetTagByIdResponse, error) {
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
	oDto, err := s.getByIdUsecase.Execute(ctx, dto.NewGetByIdInDto(in.GetId()))
	if err != nil {
		err = errors.WithStack(err)
		logger.InfoContext(ctx, "END",
			slog.Group("return",
				slog.Any("*grpc.GetTagByIdResponse", nil),
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
			// slog.Any("*grpc.GetTagByIdResponse", res),
			slog.Any("error", nil)))
	return res, nil
}

// GetAllTags is implementation of grpc.TagServiceServer#GetTagById
func (s *TagServiceServer) GetAllTags(ctx context.Context, _ *emptypb.Empty) (*grpc.GetAllTagsResponse, error) {
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
				slog.Any("*grpc.GetAllTagsResponse", nil),
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
			// slog.Any("*grpc.GetAllTagsResponse", res),
			slog.Any("error", nil)))
	return res, nil
}

// GetNextTags is implementation of grpc.TagServiceServer#GetNextTags
func (s *TagServiceServer) GetNextTags(ctx context.Context, in *grpc.GetNextTagsRequest) (*grpc.GetNextTagResponse, error) {
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
	oDto, err := s.getNextUsecase.Execute(ctx, dto.NewGetNextInDto(int(in.GetFirst()), in.After))
	if err != nil {
		err = errors.WithStack(err)
		logger.InfoContext(ctx, "END",
			slog.Group("return",
				slog.Any("*grpc.GetNextTagResponse", nil),
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
			// slog.Any("*grpc.GetNextTagResponse", res),
			slog.Any("error", nil)))
	return res, nil
}

// GetPrevTags is implementation of grpc.TagServiceServer#GetPrevTags
func (s *TagServiceServer) GetPrevTags(ctx context.Context, in *grpc.GetPrevTagsRequest) (*grpc.GetPrevTagResponse, error) {
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
	oDto, err := s.getPrevUsecase.Execute(ctx, dto.NewGetPrevInDto(int(in.GetLast()), in.Before))
	if err != nil {
		err = errors.WithStack(err)
		logger.InfoContext(ctx, "END",
			slog.Group("return",
				slog.Any("*grpc.GetPrevTagResponse", nil),
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
			// slog.Any("*grpc.GetPrevTagResponse", res),
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
