package main

import (
	"9YuBlog/middleware"
	"9YuBlog/models"
	"9YuBlog/router"
	"9YuBlog/utils"
	"github.com/gin-gonic/gin"
	"log"
)

var conf = utils.NewCfg().InitConfig() // 初始化配置文件

func main() {

	models.DB = models.InitSQL()    //初始化sql
	models.Rdb = models.InitRedis() //初始化 redis

	//初始化router
	r := gin.Default()
	router.InitRouter(r)
	//pprof.Register(r)
	//使用gin自带的异常恢复中间件，避免出现异常时程序退出
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	host := conf.GetString("ListenOn.Host")
	port := conf.GetString("ListenOn.Port")
	err := r.Run(host + ":" + port)
	if err != nil {
		log.Fatal("服务启动失败 ：", err)
	}
}
