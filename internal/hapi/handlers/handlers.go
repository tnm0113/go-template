package handlers

import (
	"192.168.205.151/vq2-go/go-template/internal/hapi"
	"192.168.205.151/vq2-go/go-template/internal/hapi/handlers/common"
	"192.168.205.151/vq2-go/go-template/internal/hapi/handlers/user"
	"github.com/labstack/echo/v4"
)

func AttackAllRoutes(s *hapi.Server) {
	s.Router.Routes = []*echo.Route{
		// GET /-/version
		common.GetVersionRoute(s),
		// GET /ready
		common.GetReadyRoute(s),
		// GET /healthy
		common.GetHealthyRoute(s),
		// GET /users?username=abc
		user.GetUserRoute(s),
		// GET /users/:id
		user.GetUserByIdRoute(s),
		// POST /users
		user.CreateUserRoute(s),
		// DELETE /users/:id
		user.DeleteUserRoute(s),
		// PUT /users/:id
		user.UpdateUserRoute(s),
	}
}
