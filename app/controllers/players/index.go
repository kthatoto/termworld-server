package players

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kthatoto/termworld-server/app/models"
	"github.com/kthatoto/termworld-server/app/services"
)

func Index(c *gin.Context) {
	var playerModel models.PlayerModel
	res, httpStatus, err := playerModel.Index(services.CurrentUser(c))
	if err != nil {
		c.JSON(httpStatus, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"players": res})
}
