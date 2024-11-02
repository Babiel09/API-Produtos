package handlers

import (
	"log"
	"net/http"

	"github.com/Babiel09/models"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	if idStr == "" {
		log.Printf("ID parameter is missing")
		http.Error(w, "ID parameter is required", http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Printf("Invalid UUID format: %v", err)
		http.Error(w, "Invalid UUID format", http.StatusBadRequest)
		return
	}

	_, err = models.DeleteRows(id)
	if err != nil {
		log.Printf("Error deleting row: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
