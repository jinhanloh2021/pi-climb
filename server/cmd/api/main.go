package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"github.com/jinhanloh2021/pi-climb/internal/auth"
	"github.com/jinhanloh2021/pi-climb/internal/config"
	"github.com/jinhanloh2021/pi-climb/internal/database"
	"github.com/jinhanloh2021/pi-climb/internal/handler"
	"github.com/jinhanloh2021/pi-climb/internal/middleware"
	"github.com/jinhanloh2021/pi-climb/internal/repository"
	"github.com/jinhanloh2021/pi-climb/internal/service"
)

func init() {
	config.LoadConfig()
	database.ConnectToDb()
}

func main() {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	// Initialize Repositories and Services
	userRepo := repository.NewUserRepository(database.DB)
	postRepo := repository.NewPostRepository(database.DB)
	followRepo := repository.NewFollowRepository(database.DB)
	likeRepo := repository.NewLikeRepository(database.DB)
	commentRepo := repository.NewCommentRepository(database.DB)

	userService := service.NewUserService(userRepo)
	postService := service.NewPostService(postRepo)
	feedService := service.NewFeedService(postRepo)
	followService := service.NewFollowService(followRepo)
	likeService := service.NewLikeService(likeRepo)
	commentService := service.NewCommentService(commentRepo)

	userHandler := handler.NewUserHandler(userService)
	postHandler := handler.NewPostHandler(postService)
	feedHandler := handler.NewFeedHandler(feedService)
	followHandler := handler.NewFollowHandler(followService)
	likeHandler := handler.NewLikeHandler(likeService)
	commentHandler := handler.NewCommentHandler(commentService)

	r.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			if origin == "http://localhost:3000" {
				return true
			}
			if origin == "https://piclimb.com" || strings.HasSuffix(origin, ".piclimb.com") {
				return true
			}

			return false
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/api/health", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Service is healthy"}) })

	apiV0 := r.Group("/api/v0")
	jwtValidator := auth.NewSupabaseJWTValidator()
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
	apiV0.POST("/posts/:id/comments", commentHandler.CreateComment)
	apiV0.GET("/posts/:id/comments", commentHandler.GetComments)

	apiV0.GET("/feed", feedHandler.GetFeed)

	apiV0.POST("/follow", followHandler.CreateFollow)
	apiV0.DELETE("/follow", followHandler.DeleteFollow)
	apiV0.GET("/followers/me", followHandler.GetFollowers)
	apiV0.GET("/following/me", followHandler.GetFollowing)

	apiV0.DELETE("/comments/:id", commentHandler.DeleteComment)

	log.Fatal(r.Run(":8080"))
}
