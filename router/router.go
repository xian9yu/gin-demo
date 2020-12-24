package router

import (
	"github.com/gin-gonic/gin"
	"jwt/ctrls/user"
	"jwt/middleware"
)

func InitRouter(router *gin.Engine) {

	router.GET("/register", user.Register)
	router.Any("/login", user.Login)
	router.GET("/token", user.GetTokenInfo)

	sv1 := router.Group("/user/")
	sv1.Use(middleware.AuthMiddleware())
	{
		sv1.GET("/id", user.FindUserById)
		sv1.GET("/name", user.FindUserByName)
		sv1.GET("/list", user.GetUserList)
		sv1.GET("/logout", user.Logout)
	}
}
