package user

import (
	"gin-demo/models"
	"gin-demo/utils"
	"gin-demo/utils/convert"
	"gin-demo/utils/encrypt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// Login 用户登录
func Login(c *gin.Context) {
	user := new(models.User)
	user.Mail = c.PostForm("mail")
	user.Password = encrypt.GetMd5String(c.PostForm("password"), models.PasswordMd5Key)
	if !user.Login(user) {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "登录失败",
		})
		c.Abort()
		return
	}

	// 获取用户信息用于生成token
	if userInfo, err := models.GetInfoBy("mail", user.Mail); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户不存在",
		})
		c.Abort()
		return
	} else {
		token := utils.NewToken(userInfo.Mail, strconv.FormatInt(int64(userInfo.Uid), 10))

		//_ = models.StrSetEX(token, userInfo.Mail+"_"+strconv.FormatInt(int64(userInfo.Uid), 10), time.Second*time.Duration(utils.ExpireTime))
		value := userInfo.Mail + "_" + convert.IntToString(userInfo.Uid)
		expiration := utils.ExpireTime
		err = models.Rdb.SetEX(models.Ctx, encrypt.GetMd5String(token, utils.TokenMd5Key), value, time.Duration(expiration)).Err()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "登录失败,ERROR=> " + err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"msg":    "登陆成功",
			"data":   token,
		})
	}

}
