package ssoRoutes

import (
	"github.com/gorilla/mux"
	"main/http/controllers/sso"
	"main/middleware"
)

func RegisterSsoRoutes(r *mux.Router) {
	r.Use(middleware.ContetTypeMiddleware)

	r.HandleFunc("", controllers.CreateUser).Methods("POST")
}
