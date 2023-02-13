package common

import (
	"net/http"

	"github.com/c4i/go-template/internal/config"
	"github.com/c4i/go-template/internal/hapi"
	"github.com/labstack/echo/v4"
)

func GetVersionRoute(s *hapi.Server) *echo.Route {
	return s.Router.Management.GET("/version", getVersionHandler(s))
}

func getVersionHandler(s *hapi.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, config.GetFormattedBuildArgs())
	}
}
