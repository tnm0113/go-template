package user

import (
	"net/http"

	"192.168.205.151/vq2-go/go-template/internal/hapi"
	"192.168.205.151/vq2-go/go-template/internal/types"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func DeleteUserRoute(s *hapi.Server) *echo.Route {
	return s.Router.Root.DELETE("/users/:id", deleteUserHandler(s))
}

func deleteUserHandler(s *hapi.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := c.Param("id")
		err := s.UserService.DeleteUser(c.Request().Context(), uid)
		if err != nil {
			log.Error().Err(err).Msg("DeleteUserById error")
			return c.JSON(http.StatusInternalServerError, err)
		}
		res := types.SucceedResponse{
			Success: true,
		}
		return c.JSON(http.StatusOK, res)
	}
}
