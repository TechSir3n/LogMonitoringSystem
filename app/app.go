package app 

import ( 
  "logs-monitoring/rabbitmq" 
)

func Run() { 
	rabbitmq.ProducerMQ()
	rabbitmq.ConsumerMQ()
}