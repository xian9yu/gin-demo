package user

import (
	"github.com/gin-gonic/gin"
	"jwt/middleware"
	"jwt/models"
	"jwt/utils/encryption"
	"net/http"
	"strconv"
	"time"
)

//var expireTime int64 = 600 // token有效期(时间戳/s)

//用户登录
func Login(c *gin.Context) {
	user := new(models.User)
	user.UserName = c.PostForm("user_name")
	user.PassWord = encryption.GetMd5String(c.PostForm("pass_word"))

	if err := user.Login(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "登录失败",
		})
	} else {
		if fuser, errs := user.FindByName(user.UserName); errs != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "用户不存在",
			})
		} else {
			j := middleware.NewJWT()
			token, err := j.GenerateToken(c, *fuser)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"status": -1,
					"msg":    err.Error(),
					"data":   nil,
				})
			}
			_ = models.StrSetEX(encryption.GetMd5String(token), strconv.FormatInt(fuser.ID, 10), time.Second*time.Duration(middleware.ExpireTime))
			c.JSON(http.StatusOK, gin.H{
				"status": 0,
				"msg":    "登陆成功",
				"data":   token,
			})
		}
	}
}
