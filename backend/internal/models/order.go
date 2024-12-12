package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Order struct {
	ID                bson.ObjectID   `json:"id" bson:"_id,omitempty"`
	UserID            bson.ObjectID   `json:"user_id" bson:"user_id,omitempty"`
	Address           Address         `json:"address" bson:"address"`
	DateTime          time.Time       `json:"date_time" bson:"date_time"`
	Price             float64         `json:"price" bson:"price"`
	NumberOfRooms     int             `json:"number_of_rooms" bson:"number_of_rooms"`
	NumberOfBathrooms int             `json:"number_of_bathrooms" bson:"number_of_bathrooms"`
	Pollution         int             `json:"pollution" bson:"pollution"`
	Area              float64         `json:"area" bson:"area"`
	Comment           string          `json:"comment,omitempty" bson:"comment"`
	StatusLogs        []StatusLog     `json:"status_log,omitempty" bson:"status_logs,omitempty"`
	Status            string          `json:"status" bson:"status"`
	Services          []bson.ObjectID `json:"services,omitempty" bson:"services,omitempty"`
	RequiredWorkers   int             `json:"required_workers" bson:"required_workers"`
	Workers           []bson.ObjectID `json:"workers,omitempty" bson:"workers"`
	CreatedAt         time.Time       `json:"created_at" bson:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
