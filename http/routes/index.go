package routes

import (
	"log"
	"main/http/routes/groups"
	"main/http/routes/posts"
	"main/http/routes/users"
	"main/http/routes/google"
	"main/http/routes/sso"
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

	// Google routes
	googleSubrouter := r.PathPrefix("/google").Subrouter()
	googleRoutes.RegisterGoogleRoutes(googleSubrouter)

	// SSO routes
	ssoSubrouter := r.PathPrefix("/sso").Subrouter()
	ssoRoutes.RegisterSsoRoutes(ssoSubrouter)

	// Apply CORS options
	corsOptions := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	handler := corsOptions(r)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
