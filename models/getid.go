package models

import (
	"github.com/Babiel09/configs"
	"github.com/google/uuid"
)

func GetForId(id uuid.UUID) (produtos Products, err error) {
	coon, err := configs.DatabaseConnection()
	if err != nil {
		return
	}
	defer coon.Close()

	result := coon.QueryRow("SELECT * FROM products WHERE id=$1", id)
	err = result.Scan(&produtos.Id, &produtos.Name, &produtos.Description, &produtos.Price, &produtos.Stock)
	return
}
