package pb

import (
	"blogapi.miyamo.today/blogging-event-service/internal/app/usecase/dto"
	"blogapi.miyamo.today/blogging-event-service/internal/if-adapter/controller/pb/presenters"
	"blogapi.miyamo.today/blogging-event-service/internal/if-adapter/controller/pb/usecase"
	grpcgen "blogapi.miyamo.today/blogging-event-service/internal/infra/grpc"
	"blogapi.miyamo.today/blogging-event-service/internal/infra/grpc/grpcconnect"
	"blogapi.miyamo.today/core/log"
	"bytes"
	"connectrpc.com/connect"
	"context"
	"github.com/cockroachdb/errors"
	"github.com/miyamo2/altnrslog"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"
	"log/slog"
	"net/url"
)

var _ grpcconnect.BloggingEventServiceHandler = (*BloggingEventServiceServer)(nil)

// BloggingEventServiceServer is implementation of grpc.BloggingEventServiceServer
type BloggingEventServiceServer struct {
	bloggingEventServiceServerConfig
	grpcconnect.UnimplementedBloggingEventServiceHandler
}

// CreateArticle is implementation of grpc.BloggingEventServiceServer.CreateArticle
func (s *BloggingEventServiceServer) CreateArticle(ctx context.Context, req *connect.Request[grpcgen.CreateArticleRequest]) (*connect.Response[grpcgen.BloggingEventResponse], error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("GetAllArticles").End()
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}

	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.String("title", req.Msg.GetTitle()), slog.String("body", req.Msg.GetBody()), slog.String("thumbnail", req.Msg.GetThumbnailUrl()), slog.Any("tagNames", req.Msg.GetTagNames())))

	inDto := dto.NewCreateArticleInDto(req.Msg.GetTitle(), req.Msg.GetBody(), req.Msg.GetThumbnailUrl(), req.Msg.GetTagNames())
	outDto, err := s.createArticleUsecase.Execute(ctx, &inDto)
	if err != nil {
		return nil, err
	}
	response, err := s.createArticleConverter.ToCreateArticleArticleResponse(ctx, outDto)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger.WarnContext(ctx, "END",
			slog.Group("return",
				slog.Any("grpc.BloggingEventResponse", nil),
				slog.Any("error", err)))
		return nil, err
	}
	logger.InfoContext(ctx, "END",
		slog.Group("return",
			slog.String("grpc.BloggingEventResponse", response.String())))
	return connect.NewResponse(response), nil
}

func (s *BloggingEventServiceServer) UpdateArticleTitle(ctx context.Context, request *connect.Request[grpcgen.UpdateArticleTitleRequest]) (*connect.Response[grpcgen.BloggingEventResponse], error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("UpdateArticleTitle").End()
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.String("article id", request.Msg.GetId()), slog.String("title", request.Msg.GetTitle())))

	inDto := dto.NewUpdateArticleTitleInDto(request.Msg.GetId(), request.Msg.GetTitle())
	outDto, err := s.updateArticleTitleUsecase.Execute(ctx, &inDto)
	if err != nil {
		return nil, err
	}
	response, err := s.updateArticleTitleConverter.ToUpdateArticleTitleResponse(ctx, outDto)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger.WarnContext(ctx, "END",
			slog.Group("return",
				slog.Any("grpc.BloggingEventResponse", nil),
				slog.Any("error", err)))
		return nil, err
	}
	logger.InfoContext(ctx, "END",
		slog.Group("return",
			slog.String("grpc.BloggingEventResponse", response.String())))
	return connect.NewResponse(response), nil
}

func (s *BloggingEventServiceServer) UpdateArticleBody(ctx context.Context, request *connect.Request[grpcgen.UpdateArticleBodyRequest]) (*connect.Response[grpcgen.BloggingEventResponse], error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("UpdateArticleBody").End()
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.String("article id", request.Msg.GetId()), slog.String("body", request.Msg.GetBody())))

	inDto := dto.NewUpdateArticleBodyInDto(request.Msg.GetId(), request.Msg.GetBody())
	outDto, err := s.updateArticleBodyUsecase.Execute(ctx, &inDto)
	if err != nil {
		return nil, err
	}
	response, err := s.updateArticleBodyConverter.ToUpdateArticleBodyResponse(ctx, outDto)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger.WarnContext(ctx, "END",
			slog.Group("return",
				slog.Any("grpc.BloggingEventResponse", nil),
				slog.Any("error", err)))
		return nil, err
	}
	logger.InfoContext(ctx, "END",
		slog.Group("return",
			slog.String("grpc.BloggingEventResponse", response.String())))
	return connect.NewResponse(response), nil
}

