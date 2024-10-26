package userRoutes

import (
	"github.com/gorilla/mux"
	"main/http/controllers/users"
	"main/middleware"
)

func RegisterUserRoutes(r *mux.Router) {
    r.Use(middleware.ContetTypeMiddleware)
    r.HandleFunc("/", controllers.Home).Methods("GET")
}
