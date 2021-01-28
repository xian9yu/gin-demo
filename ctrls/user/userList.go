package user

import (
	"github.com/gin-gonic/gin"
	"jwt/models"
	"net/http"
)

func GetUserList(c *gin.Context) {
	var user models.User
	totle, users, err := user.GetList()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":  200,
			"totle": totle,
			"user":  users,
		})
	}
}
