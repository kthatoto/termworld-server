package sessions

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kthatoto/termworld-server/app/models"
	"github.com/kthatoto/termworld-server/app/forms"
)

func TryLogin(c *gin.Context) {
	var form forms.LoginForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}


	var userModel models.UserModel
	token, httpStatus, err := userModel.TryLogin(form)
	if err != nil {
		c.JSON(httpStatus, gin.H{ "error": err.Error() })
		return
	}
	if httpStatus != http.StatusOK {
		c.Status(httpStatus)
		return
	}
	c.JSON(http.StatusOK, gin.H{ "token": token })
}
