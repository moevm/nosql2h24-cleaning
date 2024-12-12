package types

type UserFilters struct {
	UserData
	UserType  string `json:"user_type"`
	CreatedAt TimeInterval
}
