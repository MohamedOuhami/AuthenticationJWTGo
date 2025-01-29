package main

import (
	"github.com/MohamedOuhami/AuthenticationJWTGo/controllers"
	"github.com/MohamedOuhami/AuthenticationJWTGo/initializers"
	"github.com/MohamedOuhami/AuthenticationJWTGo/middleware"
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

  // Add a Post endpoint to sign up the new user
	r.POST("/signup",controllers.Signup)

  // Add a post for the login process
  r.POST("/login",controllers.Login)

  // Adding a test endpoint
  r.GET("/validate",middleware.RequireAuth,controllers.Validate)
	r.Run()

}
