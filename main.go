package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
