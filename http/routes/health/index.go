package healthRoutes

import (
	"github.com/gorilla/mux"
	"main/http/controllers/health"
)

func RegisterHealthRoutes(r *mux.Router) {
	r.HandleFunc("", controllers.Health).Methods("GET")
}
