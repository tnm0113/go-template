package user

import (
	"fmt"
	"net/http"

	"github.com/c4i/go-template/internal/hapi"
	"github.com/c4i/go-template/internal/types/user"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func GetUserRoute(s *hapi.Server) *echo.Route {
	return s.Router.Root.GET("/users", getUserHandlers(s))
}

func getUserHandlers(s *hapi.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.QueryParam("username")
		ctx := c.Request().Context()
		u, err := s.UserService.FindByUserName(ctx, username)
		fmt.Printf("Find user %v \n", u)
		if err != nil {
			log.Error().Err(err).Msg("FindByUserName error")
			return c.JSON(http.StatusNotFound, err)
		}

		res := &user.UserResponse{
			ID:        u.ID.String(),
			Username:  u.Username,
			Firstname: u.Firstname,
			Lastname:  u.Lastname,
			Age:       u.Age,
		}

		return c.JSON(http.StatusOK, res)
	}
}
