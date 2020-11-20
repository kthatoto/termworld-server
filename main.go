package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/kthatoto/termworld-server/app/middlewares"
	"github.com/kthatoto/termworld-server/app/database"
	"github.com/kthatoto/termworld-server/app/controllers/session"
)

func main() {
	router := gin.Default()

	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	}
	database.Init()

	router.Use(middlewares.LoadCurrentUser())

	router.LoadHTMLGlob("app/templates/*.html")
	sessionGroup := router.Group("")
	{
		sessionGroup.POST("/login/new", session.LoginNew)
		sessionGroup.POST("/login", session.TryLogin)
		sessionGroup.GET("/login/:token", session.AcceptToken)
		sessionGroup.DELETE("/logout", middlewares.Authentication(), session.Logout)
	}

	router.Run()
}
