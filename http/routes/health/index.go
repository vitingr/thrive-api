package healthRoutes

import (
	"main/http/controllers/health"
	"main/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterHealthRoutes(r *gin.RouterGroup) {
	r.Use(middleware.ContentTypeMiddleware())

	r.GET("", controllers.Health)
}
