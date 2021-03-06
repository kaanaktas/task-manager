package config

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewEchoEngine() *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	// Middlewares
	e.Use(middleware.RequestID())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	return e
}
