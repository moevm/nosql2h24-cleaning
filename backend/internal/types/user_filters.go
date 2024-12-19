package types

type UserFilters struct {
	UserData
	OrdersCount NumberRange
	UserType    string `json:"user_type"`
	CreatedAt   TimeInterval
}
