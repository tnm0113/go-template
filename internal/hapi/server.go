package hapi

import (
	"errors"
	"fmt"

	"192.168.205.151/vq2-go/go-template/internal/config"
	"192.168.205.151/vq2-go/go-template/internal/i18n"
	"192.168.205.151/vq2-go/go-template/internal/service"
	"github.com/labstack/echo/v4"
)

type Router struct {
	Routes     []*echo.Route
	Root       *echo.Group
	Management *echo.Group
}

type Server struct {
	Echo        *echo.Echo
	Config      config.ServiceConfig
	Router      *Router
	I18n        *i18n.Service
	UserService *service.UserService
}

func NewServer(svc *service.UserService, cfg config.ServiceConfig) *Server {
	s := &Server{
		Echo:        nil,
		Router:      nil,
		Config:      cfg,
		UserService: svc,
		I18n:        nil,
	}
	return s
}

func (s *Server) Ready() bool {
	return s.Echo != nil &&
		s.Router != nil &&
		s.UserService != nil
}

func (s *Server) InitI18n() error {
	i18nService, err := i18n.New(s.Config)

	if err != nil {
		return err
	}

	s.I18n = i18nService

	return nil
}

func (s *Server) Start(errs chan error) {
	if !s.Ready() {
		errs <- errors.New("server is not ready")
	}
	httpAddress := fmt.Sprintf("%s:%d", s.Config.HttpConfig.HttpHost, s.Config.HttpConfig.HttpPort)
	errs <- s.Echo.Start(httpAddress)
}
