package di

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	grpcgen "github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/infra/grpc"
	"github.com/newrelic/go-agent/v3/newrelic"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Dependencies struct {
	AWSConfig                  *aws.Config
	GRPCServer                 *grpc.Server
	NewRelicApp                *newrelic.Application
	BloggingEventServiceServer grpcgen.BloggingEventServiceServer
	GORMDialector              *gorm.Dialector
}

func NewDependencies(
	awsConfig *aws.Config,
	grpcServer *grpc.Server,
	newRelicApp *newrelic.Application,
	bloggingEventServiceServer grpcgen.BloggingEventServiceServer,
	gormDialector *gorm.Dialector,
) *Dependencies {
	return &Dependencies{
		AWSConfig:                  awsConfig,
		GRPCServer:                 grpcServer,
		NewRelicApp:                newRelicApp,
		BloggingEventServiceServer: bloggingEventServiceServer,
		GORMDialector:              gormDialector,
	}
}
