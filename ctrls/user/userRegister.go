package user

import (
	"gin-demo/models"
	"gin-demo/utils/convert"
	"gin-demo/utils/encrypt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Register 用户注册
func Register(c *gin.Context) {
	mail := c.PostForm("mail")
	password := c.PostForm("password")
	if len(password) < 1 || len(mail) < 1 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "邮箱或密码为空",
		})
		c.Abort()
		return
	}
	user := models.User{
		Uid:         0,
		Username:    "",
		Mail:        mail,
		Password:    encrypt.GetMd5String(password, models.PasswordMd5Key),
		Url:         "",
		Group:       "用户",
		Status:      "启用",
		CreatedTime: 0,
		UpdatedTime: 0,
		LastLogin:   0,
	}

	rowsAffected, uid, err := user.Register(user)
	if err != nil || rowsAffected < 1 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err,
		})
		c.Abort()
		return
	} else {
		if uid == 1 {

		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "注册成功,<uid>=" + convert.IntToString(uid),
		})
	}
}
