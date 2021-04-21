package user

import (
	"9YuBlog/models"
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
	"net/http"
)

// GetOnlineList 在线列表
func GetOnlineList(c *gin.Context) {
	Atoken := c.GetHeader("Authorization")
	//userAgent := c.GetHeader("User-Agent")
	userAgent := user_agent.New(c.Request.UserAgent()) // 解析 user-agent
	browser, _ := userAgent.Browser()                  // 浏览器
	os := userAgent.OS()                               // 系统
	host := c.ClientIP()                               // ip

	token, err := models.StrGetRange(Atoken)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"userId":  token,
			"token":   Atoken,
			"browser": browser,
			"os":      os,
			"host":    host,
		})
	}
}
