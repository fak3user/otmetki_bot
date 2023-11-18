package db

import (
	"context"
	"data-miner/types"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateNewUserOrCheckExist(fromUser *tgbotapi.User) (bool, error) {
	collection := GetCollection("users")
	filter := bson.M{"tg_id": fromUser.ID}
	var existingUser types.User

	err := collection.FindOne(context.TODO(), filter).Decode(&existingUser)
	if err == nil {
		return false, nil
	}

	user := types.User{
		TgID: fromUser.ID,
		Name: fromUser.FirstName + " " + fromUser.LastName,
	}
	// Insert user document into MongoDB
	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		return false, err
	}

	return true, nil
}

func GetMe(userId int64) (types.User, error) {
	collection := GetCollection("users")
	filter := bson.M{"tg_id": userId}
	var user types.User

	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err == nil {
		return user, err
	}

	return user, nil
}
