package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type UserCredentials struct {
	Email    string `json:"email" bson:"email" validate:"required,email"`
	Password string `json:"password,omitempty" bson:"password" validate:"required"`
}

type UserInfo struct {
	Name        string `json:"name" bson:"name" validate:"required"`
	Surname     string `json:"surname" bson:"surname" validate:"required"`
	Patronymic  string `json:"patronymic,omitempty" bson:"patronymic,omitempty"`
	PhoneNumber string `json:"phone_number" bson:"phone_number" validate:"required"`
}

type User struct {
	ID              bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserCredentials `bson:",inline"`
	UserInfo        `bson:",inline"`
	Addresses       []*Address `json:"addresses,omitempty" bson:"addresses,omitempty"`
	OrdersCount     int        `json:"orders_count,omitempty" bson:"orders_count,omitempty"`
	UserType        string     `json:"user_type" bson:"user_type"`
	CreatedAt       time.Time  `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type NewUser struct {
	UserCredentials
	UserInfo
}
