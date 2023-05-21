package app

import (
	_ "logs-monitoring/api/routers"
	"logs-monitoring/config"
	"logs-monitoring/mongodb"
	"logs-monitoring/rabbitmq"
	"logs-monitoring/utils"
	"time"
)

var logger = config.InitLogger()

func Run() {
	rabbitmq.ProducerMQ()
	rabbitmq.ConsumerMQ()

	if utils.ClearLogFiles() {
		logger.Info("Files log was cleared")
	}

	period := time.Hour * 12
	for {
		mongodb.DeleteLogs()
		time.Sleep(period)
	}
}
