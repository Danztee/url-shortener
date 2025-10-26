package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Danztee/url-shortener/models"
	"github.com/Danztee/url-shortener/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func ShortenUrl(c *gin.Context, client *mongo.Client) {
	var body models.CreateUrl

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	shortCode := utils.GenerateShortCode(6)

	url := models.Url{
		ShortCode:   shortCode,
		OriginalUrl: body.OriginalUrl,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	log.Printf("client: %+v\n", client)

	result, err := client.Database("url_shortener").Collection("urls").InsertOne(context.TODO(), url)
	if err != nil {
		panic(err)
	}

	logrus.WithFields(logrus.Fields{}).Info("Document inserted with ID:", result)

	c.JSON(http.StatusCreated, url)
}

func RedirectUrl() {}
