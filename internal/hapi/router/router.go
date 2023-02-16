package router

import (
	"github.com/c4i/go-template/internal/hapi"
	"github.com/c4i/go-template/internal/hapi/handlers"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Init(s *hapi.Server) {
	s.Echo = echo.New()

	s.Echo.Debug = s.Config.EchoDebug
	s.Echo.HideBanner = true
	s.Echo.Logger.SetOutput(&echoLogger{level: zerolog.Level(s.Config.RequestLevel), log: log.With().Str("component", "echo").Logger()})

	if s.Config.EnableRecoverMiddleware {
		s.Echo.Use(echoMiddleware.Recover())
	} else {
		log.Warn().Msg("Disabling recover middleware due to environment config")
	}

	if s.Config.EnableCORSMiddleware {
		s.Echo.Use(echoMiddleware.CORS())
	} else {
		log.Warn().Msg("Disabling CORS middleware due to environment config")
	}

	s.Echo.Use(echoMiddleware.Logger())

	s.Router = &hapi.Router{
		Routes:     nil,
		Root:       s.Echo.Group(""),
		Management: s.Echo.Group("/-"),
		// API:        s.Echo.Group("/api"),
	}

	handlers.AttackAllRoutes(s)
}
