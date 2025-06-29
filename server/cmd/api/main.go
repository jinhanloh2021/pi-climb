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

	r.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{}) })

	protected := r.Group("/")
	protected.Use(auth.AuthMiddleware()).Use(auth.UserAuthContextMiddleware())

	protected.GET("/myinfo", userHandler.GetMyUser)
	protected.GET("/user/:username", userHandler.GetUserByUsername)
	protected.PATCH("/dob", userHandler.TrySetDifferentUserDOB)

	log.Fatal(r.Run(":8080")) // Listen and serve on 0.0.0.0:8080 by default
}
