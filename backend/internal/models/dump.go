package models

type Dump struct {
	Users    []User    `json:"users"`
	Services []Service `json:"services"`
	Orders   []Order   `json:"orders"`
}
