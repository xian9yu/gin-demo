package user

import (
	"github.com/gin-gonic/gin"
	"jwt/middleware"
	"net/http"
)

// TODO
func Logout(c *gin.Context) {
	Authorization := c.Request.Header.Get("Authorization")
	j := middleware.NewJWT()
	token, err := j.LogoutToken(Authorization)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":  0,
			"token": token,
		})
	}
}
