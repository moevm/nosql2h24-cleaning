package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type StatusLog struct {
	ID         bson.ObjectID `json:"id" bson:"_id,omitempty"`
	PrevStatus string        `json:"prev_status" bson:"prev_status"`
	NewStatus  string        `json:"new_status" bson:"new_status"`
	CreatedAt  time.Time     `json:"created_at" bson:"created_at"`
}
