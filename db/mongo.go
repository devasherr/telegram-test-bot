package db

import (
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func MongoConnection(uri string) *mongo.Collection {
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	// TODO: make a function to create collections insted of hardcodeing it
	collections := client.Database("test-bot").Collection("users")
	return collections
}
