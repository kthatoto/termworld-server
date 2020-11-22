package players

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kthatoto/termworld-server/app/forms"
	"github.com/kthatoto/termworld-server/app/models"
)

func Create(c *gin.Context) {
	var form forms.PlayerCreateForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}

	var playerModel models.PlayerModel
	httpStatus, err := playerModel.Create(form)
	if err != nil {
		c.JSON(httpStatus, gin.H{ "error": err.Error() })
		return
	}
	c.Status(httpStatus)
}
