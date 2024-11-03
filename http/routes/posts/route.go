package postRoutes

import (
	"github.com/gorilla/mux"
	"main/http/controllers/posts"
	"main/middleware"
)

func RegisterPostRoutes(r *mux.Router) {
	r.Use(middleware.ContetTypeMiddleware)

	r.HandleFunc("", controllers.CreatePost).Methods("POST")
	r.HandleFunc("/get-all-posts/{userId}", controllers.GetAllPosts).Methods("GET")
	r.HandleFunc("/get-my-posts/{userId}", controllers.GetMyPosts).Methods("GET")
}
