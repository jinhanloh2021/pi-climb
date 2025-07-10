package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinhanloh2021/beta-blocker/internal/auth"
	"github.com/jinhanloh2021/beta-blocker/internal/config"
	"github.com/jinhanloh2021/beta-blocker/internal/database"
	"github.com/jinhanloh2021/beta-blocker/internal/handler"
	"github.com/jinhanloh2021/beta-blocker/internal/middleware"
	"github.com/jinhanloh2021/beta-blocker/internal/repository"
	"github.com/jinhanloh2021/beta-blocker/internal/service"
)

func init() {
	config.LoadConfig()
	database.ConnectToDb()
}

func main() {
	r := gin.Default()

	// Initialize Repositories and Services
	userRepo := repository.NewUserRepository(database.DB)
	postRepo := repository.NewPostRepository(database.DB)

	userService := service.NewUserService(userRepo)
	postService := service.NewPostService(postRepo)

	userHandler := handler.NewUserHandler(userService)
	postHandler := handler.NewPostHandler(postService)

	jwtValidator := auth.NewSupabaseJWTValidator()

	r.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Service is healthy"}) })

	apiV0 := r.Group("/api/v0")
	apiV0.Use(middleware.AuthMiddleware(jwtValidator)).Use(middleware.UserAuthContextMiddleware())

	apiV0.GET("/myinfo", userHandler.GetMyUser)
	apiV0.GET("/user/:username", userHandler.GetUserByUsername)
	apiV0.PATCH("/user", userHandler.UpdateUser)

	apiV0.POST("/post", postHandler.CreateNewPost)

	log.Fatal(r.Run(":8080"))
}
