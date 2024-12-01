package googleRoutes

import (
	"github.com/gorilla/mux"
	"main/http/controllers/google"
	"main/middleware"
)

func RegisterGoogleRoutes(r *mux.Router) {
	r.Use(middleware.ContetTypeMiddleware)

	r.HandleFunc("", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/get-user-by-google-id/{google_id}", controllers.GetUserByGoogleID).Methods("GET")
}
