package router

import (
	"gin-demo/ctrls"
	"gin-demo/ctrls/article"
	"gin-demo/ctrls/user"
	"gin-demo/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {

	router.GET("/register", user.Register)
	router.POST("/login", user.Login)
	router.GET("/token", user.GetTokenInfo)

	// 用户
	u := router.Group("/user/")
	u.Use(middleware.AuthMiddleware())
	{
		u.GET("/id", user.GetInfoById)
		u.GET("/name", user.GetInfoByName)
		u.GET("/list", user.GetUserList)
		u.GET("/logout", user.Logout)

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

}
