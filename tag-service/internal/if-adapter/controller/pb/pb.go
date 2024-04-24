package pb

import (
	"context"
	"fmt"
	"github.com/miyamo2/altnrslog"
	"log/slog"

	"github.com/newrelic/go-agent/v3/newrelic"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/cockroachdb/errors"
	"github.com/miyamo2/blogapi-core/util/duration"
	"github.com/miyamo2/blogapi-tag-service/internal/app/usecase/dto"
	"github.com/miyamo2/blogapi-tag-service/internal/if-adapter/controller/pb/presenter"
	"github.com/miyamo2/blogapi-tag-service/internal/if-adapter/controller/pb/usecase"
	"github.com/miyamo2/blogproto-gen/tag/server/pb"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
)

// TagServiceServer is implementation of pb.TagServiceServer
type TagServiceServer struct {
	pb.UnimplementedTagServiceServer
	getByIdUsecase usecase.GetById[dto.GetByIdInDto, dto.Article, *dto.GetByIdOutDto]
	getByIdConv    presenter.ToGetByIdConverter[dto.Article, *dto.GetByIdOutDto]
	getAllUsecase  usecase.GetAll[dto.Article, dto.Tag, *dto.GetAllOutDto]
	getAllConv     presenter.ToGetAllConverter[dto.Article, dto.Tag, *dto.GetAllOutDto]
	getNextUsecase usecase.GetNext[dto.GetNextInDto, dto.Article, dto.Tag, *dto.GetNextOutDto]
	getNextConv    presenter.ToGetNextConverter[dto.Article, dto.Tag, *dto.GetNextOutDto]
	getPrevUsecase usecase.GetPrev[dto.GetPrevInDto, dto.Article, dto.Tag, *dto.GetPrevOutDto]
	getPrevConv    presenter.ToGetPrevConverter[dto.Article, dto.Tag, *dto.GetPrevOutDto]
}

var (
	ErrConversionToGetTagByIdFailed  = errors.New("conversion to get_tag_by_id_response failed")
	ErrConversionToGetAllTagsFailed  = errors.New("conversion to get_all_tags_response failed")
	ErrConversionToGetNextTagsFailed = errors.New("conversion to get_next_tags_response failed")
	ErrConversionToGetPrevTagsFailed = errors.New("conversion to get_prev_tags_response failed")
)

// GetTagById is implementation of pb.TagServiceServer#GetTagById
func (s *TagServiceServer) GetTagById(ctx context.Context, in *pb.GetTagByIdRequest) (*pb.GetTagByIdResponse, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("GetTagById").End()
	dw := duration.Start()
	lgr, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		lgr = slog.Default()
	}
	lgr.InfoContext(ctx, "BEGIN",
		slog.Group("parameters",
			slog.Any("in", in)))
	oDto, err := s.getByIdUsecase.Execute(ctx, dto.NewGetByIdInDto(in.GetId()))
	if err != nil {
		err = errors.WithStack(err)
		lgr.InfoContext(ctx, "END",
			slog.String("duration", dw.SDuration()),
			slog.Group("return",
				slog.Any("*pb.GetTagByIdResponse", nil),
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
	lgr.InfoContext(ctx, "END",
		slog.String("duration", dw.SDuration()),
		slog.Group("return",
			// slog.Any("*pb.GetTagByIdResponse", res),
			slog.Any("error", nil)))
	return res, nil
}

// GetAllTags is implementation of pb.TagServiceServer#GetTagById
func (s *TagServiceServer) GetAllTags(ctx context.Context, _ *emptypb.Empty) (*pb.GetAllTagsResponse, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("GetAllTags").End()
	dw := duration.Start()
	lgr, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		lgr = slog.Default()
	}
	lgr.InfoContext(ctx, "BEGIN")
	oDto, err := s.getAllUsecase.Execute(ctx)
	if err != nil {
		err = errors.WithStack(err)
		lgr.InfoContext(ctx, "END",
			slog.String("duration", dw.SDuration()),
			slog.Group("return",
				slog.Any("*pb.GetAllTagsResponse", nil),
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
	lgr.InfoContext(ctx, "END",
		slog.String("duration", dw.SDuration()),
		slog.Group("return",
			// slog.Any("*pb.GetAllTagsResponse", res),
			slog.Any("error", nil)))
	return res, nil
}

// GetNextTags is implementation of pb.TagServiceServer#GetNextTags
func (s *TagServiceServer) GetNextTags(ctx context.Context, in *pb.GetNextTagsRequest) (*pb.GetNextTagResponse, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("GetTagById").End()
	dw := duration.Start()
	lgr, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		lgr = slog.Default()
	}
	lgr.InfoContext(ctx, "BEGIN",
		slog.Group("parameters",
			slog.Any("in", in)))
	oDto, err := s.getNextUsecase.Execute(ctx, dto.NewGetNextInDto(int(in.GetFirst()), in.After))
	if err != nil {
		err = errors.WithStack(err)
		lgr.InfoContext(ctx, "END",
			slog.String("duration", dw.SDuration()),
			slog.Group("return",
				slog.Any("*pb.GetNextTagResponse", nil),
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
	lgr.InfoContext(ctx, "END",
		slog.String("duration", dw.SDuration()),
		slog.Group("return",
			// slog.Any("*pb.GetNextTagResponse", res),
			slog.Any("error", nil)))
	return res, nil
}

// GetPrevTags is implementation of pb.TagServiceServer#GetPrevTags
func (s *TagServiceServer) GetPrevTags(ctx context.Context, in *pb.GetPrevTagsRequest) (*pb.GetPrevTagResponse, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("GetPrevTags").End()
	dw := duration.Start()
	lgr, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		lgr = slog.Default()
	}
	lgr.InfoContext(ctx, "BEGIN",
		slog.Group("parameters",
			slog.Any("in", in)))
	oDto, err := s.getPrevUsecase.Execute(ctx, dto.NewGetPrevInDto(int(in.GetLast()), in.Before))
	if err != nil {
		err = errors.WithStack(err)
		lgr.InfoContext(ctx, "END",
			slog.String("duration", dw.SDuration()),
			slog.Group("return",
				slog.Any("*pb.GetPrevTagResponse", nil),
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
	lgr.InfoContext(ctx, "END",
		slog.String("duration", dw.SDuration()),
		slog.Group("return",
			// slog.Any("*pb.GetPrevTagResponse", res),
			slog.Any("error", nil)))
	return res, nil
}

// NewTagServiceServer is constructor of TagServiceServer
func NewTagServiceServer(
	getByIdUsecase usecase.GetById[dto.GetByIdInDto, dto.Article, *dto.GetByIdOutDto],
	getByIdConv presenter.ToGetByIdConverter[dto.Article, *dto.GetByIdOutDto],
	getAllUsecase usecase.GetAll[dto.Article, dto.Tag, *dto.GetAllOutDto],
	getAllConv presenter.ToGetAllConverter[dto.Article, dto.Tag, *dto.GetAllOutDto],
	getNextUsecase usecase.GetNext[dto.GetNextInDto, dto.Article, dto.Tag, *dto.GetNextOutDto],
	getNextConv presenter.ToGetNextConverter[dto.Article, dto.Tag, *dto.GetNextOutDto],
	getPrevUsecase usecase.GetPrev[dto.GetPrevInDto, dto.Article, dto.Tag, *dto.GetPrevOutDto],
	getPrevConv presenter.ToGetPrevConverter[dto.Article, dto.Tag, *dto.GetPrevOutDto],
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
