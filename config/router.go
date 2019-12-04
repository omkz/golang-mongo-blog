package config

import (
	"../models"
	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/posts", models.GetAllPosts).Methods("GET", "OPTIONS")

	return router
}
