package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type version1 struct {
	*api
}

func newVersion1(a *api) *version1 {
	return &version1{a}
}

func (api *version1) attach() {
	v1 := api.engine.Group("/v1")
	{
		v1.GET("/status", api.getStatus)
	}
}

func (api *version1) getStatus(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
