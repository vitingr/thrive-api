package controllers

import (
	"encoding/json"
	"main/database"
	"main/models"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	var posts []models.Post

	result := database.DB.Preload("Creator").Where("creator_id != ?", userId).Find(&posts)

	if result.Error != nil {
		http.Error(w, "Error fetching posts", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(posts)
}

func GetMyPosts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	var posts []models.Post

	result := database.DB.Preload("Creator").Where("creator_id = ?", userId).Find(&posts)

	if result.Error != nil {
		http.Error(w, "Error fetching your posts", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(posts)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int)
	var post models.Post

	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	post.CreatorId = userID > 0

	result := database.DB.Create(&post)
	if result.Error != nil {
		http.Error(w, "Error creating post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}
