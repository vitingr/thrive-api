package routes

import (
	"log"
	"main/http/routes/groups"
	"main/http/routes/users"
	"main/http/routes/posts"
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

	// Post routes
	postSubrouter := r.PathPrefix("/posts").Subrouter()
	postRoutes.RegisterPostRoutes(postSubrouter)

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"*"}))(r)))
}
