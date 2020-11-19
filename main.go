package main

import (
	"github.com/gin-gonic/gin"

	"github.com/kthatoto/termworld-server/app/controllers"
	"github.com/kthatoto/termworld-server/app/middlewares"
)

func main() {
	router := gin.Default()

	router.Use(middlewares.LoadEnv())
	loginGroup := router.Group("/login")
	{
		loginGroup.POST("/new",   controllers.LoginNew)
		// loginGroup.POST("/",      controllers.Login)
		// loginGroup.GET("/:token", controllers.LoginToken)
	}

	router.Run()
}
