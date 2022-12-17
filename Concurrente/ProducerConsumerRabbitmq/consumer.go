package main

import (
	"log"
        "time"
        "math/rand"
	amqp "github.com/streadway/amqp"
)
const (
        Productores = 4
        Consumidores = 4
        Nmensajes = 4

)

var prodnames = [Productores]string{"Ferran","Miquel","Tomeu","Toni"}
var consnames = [Consumidores]string{"Sara", "Maria","Gloria","Xisca"}


func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func consumer(id int, name string, ch *amqp.Channel, q amqp.Queue) {
	// Consume messages
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	for d := range msgs {
		log.Printf("Consumer %d , %s Received a message: %s", id, name, d.Body)
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
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


	forever := make(chan bool)

        for i := 0; i < Consumidores; i++ {
		go consumer(i, consnames[i],ch, q)
	}


	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
