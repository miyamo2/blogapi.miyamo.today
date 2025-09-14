package pb

import (
	"context"

	"blogapi.miyamo.today/tag-service/internal/if-adapter/controller/pb/presenter/convert"
	"blogapi.miyamo.today/tag-service/internal/infra/grpc"
	"blogapi.miyamo.today/tag-service/internal/infra/grpc/grpcconnect"
	"connectrpc.com/connect"

	"github.com/newrelic/go-agent/v3/newrelic"
	"google.golang.org/protobuf/types/known/emptypb"

	"blogapi.miyamo.today/tag-service/internal/app/usecase/dto"
	"blogapi.miyamo.today/tag-service/internal/if-adapter/controller/pb/usecase"
	"github.com/cockroachdb/errors"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
)

// compatibility check
var _ grpcconnect.TagServiceHandler = (*TagServiceServer)(nil)

// TagServiceServer implements grpcconnect.TagServiceServer
type TagServiceServer struct {
	getByIdUsecase    usecase.GetById
	getByIdConverter  convert.ToGetById
	listAllUsecase    usecase.ListAll
	getAllConverter   convert.ToGetAll
	listAfterUsecase  usecase.ListAfter
	getNextConverter  convert.ToGetNext
	listBeforeUsecase usecase.ListBefore
	getPrevConverter  convert.ToGetPrev
}

var (
	ErrConversionToGetTagByIdFailed  = errors.New("conversion to get_tag_by_id_response failed")
	ErrConversionToGetAllTagsFailed  = errors.New("conversion to get_all_tags_response failed")
	ErrConversionToGetNextTagsFailed = errors.New("conversion to get_next_tags_response failed")
	ErrConversionToGetPrevTagsFailed = errors.New("conversion to get_prev_tags_response failed")
)

// GetTagById implements grpcconnect.TagServiceServer#GetTagById
func (s *TagServiceServer) GetTagById(
	ctx context.Context, in *connect.Request[grpc.GetTagByIdRequest],
) (*connect.Response[grpc.GetTagByIdResponse], error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("GetTagById").End()

	o, err := s.getByIdUsecase.Execute(ctx, dto.NewGetByIdInput(in.Msg.GetId()))
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		return nil, err
	}
	res, ok := s.getByIdConverter.ToResponse(ctx, o)
	if !ok {
		err = ErrConversionToGetTagByIdFailed
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		return nil, err
	}

	return res, nil
}

// GetAllTags implements grpcconnect.TagServiceServer#GetTagById
func (s *TagServiceServer) GetAllTags(
	ctx context.Context, _ *connect.Request[emptypb.Empty],
) (*connect.Response[grpc.GetAllTagsResponse], error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("GetAllTags").End()

	o, err := s.listAllUsecase.Execute(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		return nil, err
	}
	res, ok := s.getAllConverter.ToResponse(ctx, o)
	if !ok {
		err = ErrConversionToGetAllTagsFailed
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		return nil, err
	}
	return res, nil
}

// GetNextTags implements grpcconnect.TagServiceServer#GetNextTags
func (s *TagServiceServer) GetNextTags(
	ctx context.Context, in *connect.Request[grpc.GetNextTagsRequest],
) (*connect.Response[grpc.GetNextTagResponse], error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("GetTagById").End()

	o, err := s.listAfterUsecase.Execute(
		ctx,
		dto.NewListAfterInput(
			int(in.Msg.GetFirst()),
			dto.ListAfterInputWithCursor(in.Msg.After),
		),
	)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		return nil, err
	}
	res, ok := s.getNextConverter.ToResponse(ctx, o)
	if !ok {
		err = ErrConversionToGetNextTagsFailed
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		return nil, err
	}
	return res, nil
}

// GetPrevTags implements grpcconnect.TagServiceServer#GetPrevTags
func (s *TagServiceServer) GetPrevTags(
	ctx context.Context, in *connect.Request[grpc.GetPrevTagsRequest],
) (*connect.Response[grpc.GetPrevTagResponse], error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("GetPrevTags").End()

	o, err := s.listBeforeUsecase.Execute(
		ctx,
		dto.NewListBeforeInput(
			int(in.Msg.GetLast()),
			dto.ListBeforeInputWithCursor(in.Msg.Before),
		),
	)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		return nil, err
	}
	res, ok := s.getPrevConverter.ToResponse(ctx, o)
	if !ok {
		err = ErrConversionToGetPrevTagsFailed
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		return nil, err
	}
	return res, nil
}

// NewTagServiceServerOption sets options for NewTagServiceServer
type NewTagServiceServerOption func(*TagServiceServer)

// WithGetById sets GetById usecase and converter
func WithGetById(u usecase.GetById, conv convert.ToGetById) NewTagServiceServerOption {
	return func(s *TagServiceServer) {
		s.getByIdUsecase = u
		s.getByIdConverter = conv
	}
}

// WithListAll sets ListAll usecase and converter
func WithListAll(u usecase.ListAll, conv convert.ToGetAll) NewTagServiceServerOption {
	return func(s *TagServiceServer) {
		s.listAllUsecase = u
		s.getAllConverter = conv
	}
}

// WithListAfter sets ListAfter usecase and converter
func WithListAfter(u usecase.ListAfter, conv convert.ToGetNext) NewTagServiceServerOption {
	return func(s *TagServiceServer) {
		s.listAfterUsecase = u
		s.getNextConverter = conv
	}
}

// WithListBefore sets ListBefore usecase and converter
func WithListBefore(u usecase.ListBefore, conv convert.ToGetPrev) NewTagServiceServerOption {
	return func(s *TagServiceServer) {
		s.listBeforeUsecase = u
		s.getPrevConverter = conv
	}
}

// NewTagServiceServer constructs TagServiceServer
func NewTagServiceServer(options ...NewTagServiceServerOption) *TagServiceServer {
	v := &TagServiceServer{}
	for _, opt := range options {
		opt(v)
	}
	return v
}
