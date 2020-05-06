package main

import (
	"log"

	"github.com/gargakshit/paperplane-server/config"
	"github.com/gargakshit/paperplane-server/database"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

func main() {
	config := config.GetDefaultConfig()

	log.Println("Trying to connect to RethinkDB at", config.DatabaseConfig.RethinkDBConfig.Address)
	database.ConnectToRethink(
		config.DatabaseConfig.RethinkDBConfig.Address,
		config.DatabaseConfig.RethinkDBConfig.Database,
	)

	log.Println("Creating the Database")
	err := r.DBCreate(config.DatabaseConfig.RethinkDBConfig.Database).Exec(database.RethinkSession)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("Done!")

	log.Println("Creating the tables")
	err = r.TableCreate("directory", r.TableCreateOpts{
		PrimaryKey: "public_key",
	}).Exec(database.RethinkSession)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("Done!")
}
