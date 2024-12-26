package routes

import (
	"log"
	"main/http/routes/google"
	"main/http/routes/groups"
	"main/http/routes/health"
	"main/http/routes/posts"
	"main/http/routes/sso"
	"main/http/routes/users"

	"github.com/gin-gonic/gin"
)

func HandleRequest(r *gin.Engine) {
	userGroup := r.Group("/users"); {
		userRoutes.RegisterUserRoutes(userGroup)
	}

	groupGroup := r.Group("/groups"); {
		groupRoutes.RegisterGroupRoutes(groupGroup)
	}

	postGroup := r.Group("/posts"); {
		postRoutes.RegisterPostRoutes(postGroup)
	}

	googleGroup := r.Group("/google"); {
		googleRoutes.RegisterGoogleRoutes(googleGroup)
	}

	ssoGroup := r.Group("/sso"); {
		ssoRoutes.RegisterSsoRoutes(ssoGroup)
	}

	healthGroup := r.Group("/health"); {
		healthRoutes.RegisterHealthRoutes(healthGroup)
	}

	log.Fatal(r.Run(":8080"))
}
