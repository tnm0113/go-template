package hapi

import "github.com/labstack/echo/v4"

type Router struct {
	Routes []*echo.Route
	Root   *echo.Group
}

type Server struct {
	Echo   *echo.Echo
	Router *Router
}

func NewServer() *Server {
	
	return nil
}
