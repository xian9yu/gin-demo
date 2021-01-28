package user

import (
	"github.com/gin-gonic/gin"
	"jwt/models"
	"jwt/utils/encrypt"
	"net/http"
)

func Logout(c *gin.Context) {
	Authorization := c.Request.Header.Get("Authorization")

	res, err := models.StrDel(encrypt.GetMd5String(Authorization))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err,
			"date": res,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "退出成功",
	})

}
