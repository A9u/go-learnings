package main

import (
	"bufio"
	"github.com/streadway/amqp"
	"log"
	"os"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	queue, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	failOnError(err, "Failed to declare a queue")

	body := "hello. did I send you?"

	err = ch.Publish("", queue.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})

	failOnError(err, "Failed to publish a message")
	log.Printf("Sent first message successfuly")
	log.Printf("Use Ctrl + c to close.")

	scanner := bufio.NewScanner(os.Stdin)

	log.Println("Enter message:")
	for scanner.Scan() {
		log.Println("sending message")
		err = ch.Publish("", queue.Name, false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body:        scanner.Bytes(),
		})

		log.Println("sent message:", scanner.Text())
		log.Println("Enter message:")
	}
}
