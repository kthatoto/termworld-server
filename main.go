package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/kthatoto/termworld-server/app/controllers/players"
	"github.com/kthatoto/termworld-server/app/controllers/sessions"
	"github.com/kthatoto/termworld-server/app/database"
	"github.com/kthatoto/termworld-server/app/middlewares"
)

func main() {
	router := gin.Default()

	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	}
	database.Init()

	router.Use(middlewares.LoadCurrentUser())

	router.LoadHTMLGlob("app/templates/*.html")
	sessionsGroup := router.Group("")
	{
		sessionsGroup.POST("/login/new", sessions.LoginNew)
		sessionsGroup.POST("/login", sessions.TryLogin)
		sessionsGroup.GET("/login/:token", sessions.AcceptToken)
		sessionsGroup.DELETE("/logout", middlewares.Authentication(), sessions.Logout)
	}

	playersGroup := router.Group("/players")
	{
		playersGroup.POST("", middlewares.Authentication(), players.Create)
		playersGroup.GET("", middlewares.Authentication(), players.Index)
	}

	router.Run()
}
