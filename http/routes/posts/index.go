package postRoutes

import (
	"github.com/gorilla/mux"
	"main/http/controllers/posts"
)

func RegisterPostRoutes(r *mux.Router) {

	r.HandleFunc("", controllers.Home).Methods("GET")
	r.HandleFunc("", controllers.CreatePost).Methods("POST")
	r.HandleFunc("/get-all-posts/{userId}", controllers.GetAllPosts).Methods("GET")
	r.HandleFunc("/get-my-posts/{userId}", controllers.GetMyPosts).Methods("GET")
	r.HandleFunc("/get-post-by-id/{postId}/${userId}", controllers.GetPostsById).Methods("GET")
	r.HandleFunc("/get-posts-by-language/{userId}/{locale}", controllers.GetPostsByLanguage).Methods("GET")
	r.HandleFunc("/like-post", controllers.LikePost).Methods("POST")
	r.HandleFunc("/deslike-post", controllers.DeslikePost).Methods("POST")
}
