package service

import (
	"192.168.205.151/vq2-go/go-template/internal/config"
	"192.168.205.151/vq2-go/go-template/internal/mq"
	"github.com/qiniu/qmgo"
)

type UserService struct {
	DbClient *qmgo.Client
	db       *qmgo.Database
	rabbit   *mq.RabbitMQ
}

func New(dbClient *qmgo.Client, rb *mq.RabbitMQ, cfg config.ServiceConfig) *UserService {
	db := dbClient.Database(cfg.DbConfig.DBName)
	return &UserService{
		DbClient: dbClient,
		db:       db,
		rabbit:   rb,
	}
}
