package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Service struct {
	ID              bson.ObjectID `json:"id" bson:"_id,omitempty"`
	Name            string        `json:"name" bson:"name" validate:"required"`
	Price           int           `json:"price" bson:"price" validate:"required"`
	WorkersQuantity int           `json:"workers_quantity" bson:"workers_quantity" validate:"required"`
	Description     string        `json:"description" bson:"description"`
	Consumables     []Consumable  `json:"consumables,omitempty" bson:"consumables,omitempty" validate:"required"`
	Timestamp       `json:",inline" bson:",inline"`
}
