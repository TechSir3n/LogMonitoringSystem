package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

func ConsumerMQ() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Error("Failed to connect rabbitMq [Consumer]", err.Error())
	} else {
		log.Info("Success connected rabbitmq,[Consumer]]")
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Couldn't create channel", err.Error())
	}

	defer ch.Close() 
	
	que, err := ch.QueueDeclare(
		"logs",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Error("Failed to create queue [Consumer]", err.Error())
	}

	msgs, err := ch.Consume(
		que.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Error("Failed to create consume [Consumer]", err.Error())
	}

	for data := range msgs {
		fmt.Printf("Received  a message :%s ", data.Body)
	}
}
