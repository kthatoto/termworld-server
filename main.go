package main

import (
	"github.com/gin-gonic/gin"

	"github.com/kthatoto/termworld-server/app/controllers/session"
	"github.com/kthatoto/termworld-server/app/middlewares"
)

func main() {
	router := gin.Default()

	middlewares.LoadEnv()
	router.Use(middlewares.LoadCurrentUser())

	sessionGroup := router.Group("")
	{
		sessionGroup.POST("/login/new", session.LoginNew)
		sessionGroup.POST("/login", session.TryLogin)
		sessionGroup.GET("/login/:token", session.AcceptToken)
		sessionGroup.DELETE("/logout", session.Logout)
	}

	router.Run()
}
