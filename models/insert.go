package models

import (
	"log"

	"github.com/Babiel09/configs"
	"github.com/google/uuid"
)

func Insert(product Products) (id uuid.UUID, err error) {
	coon, err := configs.DatabaseConnection()
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	defer coon.Close()

	sql := "INSERT INTO products(name,description,price,stock) VALUES($1,$2,$3,$4) RETURNING id"
	err = coon.QueryRow(sql, product.Name, product.Description, product.Price, product.Stock).Scan(&id)
	return id, err
}
