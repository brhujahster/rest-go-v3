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

func (p *OrderHandler) GetOrders() []dto.OrderResponse {
	orders, err := p.repository.GetOrders()
	var ordersReponse []dto.OrderResponse
	if err != nil {
		// Handle the error appropriately, e.g., return an empty slice or log the error
		return []dto.OrderResponse{}
	}

	for _, order := range orders {
		ordersReponse = append(ordersReponse, p.convertToDto(order))
	}
	return ordersReponse
}

func (p *OrderHandler) GetById(id int) dto.OrderResponse {
	// Lógica para obter um pedido específico
	order, err := p.repository.GetOrderById(id)
	if err != nil {
		// Handle the error appropriately, e.g., return a zero value or log the error
		return dto.OrderResponse{}
	}

	return p.convertToDto(*order)
}

func (p *OrderHandler) CreateOrder(dto dto.OrderRequest) dto.OrderResponse {
	order := p.convertToModel(dto)
	p.repository.CreateOrder(&order)

	return p.convertToDto(order)

}

func (p *OrderHandler) convertToModel(dto dto.OrderRequest) models.Order {
	order := models.Order{
		ID:         dto.ID,
		UserID:     dto.UserID,
		Status:     dto.Status,
		TotalPrice: dto.TotalPrice,
	}

	return order
}

func (p *OrderHandler) convertToDto(model models.Order) dto.OrderResponse {
	dto := dto.OrderResponse{
		ID:         model.ID,
		UserID:     model.UserID,
		Status:     model.Status,
		TotalPrice: model.TotalPrice,
	}

	return dto
}
