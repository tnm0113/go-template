package hapi

import (
	"errors"
	"fmt"

	"github.com/c4i/go-template/internal/config"
	"github.com/c4i/go-template/internal/service"
	"github.com/labstack/echo/v4"
)

type Router struct {
	Routes     []*echo.Route
	Root       *echo.Group
	Management *echo.Group
	// API        *echo.Group
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

func (s *Server) Ready() bool {
	return s.Echo != nil &&
		s.Router != nil &&
		s.UserService != nil
}

func (s *Server) Start(errs chan error) {
	if !s.Ready() {
		errs <- errors.New("server is not ready")
	}
	httpAddress := fmt.Sprintf("%s:%d", s.Config.HttpConfig.HttpHost, s.Config.HttpConfig.HttpPort)
	errs <- s.Echo.Start(httpAddress)
}
