package common

import (
	"net/http"

	"192.168.205.151/vq2-go/go-template/internal/config"
	"192.168.205.151/vq2-go/go-template/internal/hapi"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("get-version")

func GetVersionRoute(s *hapi.Server) *echo.Route {
	return s.Router.Management.GET("/version", getVersionHandler(s))
}

func getVersionHandler(s *hapi.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, span := tracer.Start(c.Request().Context(), "getVersion")
		defer span.End()
		return c.String(http.StatusOK, config.GetFormattedBuildArgs())
	}
}
