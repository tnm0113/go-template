package gapi

import (
	"fmt"
	"net"

	"github.com/c4i/go-template/internal/config"
	"github.com/c4i/go-template/internal/gapi/pb"
	"github.com/c4i/go-template/internal/service"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.UnimplementedUserServiceServer
	Config      config.ServiceConfig
	UserService *service.UserService
}

func NewServer(cfg config.ServiceConfig, svc *service.UserService) *Server {
	s := &Server{
		Config:      cfg,
		UserService: svc,
	}
	return s
}

func (s *Server) Start(errs chan error) {
	grpcLogger := grpc.UnaryInterceptor(GrpcLogger)
	grpcServer := grpc.NewServer(grpcLogger)
	pb.RegisterUserServiceServer(grpcServer, s)
	reflection.Register(grpcServer)
	grpcAddress := fmt.Sprintf("%v:%d", s.Config.GrpcHost, s.Config.GrpcPort)
	listener, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot create grpc listener")
		errs <- err
	}
	log.Info().Msgf("start grpc server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot start grpc server")
		errs <- err
	}
}
