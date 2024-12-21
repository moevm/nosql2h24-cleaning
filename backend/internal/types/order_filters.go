package types

type OrderFilters struct {
	UserID            string `json:"user_id"`
	User              UserData
	Worker            UserData
	Address           AddressData
	Price             NumberRange
	Area              NumberRange
	NumberOfRooms     NumberRange
	NumberOfBathrooms NumberRange
	Pollution         NumberRange
	RequiredWorkers   NumberRange
	WorkersID         []string `json:"workers_id"`
	Statuses          []string `json:"status"`
	Services          []string `json:"services"`
	DateTime          TimeInterval
}
