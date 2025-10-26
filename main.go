package main

import (
	"context"
	"os"

	"github.com/Danztee/url-shortener/config"
	"github.com/Danztee/url-shortener/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.WithField("error", err.Error()).Fatal("Error loading .env file")
	}

	client := config.ConnectDB()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			logrus.WithField("error", err.Error()).Error("Error disconnecting from database")
		}
	}()

	server := gin.Default()

	routes.RegisterRoutes(server, client)

	port := os.Getenv("PORT")

	if err := server.Run(":" + port); err != nil {
		logrus.WithField("error", err.Error()).Fatal("Failed to start server")
	}
}
