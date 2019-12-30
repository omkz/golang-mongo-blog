package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/omkz/golang-mongo-rest/models"
	"log"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload, err := models.PostAll()

	if err != nil {
		log.Fatal("Error on Decoding the document", err)
	}
	json.NewEncoder(w).Encode(payload)
}
