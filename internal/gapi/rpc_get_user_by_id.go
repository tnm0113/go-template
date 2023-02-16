package gapi

import (
	"context"

	"github.com/c4i/go-template/internal/gapi/pb"
	"github.com/rs/zerolog/log"
)

func (s *Server) GetUserById(ctx context.Context, in *pb.UserId) (*pb.UserInfo, error) {
	reqUserId := in.GetValue()
	u, err := s.UserService.FindById(ctx, reqUserId)
	if err != nil {
		log.Error().Err(err).Msg("Find user by ID error")
		return &pb.UserInfo{}, err
	}
	return &pb.UserInfo{
		Id:        u.ID.String(),
		UserName:  u.Username,
		FirstName: u.Firstname,
		LastName:  u.Lastname,
		Age:       int32(u.Age),
	}, nil
}
