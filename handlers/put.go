package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Babiel09/models"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func Put(w http.ResponseWriter, r *http.Request) {

	idstr := chi.URLParam(r, "id")
	if idstr == "" {
		log.Printf("ID parameter is missing")
		http.Error(w, "ID parameter is required", http.StatusBadRequest)
		return
	}
	id, err := uuid.Parse(idstr)
	if err != nil {
		log.Printf("Unxpected error> %v", err)
		return
	}

	var product models.Products

	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Printf("Unxpected error> %v", err)
		return
	}
	updateProduct, err := models.Update(id, product)
	if err != nil {
		log.Printf("Unxpected error> %v", err)
		http.Error(w, "Failed to Update thei nformations", 500)
		return
	}
	err = json.NewEncoder(w).Encode(updateProduct)
	if err != nil {
		log.Printf("Unxpected error> %v", err)
		http.Error(w, "Failed to Update thei nformations", 500)
		return
	}

}
