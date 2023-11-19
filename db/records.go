package db

import (
	"context"
	"data-miner/types"
	"data-miner/utils"
	"time"

	"github.com/davecgh/go-spew/spew"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddRecord(lakeId string) (primitive.ObjectID, error) {
	collection := GetCollection("records", lakeId)

	record := types.Record{
		CreatedAt: primitive.Timestamp{
			T: uint32(time.Now().Unix()),
		},
	}

	inserted, err := collection.InsertOne(context.Background(), record)
	if err != nil {
		return inserted.InsertedID.(primitive.ObjectID), err
	}

	return inserted.InsertedID.(primitive.ObjectID), nil
}

func UpdateRecord(lakeId string, recordId string, updatedRecord types.Record) error {
	collection := GetCollection("records", lakeId)

	recordIdConverted, _ := primitive.ObjectIDFromHex(recordId)

	doc, _ := utils.StructToBson(updatedRecord)

	filter := bson.M{"_id": recordIdConverted}
	update := bson.M{"$set": doc}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	spew.Dump(err)

	return nil
}
