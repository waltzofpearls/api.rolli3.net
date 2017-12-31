package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type api struct {
	config *config
	logger echo.Logger
	engine *echo.Echo
}

func newAPI(c *config, l echo.Logger) *api {
	e := echo.New()
	if c.envEq(dev) {
		e.Debug = true
	}
	return &api{
		config: c,
		logger: l,
		engine: e,
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

	if c.envEq(prod) {
		return e.StartAutoTLS(c.Listen)
	} else {
		return e.StartTLS(c.Listen, c.TLSCert, c.TLSKey)
	}
}
