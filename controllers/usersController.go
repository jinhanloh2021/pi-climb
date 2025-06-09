package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jinhanloh2021/book-anything/initialisers"
	"github.com/jinhanloh2021/book-anything/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	// get email/password of request body
	var body struct {
		Email    string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fail to read body",
		})
		return
	}
	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fail to hash password",
		})
		return
	}
	// create user object
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initialisers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fail to create user",
		})
		return
	}
	// respond
	c.JSON(http.StatusOK, gin.H{})
}

func Login(c *gin.Context) {
	// Get the email and password of request body
	var body struct {
		Email    string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fail to read body",
		})
		return
	}

	// Lookup requested user
	var user models.User
	initialisers.DB.First(&user, "email = ?", body.Email)

	// Challenge password
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   string(user.ID),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 7 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	// Send it back
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*7, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}
