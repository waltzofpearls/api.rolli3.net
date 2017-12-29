package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

type api struct {
	config *config
	engine *echo.Echo
	logger *log.Logger
}

func newAPI(c *config, l *log.Logger) *api {
	return &api{
		config: c,
		logger: l,
		engine: echo.New(),
	}
}

func (a *api) run() error {
	a.engine.Use(middleware.Recover())
	a.engine.Use(middleware.Logger())
	a.engine.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	a.engine.Logger.SetLevel(log.DEBUG)
	if a.config.isEnv(prod) {
		return a.engine.StartAutoTLS(":9443")
	}
	return a.engine.StartTLS(a.config.Listen, a.config.TLSCert, a.config.TLSKey)
}
