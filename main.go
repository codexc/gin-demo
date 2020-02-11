package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"learn/topic/src/App"
	"learn/topic/src/App/services"
	"learn/topic/src/AppInit"
	"learn/topic/src/Models"
)

func main() {
	//创建路由
	router := gin.Default()

	topicRouter := router.Group("/topic")
	{
		//常规写法
		topicRouter.GET("/listold", func(context *gin.Context) {
			topics := Models.TopicList{}
			AppInit.GetDB().Limit(10).Order("id").Find(&topics)
			context.JSON(200, topics)
		})

		//三层写法
		topicService := &services.TopicService{}
		TopicHandler := App.RegisterHandler(
			services.TopicEndPoint(topicService),
			services.SetTopicListRequest(),
			services.SetTopicListResponse(),
		)
		topicRouter.GET("/list", TopicHandler)
	}

	//启动
	router.Run(AppInit.SERVER_ADDR)
}
