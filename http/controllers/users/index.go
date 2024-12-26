package controllers

import (
	"fmt"
	"main/database"
	"main/models"
	response "main/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)

	c.JSON(http.StatusOK, users)
}

func GetUserByEmail(c *gin.Context) {
	email := c.Param("email")
	var currentUser models.User

	query := fmt.Sprintf("SELECT * FROM users WHERE email = '%s' LIMIT 1", email)
	result := database.DB.Raw(query).Scan(&currentUser)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	response.SendGinResponse(c, http.StatusOK, currentUser, nil, "")
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	var currentUser models.User

	query := fmt.Sprintf("SELECT * FROM users WHERE id = %s LIMIT 1", id)
	result := database.DB.Raw(query).Scan(&currentUser)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	response.SendGinResponse(c, http.StatusOK, &currentUser, nil, "")
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	query := fmt.Sprintf(`
		UPDATE users
		SET username = '%s', firstname = '%s', lastname = '%s', email = '%s',
			profile_picture = '%s', background_picture = '%s', followers = %d,
			following = %d, locale = '%s', google_id = '%s'
		WHERE id = %s
	`,
		updatedUser.Username, updatedUser.Firstname, updatedUser.Lastname, updatedUser.Email,
		updatedUser.ProfilePicture, updatedUser.BackgroundPicture,
		updatedUser.Followers, updatedUser.Following, updatedUser.Locale, updatedUser.GoogleID, id,
	)

	result := database.DB.Exec(query)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	response.SendGinResponse(c, http.StatusOK, &updatedUser, nil, "")
}

func GetSuggestedFriends(c *gin.Context) {
	userID := c.Param("id")

	var suggestedUsers []models.User

	query := fmt.Sprintf(`
		SELECT *
		FROM users
		WHERE id != %s
		AND id NOT IN (
			SELECT following_id
			FROM followers
			WHERE follower_id = %s
		)
	`, userID, userID)

	result := database.DB.Raw(query).Scan(&suggestedUsers)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch suggested friends"})
		return
	}

	response.SendGinResponse(c, http.StatusOK, suggestedUsers, nil, "")
}
