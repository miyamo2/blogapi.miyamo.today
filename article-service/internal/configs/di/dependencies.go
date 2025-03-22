package di

import (
	"github.com/labstack/echo/v4"
	"github.com/newrelic/go-agent/v3/newrelic"
	"gorm.io/gorm"
)

type Dependencies struct {
	NewRelicApp   *newrelic.Application
	Echo          *echo.Echo
	GORMDialector *gorm.Dialector
}

func NewDependencies(
	newRelicApp *newrelic.Application,
	echoApp *echo.Echo,
	gormDialector *gorm.Dialector,
) *Dependencies {
	return &Dependencies{
		NewRelicApp:   newRelicApp,
		Echo:          echoApp,
		GORMDialector: gormDialector,
	}
}
