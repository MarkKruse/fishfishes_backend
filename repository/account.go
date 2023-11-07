package repository

import (
	"context"
	"fishfishes_backend/common"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

const User string = "user"

type UserEntity struct {
	UserId   string `bson:"_id"`
	Name     string `bson:"name"`
	Password string `bson:"password"`
}

func (r Repo) CreateAccount(ctx context.Context, user common.User) error {

	userEntity := UserEntity{
		UserId:   uuid.New().String(),
		Name:     user.Name,
		Password: user.Password,
	}

	_, err := r.db.Database.Collection(User).InsertOne(ctx, userEntity)
	if err != nil {
		return err
	}

	return nil
}

func (r Repo) CheckLogin(ctx context.Context, user common.User) (bool, string) {

	filter := bson.D{
		{
			Key:   "name",
			Value: fmt.Sprintf("%s", user.Name),
		},
		{
			Key:   "password",
			Value: fmt.Sprintf("%s", user.Password),
		},
	}

	result := r.db.Database.Collection(User).FindOne(ctx, filter)
	err := result.Err()
	if err != nil {
		return false, ""
	}

	var userEntity UserEntity
	err = result.Decode(&userEntity)
	if err != nil {
		return false, ""
	}
	return true, userEntity.UserId
}
