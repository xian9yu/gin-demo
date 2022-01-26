package user

import (
	"gin-demo/models"
	"gin-demo/utils/encrypt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Register 用户注册
func Register(c *gin.Context) {
	if len(c.Query("password")) < 1 || len(c.Query("mail")) < 1 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "邮箱或密码为空",
		})
	}
	user := new(models.User)
	user.Mail = c.Query("mail")
	user.Password = encrypt.GetMd5String(c.Query("password"))
	rowsAffected, uid, err := user.Register(user)
	if err != nil || rowsAffected < 1 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "注册成功<uid>=" + strconv.FormatUint(uid, 10),
		})
	}
}
