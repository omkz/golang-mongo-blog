package controllers

import (
	"encoding/json"
	"github.com/omkz/golang-mongo-rest/models"
	"net/http"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload, _ := models.PostAll()
	json.NewEncoder(w).Encode(payload)
}
