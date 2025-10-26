package main

import (
	"context"
	"log"
	"os"

	"github.com/Danztee/url-shortener/config"
	"github.com/Danztee/url-shortener/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	client := config.ConnectDB()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	server := gin.Default()

	routes.RegisterRoutes(server, client)

	port := os.Getenv("PORT")

	server.Run(":" + port)
}
