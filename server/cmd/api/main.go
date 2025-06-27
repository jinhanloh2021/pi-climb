package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinhanloh2021/beta-blocker/internal/auth" // Import your new auth package
	"github.com/jinhanloh2021/beta-blocker/internal/config"
	"github.com/jinhanloh2021/beta-blocker/internal/database" // Renamed controllers to handlers
	"github.com/jinhanloh2021/beta-blocker/internal/handler"
	"github.com/jinhanloh2021/beta-blocker/internal/repository" // Import repository
	"github.com/jinhanloh2021/beta-blocker/internal/service"    // Import services
)

// init runs before main
func init() {
	config.LoadConfig()
	database.ConnectToDb()
}

func main() {
	r := gin.Default()

	// Initialize Repositories and Services
	userRepo := repository.NewUserRepository(database.DB)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Routes
	r.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{}) })

	// Protected route
	r.GET("/myinfo", auth.AuthMiddleware(), userHandler.GetMyUser)

	r.GET("/user/:username", auth.AuthMiddleware(), userHandler.GetUserByUsername)

	r.PATCH("/dob", auth.AuthMiddleware(), userHandler.TrySetDifferentUserDOB)

	log.Fatal(r.Run(":8080")) // Listen and serve on 0.0.0.0:8080 by default
}
