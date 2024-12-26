package controllers

import (
	"main/database"
	"main/models"
	"main/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var newUser models.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	query := `
		INSERT INTO users (username, firstname, lastname, email, profile_picture, background_picture, followers, following, locale, password)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	result := database.DB.Exec(query,
		newUser.Username, newUser.Firstname, newUser.Lastname, newUser.Email,
		newUser.ProfilePicture, newUser.BackgroundPicture,
		newUser.Followers, newUser.Following, newUser.Locale, newUser.Password,
	)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	utils.SendGinResponse(c, http.StatusOK, &newUser, nil, "")
}
