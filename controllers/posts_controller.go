package controllers

import (
	"encoding/json"
	"time"
	"github.com/omkz/golang-mongo-rest/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/bson"
	"net/http"
	"github.com/gorilla/mux"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload, _ := models.PostAll()
	json.NewEncoder(w).Encode(payload)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	
    post.ID = primitive.NewObjectID()
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()
	
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&post)
	_ = models.PostCreate(&post)

	response, _ := json.Marshal(post)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response)
}

func GetPostById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	post, _ := models.PostFindById(id)

	response, _ := json.Marshal(post)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response)
}
