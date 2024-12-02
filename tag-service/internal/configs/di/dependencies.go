package di

import (
	grpcgen "github.com/miyamo2/blogapi.miyamo.today/tag-service/internal/infra/grpc"
	"github.com/newrelic/go-agent/v3/newrelic"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Dependencies struct {
	GRPCServer       *grpc.Server
	NewRelicApp      *newrelic.Application
	TagServiceServer grpcgen.TagServiceServer
	GORMDialector    *gorm.Dialector
}

func NewDependencies(
	grpcServer *grpc.Server,
	newRelicApp *newrelic.Application,
	tagServiceServer grpcgen.TagServiceServer,
	gormDialector *gorm.Dialector,
) *Dependencies {
	return &Dependencies{
		GRPCServer:       grpcServer,
		NewRelicApp:      newRelicApp,
		TagServiceServer: tagServiceServer,
		GORMDialector:    gormDialector,
	}
}
