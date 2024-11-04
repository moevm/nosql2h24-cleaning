package models

import "time"

type Service struct {
	ID              string
	Name            string
	Price           float64
	WorkersQuantity int
	Description     string
	Consumables     []Consumable
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
