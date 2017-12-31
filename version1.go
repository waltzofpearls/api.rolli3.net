package main

import (
	"net/http"
	"time"

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

var (
	apiVersion        = "abc123"
	apiBuildTimestamp = "1900-01-01_00:00:00am"
	apiUptimeStart    = time.Now()
)

type apiStatus struct {
	UptimeSeconds  uint64 `json:"uptime_seconds"`
	BuildTimestamp string `json:"build_timestamp"`
	Version        string `json:"version"`
}

func (api *version1) getStatus(c echo.Context) error {
	return c.JSON(http.StatusOK, &apiStatus{
		UptimeSeconds:  uint64(time.Since(apiUptimeStart) / time.Second),
		BuildTimestamp: apiBuildTimestamp,
		Version:        apiVersion,
	})
}
