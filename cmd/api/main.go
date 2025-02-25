package main

import (
	"log"

	repositories "rest-v3/internal/config/db"
	v1 "rest-v3/internal/routes/v1"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := repositories.InitDB()
	if err != nil {
		log.Fatalf("Falha ao inicializar o banco de dados: %v", err)
	}

	router := gin.Default()

	v1Group := router.Group("/api/v1")
	v1.SetupOrderRoutes(v1Group, db)
	//v1.SetupPedidosRoutes(v1Group, db)

	/*
			  	v2Group := router.Group("/api/v2")
		        v2.SetupProdutosRoutes(v2Group, db)
		        v2.SetupPedidosRoutes(v2Group, db)
	*/

	log.Println("Servidor rodando na porta 8080")
	log.Fatal(router.Run(":8080"))
}
