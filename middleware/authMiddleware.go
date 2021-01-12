package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		Authorization := c.Request.Header.Get("Authorization")

		if Authorization == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}

		j := NewJWT()
		claims, err := j.ParseToken(Authorization)
		if err != nil {
			if err == TokenExpired {
				c.JSON(http.StatusOK, gin.H{
					"code": -1,
					"msg":  "token授权已过期，请重新登录",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  err.Error(),
			})
			c.Abort()
			return
		}

		//refresh := (claims.StandardClaims.ExpiresAt - int64(time.Now().Unix())) < (ExpireTime / 2)
		//if refresh {
		//	claims.StandardClaims.ExpiresAt = (time.Now().Unix()+1000)
		//	tok, _ := j.CreateToken(claims)
		//	c.JSON(http.StatusOK, gin.H{
		//		"code": 0,
		//		"data":  tok,
		//	})
		//}

		c.Set("claims", claims)
	}
}
