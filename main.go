package main

import (
	"fmt"
	"gin-api/configs"
	"gin-api/controllers/authcontroller"
	"gin-api/routes"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load config & db
	configs.LoadConfig()
	configs.ConnectDB()

	// Use gin
	r := gin.Default()

	// Setup cors
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{configs.ENV.CLIENT_ORIGIN}
	corsConfig.AllowCredentials = true
	r.Use(cors.New(corsConfig))

	// Router
	router := r.Group("/api")
	router.GET("/health-checker", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "It's Work!"})
	})

	AuthController := authcontroller.NewAuthController(configs.DB)
	AuthRoute := routes.NewAuthRoute(AuthController)
	AuthRoute.Route(router)

	log.Fatal(r.Run(fmt.Sprintf(":%v", configs.ENV.PORT)))
}
