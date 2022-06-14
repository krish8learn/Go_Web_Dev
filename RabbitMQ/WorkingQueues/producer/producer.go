package main

import (
	"log"
	"os"
	"strings"

	"github.com/streadway/amqp"
)

func main() {
	log.Println("RabbitMQ tutorial Producer starting")

	//connect to rabbit mq server//
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	//create a channel//
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	//declare a queue//
	Q, err := ch.QueueDeclare(
		"TestWorkingQueue", // name
		true,               // durable [to stop erasing queue once RabbitMQ stops]
		false,              // delete when unused
		false,              // exclusive
		false,              // no-wait
		nil,                // arguments
	)
	failOnError(err, "Failed to declare a queue")

	//create a message//
	body := messageFromUser(os.Args)
	err = ch.Publish(
		"",     // exchange
		Q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")

	log.Printf("Sent Message %s", body)

	log.Println("Message published successfully")

}

//will be used for error handling//
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

//will be used for preaparing message from input//
func messageFromUser(args []string) string {

	var str string
	if len(args) < 2 || os.Args[1] == "" {
		str = "Default Hello"
	} else {
		str = strings.Join(args[1:], " ")
	}

	return str
}
