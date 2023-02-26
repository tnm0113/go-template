package common

import (
	"192.168.205.151/vq2-go/go-template/internal/hapi"
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetReadyRoute(s *hapi.Server) *echo.Route {
	return s.Router.Management.GET("/ready", getReadyHandler(s))
}

// Readiness check
// This endpoint returns 200 when our Service is ready to serve traffic (i.e. respond to queries).
// Does read-only probes apart from the general server ready state.
// Note that /-/ready is typically public (and not shielded by a mgmt-secret), we thus prevent information leakage here and only return `"Ready."`.
// Structured upon https://prometheus.io/docs/prometheus/latest/management_api/
func getReadyHandler(s *hapi.Server) echo.HandlerFunc {
	return func(c echo.Context) error {

		if !s.Ready() {
			// We use 521 to indicate an error state
			// same as Cloudflare: https://support.cloudflare.com/hc/en-us/articles/115003011431#521error
			return c.String(521, "Not ready.")
		}

		// General Timeout and associated context.
		ctx, cancel := context.WithTimeout(c.Request().Context(), 4)
		defer cancel()

		err := ProbeReadiness(ctx, s.UserService.DbClient)

		// Finally return the health status according to the seen states
		if err != nil {
			// We use 521 to indicate this error state
			// same as Cloudflare: https://support.cloudflare.com/hc/en-us/articles/115003011431#521error
			return c.String(521, "Not ready.")
		}

		return c.String(http.StatusOK, "Ready.")
	}
}
