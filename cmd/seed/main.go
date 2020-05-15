package main

import (
	"log"

	"github.com/gargakshit/paperplane-server/config"
	"github.com/gargakshit/paperplane-server/database"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

func main() {
	cgf := config.GetDefaultConfig()

	log.Println("Trying to connect to RethinkDB at", cgf.DatabaseConfig.RethinkDBConfig.Address)
	database.ConnectToRethink(
		cgf.DatabaseConfig.RethinkDBConfig.Address,
		cgf.DatabaseConfig.RethinkDBConfig.Database,
	)

	log.Println("Creating the Database")
	err := r.DBCreate(cgf.DatabaseConfig.RethinkDBConfig.Database).Exec(database.RethinkSession)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Done!")

	log.Println("Creating the tables")
	err = r.TableCreate("directory", r.TableCreateOpts{
		PrimaryKey: "public_key",
	}).Exec(database.RethinkSession)
	if err != nil {
		log.Println(err.Error())
	}
	err = r.TableCreate("mq", r.TableCreateOpts{
		PrimaryKey: "public_key",
	}).Exec(database.RethinkSession)
	if err != nil {
		log.Println(err.Error())
	}

	log.Println("Creating the indexes")
	r.Table("directory").IndexCreate("id").Run(database.RethinkSession)

	log.Println("Done!")
}
