package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Address struct {
	ID         bson.ObjectID `bson:"_id,omitempty"`
	City       string        `bson:"city"`
	Street     string        `bson:"street"`
	Building   string        `bson:"building"`
	Entrance   string        `bson:"entrance"`
	Floor      string        `bson:"floor"`
	DoorNumber string        `bson:"door_number"`
	CreatedAt  time.Time     `bson:"created_at"`
	UpdatedAt  time.Time     `bson:"updated_at,omitempty"`
}
