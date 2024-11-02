package models

import (
	"log"

	"github.com/Babiel09/configs"
	"github.com/google/uuid"
)

func Update(id uuid.UUID, product Products) (uuid.UUID, error) {
	coon, err := configs.DatabaseConnection()
	if err != nil {
		return uuid.Nil, err
	}
	defer coon.Close()

	res, err := coon.Exec("UPDATE products SET name=$1, description=$2, price=$3, stock=$4 WHERE id=$5",
		product.Name, product.Description, product.Price, product.Stock, id)
	if err != nil {
		log.Printf("Error: %v", err)
		return uuid.Nil, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error: %v", err)
		return uuid.Nil, err
	}
	if rowsAffected == 0 {
		log.Printf("Error: %v", err)
		return uuid.Nil, err
	}
	if rowsAffected > 1 {
		log.Printf("Error: %v", err)
		return uuid.Nil, err
	}
	return id, err
}
