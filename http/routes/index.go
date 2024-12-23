package routes

import (
	"log"
	"main/http/routes/google"
	"main/http/routes/groups"
	"main/http/routes/health"
	"main/http/routes/posts"
	"main/http/routes/sso"
	"main/http/routes/users"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()

	userSubrouter := r.PathPrefix("/users").Subrouter()
	userRoutes.RegisterUserRoutes(userSubrouter)

	groupSubrouter := r.PathPrefix("/groups").Subrouter()
	groupRoutes.RegisterGroupRoutes(groupSubrouter)

	postSubrouter := r.PathPrefix("/posts").Subrouter()
	postRoutes.RegisterPostRoutes(postSubrouter)

	googleSubrouter := r.PathPrefix("/google").Subrouter()
	googleRoutes.RegisterGoogleRoutes(googleSubrouter)

	ssoSubrouter := r.PathPrefix("/sso").Subrouter()
	ssoRoutes.RegisterSsoRoutes(ssoSubrouter)

	healthSubrouter := r.PathPrefix("/health").Subrouter()
	healthRoutes.RegisterHealthRoutes(healthSubrouter)

	// Apply CORS options
	corsOptions := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	handler := corsOptions(r)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
