package types

type AddressData struct {
	City       string        `json:"city"`
	Street     string        `json:"street"`
	Building   string        `json:"building"`
	Entrance   string        `json:"entrance"`
	Floor      string        `json:"floor"`
	DoorNumber string        `json:"door_number"`
}