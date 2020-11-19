package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func LoadEnv() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := godotenv.Load(); err != nil {
			log.Fatal(err.Error())
		}
		c.Next()
	}
}
