package types

type OrderFilters struct {
	UserID    string   `json:"user_id"`
	WorkersID []string `json:"workers_id"`
	Statuses  []string `json:"status"`
}
