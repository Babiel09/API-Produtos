package models

import (
	"log"

	"github.com/Babiel09/configs"
	"github.com/google/uuid"
)

func DeleteRows(id uuid.UUID) (uuid.UUID, error) {
	coon, err := configs.DatabaseConnection()
	if err != nil {
		return uuid.Nil, err
	}
	defer coon.Close()

	result, err := coon.Exec("DELETE FROM products WHERE id=$1", id)
	if err != nil {
		return uuid.Nil, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return uuid.Nil, err
	}
	if rowsAffected != 1 {
		log.Printf("Unxpected error, more them one row was deleted: %d", rowsAffected)
		return uuid.Nil, err
	}
	return id, err

}
