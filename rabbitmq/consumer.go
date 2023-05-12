package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
)

func ConsumerMQ() {
	connections := make([]*amqp.Connection, 3)
	for i := range connections {
		conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
		connections[i] = conn
		if err != nil {
			log.Fatal("Couldn't connect to rabbitMq [Producer]", err.Error())
		} else {
			log.Info("Connected to rabbitMq Success [Producer]")
		}
	}

	channels := make([]*amqp.Channel, len(connections)*3)
	for i, conn := range connections {
		for j := 0; j < 3; j++ {
			ch, err := conn.Channel()
			if err != nil {
				log.Errorf("Failed to create channel: %s", err)
			}
			channels[i*3+j] = ch
		}
	}

	queueNames := []string{"logs1", "logs2", "logs3"}
	for _, queueName := range queueNames {
		for _, ch := range channels {
			msgs, err := ch.Consume(
				queueName,
				"",
				true,
				false,
				false,
				false,
				nil,
			)
			if err != nil {
				log.Errorf("Failed to consume from queue %s: %s", queueName, err)
			}

			for d := range msgs {
				fmt.Printf("Received a message: %s", d.Body)
			}
		}
	}
}
