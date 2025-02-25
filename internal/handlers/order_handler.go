package handlers

import "gorm.io/gorm"

type OrderHandler struct {
	DB *gorm.DB
}

func NewOrderHandler(db *gorm.DB) OrderHandler {
	return OrderHandler{DB: db}
}

func (p *OrderHandler) ListarPedidos() {
	// Lógica para listar pedidos
}

func (p *OrderHandler) GetById(id int) {
	// Lógica para obter um pedido específico
}

func (p *OrderHandler) CriarPedido() {
	// Lógica para criar um novo pedido
}
