package user

import (
	"github.com/gin-gonic/gin"
	"jwt/middleware"
	"jwt/models"
	"jwt/utils"
	"net/http"
)

//var expireTime int64 = 600 // token有效期(时间戳/s)

//用户登录
func Login(c *gin.Context) {
	user := new(models.User)
	user.UserName = c.PostForm("user_name")
	user.PassWord = c.PostForm("pass_word")
	user.PassWord = utils.MD5V(user.PassWord) //在model处理会形成循环调用
	if err := user.Login(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "登录失败",
		})
	} else {

		if fuser, errs := user.FindByName(user.UserName); errs != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "用户不存在",
			})
		} else {
			j := middleware.NewJWT()
			token, err := j.GenerateToken(c, *fuser)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"status": -1,
					"msg":    err.Error(),
					"data":   nil,
				})
			}

			c.JSON(http.StatusOK, gin.H{
				"status": 0,
				"msg":    "登陆成功",
				"data":   token,
			})
		}
	}
}

//// token生成器
//func generateToken(c *gin.Context, user models.User) {
//	// 构造SignKey: 签名和解签名需要使用一个值
//	j := utils.NewJWT()
//
//	// 构造用户claims信息(负荷)
//	claims := utils.Claims{
//		Id:       user.ID,
//		UserName: user.UserName,
//		StandardClaims: jwt.StandardClaims{
//			ExpiresAt: int64(time.Now().Unix() + expireTime), // 签名过期时间(时间戳/s)
//			Issuer:    user.UserName,                         // 签名颁发者
//			IssuedAt:  time.Now().Unix(),                     //签名时间
//		},
//	}
//
//	// 根据claims生成token对象
//	token, err := j.CreateToken(claims)
//
//	if err != nil {
//		c.JSON(http.StatusOK, gin.H{
//			"status": -1,
//			"msg":    err.Error(),
//			"data":   nil,
//		})
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"status": 0,
//		"msg":    "登陆成功",
//		"data":   token,
//	})
//}
