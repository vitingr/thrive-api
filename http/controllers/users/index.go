package controllers

import (
	"fmt"
	"main/database"
	"main/models"
	response "main/utils/response"
	"net/http"
	"strconv"

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

func GetAllFriends(c *gin.Context) {
	userID := c.Param("userId")

	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid userId"})
		return
	}

	var friendsWithStatus []struct {
		models.Follower
		IsFriend bool `json:"is_friend"`
	}

	query := `
        SELECT 
            f.*, 
            CASE 
                WHEN f.follower_id = ? AND f.following_id = ? THEN true
                WHEN f.follower_id = ? AND f.following_id = ? THEN true
                ELSE false
            END AS is_friend
        FROM followers f
        WHERE (f.follower_id = ? OR f.following_id = ?) AND f.status = ?
    `

	if err := database.DB.Raw(query, userIDInt, userIDInt, userIDInt, userIDInt, userIDInt, userIDInt, "confirmed").Scan(&friendsWithStatus).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve confirmed friends"})
		return
	}

	response.SendGinResponse(c, http.StatusOK, &friendsWithStatus, nil, "")
}

func SendFriendRequest(c *gin.Context) {
	userID := c.Param("userId")
	friendID := c.Param("friendId")

	if userID == friendID {
		c.JSON(400, gin.H{"error": "You cannot send a friend request to yourself."})
		return
	}

	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid userId"})
		return
	}

	friendIDInt, err := strconv.Atoi(friendID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid friendId"})
		return
	}

	newFollower := models.Follower{
		FollowerId:  userIDInt,
		FollowingId: friendIDInt,
		Status:      "pending",
	}

	if err := database.DB.Create(&newFollower).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to send friend request"})
		return
	}

	response.SendGinResponse(c, http.StatusOK, &newFollower, nil, "")
}

func ConfirmFriendRequest(c *gin.Context) {
	userID := c.Param("userId")
	friendID := c.Param("friendId")

	if userID == friendID {
		c.JSON(400, gin.H{"error": "You cannot confirm a friend request to yourself."})
		return
	}

	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid userId"})
		return
	}

	friendIDInt, err := strconv.Atoi(friendID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid friendId"})
		return
	}

	var friendRequest models.Follower
	if err := database.DB.Where("follower_id = ? AND following_id = ? AND status = ?", friendIDInt, userIDInt, "pending").First(&friendRequest).Error; err != nil {
		c.JSON(404, gin.H{"error": "Friend request not found"})
		return
	}

	friendRequest.Status = "confirmed"
	if err := database.DB.Save(&friendRequest).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to confirm friend request"})
		return
	}

	response.SendGinResponse(c, http.StatusOK, &friendRequest, nil, "")
}

func RemoveFriendRequest(c *gin.Context) {
	userID := c.Param("userId")
	friendID := c.Param("friendId")

	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid userId"})
		return
	}

	friendIDInt, err := strconv.Atoi(friendID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid friendId"})
		return
	}

	var friendRequest models.Follower
	if err := database.DB.Where("follower_id = ? AND following_id = ?", userIDInt, friendIDInt).First(&friendRequest).Error; err != nil {
		c.JSON(404, gin.H{"error": "Friend request not found"})
		return
	}

	if err := database.DB.Delete(&friendRequest).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to remove friend request"})
		return
	}

	response.SendGinResponse(c, http.StatusOK, nil, nil, "")
}
