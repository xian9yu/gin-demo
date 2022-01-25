package middleware

import (
	"gin-demo/models"
	"gin-demo/utils"
	"gin-demo/utils/encrypt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// AuthMiddleware 中间件
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

		// 判断 token在 redis中是否存在
		isExist, err := models.StrExists(Authorization)
		if !isExist || err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "token失效，请重新登录",
			})
			c.Abort()
			return
		}

		// 获取 token key的有效时间
		timeVal, err := models.Ttl(encrypt.GetMd5String(Authorization))
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "获取token有效时间失败",
			})
		}
		// 判断redis中的token有效时间小于配置中设定时间的一半则更新过期时间
		if timeVal < int(utils.ExpireTime/2) {
			// 更新 token过期时间
			_ = models.StrSetExpireAt(Authorization, time.Now().Unix()+utils.ExpireTime)
		}

		_, err = models.Rdb.Get(models.Ctx, Authorization).Result()
		if err != nil { //TODO
			//if err == TokenExpired {
			//	c.JSON(http.StatusOK, gin.H{
			//		"code": -1,
			//		"msg":  "token授权已过期，请重新登录",
			//	})
			//	c.Abort()
			//	return
			//}
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  err.Error(),
			})
			c.Abort()
			return
		}

		c.Set("claims", Authorization)
	}
}
