package app

import (
	"logs-monitoring/mongodb"
	"logs-monitoring/rabbitmq"
	_ "logs-monitoring/api/routers"
	"time"
)

func Run() {
	rabbitmq.ProducerMQ()
	rabbitmq.ConsumerMQ()

	period := time.Hour * 12
	for {
		mongodb.DeleteLogs()
		time.Sleep(period)
	}
}
