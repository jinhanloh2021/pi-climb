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
	r.SetTrustedProxies([]string{"http://localhost:3000"})

	// Initialize Repositories and Services
	userRepo := repository.NewUserRepository(database.DB)
	postRepo := repository.NewPostRepository(database.DB)
	followRepo := repository.NewFollowRepository(database.DB)
	likeRepo := repository.NewLikeRepository(database.DB)

	userService := service.NewUserService(userRepo)
	postService := service.NewPostService(postRepo)
	feedService := service.NewFeedService(postRepo)
	followService := service.NewFollowService(followRepo)
	likeService := service.NewLikeService(likeRepo)

	userHandler := handler.NewUserHandler(userService)
	postHandler := handler.NewPostHandler(postService)
	feedHandler := handler.NewFeedHandler(feedService)
	followHandler := handler.NewFollowHandler(followService)
	likeHandler := handler.NewLikeHandler(likeService)

	jwtValidator := auth.NewSupabaseJWTValidator()

	r.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Service is healthy"}) })

	apiV0 := r.Group("/api/v0")
	apiV0.Use(middleware.AuthMiddleware(jwtValidator)).Use(middleware.UserAuthContextMiddleware())

	apiV0.GET("/myinfo", userHandler.GetMyUser)
	apiV0.GET("/users/username/:username", userHandler.GetUserByUsername)
	apiV0.PATCH("/users", userHandler.UpdateUser)
	apiV0.GET("/users/:id/followers", followHandler.GetFollowers)
	apiV0.GET("/users/:id/following", followHandler.GetFollowing)
	apiV0.GET("/users/:id/relationship", followHandler.GetFollowRelationship)

	apiV0.POST("/posts", postHandler.CreateNewPost)
	apiV0.POST("/posts/:id/likes", likeHandler.CreateLike)
	apiV0.DELETE("/posts/:id/likes", likeHandler.DeleteLike)
	apiV0.GET("/posts/:id/likes", likeHandler.GetPostLikes)
	apiV0.GET("/posts/:id/likes/me", likeHandler.GetMyPostLike)

	apiV0.GET("/feed", feedHandler.GetFeed)

	apiV0.POST("/follow", followHandler.CreateFollow)
	apiV0.DELETE("/follow", followHandler.DeleteFollow)
	apiV0.GET("/followers/me", followHandler.GetFollowers)
	apiV0.GET("/following/me", followHandler.GetFollowing)

	log.Fatal(r.Run(":8080"))
}
