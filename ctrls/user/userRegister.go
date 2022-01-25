package user

import (
	"gin-demo/models"
	"gin-demo/utils/encrypt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Register 用户注册
func Register(c *gin.Context) {
	user := new(models.User)
	user.Username = c.Query("user_name")
	user.Password = encrypt.GetMd5String(c.Query("pass_word"))
	err := user.Register(user.Username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "注册成功",
		})
	}
}
