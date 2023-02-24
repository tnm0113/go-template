package service

import (
	"context"
	"encoding/json"

	"192.168.205.151/vq2-go/go-template/pkg/pb"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.opentelemetry.io/otel"
)

const USER_COLLECTION = "users"

var tracer = otel.Tracer("create-user-service")

func (us *UserService) CreateUser(ctx context.Context, user *pb.UserInfo) (interface{}, error) {
	_, span := tracer.Start(ctx, "create-user-service")
	defer span.End()
	coll := us.db.Collection(USER_COLLECTION)
	result, err := coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	jsonString, _ := json.Marshal(user)
	err = us.rabbit.Publish(string(jsonString), "user.create")
	if err != nil {
		log.Error().Err(err).Msg("Publish message failed")
	}
	return result.InsertedID, nil
}

func (us *UserService) UpdateUser(ctx context.Context, user *pb.UserInfo, id string) error {
	coll := us.db.Collection(USER_COLLECTION)
	return coll.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": user})
}

func (us *UserService) DeleteUser(ctx context.Context, id string) error {
	coll := us.db.Collection(USER_COLLECTION)
	return coll.Remove(ctx, bson.M{"_id": id})
}

func (us *UserService) FindByUserName(ctx context.Context, username string) (pb.UserInfo, error) {
	coll := us.db.Collection(USER_COLLECTION)
	rs := pb.UserInfo{}
	coll.Find(ctx, bson.M{"username": username}).One(&rs)
	return rs, nil
}

func (us *UserService) FindById(ctx context.Context, id string) (pb.UserInfo, error) {
	coll := us.db.Collection(USER_COLLECTION)
	rs := pb.UserInfo{}
	coll.Find(ctx, bson.M{"_id": id}).One(&rs)
	return rs, nil
}