func (s *BloggingEventServiceServer) UpdateArticleThumbnail(ctx context.Context, request *connect.Request[grpcgen.UpdateArticleThumbnailRequest]) (*connect.Response[grpcgen.BloggingEventResponse], error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("UpdateArticleThumbnail").End()
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.String("article id", request.Msg.GetId()), slog.String("thumbnail", request.Msg.GetThumbnailUrl())))

	thumbnailUrl, err := url.Parse(request.Msg.GetThumbnailUrl())
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		return nil, err
	}
	inDto := dto.NewUpdateArticleThumbnailInDto(request.Msg.GetId(), *thumbnailUrl)
	outDto, err := s.updateArticleThumbnailUsecase.Execute(ctx, &inDto)
	if err != nil {
		return nil, err
	}
	response, err := s.updateArticleThumbnailConverter.ToUpdateArticleThumbnailResponse(ctx, outDto)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger.WarnContext(ctx, "END",
			slog.Group("return",
				slog.Any("grpc.BloggingEventResponse", nil),
				slog.Any("error", err)))
		return nil, err
	}
	logger.InfoContext(ctx, "END",
		slog.Group("return",
			slog.String("grpc.BloggingEventResponse", response.String())))
	return connect.NewResponse(response), nil
}

func (s *BloggingEventServiceServer) AttachTags(ctx context.Context, request *connect.Request[grpcgen.AttachTagsRequest]) (*connect.Response[grpcgen.BloggingEventResponse], error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("AttachTag").End()
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.String("article id", request.Msg.GetId()), slog.Any("attach_tag", request.Msg.GetTagNames())))

	inDto := dto.NewAttachTagsInDto(request.Msg.GetId(), request.Msg.GetTagNames())
	outDto, err := s.attachTagUsecase.Execute(ctx, &inDto)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		return nil, err
	}
	response, err := s.attachTagConverter.ToAttachTagsResponse(ctx, outDto)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger.WarnContext(ctx, "END",
			slog.Group("return",
				slog.Any("grpc.BloggingEventResponse", nil),
				slog.Any("error", err)))
		return nil, err
	}
	logger.InfoContext(ctx, "END",
		slog.Group("return",
			slog.String("grpc.BloggingEventResponse", response.String())))
	return connect.NewResponse(response), nil
}

func (s *BloggingEventServiceServer) DetachTags(ctx context.Context, request *connect.Request[grpcgen.DetachTagsRequest]) (*connect.Response[grpcgen.BloggingEventResponse], error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("DetachTag").End()
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.String("article id", request.Msg.GetId()), slog.Any("detach_tag", request.Msg.GetTagNames())))

	inDto := dto.NewDetachTagsInDto(request.Msg.GetId(), request.Msg.GetTagNames())
	outDto, err := s.detachTagUsecase.Execute(ctx, &inDto)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		return nil, err
	}
	response, err := s.detachTagConverter.ToDetachTagsResponse(ctx, outDto)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger.WarnContext(ctx, "END",
			slog.Group("return",
				slog.Any("grpc.BloggingEventResponse", nil),
				slog.Any("error", err)))
		return nil, err
	}
	logger.InfoContext(ctx, "END",
		slog.Group("return",
			slog.String("grpc.BloggingEventResponse", response.String())))
	return connect.NewResponse(response), nil
}

func (s *BloggingEventServiceServer) UploadImage(ctx context.Context, streamingServer *connect.ClientStream[grpcgen.UploadImageRequest]) (*connect.Response[grpcgen.UploadImageResponse], error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("DetachTag").End()
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN")

	var (
		binary      []byte
		fileName    string
		contentType string
	)
	buf := bytes.NewBuffer(binary)

	for streamingServer.Receive() {
		messege := streamingServer.Msg()
		if data := messege.GetData(); len(data) > 0 {
			logger.InfoContext(ctx, "Received data", slog.Group("data", slog.String("data", string(data))))
			buf.Write(data)
		}
		if meta := messege.GetMeta(); meta != nil {
			logger.InfoContext(ctx, "Received meta", slog.Group("meta", slog.String("name", meta.GetName())))
			fileName = meta.GetName()
			contentType = meta.GetContentType()
		}
	}

	if err := streamingServer.Err(); err != nil {
		err = errors.WithStack(err)
		logger.WarnContext(ctx, "END",
			slog.Group("return",
				slog.Any("grpc.UploadImageResponse", nil),
				slog.Any("error", err)))
		return nil, err
	}

	inDto := dto.NewUploadImageInDto(fileName, buf.Bytes(), contentType)
	outDto, err := s.uploadImageUsecase.Execute(ctx, &inDto)
	if err != nil {
		err = errors.WithStack(err)
		logger.WarnContext(ctx, "END",
			slog.Group("return",
				slog.Any("grpc.UploadImageResponse", nil),
				slog.Any("error", err)))
		return nil, err
	}
	response, err := s.uploadImageConverter.ToUploadImageResponse(ctx, outDto)
	if err != nil {
		err = errors.WithStack(err)
		logger.WarnContext(ctx, "END",
			slog.Group("return",
				slog.Any("grpc.UploadImageResponse", nil),
				slog.Any("error", err)))
		return nil, err
	}
	logger.InfoContext(ctx, "END",
		slog.Group("return",
			slog.Any("grpc.UploadImageResponse", response.String())))
	return connect.NewResponse(response), nil
}

