package hapi

import (
	"github.com/c4i/go-template/internal/config"
	"github.com/c4i/go-template/internal/service"
	"github.com/labstack/echo/v4"
)

type Router struct {
	Routes     []*echo.Route
	Root       *echo.Group
	Management *echo.Group
	User       *echo.Group
}

type Server struct {
	Echo        *echo.Echo
	Config      config.ServiceConfig
	Router      *Router
	UserService *service.UserService
}

func NewServer(svc *service.UserService, cfg config.ServiceConfig) *Server {
	s := &Server{
		Echo:        nil,
		Router:      nil,
		Config:      cfg,
		UserService: svc,
	}
	return s
}

func (s *Server) Start() error {
	return s.Echo.Start(s.Config.HttpHost)
}
