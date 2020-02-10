package src

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//练习ctx.Param
//localhost:8080/topic/id/10
func GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		topic := Topic{}
		DBConn.Table("topic").First(&topic, ctx.Param("id"))
		ctx.JSON(http.StatusOK, topic)
	}
}

//练习c.Query c.DefaultQuery
//localhost:8080/topic/list?username=lxc
func GetList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.Query("username")
		page := ctx.DefaultQuery("page", "1")
		size := ctx.DefaultQuery("size", "10")
		if username != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"username": username,
				"page":     page,
				"size":     size,
			})
		} else {
			ctx.String(http.StatusOK, "nobody的list")
		}
	}
}

//练习下参数绑定
//localhost:8080/topic/getbypage?page=1&size=10
func GetByPage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//使用绑定的形式
		query := GetByPageQuery{}
		//注意这里要传址
		if err := ctx.BindQuery(&query); err != nil {
			//todo why? 这里的200没起作用??? 看下面注释的吧, 应该用ShouldBindQuery
			ctx.JSON(200, gin.H{"error": "绑定错误"})
		} else {
			//注意返回的映射
			ctx.JSON(200, query)
		}

		//if err:=ctx.BindQuery(&query);err!=nil{
		//    //todo why? 这里的200没起作用??? 看下面注释的吧, 应该用ShouldBindQuery
		//    ctx.JSON(200,gin.H{"error":"绑定错误"})
		//}else{
		//    //注意返回的映射
		//    ctx.JSON(200,query)
		//}
	}
}

//以下post操作走MustLogin验证中间件了
//练习post 路由中间件
func DelById(ctx *gin.Context) {
	ctx.String(http.StatusOK, "delete id %s", ctx.Param("id"))
}

//练习内置的验证器 这个挺牛逼, 不用代码里各种判断
//wiki:https://godoc.org/gopkg.in/go-playground/validator.v9
func Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		topic := Topic{}
		if err := ctx.Bind(&topic); err == nil {
			ctx.JSON(http.StatusOK, topic)
		} else {
			ctx.String(400, err.Error())
		}
	}
}

//来个批量的
func Mcreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		topics := MultiTopicEntity{}
		if err := ctx.BindJSON(&topics); err == nil {
			ctx.JSON(http.StatusOK, topics)
		} else {
			ctx.String(400, err.Error())
		}
	}
}
