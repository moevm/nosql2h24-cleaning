package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Order struct {
	ID                bson.ObjectID   `bson:"_id,omitempty"`
	UserID            bson.ObjectID   `bson:"user_id,omitempty"`
	Address           Address         `bson:"address"`
	DateTime          time.Time       `bson:"date_time"`
	Price             float64         `bson:"price"`
	NumberOfRooms     int             `bson:"number_of_rooms"`
	NumberOfBathrooms int             `bson:"number_of_bathrooms"`
	Pollution         int             `bson:"pollution"`
	Area              float64         `bson:"area"`
	Comment           string          `bson:"comment"`
	StatusLogs        []StatusLog     `bson:"status_logs,omitempty"`
	Services          []bson.ObjectID `bson:"services,omitempty"`
	RequiredWorkers   int             `bson:"required_workers"`
	Workers           []bson.ObjectID `bson:"workers,omitempty"`
	CreatedAt         time.Time       `bson:"created_at"`
	UpdatedAt         time.Time       `bson:"updated_at,omitempty"`
}
