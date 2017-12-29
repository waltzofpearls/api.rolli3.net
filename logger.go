package main

import "github.com/labstack/gommon/log"

func newLogger(config *config) *log.Logger {
	logger := log.New("api")
	level := log.ERROR
	if config.isEnv(dev) {
		level = log.DEBUG
	}
	logger.SetLevel(level)
	return logger
}
