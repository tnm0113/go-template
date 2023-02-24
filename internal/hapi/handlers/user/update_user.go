package user

import (
	"net/http"

	"192.168.205.151/vq2-go/go-template/internal/hapi"
	"192.168.205.151/vq2-go/go-template/internal/types"
	"192.168.205.151/vq2-go/go-template/pkg/pb"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func UpdateUserRoute(s *hapi.Server) *echo.Route {
	return s.Router.Root.PUT("/users/:id", updateUserHandler(s))
}

func updateUserHandler(s *hapi.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := &pb.UserInfo{}
		if err := c.Bind(u); err != nil {
			log.Error().Msgf("Bind user error %v", err)
			er := types.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
			return c.JSON(http.StatusInternalServerError, er)
		}

		id := c.Param("id")
		log.Debug().Msgf("update user %v", u)
		err := s.UserService.UpdateUser(c.Request().Context(), u, id)
		if err != nil {
			log.Error().Msgf("Update user error")
			er := types.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
			return c.JSON(http.StatusInternalServerError, er)
		}

		res := types.SucceedResponse{
			Success: true,
		}
		return c.JSON(http.StatusOK, res)
	}
}
