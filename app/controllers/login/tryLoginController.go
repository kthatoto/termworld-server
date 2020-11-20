package login

import (
	"context"
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/goware/emailx"
	"go.mongodb.org/mongo-driver/bson"

	db "github.com/kthatoto/termworld-server/app/database"
)

type tryLoginRequestJson struct {
	Email string `json:"email"`
}

func TryLogin(c *gin.Context) {
	var data loginNewRequestJson
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
	if err := emailx.Validate(data.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}

	var result bson.M
	err := db.Database.Collection("users").FindOne(
		context.Background(),
		bson.M{ "email": data.Email },
	).Decode(&result)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
	fmt.Println(result)
	// fmt.Println(result.email)
	// fmt.Println(result.accepted)
}
