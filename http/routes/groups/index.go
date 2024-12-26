package groupRoutes

import (
	"github.com/gin-gonic/gin"
	groups "main/http/controllers/groups"
	"main/middleware"
)

func RegisterGroupRoutes(r *gin.RouterGroup) {
	r.Use(middleware.ContentTypeMiddleware())

	r.GET("", groups.GetAllGroups)
	r.POST("", groups.CreateGroup)
	r.GET("/get-group-by-id/:id   ", groups.GetGroupById)
}
