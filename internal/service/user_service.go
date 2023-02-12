package service

import "context"

type Service interface {
	CreateUser(ctx context.Context)
}

type UserService struct {
}
