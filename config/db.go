package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func ConnectDB() *mongo.Client {
	uri := os.Getenv("MONGO_URI")

	fmt.Println(uri, "uri")

	if uri == "" {
		log.Fatal("Set your mongodb uri")
	}

	client, err := mongo.Connect(options.Client(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}

	logrus.WithFields(logrus.Fields{}).Info("Database connected")

	return client
}
