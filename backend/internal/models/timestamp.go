package models

import "time"

type Timestamp struct {
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
