package models

import "go.mongodb.org/mongo-driver/v2/bson"

type Consumable struct {
	ID      bson.ObjectID `bson:"_id,omitempty"`
	EanCode string        `bson:"ean_code"`
	Name    string        `bson:"name"`
	Amount  int           `bson:"amount"`
}
