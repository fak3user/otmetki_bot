package db

import (
	"context"
	"data-miner/types"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddLake(title string, userId int64) (primitive.ObjectID, error) {
	collection := GetCollection("app", "lakes")

	user, _ := GetMe(userId)

	newLake := types.Lake{
		UserId: user.ID,
		Title:  title,
	}

	insertedResult, err := collection.InsertOne(context.Background(), newLake)
	if err != nil {
		return insertedResult.InsertedID.(primitive.ObjectID), err
	}

	return insertedResult.InsertedID.(primitive.ObjectID), nil
}

func GetLakes(ids []primitive.ObjectID) ([]types.Lake, error) {
	var userLakes []types.Lake
	if len(ids) == 0 {
		return userLakes, fmt.Errorf("no lakes")
	}

	collection := GetCollection("app", "lakes")

	filter := bson.M{"_id": bson.M{"$in": ids}}

	cursor, _ := collection.Find(context.Background(), filter)

	for cursor.Next(context.Background()) {
		var lake types.Lake
		err := cursor.Decode(&lake)
		if err != nil {
			return userLakes, err
		}
		userLakes = append(userLakes, lake)
	}
	return userLakes, nil
}
