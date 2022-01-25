package user

import (
	"gin-demo/models"
	"gin-demo/utils"
	"gin-demo/utils/encrypt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

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
		token := utils.NewToken(user.Username, strconv.FormatInt(int64(userInfo.Uid), 10))

		_ = models.StrSetEX(token, userInfo.Username+"_"+strconv.FormatInt(int64(userInfo.Uid), 10), time.Second*time.Duration(utils.ExpireTime))
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"msg":    "登陆成功",
			"data":   token,
		})
	}

}
