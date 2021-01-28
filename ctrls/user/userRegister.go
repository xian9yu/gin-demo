package user

import (
	"github.com/gin-gonic/gin"
	"jwt/models"
	"jwt/utils/encrypt"
	"net/http"
)

//用户注册
func Register(c *gin.Context) {
	user := new(models.User)
	user.UserName = c.Query("user_name")
	user.PassWord = encrypt.GetMd5String(c.Query("pass_word"))
	err := user.Register(user.UserName)
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
