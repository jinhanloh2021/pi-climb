package main

import (
	"log"

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
	// No /signup or /login for direct email/password via your backend anymore
	// Frontend will handle these via Supabase JS SDK directly.

	// Protected route example
	r.GET("/authenticated", auth.AuthMiddleware(), userHandler.GetAuthenticatedUser)
	r.POST("/posts", auth.AuthMiddleware(), handler.CreatePost) // Example for another handler

	log.Fatal(r.Run()) // Listen and serve on 0.0.0.0:8080 by default
}
