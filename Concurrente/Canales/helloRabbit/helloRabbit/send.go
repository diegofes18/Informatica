package main

import (
	"log"

	amqp "github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable (the queue will survive a broker restart)
		false,   // delete when unused
		false,   // exclusive (used by only one connection and the queue will be deleted when that connection closes)
		false,   // no-wait (the server will not respond to the method. The client should not wait for a reply method)
		nil,     // arguments (Those are provided by clients when they declare queues (exchanges) and control various optional features, such as queue length limit or TTL.)
	)
	failOnError(err, "Failed to declare a queue")

	body := "Hello World!"
	err = ch.Publish( //This method publishes a message to a specific exchange.
		//The message will be routed to queues as defined by the exchange
		//configuration and distributed to any active consumers when the transaction,
		// if any, is committed.
		"",     // exchange (Specifies the name of the exchange to publish to. The exchange name can be empty, meaning the default exchange)
		q.Name, // routing key (Specifies the routing key for the message. The routing key is used for routing messages depending on the exchange configuration.)
		false,  // mandatory (This flag tells the server how to react if the message cannot be routed to a queue. If this flag is set, the server will return an unroutable message with a Return method. If this flag is zero, the server silently drops the message.)
		false,  // immediate (This flag tells the server how to react if the message cannot be routed to a queue consumer immediately. If this flag is set, the server will return an undeliverable message with a Return method. If this flag is zero, the server will queue the message, but with no guarantee that it will ever be consumed.)
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}
