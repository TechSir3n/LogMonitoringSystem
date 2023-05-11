package rabbitmq

import (
	"github.com/streadway/amqp"
	"io"
	"logs-monitoring/utils"
	"os"
)

var log = utils.InitLogger()

func ProducerMQ() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("Couldn't connect to rabbitMq [Producer]", err.Error())
	} else {
		log.Info("Connected to rabbitMq Success [Producer]")
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed to open a channel,[ERROR:] ", err.Error())
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
		log.Error("Failed create Queue: ", err.Error())
	}

	file, err := os.Open("/var/log/nginx/access.log")
	if err != nil {
		log.Error("Failed to open a file", err.Error())
	}

	defer file.Close()

	var buff [512]byte

	for {
		n, err := file.Read(buff[:])
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Errorf("Failed to read file: %s ", err)
		}
		body := buff[:n]
		err = ch.Publish(
			"",
			que.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "Application/octet-stream",
				Body:        []byte(body),
			},
		)

		if err != nil {
			log.Error("Failed to publish a message: ", err.Error())
		}
	}
}
