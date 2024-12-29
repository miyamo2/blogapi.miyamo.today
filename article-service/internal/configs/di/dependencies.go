package di

import (
	grpcgen "blogapi.miyamo.today/article-service/internal/infra/grpc"
	"github.com/newrelic/go-agent/v3/newrelic"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Dependencies struct {
	GRPCServer           *grpc.Server
	NewRelicApp          *newrelic.Application
	ArticleServiceServer grpcgen.ArticleServiceServer
	GORMDialector        *gorm.Dialector
}

func NewDependencies(
	grpcServer *grpc.Server,
	newRelicApp *newrelic.Application,
	articleServiceServer grpcgen.ArticleServiceServer,
	gormDialector *gorm.Dialector,
) *Dependencies {
	return &Dependencies{
		GRPCServer:           grpcServer,
		NewRelicApp:          newRelicApp,
		ArticleServiceServer: articleServiceServer,
		GORMDialector:        gormDialector,
	}
}
