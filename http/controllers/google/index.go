package controllers

import (
	"main/database"
	"main/models"
	"main/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserByGoogleID(c *gin.Context) {
	googleID := c.Param("google_id")

	var currentUser models.User

	query := "SELECT * FROM users WHERE google_id = ? LIMIT 1"
	result := database.DB.Raw(query, googleID).Scan(&currentUser)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	utils.SendGinResponse(c, http.StatusOK, currentUser, nil, "")
}

func CreateUser(c *gin.Context) {
	var newUser models.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	query := `
		INSERT INTO users (username, firstname, lastname, email, profile_picture, background_picture, followers, following, locale, google_id)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	result := database.DB.Exec(query,
		newUser.Username, newUser.Firstname, newUser.Lastname, newUser.Email,
		newUser.ProfilePicture, newUser.BackgroundPicture,
		newUser.Followers, newUser.Following, newUser.Locale, newUser.GoogleID,
	)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	utils.SendGinResponse(c, http.StatusOK, &newUser, nil, "")
}
