package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kthatoto/termworld-server/app/controllers"
)

func main() {
	router := gin.Default()

	loginGroup := router.Group("/login")
	{
		loginGroup.POST("/new",   controllers.LoginNew)
		// loginGroup.POST("/",      controllers.Login)
		// loginGroup.GET("/:token", controllers.LoginToken)
	}

	router.Run()
}
