package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type StatusLog struct {
	ID         bson.ObjectID `bson:"_id,omitempty"`
	PrevStatus string        `bson:"prev_status"`
	NewStatus  string        `bson:"new_status"`
	CreatedAt  time.Time     `bson:"created_at"`
}
