package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/MohamedOuhami/AuthenticationJWTGo/initializers"
	"github.com/MohamedOuhami/AuthenticationJWTGo/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// The first function here will be to sign up the user in our app
// This function takes the gin context to handle the signup request

func Signup(c *gin.Context) {

	// Get the email and password off the req body

	// We first create a struct for the body that we're expecting
	var body struct {
		Email    string
		Password string
	}

	// Bind the context to the body to take the info that we need
	err := c.Bind(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		// To stop It form doing anything else
		return
	}

	// Hash the password
	// Using the bcrypt, we gonna hash the password

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash the password",
		})

		// To stop It form doing anything else
		return
	}

	// Create the user
	user := models.User{Email: body.Email, Password: string(hashedPassword)}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
	}

	// Return a 200 reponse
	c.JSON(http.StatusOK, gin.H{})
}

// Now, for the login function

func Login(c *gin.Context) {

	// Get the email and pass off of req body

	var body struct {
		Email    string
		Password string
	}

	var foundUser models.User

	err := c.Bind(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		// To stop It form doing anything else
		return
	}

	// Look up the wanted user
	initializers.DB.Where("email = ?", body.Email).First(&foundUser)

	if foundUser.ID == 0 {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})

		// To stop It form doing anything else
		return
	}

	// Compare the sent password with the hashedPassword
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(body.Password))

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})

		// To stop It form doing anything else
		return
	}

	// Generate a JWT token and send It back
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": foundUser.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})

		// To stop It form doing anything else
		return
	}

	// Return the token as a Cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}

// Validate endpoints for Authorization
func Validate(c *gin.Context) {

  user,_ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
