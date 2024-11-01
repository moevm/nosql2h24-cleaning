package models

import "time"

type StatusLog struct {
	ID         string
	PrevStatus string
	NewStatus  string
	CreatedAt  time.Time
}
