package dto

type OrderResponse struct {
	ID         int     `json:"id"`
	UserID     int     `json:"user_id"`
	Status     string  `json:"status"`
	TotalPrice float64 `json:"total_price"`
}
