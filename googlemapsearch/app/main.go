package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"googlemapsearch/app/routes"
)

func Setup() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Authorization", "Content-Type", "Access-Control-Allow-Origin"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "HEAD"}
	router.Use(cors.New(config))

	router.POST("/googlemapsearch", routes.PostGoogleMapSearch)

	return router
}

func main() {
	router := Setup()
	router.Run(":8080")
}
