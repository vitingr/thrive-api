package controllers

import (
	"bytes"
	"encoding/json"
	"io"
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
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	
	var newPost models.Post

	r.Body = io.NopCloser(bytes.NewBuffer(body))

	err = json.NewDecoder(r.Body).Decode(&newPost)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	result := database.DB.Create(&newPost)
	if result.Error != nil {
		http.Error(w, "Error creating post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPost)
}