type bloggingEventServiceServerConfig struct {
	createArticleUsecase            usecase.CreateArticle
	createArticleConverter          presenters.ToCreateArticleResponse
	updateArticleTitleUsecase       usecase.UpdateArticleTitle
	updateArticleTitleConverter     presenters.ToUpdateArticleTitleResponse
	updateArticleBodyUsecase        usecase.UpdateArticleBody
	updateArticleBodyConverter      presenters.ToUpdateArticleBodyResponse
	updateArticleThumbnailUsecase   usecase.UpdateArticleThumbnail
	updateArticleThumbnailConverter presenters.ToUpdateArticleThumbnailResponse
	attachTagUsecase                usecase.AttachTags
	attachTagConverter              presenters.ToAttachTagsResponse
	detachTagUsecase                usecase.DetachTags
	detachTagConverter              presenters.ToDetachTagsResponse
	uploadImageUsecase              usecase.UploadImage
	uploadImageConverter            presenters.ToUploadImageResponse
}

type BloggingEventServiceServerOption func(*bloggingEventServiceServerConfig)

func WithCreateArticleUsecase(createArticleUsecase usecase.CreateArticle) BloggingEventServiceServerOption {
	return func(c *bloggingEventServiceServerConfig) {
		c.createArticleUsecase = createArticleUsecase
	}
}

func WithCreateArticleConverter(createArticleConverter presenters.ToCreateArticleResponse) BloggingEventServiceServerOption {
	return func(c *bloggingEventServiceServerConfig) {
		c.createArticleConverter = createArticleConverter
	}
}

func WithUpdateArticleTitleUsecase(updateArticleTitleUsecase usecase.UpdateArticleTitle) BloggingEventServiceServerOption {
	return func(c *bloggingEventServiceServerConfig) {
		c.updateArticleTitleUsecase = updateArticleTitleUsecase
	}
}

func WithUpdateArticleTitleConverter(updateArticleTitleConverter presenters.ToUpdateArticleTitleResponse) BloggingEventServiceServerOption {
	return func(c *bloggingEventServiceServerConfig) {
		c.updateArticleTitleConverter = updateArticleTitleConverter
	}
}

func WithUpdateArticleBodyUsecase(updateArticleBodyUsecase usecase.UpdateArticleBody) BloggingEventServiceServerOption {
	return func(c *bloggingEventServiceServerConfig) {
		c.updateArticleBodyUsecase = updateArticleBodyUsecase
	}
}

func WithUpdateArticleBodyConverter(updateArticleBodyConverter presenters.ToUpdateArticleBodyResponse) BloggingEventServiceServerOption {
	return func(c *bloggingEventServiceServerConfig) {
		c.updateArticleBodyConverter = updateArticleBodyConverter
	}
}

func WithUpdateArticleThumbnailUsecase(updateArticleThumbnailUsecase usecase.UpdateArticleThumbnail) BloggingEventServiceServerOption {
	return func(c *bloggingEventServiceServerConfig) {
		c.updateArticleThumbnailUsecase = updateArticleThumbnailUsecase
	}
}

func WithUpdateArticleThumbnailConverter(updateArticleThumbnailConverter presenters.ToUpdateArticleThumbnailResponse) BloggingEventServiceServerOption {
	return func(c *bloggingEventServiceServerConfig) {
		c.updateArticleThumbnailConverter = updateArticleThumbnailConverter
	}
}

func WithAttachTagsUsecase(attachTagUsecase usecase.AttachTags) BloggingEventServiceServerOption {
	return func(c *bloggingEventServiceServerConfig) {
		c.attachTagUsecase = attachTagUsecase
	}
}

func WithAttachTagsConverter(attachTagConverter presenters.ToAttachTagsResponse) BloggingEventServiceServerOption {
	return func(c *bloggingEventServiceServerConfig) {
		c.attachTagConverter = attachTagConverter
	}
}

func WithDetachTagsUsecase(detachTagUsecase usecase.DetachTags) BloggingEventServiceServerOption {
	return func(c *bloggingEventServiceServerConfig) {
		c.detachTagUsecase = detachTagUsecase
	}
}

func WithDetachTagsConverter(detachTagConverter presenters.ToDetachTagsResponse) BloggingEventServiceServerOption {
	return func(c *bloggingEventServiceServerConfig) {
		c.detachTagConverter = detachTagConverter
	}
}

func WithUploadImageUsecase(uploadImageUsecase usecase.UploadImage) BloggingEventServiceServerOption {
	return func(c *bloggingEventServiceServerConfig) {
		c.uploadImageUsecase = uploadImageUsecase
	}
}

func WithUploadImageConverter(uploadImageConverter presenters.ToUploadImageResponse) BloggingEventServiceServerOption {
	return func(c *bloggingEventServiceServerConfig) {
		c.uploadImageConverter = uploadImageConverter
	}
}

func NewBloggingEventServiceServer(options ...BloggingEventServiceServerOption) *BloggingEventServiceServer {
	config := bloggingEventServiceServerConfig{}
	for _, option := range options {
		option(&config)
	}
	return &BloggingEventServiceServer{
		bloggingEventServiceServerConfig: config,
	}
}
