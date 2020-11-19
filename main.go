package main

import (
	"github.com/gin-gonic/gin"

	"github.com/kthatoto/termworld-server/app/controllers/login"
	"github.com/kthatoto/termworld-server/app/middlewares"
)

func main() {
	router := gin.Default()

	router.Use(middlewares.LoadEnv())
	loginGroup := router.Group("/login")
	{
		loginGroup.POST("/new", login.LoginNew)
		loginGroup.POST("",     login.TryLogin)
		// loginGroup.GET("/:token", controllers.LoginToken)
	}

	router.Run()
}
