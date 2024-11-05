package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Service struct {
	ID              bson.ObjectID `bson:"_id,omitempty"`
	Name            string        `bson:"name"`
	Price           float64       `bson:"price"`
	WorkersQuantity int           `bson:"workers_quantity"`
	Description     string        `bson:"description"`
	Consumables     []Consumable  `bson:"consumables,omitempty"`
	CreatedAt       time.Time     `bson:"created_at"`
	UpdatedAt       time.Time     `bson:"updated_at,omitempty"`
}
