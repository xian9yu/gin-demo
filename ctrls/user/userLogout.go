package user

import (
	"gin-demo/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Logout 用户退出
func Logout(c *gin.Context) {
	Authorization := c.Request.Header.Get("Authorization")

	res, err := models.StrDel(Authorization)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err,
			"date": res,
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "退出成功",
	})

}
