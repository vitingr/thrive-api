package userRoutes

import (
	"github.com/gorilla/mux"
	"main/http/controllers/users"
	"main/middleware"
)

func RegisterUserRoutes(r *mux.Router) {
	r.Use(middleware.ContetTypeMiddleware)
	
	r.HandleFunc("", controllers.GetAllUsers).Methods("GET")
	r.HandleFunc("", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/get-user-by-email/{email}", controllers.GetUserByEmail).Methods("GET")
	r.HandleFunc("/get-user-by-google-id/{google_id}", controllers.GetUserById).Methods("GET")
	r.HandleFunc("/get-user-by-id/{id}", controllers.GetUserById).Methods("GET")
	r.HandleFunc("/update-user/{id}", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/get-suggest-friends/{id}", controllers.GetSuggestedFriends).Methods("GET")
}
