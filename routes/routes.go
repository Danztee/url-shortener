package routes

import (
	"github.com/Danztee/url-shortener/controllers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func RegisterRoutes(server *gin.Engine, client *mongo.Client) {
	server.POST("/shorten", func(c *gin.Context) {
		controllers.ShortenUrl(c, client)
	})
}
