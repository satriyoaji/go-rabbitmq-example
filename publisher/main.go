package main

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"math/rand"
	rabbitmq_example "rabbitmq-example"
	"rabbitmq-example/helpers"
	"time"
)

func main() {
	conn, err := amqp.Dial(rabbitmq_example.Config.AMQPConnectionURL)
	helpers.HandleError(err, "Can't connect to AMQP")
	defer conn.Close()

	amqpChannel, err := conn.Channel()
	helpers.HandleError(err, "Can't create a amqpChannel")

	defer amqpChannel.Close()

	queue, err := amqpChannel.QueueDeclare("add", true, false, false, false, nil)
	helpers.HandleError(err, "Could not declare `add` queue")

	rand.Seed(time.Now().UnixNano())

	addTask := rabbitmq_example.AddTask{Number1: rand.Intn(999), Number2: rand.Intn(999)}
	body, err := json.Marshal(addTask)
	if err != nil {
		helpers.HandleError(err, "Error encoding JSON")
	}

	err = amqpChannel.Publish("", queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})

	if err != nil {
		log.Fatalf("Error publishing message: %s", err)
	}

	log.Printf("AddTask: %d+%d", addTask.Number1, addTask.Number2)

}
