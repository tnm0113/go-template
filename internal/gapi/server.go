package gapi

import (
	"fmt"
	"net"

	"192.168.205.151/vq2-go/go-template/internal/config"
	"192.168.205.151/vq2-go/go-template/internal/gapi/user"
	"192.168.205.151/vq2-go/go-template/internal/service"
	"192.168.205.151/vq2-go/go-template/pkg/pb"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
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
	// embedded logger to grpc server
	grpcServer := grpc.NewServer(grpcLogger)
	// Register pb user service with grpc server
	userHandler := user.NewHandler(s.UserService)
	pb.RegisterUserServiceServer(grpcServer, userHandler)
	reflection.Register(grpcServer)
	grpcAddress := fmt.Sprintf("%v:%d", s.Config.GrpcConfig.GrpcHost, s.Config.GrpcConfig.GrpcPort)
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
