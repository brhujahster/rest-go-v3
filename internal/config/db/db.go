package database

import (
	"rest-v3/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=username password=password dbname=default_database port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Este comando criar as tabelas no banco de dados, só deve ser usado em ambiente de desenvolvimento
	// Ao se criar um novo modelo, é necessário adicionar a estrutura do modelo aqui
	// Exemplo: db.AutoMigrate(&models.User{})
	db.AutoMigrate(models.Order{})

	return db, nil
}
