package session

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kthatoto/termworld-server/app/services"
	// db "github.com/kthatoto/termworld-server/app/database"
)

func Logout(c *gin.Context) {
	if err := services.Authentication(c); err != nil {
		return
	}
	currentUser := services.CurrentUser(c)
	fmt.Println(currentUser)

	c.Status(http.StatusOK)
}
