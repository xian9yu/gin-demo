package user

import (
	"github.com/gin-gonic/gin"
	"jwt/models"
	"net/http"
)

func GetOnlineList(c *gin.Context) {
	Atoken := c.GetHeader("Authorization")
	token, err := models.StrGetRange(Atoken)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":   0,
			"userId": token,
			"token":  Atoken,
		})
	}
}
