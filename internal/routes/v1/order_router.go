package v1

import (
	"rest-v3/internal/handlers"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var pedidoHandler handlers.OrderHandler

func SetupOrderRoutes(router *gin.RouterGroup, db *gorm.DB) {
	pedidoHandler = handlers.NewOrderHandler(db)
	router.GET("/order", GetOrders)
	router.GET("/order/:id", GetById)
	router.POST("/order", CreateOrder)
}

func GetOrders(c *gin.Context) {
	// Lógica para listar pedidos
	pedidoHandler.ListarPedidos()
}

func GetById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "ID do pedido não fornecido"})
		return
	}
	idPedido, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID do pedido inválido"})
		return
	}

	// Lógica para obter um pedido específico
	pedidoHandler.GetById(idPedido)
}

func CreateOrder(c *gin.Context) {
	// Lógica para criar um novo pedido
}
