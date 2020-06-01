package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	// MongoConnection is the global connection of mongodb
	MongoConnection *mongo.Client
)

// ConnectToMongo connects to the mongodb server
func ConnectToMongo(uri string) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatalln(err.Error())
	} else {
		log.Println("Connected to MongoDB")
	}

	MongoConnection = client
}

// DisconnectMongo disconnects from the mongodb server
func DisconnectMongo() {
	MongoConnection.Disconnect(context.TODO())
}
