package router

import (
	"192.168.205.151/vq2-go/go-template/internal/hapi"
	"192.168.205.151/vq2-go/go-template/internal/hapi/handlers"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	_ "192.168.205.151/vq2-go/go-template/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
)

func Init(s *hapi.Server) {
	s.Echo = echo.New()

	s.Echo.GET("/swagger/*", echoSwagger.WrapHandler)

	s.Echo.Debug = s.Config.HttpConfig.EchoDebug
	s.Echo.HideBanner = true
	s.Echo.Logger.SetOutput(&echoLogger{level: zerolog.Level(s.Config.LoggerConfig.RequestLevel), log: log.With().Str("component", "echo").Logger()})

	if s.Config.HttpConfig.EnableRecoverMiddleware {
		s.Echo.Use(echoMiddleware.Recover())
	} else {
		log.Warn().Msg("Disabling recover middleware due to environment config")
	}

	if s.Config.HttpConfig.EnableCORSMiddleware {
		s.Echo.Use(echoMiddleware.CORS())
	} else {
		log.Warn().Msg("Disabling CORS middleware due to environment config")
	}

	s.Echo.Use(echoMiddleware.Logger())
	s.Echo.Use(otelecho.Middleware("init-router"))

	s.Router = &hapi.Router{
		Routes:     nil,
		Root:       s.Echo.Group(""),
		Management: s.Echo.Group("/-"),
	}

	handlers.AttackAllRoutes(s)
}
