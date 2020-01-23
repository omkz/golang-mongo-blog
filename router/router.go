package router

import (
	"github.com/gorilla/mux"
	"github.com/omkz/golang-mongo-rest/controllers"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/posts", controllers.GetAllPosts).Methods("GET")
	router.HandleFunc("/posts", controllers.CreatePost).Methods("POST")
	router.HandleFunc("/posts/{id}", controllers.GetPostById).Methods("GET")
	router.HandleFunc("/posts/{id}/delete", controllers.DeletePost).Methods("DELETE")

	return router
}
