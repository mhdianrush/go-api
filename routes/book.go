package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mhdianrush/go-api/controllers/bookcontroller"
)

func BookRoutes(r *mux.Router) {
	router := r.PathPrefix("/books").Subrouter()

	router.HandleFunc("", bookcontroller.Index).Methods(http.MethodGet)
	router.HandleFunc("", bookcontroller.Create).Methods(http.MethodPost)
	router.HandleFunc("/{id}/detail", bookcontroller.Detail).Methods(http.MethodGet)
	router.HandleFunc("/{id}/update", bookcontroller.Update).Methods(http.MethodPut)
	router.HandleFunc("/{id}/delete", bookcontroller.Delete).Methods(http.MethodDelete)
}
