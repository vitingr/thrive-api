package googleRoutes

import (
	"main/http/controllers/google"
	"main/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterGoogleRoutes(r *gin.RouterGroup) {
	r.Use(middleware.ContentTypeMiddleware())

	r.POST("", controllers.CreateUser)
	r.GET("/get-user-by-google-id/:google_id", controllers.GetUserByGoogleID)
}
