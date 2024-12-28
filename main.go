package main

import (
	"fmt"
	"main/database"
	"main/http/routes"
	"main/middleware"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/gin-contrib/cors"
)

func main() {
	r := gin.Default()
	middleware.PrometheusInit()
	r.Use(middleware.TrackMetrics())

	r.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			return strings.HasPrefix(origin, "http://localhost:3000/")
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	database.ConnectionWithDatabase()
	fmt.Println("Starting Thrive rest server :D")
	routes.HandleRequest(r)
}
