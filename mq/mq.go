package mq

import (
	"log"

	"github.com/streadway/amqp"
)

var (
	// MQConnection is the global AMQP connection
	MQConnection *amqp.Connection

	// MQChannel is the global AMQP channel
	MQChannel *amqp.Channel
)

// ConnectToMQ connects to the AMQP server
func ConnectToMQ(uri string) {
	conn, err := amqp.Dial(uri)
	if err != nil {
		log.Fatalln(err.Error())
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalln(err.Error())
	}

	MQConnection = conn
	MQChannel = ch
}

// CleanupMQ cleans up the MQ connection
func CleanupMQ() {
	MQChannel.Close()
	MQConnection.Close()
}
