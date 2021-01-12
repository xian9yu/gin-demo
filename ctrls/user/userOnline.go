package user

import (
	"github.com/gin-gonic/gin"
	"jwt/models"
	"net/http"
)

func GetOnlineList(c *gin.Context) {
	var user models.User
	user.UserName = c.Query("user_name")
	token, err := models.StrGetRange(user.UserName)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":  0,
			"userName": user.UserName,
			"token":  token,
		})
	}
}