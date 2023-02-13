package service

import (
	"context"
	"fmt"
	"os"

	"github.com/c4i/go-template/internal/config"
	"github.com/c4i/go-template/internal/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserService interface {
	CreateUser(ctx context.Context, user *db.UserModel) (interface{}, error)

	UpdateUser(ctx context.Context, user *db.UserModel) (bool, error)

	DeleteUser(ctx context.Context, id string) (int, error)
}

type userService struct {
	users db.UserRepository
}

var _ UserService = (*userService)(nil)

func New(dbConfig config.MongoDB) UserService {
	mongo := connectToMongoDB(dbConfig)
	userRepo := db.NewUserRepository(mongo.Collection(db.USER_COLLECTION))
	return &userService{
		users: userRepo,
	}
}

func connectToMongoDB(config config.MongoDB) *mongo.Database {
	addr := fmt.Sprintf("mongodb://%s:%d/?replicaset=%s", config.Host, config.Port, config.Replica)
	credential := options.Credential{
		Username: config.Username,
		Password: config.Password,
	}
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(addr).SetAuth(credential))
	if err != nil {
		os.Exit(1)
	}

	return client.Database(config.DbName)
}

func (us *userService) CreateUser(ctx context.Context, user *db.UserModel) (interface{}, error) {
	insertedID, err := us.users.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return insertedID, nil
}

func (us *userService) UpdateUser(ctx context.Context, user *db.UserModel) (bool, error) {
	return us.users.UpdateByID(ctx, user, user.ID)
}

func (us *userService) DeleteUser(ctx context.Context, id string) (int, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return -1, err
	}
	return us.users.DeleteByID(ctx, objID)
}
