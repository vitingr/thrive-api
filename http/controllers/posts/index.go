package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"main/database"
	"main/models"
	"main/utils/response"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Thrive default API route")
	fmt.Println("testando")
}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	var posts []struct {
		models.Post
		IsLikedByUser bool `json:"is_liked_by_user"`
	}

	query := `
		SELECT 
			p.*,
			CASE 
				WHEN l.id IS NOT NULL THEN true 
				ELSE false 
			END AS is_liked_by_user
		FROM posts p
		LEFT JOIN likes l ON p.id = l.post_id AND l.user_id = ?
		WHERE p.creator_id != ?
	`

	if err := database.DB.Raw(query, userId, userId).Scan(&posts).Error; err != nil {
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
	if err := database.DB.Where("creator_id = ?", userId).Find(&posts).Error; err != nil {
		http.Error(w, "Error fetching your posts", http.StatusInternalServerError)
		return
	}

	for i := range posts {
		if err := database.DB.Where("id = ?", posts[i].CreatorId).First(&posts[i].Creator).Error; err != nil {
			http.Error(w, "Error fetching post creator", http.StatusInternalServerError)
			return
		}
	}

	utils.SendResponse(w, http.StatusOK, &posts, nil, "")
}

func GetPostsByLanguage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	locale := vars["locale"]

	query := `
		SELECT 
			p.*, 
			CASE 
				WHEN l.id IS NOT NULL THEN true 
				ELSE false 
			END AS is_liked_by_user
		FROM posts p
		LEFT JOIN likes l ON p.id = l.post_id AND l.user_id = ?
		WHERE p.creator_id != ? AND p.locale = ?
	`

	var posts []struct {
		models.Post
		IsLikedByUser bool `json:"is_liked_by_user"`
	}

	if err := database.DB.Raw(query, userId, userId, locale).Scan(&posts).Error; err != nil {
		http.Error(w, "Error fetching posts", http.StatusInternalServerError)
		return
	}

	for i := range posts {
		if err := database.DB.Where("id = ?", posts[i].CreatorId).First(&posts[i].Creator).Error; err != nil {
			http.Error(w, "Error fetching post creator", http.StatusInternalServerError)
			return
		}
	}

	utils.SendResponse(w, http.StatusOK, &posts, nil, "")
}

func GetPostsById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	postId := vars["postId"]

	query := `
		SELECT 
			p.*, 
			CASE 
				WHEN l.id IS NOT NULL THEN true 
				ELSE false 
			END AS is_liked_by_user
		FROM posts p
		LEFT JOIN likes l ON p.id = l.post_id AND l.user_id = ?
		WHERE p.creator_id != ? AND p.id = ?
	`

	var posts []struct {
		models.Post
		IsLikedByUser bool `json:"is_liked_by_user"`
	}

	if err := database.DB.Raw(query, userId, userId, postId).Scan(&posts).Error; err != nil {
		http.Error(w, "Error fetching posts", http.StatusInternalServerError)
		return
	}

	for i := range posts {
		if err := database.DB.Where("id = ?", posts[i].CreatorId).First(&posts[i].Creator).Error; err != nil {
			http.Error(w, "Error fetching post creator", http.StatusInternalServerError)
			return
		}
	}

	utils.SendResponse(w, http.StatusOK, &posts, nil, "")
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	var newPost models.Post
	err = json.Unmarshal(body, &newPost)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := database.DB.Create(&newPost).Error; err != nil {
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}

	utils.SendResponse(w, http.StatusOK, &newPost, nil, "")
}

func LikePost(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		UserID int `json:"userId"`
		PostID int `json:"postId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	body, _ := io.ReadAll(r.Body)
	fmt.Println("Raw Request Body:", string(body))

	if payload.UserID == 0 || payload.PostID == 0 {
		http.Error(w, "Missing userId or postId", http.StatusBadRequest)
		return
	}

	var post models.Post
	if err := database.DB.Where("id = ?", payload.PostID).First(&post).Error; err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	var existingLike models.Like
	if err := database.DB.Where("user_id = ? AND post_id = ?", payload.UserID, payload.PostID).First(&existingLike).Error; err == nil {
		http.Error(w, "User has already liked this post", http.StatusBadRequest)
		return
	}

	like := models.Like{UserId: payload.UserID, PostId: payload.PostID}
	if err := database.DB.Create(&like).Error; err != nil {
		http.Error(w, "Failed to like post", http.StatusInternalServerError)
		return
	}

	if err := database.DB.Model(&post).Update("number_likes", post.NumberLikes+1).Error; err != nil {
		http.Error(w, "Failed to update post like count", http.StatusInternalServerError)
		return
	}

	utils.SendResponse(w, http.StatusOK, &post, nil, "Post liked successfully")
}

func DeslikePost(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		UserID int `json:"userId"`
		PostID int `json:"postId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var post models.Post
	if err := database.DB.Where("id = ?", payload.PostID).First(&post).Error; err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	var existingLike models.Like
	if err := database.DB.Where("user_id = ? AND post_id = ?", payload.UserID, payload.PostID).First(&existingLike).Error; err != nil {
		http.Error(w, "User has not liked this post", http.StatusBadRequest)
		return
	}

	if err := database.DB.Delete(&existingLike).Error; err != nil {
		http.Error(w, "Failed to remove like", http.StatusInternalServerError)
		return
	}

	if post.NumberLikes > 0 {
		if err := database.DB.Model(&post).Update("number_likes", post.NumberLikes-1).Error; err != nil {
			http.Error(w, "Failed to update post like count", http.StatusInternalServerError)
			return
		}
	}

	utils.SendResponse(w, http.StatusOK, &post, nil, "Post disliked successfully")
}

func HasUserLikedPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	postId := vars["postId"]

	var like models.Like
	if err := database.DB.Where("user_id = ? AND post_id = ?", userId, postId).First(&like).Error; err != nil {
		if err.Error() == "record not found" {
			utils.SendResponse(w, http.StatusOK, map[string]bool{"liked": false}, nil, "")
			return
		}
		http.Error(w, "Error checking like status", http.StatusInternalServerError)
		return
	}

	utils.SendResponse(w, http.StatusOK, map[string]bool{"liked": true}, nil, "")
}
