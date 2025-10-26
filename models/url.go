package models

import "time"

type URL struct {
	ID        string    `bson:"_id,omitempty"`
	ShortCode string    `bson:"short_code"`
	Original  string    `bson:"original"`
	CreatedAt time.Time `bson:"created_at"`
}
