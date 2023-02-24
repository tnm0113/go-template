package user

import (
	"context"

	"192.168.205.151/vq2-go/go-template/pkg/pb"
	"github.com/rs/zerolog/log"
)

func (h *UserHandler) GetUserById(ctx context.Context, in *pb.UserId) (*pb.UserInfo, error) {
	reqUserId := in.GetValue()
	u, err := h.UserService.FindById(ctx, reqUserId)
	if err != nil {
		log.Error().Err(err).Msg("Find user by ID error")
		return &pb.UserInfo{}, err
	}
	return &u, nil
}
