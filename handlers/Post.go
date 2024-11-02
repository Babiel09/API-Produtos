package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Babiel09/models"
)

func Post(w http.ResponseWriter, r *http.Request) {
	var product models.Products
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Printf("This element does not exist, we catch an error: %v", err)
		http.Error(w, http.StatusText(500), 500)
	}

	insertProduct, err := models.Insert(product)
	if err != nil {
		log.Printf("This element does not exist, we catch an error: %v", err)
		http.Error(w, http.StatusText(500), 500)
	}
	json.NewEncoder(w).Encode(insertProduct)
}
