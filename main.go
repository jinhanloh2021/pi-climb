package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinhanloh2021/book-anything/initialisers"
)

func init() {
	initialisers.LoadEnvVariables()
	initialisers.ConnectToDb()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
