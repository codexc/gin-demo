package App

import (
	"context"
	"github.com/gin-gonic/gin"
)

type GinHandler func(ginCtx *gin.Context)

//业务处理原型函数
type EndPoint func(context context.Context, request interface{}) (response interface{}, err error)

//参数处理原型(gin的context拿参数, 返回interface和err)
type RequestHandler func(ginCtx *gin.Context) (request interface{}, err error)

//返回处理原型
type ResponseHandler func(ginCtx *gin.Context, response interface{}) (err error)

//业务处理流程(参数 业务 返回)
func RegisterHandler(endPoint EndPoint, reqHandler RequestHandler, rspHandler ResponseHandler) func(context *gin.Context) {
	return func(ginCtx *gin.Context) {
		//取参数
		request, err := reqHandler(ginCtx)
		if err != nil {
			ginCtx.JSON(500, gin.H{"error": "params error" + err.Error()})
			return
		}
		//处理业务
		rsp, err := endPoint(ginCtx, request)
		if err != nil {
			ginCtx.JSON(500, gin.H{"error": "business process error" + err.Error()})
			return
		}

		//返回结果
		err = rspHandler(ginCtx, rsp)
		if err != nil {
			ginCtx.JSON(500, gin.H{"error": "response error" + err.Error()})
			return
		}
	}
}
