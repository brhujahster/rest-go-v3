package main

import (
	"log"
	database "rest-v3/internal/config/db"
	"rest-v3/internal/handlers"
	"rest-v3/internal/repositories"
	v1 "rest-v3/internal/routes/v1"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Falha ao inicializar o banco de dados: %v", err)
	}

	orderRepository := repositories.NewOrderRepository(db)
	orderHandler := handlers.NewOrderHandler(orderRepository)

	router := gin.Default()

	v1Group := router.Group("/api/v1")

	v1.SetupOrderRoutes(v1Group, orderHandler)
	//v1.SetupPedidosRoutes(v1Group, db)

	/*
			  	v2Group := router.Group("/api/v2")
		        v2.SetupProdutosRoutes(v2Group, db)
		        v2.SetupPedidosRoutes(v2Group, db)
	*/

	log.Println("Servidor rodando na porta 8080")
	log.Fatal(router.Run(":8080"))
}
