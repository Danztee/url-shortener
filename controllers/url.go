package controllers

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/Danztee/url-shortener/models"
	"github.com/Danztee/url-shortener/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const (
	DATABASE   = "url_shortener"
	COLLECTION = "urls"
)

func ShortenUrl(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 100*time.Second)
		defer cancel()

		var body models.CreateUrl

		err := c.ShouldBindJSON(&body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request data"})
			return
		}

		shortCode := utils.GenerateShortCode(6)

		url := models.Url{
			ShortCode:   shortCode,
			OriginalUrl: body.OriginalUrl,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		_, err = client.Database(DATABASE).Collection(COLLECTION).InsertOne(ctx, url)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Error("Failed to insert URL into database")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create short URL"})
			return
		}

		baseUrl := os.Getenv("BASE_URL")

		c.JSON(http.StatusCreated, gin.H{
			"short_url": baseUrl + "/" + shortCode,
		})
	}
}

func RedirectUrl(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 100*time.Second)
		defer cancel()

		code := c.Param("code")

		var urlDoc struct {
			OriginalUrl string `bson:"original" json:"original_url"`
		}

		err := client.Database(DATABASE).
			Collection(COLLECTION).
			FindOne(ctx, bson.M{"short_code": code}).Decode(&urlDoc)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"error":      err.Error(),
				"short_code": code,
			}).Warning("Short code not found")
			c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
			return
		}

		c.Redirect(http.StatusFound, urlDoc.OriginalUrl)
	}
}
