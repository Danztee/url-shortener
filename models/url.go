package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Url struct {
	ID          bson.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	ShortCode   string        `bson:"short_code" json:"short_code"`
	OriginalUrl string        `bson:"original" json:"original_url"`
	CreatedAt   time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time     `bson:"updated_at" json:"updated_at"`
}

type CreateUrl struct {
	OriginalUrl string `bson:"original_url" json:"original_url"`
}
