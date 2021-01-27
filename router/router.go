package router

import (
	"github.com/gin-gonic/gin"
	"jwt/ctrls"
	"jwt/ctrls/user"
	"jwt/middleware"
)

func InitRouter(router *gin.Engine) {

	router.GET("/register", user.Register)
	router.POST("/login", user.Login)
	router.GET("/token", user.GetTokenInfo)

	//用户
	sv1 := router.Group("/user/")
	sv1.Use(middleware.AuthMiddleware())
	{
		sv1.GET("/id", user.FindUserById)
		sv1.GET("/name", user.FindUserByName)
		sv1.GET("/list", user.GetUserList)
		sv1.GET("/logout", user.Logout)
		sv1.GET("/onlineList", user.GetOnlineList)

	}

	//服务器
	sv2 := router.Group("/server/")
	sv2.Use(middleware.AuthMiddleware())
	{
		sv2.GET("/info", ctrls.ServerInfo)
	}
}
