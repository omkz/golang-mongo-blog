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

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&post)
	_ = models.PostCreate(&post)

	response, _ := json.Marshal(post)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response)
}
