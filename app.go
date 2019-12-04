package main

import (
	"fmt"
	"log"
	"net/http"

	"./config"
)

func main() {
	router := config.Router()
	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", router))
}
