package router

import (
	"github.com/gin-gonic/gin"
	"jwt/ctrls"
	"jwt/ctrls/article"
	"jwt/ctrls/user"
	"jwt/middleware"
)

func InitRouter(router *gin.Engine) {

	router.GET("/register", user.Register)
	router.POST("/login", user.Login)
	router.GET("/token", user.GetTokenInfo)

	// 用户
	u := router.Group("/user/")
	u.Use(middleware.AuthMiddleware())
	{
		u.GET("/id", user.FindUserById)
		u.GET("/name", user.FindUserByName)
		u.GET("/list", user.GetUserList)
		u.GET("/logout", user.Logout)
		u.GET("/onlineList", user.GetOnlineList)

	}
	// 文章
	a := router.Group("/article/")
	a.Use(middleware.AuthMiddleware())
	{
		a.POST("/add", article.Add)

	}
	//服务器
	s := router.Group("/server/")
	s.Use(middleware.AuthMiddleware())
	{
		s.GET("/info", ctrls.ServerInfo)
	}
	//file
	file := router.Group("/files/")
	file.Use(middleware.AuthMiddleware())
	{
		//file.POST("/upload", ctrls.Upload)
		//file.GET("/download", ctrls.Download)
	}
}
