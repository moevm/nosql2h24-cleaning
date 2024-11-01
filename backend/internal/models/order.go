package models

import "time"

type Order struct {
	ID                string
	UserID            string
	Address           Address
	DateTime          time.Time
	Price             float64
	NumberOfRooms     int
	NumberOfBathrooms int
	Pollution         int
	Area              float64
	Comment           string
	StatusLogs        []StatusLog
	Services          []string
	RequiredWorkers   int
	Workers           []string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
