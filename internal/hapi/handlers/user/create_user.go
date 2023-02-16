package user

import (
	"errors"
	"net/http"

	"github.com/c4i/go-template/internal/db"
	"github.com/c4i/go-template/internal/hapi"
	"github.com/c4i/go-template/internal/types/user"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUserRoute(s *hapi.Server) *echo.Route {
	return s.Router.Root.POST("/users", createUserHandler(s))
}

func createUserHandler(s *hapi.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := &db.UserModel{}
		if err := c.Bind(u); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		insertedId, err := s.UserService.CreateUser(c.Request().Context(), u)
		if err != nil {
			return err
		}

		if oid, ok := insertedId.(primitive.ObjectID); ok {
			u.ID = oid
			res := &user.UserResponse{
				ID:        u.ID.String(),
				Username:  u.Username,
				Firstname: u.Firstname,
				Lastname:  u.Lastname,
				Age:       u.Age,
			}
			return c.JSON(http.StatusCreated, res)
		} else {
			return c.JSON(http.StatusInternalServerError, errors.New("invalid insertedId from mongodb"))
		}
	}
}
