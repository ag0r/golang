package main

import (
	"log"
	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if (err != nil) {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	//conn, err := amqp.Dial("amqp://guest:guest@rabbit:5672/")
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	body := "Hello world!"
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing {
			ContentType: "text/plain",
			Body:	     []byte(body),
		})
	failOnError(err, "Failed to publish message")

	log.Printf(" [x] Sent %s\n", body)
}
