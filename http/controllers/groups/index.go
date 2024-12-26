package controllers

import (
	"fmt"
	"main/database"
	"main/models"
	"main/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllGroups(c *gin.Context) {
	var groups []models.Group

	query := "SELECT * FROM groups"
	result := database.DB.Raw(query).Scan(&groups)
	if result.Error != nil {
		utils.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to fetch groups")
		return
	}

	utils.SendGinResponse(c, http.StatusOK, groups, nil, "")
}

func GetGroupById(c *gin.Context) {
	id := c.Param("id")

	var currentGroup models.Group
	query := fmt.Sprintf("SELECT * FROM groups WHERE id = %s LIMIT 1", id)
	result := database.DB.Raw(query).Scan(&currentGroup)
	if result.Error != nil {
		utils.SendGinResponse(c, http.StatusNotFound, nil, nil, "Group not found")
		return
	}

	utils.SendGinResponse(c, http.StatusOK, currentGroup, nil, "")
}

func CreateGroup(c *gin.Context) {
	var newGroup models.Group
	if err := c.ShouldBindJSON(&newGroup); err != nil {
		utils.SendGinResponse(c, http.StatusBadRequest, nil, nil, "Invalid JSON")
		return
	}

	query := fmt.Sprintf(`
		INSERT INTO groups (name, description, activities, group_picture, background_picture, is_private, followers, locale, members)
		VALUES ('%s', '%s', '%s', '%s', '%s', %t, %d, '%s', %d)
	`,
		newGroup.Name, newGroup.Description, newGroup.Activities,
		newGroup.GroupPicture, newGroup.BackgroundPicture,
		newGroup.IsPrivate, newGroup.Followers, newGroup.Locale, newGroup.Members,
	)

	result := database.DB.Exec(query)
	if result.Error != nil {
		utils.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to create group")
		return
	}

	utils.SendGinResponse(c, http.StatusCreated, newGroup, nil, "")
}
