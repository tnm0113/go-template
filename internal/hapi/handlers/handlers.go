package handlers

import (
	"github.com/c4i/go-template/internal/hapi"
	"github.com/c4i/go-template/internal/hapi/handlers/common"
	"github.com/labstack/echo/v4"
)

func AttackAllRoutes(s *hapi.Server) {
	s.Router.Routes = []*echo.Route{
		common.GetVersionRoute(s),
	}
}
