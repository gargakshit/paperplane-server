package database

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

var (
	// RethinkSession is the global session of rethinkdb
	RethinkSession *r.Session

	// MongoConnection is the global connection of mongodb
	MongoConnection *mongo.Client
)

// ConnectToRethink connects to the rethinkdb server
func ConnectToRethink(rethinkAddress string, database string) {
	session, err := r.Connect(r.ConnectOpts{
		Address:  rethinkAddress,
		Database: database,
	})

	if err != nil {
		log.Fatalln(err.Error())
	} else {
		log.Println("Connected to RethinkDB")
	}

	RethinkSession = session
}

// ConnectToMongo connects to the mongodb server
func ConnectToMongo(uri string) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatalln(err.Error())
	} else {
		log.Println("Connected to MongoDB")
	}

	MongoConnection = client
}
