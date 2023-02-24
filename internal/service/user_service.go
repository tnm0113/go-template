package service

import (
	"192.168.205.151/vq2-go/go-template/internal/mq"
	"github.com/qiniu/qmgo"
)

type UserService struct {
	db     *qmgo.Database
	rabbit *mq.RabbitMQ
}

func New(db *qmgo.Database, rb *mq.RabbitMQ) *UserService {
	return &UserService{
		db:     db,
		rabbit: rb,
	}
}
