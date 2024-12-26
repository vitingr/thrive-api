package postRoutes

import (
	"main/http/controllers/posts"
	"main/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterPostRoutes(r *gin.RouterGroup) {
	r.Use(middleware.ContentTypeMiddleware())

	r.POST("", controllers.CreatePost)
	r.GET("/get-all-posts/:userId", controllers.GetAllPosts)
	r.GET("/get-my-posts/:userId", controllers.GetMyPosts)
	r.GET("/get-post-by-id/:postId/:userId", controllers.GetPostsById)
	r.GET("/get-posts-by-language/:userId/:locale", controllers.GetPostsByLanguage)
	r.POST("/like-post", controllers.LikePost)
	r.POST("/deslike-post", controllers.DeslikePost)
}
