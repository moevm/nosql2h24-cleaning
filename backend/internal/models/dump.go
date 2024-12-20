package models

import "go.mongodb.org/mongo-driver/v2/bson"

type UserDump struct {
	ID           bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Email        string        `json:"email" bson:"email"`
	PasswordHash string        `json:"password" bson:"password"`
	UserInfo     `bson:",inline"`
	Addresses    []*AddressWithTimestamp `json:"addresses,omitempty" bson:"addresses,omitempty"`
	OrdersCount  int                     `json:"orders_count,omitempty" bson:"orders_count,omitempty"`
	UserType     string                  `json:"user_type" bson:"user_type"`
	Timestamp    `json:",inline" bson:",inline"`
}

type Dump struct {
	Users    []UserDump `json:"users"`
	Services []Service  `json:"services"`
	Orders   []Order    `json:"orders"`
}
