package controllers

import (
	"fmt"

	db "github.com/kthatoto/termworld-server/app/database"
	"github.com/gin-gonic/gin"
)

func LoginNew(c *gin.Context) {
	fmt.Println(db.Client)
}
