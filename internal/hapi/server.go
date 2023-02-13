package hapi

import (
	"github.com/c4i/go-template/internal/config"
	"github.com/labstack/echo/v4"
)

type Router struct {
	Routes     []*echo.Route
	Root       *echo.Group
	Management *echo.Group
}

type Server struct {
	Echo   *echo.Echo
	Config config.HttpServer
	Router *Router
}

func NewServer() *Server {

	return nil
}
