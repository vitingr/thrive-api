package main

import (
	"fmt"
	"main/database"
	"main/http/routes"
	"main/middleware"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	r := gin.Default()
	middleware.PrometheusInit()
	r.Use(middleware.TrackMetrics())
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	database.ConnectionWithDatabase()
	fmt.Println("Starting Thrive rest server :D")
	routes.HandleRequest(r)
}