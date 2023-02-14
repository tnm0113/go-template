package service

import (
	"context"

	"github.com/c4i/go-template/internal/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService interface {
	CreateUser(ctx context.Context, user *db.UserModel) (interface{}, error)

	UpdateUser(ctx context.Context, user *db.UserModel) (bool, error)

	DeleteUser(ctx context.Context, id string) (int, error)

	FindByUserName(ctx context.Context, username string) (*db.UserModel, error)
}

type userService struct {
	users db.UserRepository
}

var _ UserService = (*userService)(nil)

func New(mongo *mongo.Database) UserService {
	userRepo := db.NewUserRepository(mongo.Collection(db.USER_COLLECTION))
	return &userService{
		users: userRepo,
	}
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

func (us *userService) FindByUserName(ctx context.Context, username string) (*db.UserModel, error) {
	return us.users.FindByUsername(ctx, username)
}
