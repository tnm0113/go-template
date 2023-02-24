package user

import (
	"net/http"

	"192.168.205.151/vq2-go/go-template/internal/hapi"
	"192.168.205.151/vq2-go/go-template/internal/types"
	"192.168.205.151/vq2-go/go-template/pkg/pb"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("create-user-handler")

func CreateUserRoute(s *hapi.Server) *echo.Route {
	return s.Router.Root.POST("/users", createUserHandler(s))
}

// CreateUser godoc
//
//	@Summary		Create a user
//	@Description	Create a new user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		pb.UserInfo	true	"User"
//	@Success		201		{object}	db.UserModel
//	@Failure		400		{object}	types.ErrorResponse
//	@Router			/users [post]
func createUserHandler(s *hapi.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, span := tracer.Start(c.Request().Context(), "create-user-handler")
		defer span.End()
		u := &pb.UserInfo{}
		if err := c.Bind(u); err != nil {
			log.Error().Msgf("Bind user error %v", err)
			er := types.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
			return c.JSON(http.StatusInternalServerError, er)
		}

		u.ID = uuid.NewString()

		_, err := s.UserService.CreateUser(c.Request().Context(), u)
		if err != nil {
			er := types.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
			return c.JSON(http.StatusInternalServerError, er)
		}

		return c.JSON(http.StatusCreated, u)

	}
}
