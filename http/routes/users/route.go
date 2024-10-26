package userRoutes

import (
	"github.com/gorilla/mux"
	"main/http/controllers/users"
	"main/middleware"
)

func RegisterUserRoutes(r *mux.Router) {
	r.Use(middleware.ContetTypeMiddleware)
	
	r.HandleFunc("/", controllers.GetAllUsers).Methods("GET")
	r.HandleFunc("/", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/get-user-by-email/{email}", controllers.GetUserByEmail).Methods("GET")
	r.HandleFunc("/get-user-by-id/{id}", controllers.GetUserById).Methods("GET")
}
