package models

import "time"

type User struct {
	ID          string
	Email       string
	Password    string
	Name        string
	Surname     string
	Patronymic  string
	PhoneNumber string
	Addresses   []Address
	UserType    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
