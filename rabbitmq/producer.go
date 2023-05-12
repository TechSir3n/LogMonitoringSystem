package rabbitmq

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"logs-monitoring/models"
	"logs-monitoring/smtp"
	"logs-monitoring/utils"
	"os"
)

var log = utils.InitLogger()

func ProducerMQ() {
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
				log.Fatal("Failed to  create channel,", err.Error())
			}
			channels[i*3+j] = ch
		}
	}

	queueNames := []string{"logs1", "logs2", "logs3"}
	for _, queuName := range queueNames {
		for _, ch := range channels {
			_, err := ch.QueueDeclare(
				queuName,
				true,
				false,
				false,
				false,
				nil,
			)

			if err != nil {
				log.Error("Failed to create queueu", err.Error())
			}
		}
	}

	file, err := os.Open("/var/log/nginx/error.log")
	if err != nil {
		log.Error("Failed to open a file", err.Error())
	}

	defer file.Close()

	currentIndex := 0
	var obj models.LogEntry
	var entry models.LogEntry
	decoder := json.NewDecoder(file)

	for {
		if err := decoder.Decode(&entry); err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Error("Failed to decode struct", err.Error())
			continue
		}

		if entry.Level == "error" {
			smtp.SendNotification(obj.Level, entry.Message)
		} else if entry.Level == "critical" {
			smtp.SendNotification(obj.Level, entry.Message)
		}

		ch := channels[currentIndex%len(channels)]
		queName := queueNames[currentIndex%len(queueNames)]
		jsonData, err := json.Marshal(entry)
		if err != nil {
			log.Fatal("Failed to serialize the entry to JSON:", err.Error())
		}

		err = ch.Publish(
			"",
			queName,
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        jsonData,
			},
		)

		if err != nil {
			log.Error("Failed to publish a message: ", err.Error())
			continue
		}
		currentIndex++
	}
}
