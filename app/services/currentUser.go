package services

import (
	"github.com/gin-gonic/gin"

	"github.com/kthatoto/termworld-server/app/models"
)

func CurrentUser(c *gin.Context) models.User {
	user, ok := c.Get("currentUser")
	if !ok {
		var emptyUser models.User
		return emptyUser
	}
	return user.(models.User)
}
