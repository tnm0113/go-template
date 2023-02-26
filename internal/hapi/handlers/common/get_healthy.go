package common

import (
	"192.168.205.151/vq2-go/go-template/internal/hapi"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func GetHealthyRoute(s *hapi.Server) *echo.Route {
	return s.Router.Management.GET("/healthy", getHealthyHandler(s))
}

func getHealthyHandler(s *hapi.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		if !s.Ready() {
			// We use 521 to indicate an error state
			// same as Cloudflare: https://support.cloudflare.com/hc/en-us/articles/115003011431#521error
			return c.String(521, "Not ready.")
		}

		var str strings.Builder

		ctx, cancel := context.WithTimeout(c.Request().Context(), 5)
		defer cancel()

		err := ProbeLiveness(ctx, s.UserService.DbClient)
		if err != nil {
			return c.String(521, "Not ready.")
		}

		fmt.Fprintln(&str, "Probes succeeded.")

		return c.String(http.StatusOK, str.String())
	}
}
