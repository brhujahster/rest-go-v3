package v1

import (
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
	router.PUT("/order/:id", UpdateOrder)
	router.DELETE("/order/:id", DeleteOrder)
}

func GetOrders(c *gin.Context) {
	ordersResponse := orderHandler.GetOrders()
	c.JSON(http.StatusOK, ordersResponse)
}

func GetById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do pedido não fornecido"})
		return
	}
	idPedido, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do pedido inválido"})
		return
	}

	// Lógica para obter um pedido específico
	p := orderHandler.GetById(idPedido)
	c.JSON(http.StatusOK, p)
}

func CreateOrder(c *gin.Context) {
	var produtoDto dto.OrderRequest

	err := c.BindJSON(&produtoDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao analisar JSON"})
		return
	}
	orderResponse := orderHandler.CreateOrder(produtoDto)
	c.JSON(http.StatusCreated, orderResponse)
}

func UpdateOrder(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do pedido não fornecido"})
		return
	}
	idPedido, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do pedido inválido"})
		return
	}
	var orderRequest dto.OrderRequest
	err = c.BindJSON(&orderRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao analisar JSON"})
		return
	}
	orderResponse := orderHandler.UpdateOrder(idPedido, orderRequest)
	c.JSON(http.StatusOK, orderResponse)
}

func DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do pedido não fornecido"})
		return
	}
	idPedido, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do pedido inválido"})
		return
	}
	err = orderHandler.DeleteOrder(idPedido)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir o pedido"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
