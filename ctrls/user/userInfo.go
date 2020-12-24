package user

import (
	"github.com/gin-gonic/gin"
	"jwt/middleware"
	"jwt/models"
	"net/http"
	"strconv"
)

func FindUserById(c *gin.Context) {
	user := new(models.User)
	ids := c.Query("id")
	idInt64, _ := strconv.Atoi(ids)

	user, err := user.FindById(idInt64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户不存在",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"user": user,
		})
	}
}

func FindUserByName(c *gin.Context) {
	user := new(models.User)
	user.UserName = c.Query("user_name")

	user, err := user.FindByName(user.UserName)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户不存在",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"user": user,
		})
	}
}

func GetTokenInfo(c *gin.Context) {
	Authorization := c.Request.Header.Get("Authorization")

	if Authorization == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "请求未携带token",
		})
		return
	}
	j := middleware.NewJWT()
	claims, err := j.ParseToken(Authorization)
	switch err {
	case middleware.TokenExpired:
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "该token已过期",
		})
	case nil:
		c.JSON(http.StatusOK, gin.H{
			"code":  0,
			"token": claims,
		})
	default:
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Invalid token",
		})
	}
}