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
		v1.GET("/projects", api.getProjects)
		v1.GET("/contributions", api.getContributions)
		v1.GET("/activities", api.getActivities)
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

func (api *version1) getProjects(c echo.Context) error {
	repos, err := api.github.getRepos()
	if err != nil {
		return err
	}
	c.JSON(http.StatusOK, repos)
	return nil
}

func (api *version1) getContributions(c echo.Context) error {
	return nil
}

func (api *version1) getActivities(c echo.Context) error {
	return nil
}
