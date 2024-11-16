package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"main/database"
	"main/models"
	"main/utils/response"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Thrive default API route")
	fmt.Println("testando")
}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	var posts []models.Post

	result := database.DB.Preload("Creator").Where("creator_id != ?", userId).Find(&posts)

	if result.Error != nil {
		http.Error(w, "Error fetching posts", http.StatusInternalServerError)
		return
	}

	utils.SendResponse(w, http.StatusOK, &posts, nil, "")
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

	utils.SendResponse(w, http.StatusOK, &posts, nil, "")
}

func GetPostsByLanguage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	locale := vars["locale"]

	var posts []models.Post

	result := database.DB.Preload("Creator").
		Where("creator_id != ? AND locale = ?", userId, locale).
		Find(&posts)

	if result.Error != nil {
		http.Error(w, "Error fetching posts", http.StatusInternalServerError)
		return
	}

	utils.SendResponse(w, http.StatusOK, &posts, nil, "")
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("teste")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	var newPost models.Post

	r.Body = io.NopCloser(bytes.NewBuffer(body))

	err = json.NewDecoder(r.Body).Decode(&newPost)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	result := database.DB.Create(&newPost)

	if result.Error != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return 
	}

	utils.SendResponse(w, http.StatusOK, &newPost, nil, "")
}
