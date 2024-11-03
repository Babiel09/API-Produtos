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
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		log.Printf("Unxpected error: %v", err)
		http.Error(w, http.StatusText(500), 500)
	}
}
