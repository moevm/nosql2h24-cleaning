package models

import "time"

type Address struct {
	ID         string
	City       string
	Street     string
	Building   string
	Entrance   string
	Floor      string
	DoorNumber string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
