package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

type api struct {
	config *config
	logger *log.Logger
	engine *echo.Echo
}

func newAPI(c *config, l *log.Logger) *api {
	return &api{
		config: c,
		logger: l,
		engine: echo.New(),
	}
}

func (a *api) run() error {
	newVersion1(a).attach()
	newVersion2(a).attach()
	return startAPIWith(a.config, a.engine)
}

type engineProvider interface {
	Use(middleware ...echo.MiddlewareFunc)
	StartAutoTLS(address string) error
	StartTLS(address string, certFile, keyFile string) error
}

var startAPIWith = func(c *config, e engineProvider) error {
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	if c.isEnv(prod) {
		return e.StartAutoTLS(c.Listen)
	} else {
		return e.StartTLS(c.Listen, c.TLSCert, c.TLSKey)
	}
}
