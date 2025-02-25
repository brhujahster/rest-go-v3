package v1

import (
	"fmt"
	"net/http"
	"rest-v3/internal/handlers"
	"rest-v3/internal/models/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

var orderHandler handlers.OrderHandler

func SetupOrderRoutes(router *gin.RouterGroup, handler handlers.OrderHandler) {
	orderHandler = handler
	router.GET("/order", GetOrders)
	router.GET("/order/:id", GetById)
	router.POST("/order", CreateOrder)
}

func GetOrders(c *gin.Context) {
	// Lógica para listar pedidos
	orderHandler.GetOrders()
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
	p := orderHandler.GetById(idPedido)
	fmt.Print(p)
}

func CreateOrder(c *gin.Context) {
	var produtoDto dto.OrderDTO

	err := c.BindJSON(&produtoDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao analisar JSON"})
		return
	}
	orderHandler.CreateOrder(produtoDto)
}
