package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const USER_COLLECTION = "users"

type UserModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username  string             `bson:"username" json:"username"`
	Firstname string             `bson:"firstname" json:"firstname"`
	Lastname  string             `bson:"lastname" json:"lastname"`
	Age       int                `bson:"age" json:"age"`
}

//go:generate repogen -src=user.go -dest=user_repo.go -model=UserModel -repo=UserRepository

// UserRepository is an interface that describes the specification of querying
// user data in the database.
type UserRepository interface {
	// InsertOne stores userModel into the database and returns inserted ID
	// if insertion succeeds and returns error if insertion fails.
	InsertOne(ctx context.Context, userModel *UserModel) (interface{}, error)

	// FindByUsername queries user by username. If a user with specified
	// username exists, the user will be returned. Otherwise, error will be
	// returned.
	FindByUsername(ctx context.Context, username string) (*UserModel, error)

	// UpdateByID updates a single document by ID
	UpdateByID(ctx context.Context, model *UserModel, id primitive.ObjectID) (bool, error)

	// DeleteByID deletes users that have `id` value match the parameter
	// and returns the match count. The error will be returned only when
	// error occurs while accessing the database. This is a MANY mode
	// because the first return type is an integer.
	DeleteByID(ctx context.Context, id primitive.ObjectID) (int, error)
}
