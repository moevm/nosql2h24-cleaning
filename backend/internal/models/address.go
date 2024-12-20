package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Address struct {
	ID         bson.ObjectID `json:"id" bson:"_id,omitempty"`
	City       string        `json:"city" bson:"city"`
	Street     string        `json:"street" bson:"street"`
	Building   string        `json:"building" bson:"building"`
	Entrance   string        `json:"entrance" bson:"entrance"`
	Floor      string        `json:"floor" bson:"floor"`
	DoorNumber string        `json:"door_number" bson:"door_number"`
}

type AddressWithTimestamp struct {
	Address   `json:",inline" bson:",inline"`
	Timestamp `json:",inline" bson:",inline"`
}
