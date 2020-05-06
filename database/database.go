package database

import (
	"log"

	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

var (
	// RethinkSession is the global session of rethinkdb
	RethinkSession *r.Session
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
