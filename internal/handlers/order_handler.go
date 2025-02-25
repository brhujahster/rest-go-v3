package handlers

import (
	"rest-v3/internal/models"
	"rest-v3/internal/models/dto"
	"rest-v3/internal/repositories"
)

type OrderHandler struct {
	repository *repositories.OrderRepository
}

func NewOrderHandler(r *repositories.OrderRepository) OrderHandler {
	return OrderHandler{repository: r}
}

func (p *OrderHandler) GetOrders() {
	// Lógica para listar pedidos
}

func (p *OrderHandler) GetById(id int) {
	// Lógica para obter um pedido específico
}

func (p *OrderHandler) CreateOrder(dto dto.OrderDTO) {
	order := p.convertToModel(dto)
	p.repository.CreateOrder(&order)
}

func (p *OrderHandler) convertToModel(dto dto.OrderDTO) models.Order {
	order := models.Order{
		ID:         dto.ID,
		UserID:     dto.UserID,
		Status:     dto.Status,
		TotalPrice: dto.TotalPrice,
	}

	return order
}
