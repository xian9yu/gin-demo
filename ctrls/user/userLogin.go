package user

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"jwt/middleware"
	"jwt/models"
	"jwt/utils"
	"net/http"
	"time"
)

//var expireTime int64 = 600 // token有效期(时间戳/s)

//用户登录
func Login(c *gin.Context) {
	user := new(models.User)
	user.UserName = c.PostForm("user_name")
	user.PassWord = c.PostForm("pass_word")
	user.PassWord = utils.MD5V(user.PassWord) //在model处理会形成循环调用
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
			_ = models.StrSet(GetMd5String(token), user.UserName, time.Hour*24)
			c.JSON(http.StatusOK, gin.H{
				"status": 0,
				"msg":    "登陆成功",
				"data":   token,
			})
		}
	}
}

func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
