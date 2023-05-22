package config

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var logger = InitLogger()

// i don't think that will be useful to return error in some function because all errors i see in logs file)
func ConnectToMongoDB() *mongo.Client {
	optionClient := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), optionClient)
	if err != nil {
		logger.Fatal("Unable add mongodb", err.Error())
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		logger.Error("Failed connect to mongodb", err.Error())
	} else {
		logger.Info("Success connected to mongoDB")
	}

	return client
}
