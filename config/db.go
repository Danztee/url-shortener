package config

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func ConnectDB() *mongo.Client {
	uri := os.Getenv("MONGO_URI")

	if uri == "" {
		logrus.Fatal("MONGO_URI environment variable is not set")
	}

	client, err := mongo.Connect(options.Client(), options.Client().ApplyURI(uri))
	if err != nil {
		logrus.WithField("error", err.Error()).Fatal("Failed to connect to MongoDB")
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		logrus.WithField("error", err.Error()).Fatal("Failed to ping MongoDB")
		panic(err)
	}

	logrus.Info("Successfully connected to MongoDB")

	return client
}
