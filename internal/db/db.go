package db

import (
	"context"
	"fmt"
	"os"

	"github.com/c4i/go-template/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongoDB(config config.ServiceConfig) (*mongo.Database, error) {
	addr := fmt.Sprintf("mongodb://%s:%d/?replicaset=%s", config.DBHost, config.DBPort, config.DBReplica)
	credential := options.Credential{
		Username: config.DBUser,
		Password: config.DBPass,
	}
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(addr).SetAuth(credential))
	if err != nil {
		os.Exit(1)
	}

	return client.Database(config.DBName), nil
}
