package user

import (
	"9YuBlog/middleware"
	"9YuBlog/models"
	"9YuBlog/utils/encrypt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

//var expireTime int64 = 600 // token有效期(时间戳/s)

// Login 用户登录
func Login(c *gin.Context) {
	user := new(models.User)
	user.Username = c.PostForm("user_name")
	user.Password = encrypt.GetMd5String(c.PostForm("pass_word"))

	if err := user.Login(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "登录失败",
		})
	}

	// 获取用户信息用于生成token
	if userInfo, errs := user.FindByName(user.Username); errs != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户不存在",
		})
	} else {
		j := middleware.NewJWT()
		token, err := j.GenerateToken(c, *userInfo)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    err.Error(),
				"data":   nil,
			})
		}

		_ = models.StrSetEX(encrypt.GetMd5String(token), strconv.FormatInt(userInfo.Id, 10), time.Second*time.Duration(middleware.ExpireTime))
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"msg":    "登陆成功",
			"data":   token,
		})
	}

}
