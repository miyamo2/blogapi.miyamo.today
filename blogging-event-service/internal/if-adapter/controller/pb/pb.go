package pb

import (
	"context"
	"github.com/cockroachdb/errors"
	"github.com/miyamo2/altnrslog"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/app/usecase/dto"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/if-adapter/controller/pb/presenters"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/if-adapter/controller/pb/usecase"
	grpcgen "github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/infra/grpc"
	"github.com/miyamo2/blogapi.miyamo.today/core/log"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"
	"google.golang.org/grpc"
	"log/slog"
	"net/url"
)

var _ grpcgen.BloggingEventServiceServer = (*BloggingEventServiceServer)(nil)

// BloggingEventServiceServer is implementation of grpc.BloggingEventServiceServer
type BloggingEventServiceServer struct {
	bloggingEventServiceServerConfig
	grpcgen.UnimplementedBloggingEventServiceServer
}

// CreateArticle is implementation of grpc.BloggingEventServiceServer.CreateArticle
func (s *BloggingEventServiceServer) CreateArticle(ctx context.Context, req *grpcgen.CreateArticleRequest) (*grpcgen.BloggingEventResponse, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("GetAllArticles").End()
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.String("title", req.GetTitle()), slog.String("body", req.GetBody()), slog.String("thumbnail", req.GetThumbnailUrl()), slog.Any("tagNames", req.GetTagNames())))

	inDto := dto.NewCreateArticleInDto(req.GetTitle(), req.GetBody(), req.GetThumbnailUrl(), req.GetTagNames())
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
			slog.Any("grpc.BloggingEventResponse", *response)))
	return response, nil
}

func (s *BloggingEventServiceServer) UpdateArticleTitle(ctx context.Context, request *grpcgen.UpdateArticleTitleRequest) (*grpcgen.BloggingEventResponse, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("UpdateArticleTitle").End()
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.String("article id", request.GetId()), slog.String("title", request.GetTitle())))

	inDto := dto.NewUpdateArticleTitleInDto(request.GetId(), request.GetTitle())
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
			slog.Any("grpc.BloggingEventResponse", *response)))
	return response, nil
}

func (s *BloggingEventServiceServer) UpdateArticleBody(ctx context.Context, request *grpcgen.UpdateArticleBodyRequest) (*grpcgen.BloggingEventResponse, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("UpdateArticleBody").End()
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.String("article id", request.GetId()), slog.String("body", request.GetBody())))

	inDto := dto.NewUpdateArticleBodyInDto(request.GetId(), request.GetBody())
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
			slog.Any("grpc.BloggingEventResponse", *response)))
	return response, nil
}

func (s *BloggingEventServiceServer) UpdateArticleThumbnail(ctx context.Context, request *grpcgen.UpdateArticleThumbnailRequest) (*grpcgen.BloggingEventResponse, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("UpdateArticleThumbnail").End()
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.String("article id", request.GetId()), slog.String("thumbnail", request.GetThumbnailUrl())))

	thumbnailUrl, err := url.Parse(request.GetThumbnailUrl())
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		return nil, err
	}
	inDto := dto.NewUpdateArticleThumbnailInDto(request.GetId(), *thumbnailUrl)
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
			slog.Any("grpc.BloggingEventResponse", *response)))
	return response, nil
}

func (s *BloggingEventServiceServer) AttachTags(ctx context.Context, request *grpcgen.AttachTagsRequest) (*grpcgen.BloggingEventResponse, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("AttachTag").End()
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN",
		slog.Group("parameters", slog.String("article id", request.GetId()), slog.Any("attach_tag", request.GetTagNames())))

	inDto := dto.NewAttachTagsInDto(request.GetId(), request.GetTagNames())
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
			slog.Any("grpc.BloggingEventResponse", *response)))
	return response, nil
}

func (s *BloggingEventServiceServer) DetachTags(ctx context.Context, request *grpcgen.DetachTagsRequest) (*grpcgen.BloggingEventResponse, error) {
	return s.UnimplementedBloggingEventServiceServer.DetachTags(ctx, request)
}

func (s *BloggingEventServiceServer) UploadImage(streamingServer grpc.ClientStreamingServer[grpcgen.UploadImageRequest, grpcgen.UploadImageResponse]) error {
	return s.UnimplementedBloggingEventServiceServer.UploadImage(streamingServer)
}

func (s *BloggingEventServiceServer) mustEmbedUnimplementedBloggingEventServiceServer() {}

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

func NewBloggingEventServiceServer(options ...BloggingEventServiceServerOption) *BloggingEventServiceServer {
	config := bloggingEventServiceServerConfig{}
	for _, option := range options {
		option(&config)
	}
	return &BloggingEventServiceServer{
		bloggingEventServiceServerConfig: config,
	}
}
