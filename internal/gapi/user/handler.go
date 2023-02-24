package user

import (
	"192.168.205.151/vq2-go/go-template/internal/service"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewHandler(svc *service.UserService) *UserHandler {
	s := &UserHandler{
		UserService: svc,
	}
	return s
}
