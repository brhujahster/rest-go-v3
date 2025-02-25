package models

import "gorm.io/gorm"

type Order struct {
	Model      *gorm.Model
	ID         int
	UserID     int
	Status     string
	TotalPrice float64
}
