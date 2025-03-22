package di

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/labstack/echo/v4"
	"github.com/newrelic/go-agent/v3/newrelic"
	"gorm.io/gorm"
)

type Dependencies struct {
	AWSConfig     *aws.Config
	NewRelicApp   *newrelic.Application
	Echo          *echo.Echo
	GORMDialector *gorm.Dialector
}

func NewDependencies(
	awsConfig *aws.Config,
	newRelicApp *newrelic.Application,
	echoApp *echo.Echo,
	gormDialector *gorm.Dialector,
) *Dependencies {
	return &Dependencies{
		AWSConfig:     awsConfig,
		NewRelicApp:   newRelicApp,
		Echo:          echoApp,
		GORMDialector: gormDialector,
	}
}
