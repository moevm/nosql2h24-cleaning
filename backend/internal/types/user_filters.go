package types

import "time"

type UserFilters struct {
	Name           string `json:"name"`
	Surname        string `json:"surname"`
	Email          string `json:"email"`
	UserType       string `json:"user_type"`
	CreatedAtBegin *time.Time
	CreatedAtEnd   *time.Time
}
