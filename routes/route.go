package routes

import (
	"net/http"

	"github.com/Babiel09/handlers"
	"github.com/Babiel09/middlewares"
	"github.com/go-chi/chi/v5"
)

func HandlerRequests() {
	r := chi.NewRouter()
	r.Use(middlewares.ConextexType)
	r.Get("/products", handlers.GetAll)
	r.Get("/products/{id}", handlers.Getid)
	r.Post("/products", handlers.Post)
	r.Delete("/products/{id}", handlers.Delete)
	r.Put("/products/{id}", handlers.Put)

	http.ListenAndServe(":8000", r)
}
