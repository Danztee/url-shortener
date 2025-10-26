package models

import "time"

type Url struct {
	ID          string    `bson:"_id,omitempty" json:"id"`
	ShortCode   string    `bson:"short_code" json:"short_code"`
	OriginalUrl string    `bson:"original" json:"original_url"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updated_at"`
}

type CreateUrl struct {
	OriginalUrl string `bson:"original_url" json:"original_url"`
}
