package repositories

import (
	"rest-v3/internal/models"

	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (r *OrderRepository) GetOrders() ([]models.Order, error) {
	var pedidos []models.Order
	result := r.DB.Find(&pedidos)
	if result.Error != nil {
		return nil, result.Error
	}
	return pedidos, nil
}
