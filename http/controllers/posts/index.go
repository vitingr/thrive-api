package controllers

import (
	"main/database"
	"main/models"
	"main/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllPosts(c *gin.Context) {
	userId := c.Param("userId")

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
		utils.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Error fetching posts")
		return
	}

	utils.SendGinResponse(c, http.StatusOK, posts, nil, "")
}

func GetMyPosts(c *gin.Context) {
	userId := c.Param("userId")

	var posts []models.Post
	if err := database.DB.Where("creator_id = ?", userId).Find(&posts).Error; err != nil {
		utils.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Error fetching your posts")
		return
	}

	for i := range posts {
		if err := database.DB.Where("id = ?", posts[i].CreatorId).First(&posts[i].Creator).Error; err != nil {
			utils.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Error fetching post creator")
			return
		}
	}

	utils.SendGinResponse(c, http.StatusOK, posts, nil, "")
}

func GetPostsByLanguage(c *gin.Context) {
	userId := c.Param("userId")
	locale := c.Param("locale")

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
		utils.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Error fetching posts")
		return
	}

	for i := range posts {
		if err := database.DB.Where("id = ?", posts[i].CreatorId).First(&posts[i].Creator).Error; err != nil {
			utils.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Error fetching post creator")
			return
		}
	}

	utils.SendGinResponse(c, http.StatusOK, posts, nil, "")
}

func GetPostsById(c *gin.Context) {
	userId := c.Param("userId")
	postId := c.Param("postId")

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
		utils.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Error fetching posts")
		return
	}

	for i := range posts {
		if err := database.DB.Where("id = ?", posts[i].CreatorId).First(&posts[i].Creator).Error; err != nil {
			utils.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Error fetching post creator")
			return
		}
	}

	utils.SendGinResponse(c, http.StatusOK, posts, nil, "")
}

func CreatePost(c *gin.Context) {
	var newPost models.Post
	if err := c.ShouldBindJSON(&newPost); err != nil {
		utils.SendGinResponse(c, http.StatusBadRequest, nil, nil, "Invalid JSON")
		return
	}

	if err := database.DB.Create(&newPost).Error; err != nil {
		utils.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to create post")
		return
	}

	utils.SendGinResponse(c, http.StatusCreated, newPost, nil, "")
}

func LikePost(c *gin.Context) {
	var payload struct {
		UserID int `json:"userId"`
		PostID int `json:"postId"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.SendGinResponse(c, http.StatusBadRequest, nil, nil, "Invalid request payload")
		return
	}

	var post models.Post
	if err := database.DB.Where("id = ?", payload.PostID).First(&post).Error; err != nil {
		utils.SendGinResponse(c, http.StatusNotFound, nil, nil, "Post not found")
		return
	}

	var existingLike models.Like
	if err := database.DB.Where("user_id = ? AND post_id = ?", payload.UserID, payload.PostID).First(&existingLike).Error; err == nil {
		utils.SendGinResponse(c, http.StatusBadRequest, nil, nil, "User has already liked this post")
		return
	}

	like := models.Like{UserId: payload.UserID, PostId: payload.PostID}
	if err := database.DB.Create(&like).Error; err != nil {
		utils.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to like post")
		return
	}

	if err := database.DB.Model(&post).Update("number_likes", post.NumberLikes+1).Error; err != nil {
		utils.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to update post like count")
		return
	}

	utils.SendGinResponse(c, http.StatusOK, post, nil, "Post liked successfully")
}

func DeslikePost(c *gin.Context) {
	var payload struct {
		UserID int `json:"userId"`
		PostID int `json:"postId"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.SendGinResponse(c, http.StatusBadRequest, nil, nil, "Invalid request payload")
		return
	}

	var post models.Post
	if err := database.DB.Where("id = ?", payload.PostID).First(&post).Error; err != nil {
		utils.SendGinResponse(c, http.StatusNotFound, nil, nil, "Post not found")
		return
	}

	var existingLike models.Like
	if err := database.DB.Where("user_id = ? AND post_id = ?", payload.UserID, payload.PostID).First(&existingLike).Error; err != nil {
		utils.SendGinResponse(c, http.StatusBadRequest, nil, nil, "User has not liked this post")
		return
	}

	if err := database.DB.Delete(&existingLike).Error; err != nil {
		utils.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to remove like")
		return
	}

	if post.NumberLikes > 0 {
		if err := database.DB.Model(&post).Update("number_likes", post.NumberLikes-1).Error; err != nil {
			utils.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to update post like count")
			return
		}
	}

	utils.SendGinResponse(c, http.StatusOK, post, nil, "Post disliked successfully")
}

func HasUserLikedPost(c *gin.Context) {
	userId := c.Param("userId")
	postId := c.Param("postId")

	var like models.Like
	if err := database.DB.Where("user_id = ? AND post_id = ?", userId, postId).First(&like).Error; err != nil {
		if err.Error() == "record not found" {
			utils.SendGinResponse(c, http.StatusOK, map[string]bool{"liked": false}, nil, "")
			return
		}
		utils.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Error checking like status")
		return
	}

	utils.SendGinResponse(c, http.StatusOK, map[string]bool{"liked": true}, nil, "")
}
