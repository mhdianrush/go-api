package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mhdianrush/go-api/controllers/authorcontroller"
)

func AuthorRoutes(r *mux.Router) {
	router := r.PathPrefix("/authors").Subrouter()

	router.HandleFunc("", authorcontroller.Index).Methods(http.MethodGet)
	router.HandleFunc("", authorcontroller.Create).Methods(http.MethodPost)
	router.HandleFunc("/{id}/detail", authorcontroller.Detail).Methods(http.MethodGet)
	router.HandleFunc("/{id}/update", authorcontroller.Update).Methods(http.MethodPut)
	router.HandleFunc("/{id}/delete", authorcontroller.Delete).Methods(http.MethodDelete)
}
