package routes

import (
	"log"
	"main/http/routes/groups"
	"main/http/routes/users"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()

	// User routes
	userSubrouter := r.PathPrefix("/users").Subrouter()
	userRoutes.RegisterUserRoutes(userSubrouter)

	// Group routes
	groupSubrouter := r.PathPrefix("/groups").Subrouter()
	groupRoutes.RegisterGroupRoutes(groupSubrouter)

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"*"}))(r)))
}
