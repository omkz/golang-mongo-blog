package main

import (
	"fmt"

	"log"
	"net/http"
	"github.com/omkz/golang-mongo-rest/config"
	"github.com/omkz/golang-mongo-rest/router"
)

func init(){
	config.ConnectDB("mongodb://localhost:27017")
}

func main() {

	router := router.Router()

	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", router))
}
