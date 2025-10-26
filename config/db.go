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

func ConnectDB() {
	uri := os.Getenv("MONGO_URI")

	fmt.Println(uri, "uri")

	if uri == "" {
		log.Fatal("Set your mongodb uri")
	}

	client, err := mongo.Connect(options.Client(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	logrus.WithFields(logrus.Fields{}).Info("Database connected")

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
}
