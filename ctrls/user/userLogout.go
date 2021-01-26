package user

import (
	"github.com/gin-gonic/gin"
	"jwt/models"
	"jwt/utils/encryption"
	"net/http"
)

func Logout(c *gin.Context) {
	Authorization := c.Request.Header.Get("Authorization")

	res, err := models.StrDel(encryption.GetMd5String(Authorization))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err,
			"date": res,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "退出成功",
	})

}
