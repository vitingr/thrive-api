package groupRoutes

import (
	"github.com/gorilla/mux"
	"main/http/controllers/groups"
	"main/middleware"
)

func RegisterGroupRoutes(r *mux.Router) {
	r.Use(middleware.ContetTypeMiddleware)

	r.HandleFunc("", controllers.GetAllGroups).Methods("GET")
	r.HandleFunc("", controllers.CreateGroup).Methods("POST")
	r.HandleFunc("/get-group-by-id/{id}", controllers.GetGroupById).Methods("GET")
}
