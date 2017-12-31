package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func newLogger(c *config) echo.Logger {
	logger := log.New("api")
	if c.envEq(dev) {
		logger.SetLevel(log.DEBUG)
	}
	return logger
}
