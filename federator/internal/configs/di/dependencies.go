package di

import "github.com/labstack/echo/v4"

type Dependencies struct {
	Echo *echo.Echo
}

func NewDependencies(e *echo.Echo) *Dependencies {
	return &Dependencies{
		Echo: e,
	}
}
