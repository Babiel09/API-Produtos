package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Babiel09/models"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func Getid(w http.ResponseWriter, r *http.Request) {
	idstr := chi.URLParam(r, "id")
	if idstr == "" {
		log.Printf("Unxpected error, %v", idstr)
		return
	}
	id, err := uuid.Parse(idstr)
	if err != nil {
		log.Printf("Unxpected error: %v", err)
		return
	}

	products, err := models.GetForId(id)
	if err != nil {
		log.Printf("Unxpected error: %v", err)
		http.Error(w, http.StatusText(500), 500)
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		log.Printf("Unxpected error: %v", err)
		http.Error(w, http.StatusText(500), 500)
	}

}
