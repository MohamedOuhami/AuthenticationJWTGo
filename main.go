package main

import (
	"time"

	"github.com/MohamedOuhami/AuthenticationJWTGo/controllers"
	"github.com/MohamedOuhami/AuthenticationJWTGo/initializers"
	"github.com/MohamedOuhami/AuthenticationJWTGo/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// This function runs directly before the main function
func init() {

	// Load the env variables
	initializers.LoadEnvVars()

	// Connect to the database
	initializers.ConnectToDB()

	// Sync the database
	initializers.SyncDatabase()

}
func main() {

	r := gin.Default()

	// Enable CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"}, // Allow frontend origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// Add a Post endpoint to sign up the new user
	r.POST("/signup", controllers.Signup)

	// Add a post for the login process
	r.POST("/login", controllers.Login)

	// Adding a test endpoint
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.Run()

}
