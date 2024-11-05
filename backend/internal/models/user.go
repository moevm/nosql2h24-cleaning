package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID          bson.ObjectID `bson:"_id,omitempty"`
	Email       string        `bson:"email"`
	Password    string        `bson:"password"`
	Name        string        `bson:"name"`
	Surname     string        `bson:"surname"`
	Patronymic  string        `bson:"patronymic,omitempty"`
	PhoneNumber string        `bson:"phone_number"`
	Addresses   []Address     `bson:"addresses,omitempty"`
	UserType    string        `bson:"user_type"`
	CreatedAt   time.Time     `bson:"created_at"`
	UpdatedAt   time.Time     `bson:"updated_at,omitempty"`
}
