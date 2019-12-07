package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/omkz/golang-mongo-rest/models"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := models.PostAll()
	json.NewEncoder(w).Encode(payload)
}
