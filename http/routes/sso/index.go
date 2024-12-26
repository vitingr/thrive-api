package ssoRoutes

import (
	"main/http/controllers/sso"
	"main/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterSsoRoutes(r *gin.RouterGroup) {
	r.Use(middleware.ContentTypeMiddleware())

	r.POST("", controllers.CreateUser)
}
