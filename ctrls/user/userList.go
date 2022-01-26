package user

import (
	"gin-demo/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUserList 获取用户列表
func GetUserList(c *gin.Context) {
	var user models.User
	total, users, err := user.GetAll()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		c.Abort()
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":  200,
			"total": total,
			"user":  users,
		})
	}
}
