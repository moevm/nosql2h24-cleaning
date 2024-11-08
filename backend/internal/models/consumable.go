package models

import "go.mongodb.org/mongo-driver/v2/bson"

type Consumable struct {
	ID      bson.ObjectID `json:"id" bson:"_id,omitempty"`
	EanCode string        `json:"ean_code" bson:"ean_code" validate:"required"`
	Name    string        `json:"name" bson:"name" validate:"required"`
	Amount  int           `json:"amount" bson:"amount" validate:"required,gte=1"`
}
