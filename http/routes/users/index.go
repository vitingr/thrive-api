package userRoutes

import (
	"github.com/gin-gonic/gin"
	users "main/http/controllers/users"
	"main/middleware"
)

func RegisterUserRoutes(r *gin.RouterGroup) {
	r.Use(middleware.ContentTypeMiddleware())

	r.GET("", users.GetAllUsers)
	r.GET("/get-user-by-email/:email", users.GetUserByEmail)
	r.GET("/get-user-by-id/:id", users.GetUserById)
	r.PUT("/update-user/:id", users.UpdateUser)
	r.GET("/get-suggest-friends/:id", users.GetSuggestedFriends)
	r.POST("/send-friend-request/:userId/:friendId", users.SendFriendRequest)
	r.POST("/confirm-friend-request/:userId/:friendId", users.ConfirmFriendRequest)
	r.POST("/cancel-friend-request/:userId/:friendId", users.RemoveFriendRequest)
	r.GET("/get-friends/:userId", users.GetAllFriends)
	r.GET("/get-pending-friends/:userId", users.GetPendingFriendRequests)
}
