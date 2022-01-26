package main

import (
	"Project1/handler"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	router := chi.NewRouter()
	router.Post("/text", handler.CountWords)
	http.ListenAndServe(":8080", router)
}
