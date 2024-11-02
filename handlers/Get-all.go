package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Babiel09/models"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := models.GetAll()
	if err != nil {
		log.Printf("Unxpected error: %v", err)
	}
	json.NewEncoder(w).Encode(products)
}
