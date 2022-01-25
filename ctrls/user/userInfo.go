package user

import (
	"gin-demo/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetInfoById 通过id获取用户info
func GetInfoById(c *gin.Context) {
	ids := c.Query("id")
	idInt64, _ := strconv.Atoi(ids)

	userInfo, err := models.GetInfoBy("id", idInt64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户不存在",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"user": userInfo,
		})
	}
}

// GetInfoByName 通过用户名获取用户info
func GetInfoByName(c *gin.Context) {
	user := new(models.User)
	user.Username = c.Query("username")

	userInfo, err := models.GetInfoBy("username", user.Username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户不存在",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"user": userInfo,
		})
	}
}

// GetTokenInfo 解析toke
func GetTokenInfo(c *gin.Context) {
	Authorization := c.Request.Header.Get("Authorization")

	if len(Authorization) < 1 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "请求未携带token",
		})
		return
	}

	userId, err := models.Rdb.Get(models.Ctx, Authorization).Result()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Invalid token",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"token":  Authorization,
		"userId": userId,
	})
}
