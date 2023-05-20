package mongodb

import (
	"context"
	//"go.mongodb.org/mongo-driver/bson"
	"logs-monitoring/config"
)

var logger = config.InitLogger()

func AddLog(log []byte) error {
	ctx := context.TODO()
	client := config.ConnectToMongoDB()
	collection := client.Database("LogsMonitoring").Collection("Logs")
	_, err := collection.InsertOne(ctx, log)
	if err != nil {
		logger.Error(err.Error())
	}

	defer client.Disconnect(ctx)

	return nil
}
