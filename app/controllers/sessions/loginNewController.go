package sessions

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kthatoto/termworld-server/app/forms"
	"github.com/kthatoto/termworld-server/app/models"
)

func LoginNew(c *gin.Context) {
	var form forms.LoginForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}

	var userModel models.UserModel
	httpStatus, err := userModel.LoginNew(form)
	if err != nil {
		c.JSON(httpStatus, gin.H{ "error": err.Error() })
		return
	}

	c.Status(httpStatus)
}
