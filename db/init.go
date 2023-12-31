package db

import (
	"context"
	"data-miner/utils"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func GetClientOptions() *options.ClientOptions {
	envs, err := utils.GetEnvs()
	if err != nil {
		log.Fatal(err)
	}

	dburi := envs.DbUri

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(dburi).
		SetServerAPIOptions(serverAPIOptions)

	return clientOptions
}

func GetCollection(db string, collection string) *mongo.Collection {
	client := GetMongoClient()

	return client.Database(db).Collection(collection)
}

func InitDb() {
	clientOptions := GetClientOptions()

	newClient, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	} else {
		client = newClient
	}
}

func GetMongoClient() mongo.Client {
	return *client
}
