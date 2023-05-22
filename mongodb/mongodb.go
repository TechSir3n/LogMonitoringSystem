package mongodb

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"logs-monitoring/config"
	"logs-monitoring/models"
)

var logger = config.InitLogger()

func AddLog(log []byte) {
	ctx := context.TODO()
	client := config.ConnectToMongoDB()
	collection := client.Database("LogsMonitoring").Collection("Logs")

	var logs models.LogFormatJSON
	if err := json.Unmarshal(log, &logs); err != nil {
		logger.Error("Couldn't unmarshal log's structure", err.Error())
	}

	result, err := collection.InsertOne(ctx, logs)
	if err != nil {
		logger.Error(err.Error())
	} else {
		logger.Info("Success inserted with id", result.InsertedID)
	}

	defer client.Disconnect(ctx)
}

func DeleteLogs() {
	client := config.ConnectToMongoDB()
	collection := client.Database("LogsMonitoring").Collection("Logs")

	deleteResult, err := collection.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		logger.Error("Delete Error: ", err.Error())
	} else {
		logger.Info("Success deleted logs", deleteResult.DeletedCount)
	}
}
