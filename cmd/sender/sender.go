package main

import (
	"context"
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	// Opening rabbitMQ connection and deferring it's closure
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	// Initializing RMQ channel for messaging
	ch, err := conn.Channel()
	failOnError(err, "Failed to open RabbitMQ channel")
	defer func() {
		_ = ch.Close()
	}()

	// Initializing queue
	q, err := ch.QueueDeclare(
		"senderQueue", // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	failOnError(err, "Failed to declare a queue")

	_, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s, %s", msg, err)
	}

	failOnError(err, "Failed to")
}
