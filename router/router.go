package router

import (
	"github.com/gorilla/mux"
	"github.com/omkz/golang-mongo-blog/models"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/posts", models.GetAllPosts).Methods("GET", "OPTIONS")

	return router
}
