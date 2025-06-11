package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinhanloh2021/beta-blocker/controllers"
	"github.com/jinhanloh2021/beta-blocker/initialisers"
)

func init() {
	initialisers.LoadEnvVariables()
	initialisers.ConnectToDb()
	initialisers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.Run()
}
