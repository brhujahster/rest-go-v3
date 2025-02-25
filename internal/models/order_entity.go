package models

type Order struct {
	ID         int     `gorm:"primaryKey" json:"id"`
	UserID     int     `gorm:"foreignKey:UserID" json:"user_id"`
	Status     string  `gorm:"not null" json:"status"`
	TotalPrice float64 `gorm:"not null" json:"total_price"`
}
