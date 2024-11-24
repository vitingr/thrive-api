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

	w.Header().Set("Content-Type", "application/json")
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

	utils.SendResponse(w, http.StatusOK, &posts, nil, "")
}

func GetPostsByLanguage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	locale := vars["locale"]

	if database.DB == nil {
		http.Error(w, "Database connection is nil", http.StatusInternalServerError)
		return
	}

	query := `
    SELECT * 
    FROM posts 
    WHERE creator_id != $1 
    AND locale = $2
`

	var posts []models.Post
	result := database.DB.Raw(query, userId, locale).Preload("Creator").Scan(&posts)
	if result.Error != nil {
		http.Error(w, "Error fetching posts", http.StatusInternalServerError)
		return
	}

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

func LikePost(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		UserId int `json:"user"`
		PostId int `json:"postId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var post models.Post
	result := database.DB.First(&post, payload.PostId)
	if result.Error != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	var existingLike models.Like
	likeCheck := database.DB.Where("user_id = ? AND post_id = ?", payload.UserId, payload.PostId).First(&existingLike)
	if likeCheck.Error == nil {
		http.Error(w, "User has already liked this post", http.StatusBadRequest)
		return
	}

	newLike := models.Like{
		UserId: payload.UserId,
		PostId: payload.PostId,
	}

	if err := database.DB.Create(&newLike).Error; err != nil {
		http.Error(w, "Failed to like post", http.StatusInternalServerError)
		return
	}

	post.NumberLikes++
	if err := database.DB.Save(&post).Error; err != nil {
		http.Error(w, "Failed to update post like count", http.StatusInternalServerError)
		return
	}

	utils.SendResponse(w, http.StatusOK, &post, nil, "Post liked successfully")
}

func DeslikePost(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		UserID int `json:"user"`
		PostID int `json:"postId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var post models.Post
	result := database.DB.First(&post, payload.PostID)
	if result.Error != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	var existingLike models.Like
	likeCheck := database.DB.Where("user_id = ? AND post_id = ?", payload.UserID, payload.PostID).First(&existingLike)
	if likeCheck.Error != nil {
		http.Error(w, "User has not liked this post", http.StatusBadRequest)
		return
	}

	if err := database.DB.Delete(&existingLike).Error; err != nil {
		http.Error(w, "Failed to remove like", http.StatusInternalServerError)
		return
	}

	if post.NumberLikes > 0 {
		post.NumberLikes--
		if err := database.DB.Save(&post).Error; err != nil {
			http.Error(w, "Failed to update post like count", http.StatusInternalServerError)
			return
		}
	}

	utils.SendResponse(w, http.StatusOK, &post, nil, "Post disliked successfully")
}
