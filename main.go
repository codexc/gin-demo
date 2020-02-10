package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	. "learn/topic/src"
)

func main() {
	//创建路由
	router := gin.Default()

	topicRouter := router.Group("/topic")
	{
		topicRouter.Handle("GET", "/id/:id", GetById())
		topicRouter.GET("/list", GetList())
		topicRouter.GET("/getbypage", GetByPage())
		//需要使用鉴权的
		topicRouter.Use(MustLogin())
		topicRouter.POST("/del/:id", DelById) //这个方法为啥没加()?
		topicRouter.POST("create", Create())
		topicRouter.POST("mcreate", Mcreate())
	}

	//启动
	router.Run()
}
